package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"os/exec"
)

func main() {

	//done := make(chan map[string]bool, 8)

	origin := "https://github.com/open-set/forked-sync.git"

	upstream := "https://github.com/openset/forked-sync.git"

	mapOrigin := map[string]string{origin: upstream}

	for origin, upstream := range mapOrigin {
		sync(origin, upstream)
	}

}

//同步master分支
func sync(origin, upstream string) {

	tempDir := os.TempDir()

	println(tempDir)

	os.Chdir(tempDir)

	tempFolder := getFolderName(origin)

	println(tempFolder)

	exec.Command("git", "clone", origin, tempFolder).Run()

	os.Chdir(tempFolder)

	exec.Command("git", "config", "user.name", "Openset").Run()
	exec.Command("git", "config", "user.email", "openset.wang@gmail.com").Run()

	//git remote add upstream https://github.com/openset/forked-sync.git
	exec.Command("git", "remote", "add", "upstream", upstream).Run()

	exec.Command("git", "pull", "upstream").Run()

	exec.Command("git", "merge", "upstream/master").Run()

	exec.Command("git", "push", "origin").Run()

	println(tempDir + tempFolder)

	//os.RemoveAll(filepath.Join(tempDir, tempFolder))

	//done <- map[string]bool{origin: true}
	println(origin)
}

//获取MD5值
func getMD5(str string) string {

	res := md5.Sum([]byte(str))

	return hex.EncodeToString(res[:])
}

//获取临时文件夹名
func getFolderName(origin string) string {

	//timestamp := time.Now().Unix()

	//tempFolder := strconv.Itoa(int(timestamp))

	folderMD5 := getMD5(origin)

	return "sync_tmp_" + folderMD5
}
