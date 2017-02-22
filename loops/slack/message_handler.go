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
	name         string
	allotCommand allot.Command
	nlpSamples   []string
	msgFoo       func(ev *slack.MessageEvent, match allot.MatchInterface) (string, error)
}

var (
	message_types = []MessageHandler{
		sshAccessMessageHandler,
		dbAccessMessageHandler,
		serverListMessageHandler,
		nslookupMessageHandler,
		helpMessageHandler,
	}
)

func processMessageByAllot(ev *slack.MessageEvent, msg string) bool {
	for _, message_type := range message_types {
		match, err := message_type.allotCommand.Match(msg)
		if err != nil {
			continue
		}

		reply, axn_err := message_type.msgFoo(ev, match)
		if axn_err != nil {
			log.Println("[ERROR] running task failed,", axn_err, match)
		}
		axn_err = Reply(ev, reply)
		if axn_err != nil {
			log.Println("[ERROR] sending reply failed,", axn_err, match)
		}
		return true
	}
	return false
}

func processMessage(ev *slack.MessageEvent, msg string) {
	msg = strings.Join(strings.Fields(msg), " ")
	if processMessageByAllot(ev, msg) {
		return
	}

	err := Reply(ev, doloresWrongCmdMessage)
	if err != nil {
		log.Println("[ERROR] sending reply failed,", err)
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
