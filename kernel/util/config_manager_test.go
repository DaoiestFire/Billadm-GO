package util

import (
	"testing"
)

func Test_newConfigManager(t *testing.T) {
	_, err := newConfigManager()
	if err != nil {
		t.Fatal("newConfigManager failed")
	}
}
