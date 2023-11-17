package main

import (
	"fmt"
	"github.com/huntelaar112/goutils/sched"
	"github.com/huntelaar112/goutils/timeutils"
	"github.com/huntelaar112/goutils/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
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
	returnMD5String, err := utils.FileHashMd5("testdir/file1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(returnMD5String)
	}

	/*test file wait content*/
	content, err := utils.FileWaitContentsAndRead("testdir/file1", 10000)
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

	err = utils.FileInsertStringAtLine("testdir/file1", "pro vip", 30)
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
	_, err = utils.FileCreateWithContent("testdir/testfile.txt", []byte("you know the rest\n"))
	if err != nil {
		log.Error(err)
	}
	err = utils.Cat("testdir/testfile.txt", "gitupdate.sh")
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

	multiline := `line \n
by line
and line
after line`
	strarray := utils.String2lines(multiline)
	for i, str := range strarray {
		fmt.Println("index", i, ":", str)
	}

	fmt.Println("************************************* time test")
	fmt.Println("get time stamp: ", timeutils.GetTimeStamp("Asia/Ho_Chi_Minh"))
	fmt.Println("- convert time stamp to local")
	timestamp := timeutils.GetTimeStamp("Asia/Ho_Chi_Minh")
	fmt.Println(timeutils.ConvetTimestamsToLocalTime(utils.String2Int64(timestamp)))
	fmt.Println("- time now UTC: ", timeutils.TimeNowUTC())
	fmt.Println("- get to day date:", timeutils.GetTodaysDate("Asia/Ho_Chi_Minh"))
	fmt.Println("get today datetime format: ", timeutils.GetTodaysDateTimeFormatted("Asia/Ho_Chi_Minh"))

	fmt.Println("************************************* sched test")
	_, err = sched.Every(5).ESeconds().Run(PrintHello)
	if err != nil {
		log.Error(err)
	}
	_, err = sched.Every(10).ESeconds().Run(PrintBille)
	if err != nil {
		log.Error(err)
	}
	/*	_, err = sched.Every(10).WFriday().At("14:51").Run(PrintYes)
		if err != nil {
			log.Error(err)
		}*/
	runtime.Goexit()
}

func PrintHello(job *sched.Job) {
	log.Info("hmm...")
	//fmt.Println("Hmm...")
}

func PrintBille(job *sched.Job) {
	log.Info("bedder than ever...")
}

func PrintYes(job *sched.Job) {
	log.Info("yes...")
	//fmt.Println("Hmm...")
}
