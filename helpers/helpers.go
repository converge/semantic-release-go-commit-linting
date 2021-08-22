// Package helpers group common functions for the Semantic Release linting git hook
package helpers

import (
	"bufio"
	"fmt"
	"os"
)

type CommitMessage struct {
	typeScope string
	body      string
	footer    string
}

// BufferLastCommit load last commit message and return CommitMessage object
func BufferLastCommit() CommitMessage {
	f, err := os.Open(".git/COMMIT_EDITMSG")
	if err != nil {
		panic(err)
	}
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	var cMsg CommitMessage
	for i, line := range text {
		if i == 0 {
			cMsg.typeScope = line
		}
		cMsg.body = line

	}
	err = f.Close()
	if err != nil {
		panic(err)
	}

	return cMsg
}
