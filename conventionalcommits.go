package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var examples = `
+ 61c8ca9 fix: navbar not responsive on mobile
+ 479c48b test: prepared test cases for user authentication
+ dfdc715 feat(auth): added social login using twitter

"fixup!" and "Initial commit" are withelisted"
`

func main() {
	log.SetFlags(0)

	// Get commit message file path
	if len(os.Args) < 2 {
		log.Fatalln("Please specify a message commit file path")
	}

	// Read it
	commitMessagePath := os.Args[1]
	data, err := ioutil.ReadFile(commitMessagePath)
	if err != nil {
		log.Fatalln(err)
	}

	var re = regexp.MustCompile(`((build|ci|docs|feat|fix|refactor|style|test|revert)(\([\w\-]+\))?:\s.*)|fixup! |^Initial commit\n$|^BREAKING CHANGE:|(^Merge )`)
	if !re.MatchString(string(data)) {
		log.Println("🤬 Please spectify a valid commit message")
		log.Fatalf(examples)
	}
}
