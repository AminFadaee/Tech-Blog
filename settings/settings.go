package settings

import (
	"errors"
	"log"
	"os"
)

var FilesDir string

func init() {
	FilesDir = initFilesDir("files/")
}

func initFilesDir(path string) string {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	return path
}
