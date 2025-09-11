package workspace

import "sync"

func init() {
	Manager = &WsManager{
		workspaces: make(map[string]*Workspace),
	}
}

var Manager *WsManager

type WsManager struct {
	lock       sync.Mutex
	workspaces map[string]*Workspace
	homeDir    string
}

func (wm *WsManager) InitFromConfig(homeDir string) error {

}

func (wm *WsManager) OpenWorkspace() error {

}

func (wm *WsManager) GetWorkspaceByPath() (*Workspace, error) {

}

func (wm *WsManager) RemoveWorkspaceByPath() (*Workspace, error) {

}

// 定时任务：删除不存在的工作空间；刷新用户配置文件
func (wm *WsManager) refreshWorkspace() error {

}

func (wm *WsManager) saveToConfig() error {

}

func (wm *WsManager) Close() error {

}
