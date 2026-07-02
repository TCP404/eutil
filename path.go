package eutil

import (
	"os"
	"path/filepath"
	"strings"
)

func ParseFilePath(path string) (dir, filename, filetype string, err error) {
	// 转换为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", "", "", err
	}

	info, err := os.Stat(absPath)
	if err == nil && info.IsDir() {
		return absPath, "", "", nil
	}

	base := filepath.Base(absPath)

	if strings.HasPrefix(base, ".") && strings.Count(base, ".") == 1 {
		return filepath.Dir(absPath), base, "", nil
	}

	ext := filepath.Ext(base)
	filename = strings.TrimSuffix(base, ext)
	return filepath.Dir(absPath), filename, ext, nil
}
