package dolores_slack

import (
	"fmt"
	"log"
	"strings"

	"github.com/dchest/stemmer/porter2"
	"github.com/nlopes/slack"
	"github.com/sbstjn/allot"
)

type MessageHandler struct {
	name    string
	command allot.Command
	msgFoo  func(ev *slack.MessageEvent, match allot.MatchInterface) (string, error)
}

var (
	message_types = []MessageHandler{
		helpMessageHandler,
		sshAccessMessageHandler,
	}
)

func processMessage(ev *slack.MessageEvent, msg string) {
	msg = strings.Join(strings.Fields(msg), " ")

	for _, message_type := range message_types {
		match, err := message_type.command.Match(msg)
		if err != nil {
			continue
		}

		reply, axn_err := message_type.msgFoo(ev, match)
		if axn_err != nil {
			log.Println("[ERROR]", axn_err, match)
		}
		axn_err = Reply(ev, reply)
		if axn_err != nil {
			log.Println("[ERROR]", axn_err, match)
		}
		return
	}

	err := Reply(ev, "What do you mean by that?\n\tTry 'help' for how-to.\n\tIf you think your command was correct, talk to systems team about the issue.")
	if err != nil {
		log.Println("[ERROR]", err)
	}
	log.Println("Request did not match command.")
}

func stemSentence(sentence string) (stemText string) {
	eng := porter2.Stemmer
	sentenceTokens := strings.Fields(sentence)
	stemText = eng.Stem(sentenceTokens[0])

	for _, word := range sentenceTokens[1:] {
		stemText = fmt.Sprintf("%s %s", stemText, eng.Stem(word))
	}
	return
}
