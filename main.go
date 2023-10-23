package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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

	/* test dir remove content */
	err = utils.DirRemoveContents("testdir1")
	if err != nil {
		fmt.Println(err)
	}

	/*GenerateRandomBytes*/
	randbytes, err := utils.GenerateRandomBytes(5)
	fmt.Println(randbytes)

	/* linux cat*/
	_, err = utils.FileCreateWithContent("testfile.txt", []byte("you know the rest\n"))
	if err != nil {
		log.Error(err)
	}
	err = utils.Cat("testfile.txt", "gitupdate.sh")
	if err != nil {
		log.Error(err)
	}

	/*ID*/
	teststr := utils.GenerateID()
	log.Info(teststr)
	teststr, err = utils.GenerateIdv4("")
	log.Info(teststr)

	/* which */
	fipath, isexist := utils.Which("/usr/bin/ssh", os.Getenv("PATH"))
	log.Info(fipath, ", status: ", isexist)

	/*PathGetEnvPathValue*/
	log.Info(utils.PATHGetEnvPathValue())

	_, err = os.Create(utils.FileTempCreateInNewTemDir("file.txt"))
	log.Info("Create temp file name: ", utils.FileTempCreateInNewTemDir("file.txt"))

	log.Info("Sha1 of \"mannk-dz\" is: ", utils.GenerateSha1Bytes([]byte("mannk-dz")))

	tokensha1 := utils.GenerateTokenSha1(1998)
	log.Info("Create toke sha1 timestamp: ", tokensha1)
	//time.Sleep(15 * time.Second)
	log.Info(utils.TokenSha1IsMatch(1998, tokensha1))

	check := utils.Cp("gitupdate.sh", filepath.Join(".", "testdir", "gitupdate.sh"))
	if check {
		log.Info("Cp success.")
	}

	lines := utils.String2lines("we do thee right thing\n and that is")
	log.Info(lines)

	stringarr, _ := utils.File2lines("gitupdate.sh")
	log.Info(stringarr)
	log.Info(stringarr[1])
}
