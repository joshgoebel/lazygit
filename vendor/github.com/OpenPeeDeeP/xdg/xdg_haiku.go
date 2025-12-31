//go:build haiku
package xdg

import (
	"os"
	"path/filepath"
)

//type osDefaulter struct{}

func (o *osDefaulter) defaultDataHome() string {
	return filepath.Join(os.Getenv("HOME"), ".local", "share")
}

func (o *osDefaulter) defaultDataDirs() []string {
	return []string{"/boot/system/data", "/boot/common/data"}
}

func (o *osDefaulter) defaultConfigHome() string {
	return filepath.Join(os.Getenv("HOME"), ".config")
}

func (o *osDefaulter) defaultConfigDirs() []string {
	return []string{"/boot/system/settings", "/etc/xdg"}
}

func (o *osDefaulter) defaultCacheHome() string {
	return filepath.Join(os.Getenv("HOME"), ".cache")
}

