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
	var done bool
	var which string
	if len(os.Args) == 2 {
		which = os.Args[1]
		if which == "help" {
			showHelp()
		} else if which == "help-emoji" {
			allEmoji := emoji.CodeMap()
			fmt.Println("== Supported Emoji ==")
			for k, v := range allEmoji {
				fmt.Println(k + " " + v)
			}
			os.Exit(0)
		}
		// it might be a one-word emoji request
		allEmoji := emoji.CodeMap()
		if _, done = allEmoji[which]; done {
			tst1 := ":" + which + ":"
			tst2 := emoji.Sprint(tst1)
			if tst1 != tst2 {
				fmt.Println(tst2)
				done = true
			}
		} else {
			// Search for any emoji that contain this string
			var cnt int
			for k, v := range allEmoji {
				if strings.Contains(k, which) {
					fmt.Println(k + " " + v)
					cnt++
				}
			}
			if cnt > 0 {
				done = true
			} else {
				fmt.Println("No emoji for " + which + " found.")
			}
		}
	}
	if !done {
		var i string
		if i, done = emoticons[which]; done {
			fmt.Println(i)
		}
	}
	if !done {
		showHelp()
	}
}

func showHelp() {
	fmt.Println("There are a *ton* of emoji supported, to list them, do \"help-emoji\"")
	for k, v := range emoticons {
		fmt.Println(k, v)
	}
	os.Exit(0)
}

func initEmoticons() {
	emoticons = make(map[string]string)
	emoticons["ascii-bat"] = "～°·_·°～"
	emoticons["ascii-bomb"] = "●～*"
	emoticons["ascii-fish"] = ">°)))彡"
	emoticons["ascii-flip"] = "(╯°□°）╯︵ ┻━┻"
	emoticons["ascii-lenny"] = "( ͡° ͜ʖ ͡°)"
	emoticons["ascii-octopus"] = "<コ:彡"
	emoticons["ascii-shrug"] = "¯\\_(ツ)_/¯"
	emoticons["ascii-smirk"] = "( ͡° ͜ʖ ͡°)"
	emoticons["ascii-snake"] = "~>°)～～～"
	emoticons["ascii-star"] = "☆彡"
	emoticons["ascii-tadpole"] = "(°°)～"
}
