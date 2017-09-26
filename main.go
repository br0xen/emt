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
	if len(os.Args) > 1 {
		which := os.Args[1]
		if which == "help" {
			for k, v := range emoticons {
				fmt.Println(k, v)
			}
			os.Exit(0)
		} else if which == "emoji" {
			rest := strings.Join(os.Args[2:], " ")
			emt = emoji.Sprint(rest)
		} else {
			if i, ok := emoticons[which]; ok {
				emt = i
			}
		}
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
