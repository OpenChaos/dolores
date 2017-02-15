package dolores_slack

import (
	"github.com/nlopes/slack"
	"github.com/sbstjn/allot"

	dolores_drives "dolores/drives"
)

var (
	helpMessageHandler = MessageHandler{name: "help",
		allotCommand: helpAllotCommand,
		nlpSamples:   helpNlpSamples,
		msgFoo:       helpMessage}

	sshAccessMessageHandler = MessageHandler{name: "ssh-access",
		allotCommand: sshAccessAllotCommand,
		nlpSamples:   sshAccessNlpSamples,
		msgFoo:       sshAccess}
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

	Reply(ev, sshAccessReplyDeferMessage)
	if axn == "give" && (prep == "to" || prep == "for") {
		reply, err = dolores_drives.GiveSshAccess(machinePattern, user, isAdmin)
		if err != nil && sshAccessReplyFailureMessage != "" {
			reply = sshAccessReplyFailureMessage
		} else if err == nil && sshAccessReplySuccessMessage != "" {
			reply = sshAccessReplySuccessMessage
		}
	}
	return
}
