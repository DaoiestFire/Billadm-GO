package logger

import "testing"

func TestLogger(t *testing.T) {
	Debug("test debug")
	Debug("test debug: %s", "string")

	Warn("test warn")
	Warn("test warn: %s", "string")

	Info("test info")
	Info("test info: %s", "string")

	Error("test error")
	Error("test error: %s", "string")
}
