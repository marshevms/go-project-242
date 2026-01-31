package du

import (
	"os"
)

// GetSize returns size of a file or directory by given path in bytes.
func GetSize(path string) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		info, err := file.Info()
		if err != nil {
			return 0, err
		}

		size += info.Size()
	}

	return size, nil
}
