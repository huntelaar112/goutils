package goutils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func PathBaseName(filePath string) string {
	return filepath.Base(filePath)
}

func PathIsExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func FileCreateWithPath(fullPath string) error {
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return err
}

func FileCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file ", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if err == nil {
		os.Chmod(dst, sourceFileStat.Mode())
		os.Chtimes(dst, sourceFileStat.ModTime(), sourceFileStat.ModTime())
	}
	return nBytes, err
}

func DirCreate(dirPath string, permission fs.FileMode) error {
	dirFullPath := dirPath
	if _, err := os.Stat(dirFullPath); err == nil {
		if err := os.RemoveAll(dirFullPath); err != nil {
			//fmt.Println("Error removing existing directory:", err)
			return err
		}
	}
	/*err := os.MkdirAll(dirFullPath, os.ModePerm)*/
	err := os.MkdirAll(dirFullPath, permission)
	if err != nil {
		//fmt.Println("Error creating the directory:", err)
		return err
	}
	return err
	//fmt.Println("Directory created successfully at:", fullPath)
}

func DirRemove(dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}
	return err
}
