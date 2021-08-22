package helpers

import (
	"log"
	"strings"
)

// IsTypeValid verify if the commit message follows Semantic Release format
func IsTypeValid(cMsg CommitMessage) bool {
	possibleTypes := []string{
		"chore:",
		"docs:",
		"build:",
		"feat",
		"refactor:",
		"ci:",
		"style:",
		"test",
		"perf:",
		"BREAKING CHANGE",
	}
	// validate typeScope
	firstWord := strings.Split(cMsg.typeScope, " ")

	for _, p := range possibleTypes {
		if firstWord[0] == p {
			return true
		}
	}
	log.Fatal("Commit message doesnt have a valid Semantic Release Type")
	return false
}
