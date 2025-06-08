package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	logrus.Debug("test debug")
	logrus.Debugf("test debug: %s", "string")

	logrus.Warn("test warn")
	logrus.Warnf("test warn: %s", "string")

	logrus.Info("test info")
	logrus.Infof("test info: %s", "string")

	logrus.Error("test error")
	logrus.Errorf("test error: %s", "string")
}
