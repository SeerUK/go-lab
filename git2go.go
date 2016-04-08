package main

import (
	"fmt"

	"gopkg.in/libgit2/git2go.v23"
)

func main() {
	repo, err := git.Clone("git://github.com/eidolon/console.git", "console", &git.CloneOptions{
		FetchOptions: &git.FetchOptions{
			RemoteCallbacks: git.RemoteCallbacks{
				TransferProgressCallback: func(stats git.TransferProgress) git.ErrorCode {
					fmt.Println(stats.ReceivedBytes)

					return git.ErrOk
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	println(repo)
}
