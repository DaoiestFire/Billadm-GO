package util

import (
	"math/rand"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func GetUUID() string {
	uuidObj, _ := uuid.NewV7()
	return uuidObj.String()
}

func GetTestDir() string {
	dir, _ := os.Getwd()
	dir = filepath.Dir(dir)
	return filepath.Join(dir, "test")
}

func GetRandomString(l int) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
