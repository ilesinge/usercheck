package main

import (
	"fmt"

	"github.com/ilesinge/usercheck/github"
	"github.com/ilesinge/usercheck/twitter"
)

func main() {
	var twitter twitter.Twitter
	avail, err := twitter.IsAvailable("ilesingezzz")
	fmt.Println(avail, err)
	avail, err = twitter.IsAvailable("ilesinge")
	fmt.Println(avail, err)
	var github github.Github
	avail, err = github.IsAvailable("ilesingezzz")
	fmt.Println(avail, err)
	avail, err = github.IsAvailable("ilesinge")
	fmt.Println(avail, err)
}
