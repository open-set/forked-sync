package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	debug := true

	mapOrigin := [...]map[string]string{
		{
			"origin":   "https://github.com/open-set/forked-sync.git",
			"upstream": "https://github.com/openset/forked-sync.git",
		},
	}

	jobs := make(chan map[string]string, 100)

	results := make(chan string, 100)

	for w := 0; w < 8; w++ {
		go worker(jobs, results, debug)
	}

	for _, j := range mapOrigin {
		jobs <- j
	}
	close(jobs)

	for r := 0; r < len(mapOrigin); r++ {
		println(<-results)
	}
}

//工作池
func worker(jobs <-chan map[string]string, results chan<- string, debug bool) {
	for j := range jobs {
		sync(j["origin"], j["upstream"], debug)
		results <- j["origin"]
	}
}

//同步master分支
func sync(origin, upstream string, debug bool) {

	tempDir := os.TempDir()

	if debug {
		println("tempDir: " + tempDir)
	}

	os.Chdir(tempDir)

	tempFolder := getFolderName(origin, debug)

	exec.Command("git", "clone", origin, tempFolder).Run()

	os.Chdir(tempFolder)

	exec.Command("git", "config", "user.name", "Openset").Run()
	exec.Command("git", "config", "user.email", "openset.wang@gmail.com").Run()

	exec.Command("git", "remote", "add", "upstream", upstream).Run()

	exec.Command("git", "pull", "upstream").Run()

	exec.Command("git", "merge", "upstream/master").Run()

	exec.Command("git", "push", "origin").Run()

	if debug {
		println("path: " + tempDir + tempFolder)
		println("origin: " + origin)
	}

}

//获取MD5值
func getMD5(str string, debug bool) string {

	byte := md5.Sum([]byte(str))
	string := hex.EncodeToString(byte[:])

	if debug {
		println("string: " + str)
		println("MD5: " + string)
	}

	return string
}

//获取临时文件夹名
func getFolderName(origin string, debug bool) string {

	//timestamp := time.Now().Unix()

	//tempFolder := strconv.Itoa(int(timestamp))

	folderMD5 := getMD5(origin, debug)

	tempFolder := "sync_tmp_" + folderMD5

	if debug {
		println("tempFolder: " + tempFolder)
	}

	return tempFolder
}

func cleanup(origin string, debug bool) {

	tempDir := os.TempDir()

	tempFolder := getFolderName(origin, debug)

	os.RemoveAll(filepath.Join(tempDir, tempFolder))

}
