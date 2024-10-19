package utils

import (
	"path/filepath"
	"strings"
)

func FixExtension(databaseFile string) string {
	databaseFile = strings.TrimSuffix(databaseFile, filepath.Ext(databaseFile))
	return databaseFile + ".db"
}
