package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	GOOS       = runtime.GOARCH
	AppName    string
	AppVersion string
	NEWLINE    = "\n"
	Ipv4Regex  = `([0-9]+\.){3}[0-9]+`
)

func init() {
	GOOS := runtime.GOOS
	if GOOS != "windows" {
		NEWLINE = "\r\n"
	} else if GOOS != "darwin" {
		NEWLINE = "\r"
	}
}

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

/* add data to PATH , vd: /home/mannk:path */
func PATHJointList(PATH, data string) string {
	//	data = data + string(os.PathSeparator)
	if len(PATH) == 0 {
		return data
	}
	return PATH + string(os.PathListSeparator) + data
	//	filepath.ListSeparator
}

/* remove addpath from path*/
func PATHRemove(PATH, addpath string) string {
	if len(PATH) == 0 {
		return ""
	}
	newpath := ""
	for i, val := range strings.Split(PATH, string(os.PathListSeparator)) {
		if !(val == addpath) {
			if i == 0 {
				newpath = val
			} else {
				newpath = newpath + string(os.PathListSeparator) + val
			}
		}
	}
	return newpath
	//	filepath.ListSeparator
}

func PATHGetEnvPathValue() string {
	for _, pathname := range []string{"PATH", "path"} {
		path := os.Getenv(pathname)
		if len(path) != 0 {
			return path
		}
	}
	return ""
}

/* return PATH as array (:)*/
func PATHArr() []string {
	envs := PATHGetEnvPathValue()
	if len(envs) != 0 {
		return strings.Split(envs, string(os.PathListSeparator))
	}
	return []string{}
}

/* retur path or PATH*/
func PathGetEnvPathKey() string {
	for _, pathname := range []string{"PATH", "path"} {
		path := os.Getenv(pathname)
		if len(path) != 0 {
			return pathname
		}
	}
	return ""
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
	if err != nil {
		return 0, err
	}
	defer file.Close()
	if err != nil {
		return 0, err
	}
	n, err := file.Write(data)
	return n, err
}

// open and attend to file
func FileOpen2Write(fullPath string) (*os.File, error) {
	// create dir to path if it is not exist
	err := DirCreate(filepath.Dir(fullPath), 0775)
	if err != nil {
		return nil, err
	}
	// create file if it not existed
	if os.Stat(fullPath); err != nil {
		err = FileCreate(fullPath)
		if err != nil {
			return nil, err
		}
	}

	logf, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return logf, err
}

func FileReadAll(fullPath string) (string, error) {
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(content), err
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

/* only read filesize, not dir*/
func FileGetSize(filepath string) (int64, error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}

/* write content to file if it different with exist file content*/
func FileWriteStringIfChange(pathfile string, contents []byte) (bool, error) {

	oldContents := []byte{}
	if _, err := os.Stat(pathfile); err == nil {
		oldContents, _ = ioutil.ReadFile(pathfile)
	}

	//if bytes.Compare(oldContents, contents) != 0 {
	if !bytes.Equal(oldContents, contents) {
		return true, ioutil.WriteFile(pathfile, contents, 0644)
	} else {
		return false, nil
	}
}

/* cp source to dir, keep modify time */
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

/* insert string at index lines of file, if this line have text, push it one line down. */
func FileInsertStringAtLine(filePath, str string, index int) error {
	NEWLINE := "\n"
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	str = str + NEWLINE //add newline
	scanner := bufio.NewScanner(f)
	lines := ""
	linenum := 0
	inserted := false
	for scanner.Scan() {
		linenum = linenum + 1
		if linenum == index {
			inserted = true
			lines = lines + str
		}
		lines = lines + scanner.Text() + NEWLINE
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if !inserted {
		if index == -1 {
			index = linenum + 1
		}
		for i := linenum + 1; i < index; i++ {
			lines = lines + NEWLINE
		}
		lines = lines + str
	}
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, []byte(lines), info.Mode().Perm())
}

/* create tempdir and return tempfile in tempdir (not create yet)*/
func FileTempCreateInNewTemDir(filename string) string {

	rootdir, err := ioutil.TempDir("", "system")
	if err != nil {
		return ""
	} else {
		//			defer os.RemoveAll(dir)
	}
	return filepath.Join(rootdir, filename)
}

func FileTempCreateInNewTemDirWithContent(filename string, data []byte) string {
	rootdir, err := ioutil.TempDir("", "system")
	if err != nil {
		return ""
	}
	fPath := filepath.Join(rootdir, filename)
	err = os.WriteFile(fPath, data, 0755)
	if err != nil {
		os.RemoveAll(rootdir)
		return ""
	}
	return fPath
}

func walk(filename string, linkDirname string, walkFn filepath.WalkFunc) error {
	symWalkFunc := func(path string, info os.FileInfo, err error) error {

		if fname, err := filepath.Rel(filename, path); err == nil {
			path = filepath.Join(linkDirname, fname)
		} else {
			return err
		}

		if err == nil && info.Mode()&os.ModeSymlink == os.ModeSymlink {
			finalPath, err := filepath.EvalSymlinks(path)
			if err != nil {
				return err
			}
			info, err := os.Lstat(finalPath)
			if err != nil {
				return walkFn(path, info, err)
			}
			if info.IsDir() {
				return walk(finalPath, path, walkFn)
			}
		}

		return walkFn(path, info, err)
	}
	return filepath.Walk(filename, symWalkFunc)
}

// Walk extends filepath.Walk to also follow symlinks
func SymWalk(path string, walkFn filepath.WalkFunc) error {
	return walk(path, path, walkFn)
}

/*
Dir
*/
func DirCreate(dirPath string, permission fs.FileMode) error {
	dirFullPath := dirPath
	// if _, err := os.Stat(dirFullPath); err == nil {
	// 	if err := os.RemoveAll(dirFullPath); err != nil {
	// 		//fmt.Println("Error removing existing directory:", err)
	// 		return err
	// 	}
	// }
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

/* Remove all content of dir */
func DirRemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
