package util

import (
	"math/rand"
	"path/filepath"
	"runtime"

	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func GetUUID() string {
	uuidObj, _ := uuid.NewV7()
	return uuidObj.String()
}

func GetTestDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filepath.Dir(filename))
	return filepath.Join(dir, "test")
}

func GetRandomString(l int) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
