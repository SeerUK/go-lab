package main

import "gopkg.in/libgit2/git2go.v23"

func main() {
	repo, err := git.Clone("git://github.com/eidolon/console.git", "console", &git.CloneOptions{})
	if err != nil {
		panic(err)
	}

	println(repo)
}
