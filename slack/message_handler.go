package dolores_slack

import (
	"fmt"
	"strings"

	"github.com/dchest/stemmer/porter2"
	"github.com/sbstjn/allot"
)

func messageHandler(msg string) {
	eng := porter2.Stemmer
	fmt.Println("*****", msg)
	for _, w := range strings.Split(msg, " ") {
		fmt.Println("*****", eng.Stem(w))
	}
	cmd := allot.New("(give|remove) access (to|for|from|of) <user:string> with \"<sshKeys:string>\"")
	match, err := cmd.Match(msg)

	if err == nil {
		axn, _ := match.Match(0)
		prep, _ := match.Match(1)
		user, _ := match.Match(2)
		sshKeys, _ := match.Match(3)

		fmt.Printf("%s %s \"%s\" using \"%s\"", axn, prep, user, sshKeys)

	} else {
		fmt.Println("Request did not match command.", match)
	}
}
