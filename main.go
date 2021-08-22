package main

import (
	"os"
	"semantic-release-go-commit-linting/helpers"
)

func main() {
	// get commit msg
	var commitMsg helpers.CommitMessage
	commitMsg = helpers.BufferLastCommit()
	if helpers.IsTypeValid(commitMsg) {
		os.Exit(0)
	}
	os.Exit(1)
}
