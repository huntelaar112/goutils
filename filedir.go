package goutils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

/* Path */
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

/* File */
func FileCreateWithPath(fullPath string) error {
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return err
}

func FileIsWriteable(path string) (isWritable bool) {
	isWritable = false

	if file, err := os.OpenFile(path, os.O_WRONLY, 0666); err == nil {
		defer file.Close()
		isWritable = true
	} else {
		if os.IsPermission(err) {
			return false
		}
	}
	return
}

/* cp date-time modify from dst file to src file */
func FileCloneDate(dst, src string) bool {
	var err error
	var srcinfo os.FileInfo
	if srcinfo, err = os.Stat(src); err == nil {
		if err = os.Chtimes(dst, srcinfo.ModTime(), srcinfo.ModTime()); err == nil {
			return true
		}
	}
	//	fmt.Errorf("Cannot clone date file ", err)
	return false
}

/* list all dir child (1 level) */
func FileAllChild(directory string) (err error) {
	dirRead, _ := os.Open(directory)
	dirFiles, _ := dirRead.Readdir(0)
	for index := range dirFiles {
		fileHere := dirFiles[index]

		// Get name of file and its full path.
		nameHere := fileHere.Name()
		fullPath := directory + nameHere

		// Remove the file.
		os.Remove(fullPath)
		fmt.Println("Removed file:", fullPath)
	}
	return nil
}

// removeFile removes the specified file. Errors are ignored.
func FileRemoveFile(path string) error {
	return os.Remove(path)
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

/* Dir */
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
