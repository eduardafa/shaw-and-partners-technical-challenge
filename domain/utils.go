package domain

import (
	"path/filepath"
	"runtime"
)

type Utils struct{}

// RootPath returns the project root path
func (u Utils) RootPath() string {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Dir(b)
	root := filepath.Dir(path)
	return root
}
