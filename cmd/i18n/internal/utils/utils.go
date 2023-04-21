package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// IsDirectory reports whether the named file is a directory.
func IsDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

func CreateFolder(folderPath string) error {
	// log.Info().Str("path", folderPath).Msg("create folder")
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		// log.Error().Err(err).Msg("create folder failed")
		return err
	}
	return nil
}

// GetFilesBySuffix .
func GetFilesBySuffix(dirPth, suffix string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	for _, fi := range dir {
		if fi.IsDir() {
			subFiles, err := GetFilesBySuffix(dirPth+"/"+fi.Name(), suffix)
			if err == nil {
				files = append(files, subFiles...)
			}
		} else {
			ok := strings.HasSuffix(fi.Name(), suffix)
			if ok {
				files = append(files, dirPth+"/"+fi.Name())
			}
		}
	}
	return files, nil
}
