package dolores_slack

import (
	"fmt"

	"github.com/sbstjn/allot"
)

var (
	sshAccessMessageHandler = MessageHandler{name: "ssh-access",
		command: allot.New("(give|remove) access (to|for|from|of) <user:string> for <machinePattern:string> with \"<sshKeys:string>\""),
		msgFoo:  sshAccess}
)

func sshAccess(match allot.MatchInterface) (status bool) {
	axn, _ := match.Match(0)
	prep, _ := match.Match(1)
	user, _ := match.Match(2)
	machinePattern, _ := match.Match(3)
	sshKeys, _ := match.Match(4)

	status = true
	fmt.Printf("%s %s \"%s\" for \"%s\" using \"%s\"", axn, prep, user, machinePattern, sshKeys)
	return
}
