package util

import (
	"testing"
)

func Test_newConfigManager(t *testing.T) {
	manager := newConfigManager()
	if manager == nil {
		t.Fatal("newConfigManager failed")
	}
}
