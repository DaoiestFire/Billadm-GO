package workspace

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/billadm/util"
)

func init() {
	Manager = &WsManager{
		workspaces: make(map[string]*Workspace),
		closeCh:    make(chan struct{}),
	}
}

var Manager *WsManager

type billadmConfig struct {
	OpenedWorkspace string   `json:"openedWorkspace"`
	WorkspaceDirs   []string `json:"workspaceDirs"`
}
type WsManager struct {
	lock            sync.Mutex
	workspaces      map[string]*Workspace
	homeDir         string
	openedWorkspace string
	closeCh         chan struct{}
}

func (wm *WsManager) InitFromConfig(homeDir string) error {
	wm.homeDir = homeDir

	configFilePath := path.Join(homeDir, ".billadm")
	if util.IsFileExists(configFilePath) {
		bytes, err := os.ReadFile(configFilePath)
		if err != nil {
			return err
		}

		var config billadmConfig
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			return err
		}

		for _, dir := range config.WorkspaceDirs {
			err = wm.OpenWorkspace(dir)
			if err != nil {
				logrus.Errorf("open workspace %s error: %v", dir, err)
			}
		}
		wm.openedWorkspace = config.OpenedWorkspace
	}

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-wm.closeCh:
				logrus.Debug("退出用户配置刷新定时任务")
				return
			case <-ticker.C:
				func() {
					logrus.Debug("刷新用户配置: .billadm")
					wm.lock.Lock()
					defer wm.lock.Unlock()
					if err := wm.refreshWorkspace(); err != nil {
						logrus.Errorf("refresh workspace error: %v", err)
					}
					if err := wm.saveToConfig(); err != nil {
						logrus.Errorf("save workspace error: %v", err)
					}
				}()
			}
		}
	}()

	return nil
}

func (wm *WsManager) OpenWorkspace(directory string) error {
	wm.lock.Lock()
	defer wm.lock.Unlock()

	_, ok := wm.workspaces[directory]
	if ok {
		return fmt.Errorf("workspace %s already exists", directory)
	}

	ws, err := NewWorkspace(directory)
	if err != nil {
		return err
	}
	wm.workspaces[directory] = ws

	return nil
}

func (wm *WsManager) SetOpenedWorkspace(directory string) error {
	wm.lock.Lock()
	defer wm.lock.Unlock()

	_, ok := wm.workspaces[directory]
	if !ok {
		wm.openedWorkspace = ""
		return fmt.Errorf("workspace %s not exists", directory)
	}
	wm.openedWorkspace = directory

	return nil
}

func (wm *WsManager) GetWorkspaceByPath(directory string) (*Workspace, error) {
	wm.lock.Lock()
	defer wm.lock.Unlock()

	ws, ok := wm.workspaces[directory]
	if !ok {
		return nil, fmt.Errorf("workspace %s not exists", directory)
	}

	return ws, nil
}

func (wm *WsManager) RemoveWorkspaceByPath(directory string) (*Workspace, error) {
	wm.lock.Lock()
	defer wm.lock.Unlock()

	ws, ok := wm.workspaces[directory]
	if !ok {
		return nil, fmt.Errorf("workspace %s not exists", directory)
	}

	delete(wm.workspaces, directory)
	ws.Close()

	return ws, nil
}

func (wm *WsManager) OpenedWorkspace() *Workspace {
	wm.lock.Lock()
	defer wm.lock.Unlock()

	ws, ok := wm.workspaces[wm.openedWorkspace]
	if !ok {
		return nil
	}

	return ws
}

func (wm *WsManager) AllWorkspaces() []string {
	wm.lock.Lock()
	defer wm.lock.Unlock()

	var list []string
	for _, ws := range wm.workspaces {
		list = append(list, ws.directory)
	}
	return list
}

// 定时任务：删除不存在的工作空间；刷新用户配置文件
func (wm *WsManager) refreshWorkspace() error {
	var needDelete []string
	for dir, ws := range wm.workspaces {
		if util.IsDirectoryExists(ws.directory) {
			continue
		}
		needDelete = append(needDelete, dir)
	}

	for _, dir := range needDelete {
		delete(wm.workspaces, dir)
	}

	return nil
}

func (wm *WsManager) saveToConfig() error {
	var workspaceDirs []string
	for _, ws := range wm.workspaces {
		workspaceDirs = append(workspaceDirs, ws.directory)
	}

	res := billadmConfig{
		OpenedWorkspace: wm.openedWorkspace,
		WorkspaceDirs:   workspaceDirs,
	}

	configFilePath := path.Join(wm.homeDir, ".billadm")
	bytes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return os.WriteFile(configFilePath, bytes, 0644)
}

func (wm *WsManager) Close() error {
	wm.lock.Lock()
	defer wm.lock.Unlock()

	close(wm.closeCh)

	if err := wm.refreshWorkspace(); err != nil {
		return err
	}

	if err := wm.saveToConfig(); err != nil {
		return err
	}

	for _, ws := range wm.workspaces {
		ws.Close()
	}
	return nil
}
