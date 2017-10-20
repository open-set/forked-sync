package main

import "os/exec"

func main() {

	origin := "https://github.com/open-set/forked-sync.git"

	upstream := "https://github.com/openset/forked-sync.git"

	exec.Command("git", "clone", origin).Start()

	exec.Command("cd", "forked-sync").Start()

	//git remote add upstream <path/to/original/repo>
	//
	//git fetch upstream
	//
	//git merge upstream/master master
	//
	//git push origin master

	exec.Command("git", "remote", "add", upstream).Start()

	exec.Command("git", "pull", "upstream").Start()

	exec.Command("git", "push", "origin").Start()

}
