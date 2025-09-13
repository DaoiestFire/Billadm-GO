package server

import (
	"github.com/billadm/workspace"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

var manager *ExitManager

// ExitManager 简单的程序退出管理器
type ExitManager struct {
	once sync.Once
	done chan struct{} // 用于通知已关闭
}

// NewExitManager 创建一个新的 ExitManager
func NewExitManager() *ExitManager {
	if manager != nil {
		return manager
	}
	manager = &ExitManager{
		done: make(chan struct{}),
	}

	return manager
}

// Start 启动信号监听
func (em *ExitManager) Start() {
	logrus.Info("启动程序退出处理协程")
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Kill, os.Interrupt, syscall.SIGTERM)
		select {
		case sig := <-sigCh:
			logrus.Warnf("billadm exit manager shutdown by signal %s", sig.String())
			em.shutdown()
			return
		}
	}()
}

// Exit 显式触发关闭流程
func (em *ExitManager) Exit() {
	logrus.Warn("主动关闭Billadm")
	em.shutdown()
}

// shutdown 执行实际的关闭逻辑（只执行一次）
func (em *ExitManager) shutdown() {
	em.once.Do(func() {
		workspace.Manager.Close()
		os.Exit(0)
	})
}
