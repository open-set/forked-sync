package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {

	done := make(chan bool)

	origin := "https://github.com/open-set/forked-sync.git"

	upstream := "https://github.com/openset/forked-sync.git"

	go sync(origin, upstream, done)

	<-done
}

//同步master分支
func sync(origin, upstream string, done chan bool) {

	tempDir := os.TempDir()

	os.Chdir(tempDir)

	println(tempDir)

	timestamp := time.Now().Unix()

	tempFolder := "tmp_" + strconv.Itoa(int(timestamp))

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

	os.RemoveAll(filepath.Join(tempDir, tempFolder))

	done <- true
}
