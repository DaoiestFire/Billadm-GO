package workspace

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWsManager_InitFromConfig_init(t *testing.T) {
	pwd, _ := os.Getwd()
	testDir := filepath.Join(filepath.Dir(pwd), "test")
	t.Log(testDir)
	err := Manager.InitFromConfig(testDir)
	if err != nil {
		t.Fatal(err)
	}

	ws1 := filepath.Join(testDir, "workspace1")
	err = Manager.OpenWorkspace(ws1)
	if err != nil {
		t.Fatal(err)
	}

	ws2 := filepath.Join(testDir, "workspace2")
	err = Manager.OpenWorkspace(ws2)
	if err != nil {
		t.Fatal(err)
	}

	err = Manager.SetOpenedWorkspace(ws2)
	if err != nil {
		t.Fatal(err)
	}

	err = Manager.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestWsManager_InitFromConfig_read(t *testing.T) {
	pwd, _ := os.Getwd()
	testDir := filepath.Join(filepath.Dir(pwd), "test")
	t.Log(testDir)
	err := Manager.InitFromConfig(testDir)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(Manager.workspaces)
	t.Log(Manager.openedWorkspace)

	ws3 := filepath.Join(testDir, "workspace3")
	err = Manager.OpenWorkspace(ws3)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(15 * time.Second)
}
