package main

import (
	"os"
	"os/exec"
)

func main() {

	origin := "https://github.com/open-set/forked-sync.git"

	upstream := "https://github.com/openset/forked-sync.git"

	exec.Command("git", "clone", origin).Start()

	os.Chdir("forked-sync")

	exec.Command("git", "remote", "add", "upstream", upstream).Start()

	exec.Command("git", "pull", "upstream").Start()

	exec.Command("git", "merge", "upstream/master").Start()

	exec.Command("git", "push", "origin").Start()

}
