package utils

import "os"

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func IsDir(dirname string) bool {
	f, err := os.Stat(dirname)
	if err != nil {
		return false
	}
	if f.IsDir() {
		return true
	}
	return false
}

func IsFile(filename string) bool {
	f, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if f.IsDir() {
		return false
	}
	return true
}
