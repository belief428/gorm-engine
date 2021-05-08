package tools

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDir(path string) bool {
	s, err := os.Stat(path)

	if err != nil {
		return false
	}
	return s.IsDir()
}

func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
