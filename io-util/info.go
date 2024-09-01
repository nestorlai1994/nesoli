package ioutil

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	cmdExecutor "github.com/nestorlai1994/nesoli/cmd-executor"
)

func CheckOrCreate(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func CheckOrCreateFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func IsFolderEmpty(path string) (bool, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return len(files) == 0, nil
}

func Tree(path string, level int) {
	// fmt.Println(path)
	if level < 1 {
		cmdExecutor.ExecuteCommand("tree", path)
	} else {
		cmdExecutor.ExecuteCommand("tree", path, "-L", strconv.Itoa(level))
	}
}

func FileNameOnly(path string) string {
	_, file := filepath.Split(path)
	return file[:len(file)-len(filepath.Ext(file))]
}
