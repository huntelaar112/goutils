package main

import (
	"fmt"
	"goutils/utils"
	"os"
	"path/filepath"
)

func main() {
	/* test write able*/
	path := filepath.Join(".", "README.md")
	if utils.FileIsWriteable(path) {
		fmt.Println("write able yes")
	} else {
		fmt.Println("writeable no")
	}

	/*test file all child*/
	dirpath := filepath.Join(".", "testdir")
	files, err := utils.DirAllChild(dirpath)
	fmt.Println("list file: ", files)
	if err != nil {
		fmt.Println(err)
	}
	for _, path := range files {
		os.Chmod(path, 0777)
	}

	/*file hast md5*/
	returnMD5String, err := utils.FileHashMd5("file1.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(returnMD5String)
	}

	/*test file wait content*/
	content, err := utils.FileWaitContentsAndRead("testfile.txt", 10000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}

	/* test file size*/
	size, err := utils.FileGetSize("gitupdate.sh")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(size)
	}

	err = utils.FileInsertStringAtLine("file1.txt", "pro vip", 30)
	if err != nil {
		fmt.Println(err)
	}

}
