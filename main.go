package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kyokomi/emoji"
)

var emoticons map[string]string

func main() {
	initEmoticons()
	// We default to "shrug"
	emt := "¯\\_(ツ)_/¯"
	var done bool
	if len(os.Args) == 2 {
		which := os.Args[1]
		if which == "help" {
			fmt.Println("There are a *ton* of emoji supported, to list them, do \"help-emoji\"")
			for k, v := range emoticons {
				fmt.Println(k, v)
			}
			os.Exit(0)
		} else if which == "help-emoji" {
			allEmoji := emoji.CodeMap()
			fmt.Println("== Supported Emoji ==")
			for k, v := range allEmoji {
				fmt.Println(k + " " + v)
			}
			os.Exit(0)
		} else {
			var i string
			if i, done = emoticons[which]; done {
				emt = i
			}
		}
		if !done {
			// it might be a one-word emoji request
			tst1 := ":" + which + ":"
			tst2 := emoji.Sprint(tst1)
			if tst1 != tst2 {
				emt = tst2
				done = true
			}
		}
	}
	if !done {
		rest := strings.Join(os.Args[1:], " ")
		emt = emoji.Sprint(rest)
	}
	fmt.Println(emt)
}

func initEmoticons() {
	emoticons = make(map[string]string)
	emoticons["bat"] = "～°·_·°～"
	emoticons["bomb"] = "●～*"
	emoticons["fish"] = ">°)))彡"
	emoticons["flip"] = "(╯°□°）╯︵ ┻━┻"
	emoticons["lenny"] = "( ͡° ͜ʖ ͡°)"
	emoticons["octopus"] = "<コ:彡"
	emoticons["shrug"] = "¯\\_(ツ)_/¯"
	emoticons["smirk"] = "( ͡° ͜ʖ ͡°)"
	emoticons["snake"] = "~>°)～～～"
	emoticons["squid"] = "Ｃ:.ミ"
	emoticons["star"] = "☆彡"
	emoticons["tadpole"] = "(°°)～"
}
