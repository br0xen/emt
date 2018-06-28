package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kyokomi/emoji"
)

func main() {
	if len(os.Args) < 2 || strings.HasSuffix(os.Args[1], "-help") {
		showUsage()
	}

	var getApprox bool
	var which string
	var matchWhich string
	for _, v := range os.Args {
		if v == "-a" {
			getApprox = true
		} else {
			which = v
		}
	}

	allEmoji := emoji.CodeMap()
	if !strings.HasPrefix(which, ":") && !strings.HasSuffix(which, ":") {
		matchWhich = ":" + which + ":"
	}
	if v, ok := allEmoji[matchWhich]; ok {
		fmt.Println(v)
		os.Exit(0)
	} else if getApprox {
		for k, v := range allEmoji {
			if strings.Contains(k, which) {
				fmt.Println(v)
				os.Exit(0)
			}
		}
	}
	os.Exit(1)
}

func showUsage() {
	allEmoji := emoji.CodeMap()
	fmt.Println("== Supported Emoji ==")
	for k, v := range allEmoji {
		fmt.Println(k + " " + v)
	}
	os.Exit(0)
}
