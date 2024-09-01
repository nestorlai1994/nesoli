package ioutil

import (
	"io"
	"os"
	"path/filepath"

	"github.com/nestorlai1994/nesoli/logger"
)

func Copy(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		logger.Error("failed to open source file: "+src, err)
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		logger.Error("failed to create destination file: "+dest, err)
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		logger.Error("failed to copy file: ", err)
		return err
	}

	return nil
}

func CopyDir(src, dest string) error {
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		logger.Error("failed to create destination directory: "+dest, err)
		return err
	}

	files, err := os.ReadDir(src)
	if err != nil {
		logger.Error("failed to read source directory: "+src, err)
		return err
	}

	for _, file := range files {
		srcFile := filepath.Join(src, file.Name())
		destFile := filepath.Join(dest, file.Name())

		if file.IsDir() {
			err = CopyDir(srcFile, destFile)
			if err != nil {
				return err
			}
		} else {
			err = Copy(srcFile, destFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
