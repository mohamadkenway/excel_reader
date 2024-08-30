package util

import (
	"os"
	"strings"
)

type FileManager struct{}

func (_ FileManager) GetFilePathsFromDir(path string) ([]string, error) {

	var list []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return list, err
	}

	for _, e := range entries {
		if e.IsDir() {
			_list, err := FileManager{}.GetFilePathsFromDir(path + "/" + e.Name())
			if err != nil {
				return list, err
			}
			list = append(list, _list...)
		} else {
			if strings.Contains(e.Name(), "xlsx") == false || strings.Contains(e.Name(), "._") || strings.Contains(e.Name(), ".~") {
				continue
			}
			list = append(list, path+"/"+e.Name())
		}
	}

	return list, nil
}
