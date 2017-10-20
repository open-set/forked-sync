package main

import (
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	origin := "https://github.com/open-set/forked-sync.git"

	upstream := "https://github.com/openset/forked-sync.git"

	tempDir := os.TempDir()

	os.Chdir(tempDir)

	println(tempDir)

	timestamp := time.Now().Unix()

	tempFolder := "tmp_" + strconv.Itoa(int(timestamp))

	println(tempFolder)

	exec.Command("git", "clone", origin, tempFolder).Run()

	os.Chdir(tempFolder)

	//git remote add upstream https://github.com/openset/forked-sync.git
	exec.Command("git", "remote", "add", "upstream", upstream).Run()

	exec.Command("git", "pull", "upstream").Run()

	exec.Command("git", "merge", "upstream/master").Run()

	exec.Command("git", "push", "origin").Run()

	println(tempDir + tempFolder)

	os.RemoveAll(tempDir + tempFolder)

}
