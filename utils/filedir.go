package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
Path
*/
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

func PathIsFile(path string) bool {
	if finfo, err := os.Stat(path); err == nil {
		if !finfo.IsDir() {
			return true
		}
	}
	return false
}

func PathIsDir(path string) bool {
	if finfo, err := os.Stat(path); err == nil {
		if finfo.IsDir() {
			return true
		}
	}
	return false
}

/*
File
*/
func FileCreate(fullPath string) error {
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return err
}

func FileCreateWithContent(fullPath string, data []byte) (bytewrite int, err error) {
	file, err := os.Create(fullPath)
	defer file.Close()
	if err != nil {
		return 0, err
	}
	n, err := file.Write(data)
	return n, err
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

/* get Md5 string of file */
func FileHashMd5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

/*
	waitForFile waits for the specified file to exist before returning. If the an

error, other than the file not existing, occurs, the error is returned. If,
after 100 attempts, the file does not exist, an error is returned.
*/
func FileWaitForFileExist(path string, timeoutms int) error {
	if timeoutms < 50 && timeoutms != 0 {
		timeoutms = 50
	}

	for i := 0; i < timeoutms/50; i++ {
		_, err := os.Stat(path)
		if err == nil || !os.IsNotExist(err) {
			return err
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("file does not exist: %s", path)
}

/*
	readFile waits for the specified file to contain contents, and then returns

those contents as a string. If an error occurs while reading the file, the
error is returned. If the file has no content after 100 attempts, an error is
returned.
*/
func FileWaitContentsAndRead(path string, timeoutms int) (string, error) {
	if timeoutms < 50 && timeoutms != 0 {
		timeoutms = 50
	}
	for i := 0; i < timeoutms/50; i++ {
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}
		if len(bytes) > 0 {
			return strings.TrimSpace(string(bytes)), err
		}
		time.Sleep(50 * time.Millisecond)
	}
	return "", fmt.Errorf("file is empty: %s", path)
}

/* removeFile removes the specified file. Errors are ignored.*/
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

/* only read filesize, not dir*/
func FileGetSize(filepath string) (int64, error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}

/*
Dir
*/
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

/* list all file and dir int level1 of dir*/
func DirAllChild(directory string) (files []string, err error) {
	dirRead, err := os.Open(directory)
	if err != nil {
		return nil, err
	}
	dirFiles, err := dirRead.Readdir(0)
	if err != nil {
		return nil, err
	}
	for index := range dirFiles {
		fileHere := dirFiles[index]

		// Get name of file and its full path.
		nameHere := fileHere.Name()
		/*		fmt.Println(nameHere)*/
		/*fullPath := directory + nameHere*/
		fullPath := filepath.Join(directory, nameHere)
		files = append(files, fullPath)
		// Remove the file.
	}
	return files, err
}
