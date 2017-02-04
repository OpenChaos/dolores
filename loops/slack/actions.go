package dolores_slack

import (
	"github.com/nlopes/slack"
	"github.com/sbstjn/allot"

	dolores_drives "dolores/drives"
)

var (
	helpMessageHandler = MessageHandler{name: "help",
		command: allot.New("(help|sos)"),
		msgFoo:  helpMessage}

	sshAccessMessageHandler = MessageHandler{name: "ssh-access",
		command: allot.New("(give|remove) access (to|for|from|of) <user:string> for <machinePattern:string>"),
		msgFoo:  sshAccess}
)

func helpMessage(ev *slack.MessageEvent, match allot.MatchInterface) (reply string, err error) {
	reply = replyHelpMessage
	return
}

func sshAccess(ev *slack.MessageEvent, match allot.MatchInterface) (reply string, err error) {
	axn, _ := match.Match(0)
	prep, _ := match.Match(1)
	user, _ := match.Match(2)
	machinePattern, _ := match.Match(3)

	requestedBy := SenderEmail(ev)
	isAdmin := "no"
	if IsAdmin(requestedBy) {
		isAdmin = "yes"
	}

	Reply(ev, "sure, let me check if it's allowed as of now :)")
	if axn == "give" && (prep == "to" || prep == "for") {
		reply, err = dolores_drives.GiveSshAccess(machinePattern, user, isAdmin)
	}
	return
}
