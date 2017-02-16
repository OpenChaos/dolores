package dolores_slack

import (
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/nlopes/slack"
)

func HelloEvent(ev *slack.HelloEvent) {
	fmt.Println(ev)
}

func ConnectedEvent(ev *slack.ConnectedEvent) {
	BotID = ev.Info.User.ID
}

func MessageEvent(ev *slack.MessageEvent) {
	if ev.Msg.Type == "message" && ev.Msg.User != BotID && ev.Msg.SubType != "message_deleted" {
		fmt.Printf("Message: %+v\n", ev.Msg)
		// strip out bot's name and spaces
		parsedMessage := strings.TrimSpace(strings.Replace(ev.Msg.Text, "<@"+BotID+">", "", -1))
		r, n := utf8.DecodeRuneInString(parsedMessage)
		parsedMessage = string(unicode.ToLower(r)) + parsedMessage[n:]
		go processMessage(ev, parsedMessage)
	}
}

func PresenceChangeEvent(ev *slack.PresenceChangeEvent) {
	log.Printf("Presence Change: %+v\n", ev)
}

func LatencyReport(ev *slack.LatencyReport) {
	API.GetUserInfo(BotID)
	log.Printf("Current latency: %+v\n", ev.Value)
}

func RTMError(ev *slack.RTMError) {
	log.Printf("Error: %s\n", ev.Error())
}

func InvalidAuthEvent(ev *slack.InvalidAuthEvent) {
	fmt.Println("Error: Invalid credentials")
}

func DefaultEvent(msg slack.RTMEvent) {
	// via builtin channel, here we check for custom events & act accordingly
	log.Println("[skipped-event]", msg)
}
