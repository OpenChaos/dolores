package dolores_slack

import (
	"fmt"
	"strings"

	"github.com/dchest/stemmer/porter2"
	"github.com/sbstjn/allot"
)

type MessageHandler struct {
	name    string
	command allot.Command
	msgFoo  func(allot.MatchInterface) bool
}

var (
	message_types = []MessageHandler{
		sshAccessMessageHandler,
	}
)

func processMessage(msg string) {
	for _, message_type := range message_types {
		match, err := message_type.command.Match(msg)
		if err == nil {
			if message_type.msgFoo(match) == false {
				fmt.Printf("Error: %q", match)
				panic(err)
			}
			return
		}
	}

	fmt.Println("Request did not match command.")
}

func stemSentence(sentence string) (stemText string) {
	eng := porter2.Stemmer
	sentenceTokens := strings.Split(sentence, " ")
	stemText = eng.Stem(sentenceTokens[0])

	for _, word := range sentenceTokens[1:] {
		stemText = fmt.Sprintf("%s %s", stemText, eng.Stem(word))
	}
	return
}
