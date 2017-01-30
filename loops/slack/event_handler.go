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
	//rtm.SendMessage(rtm.NewOutgoingMessage("I'm back baby!", generalChannel))
}

func MessageEvent(ev *slack.MessageEvent) {
	if ev.Msg.Type == "message" && ev.Msg.User != BotID && ev.Msg.SubType != "message_deleted" {
		//&& (strings.Contains(ev.Msg.Text, "<@"+BotID+">") || strings.HasPrefix(ev.Msg.Channel, "infra")) {
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
	// the gocron scheduler above communicates with the RTMbot subroutine
	// via it's builtin channel. here we check for custom events and act
	// accordingly
	if msg.Type == "ListStaging" || msg.Type == "ListProduction" ||
		msg.Type == "ListUAT" ||
		msg.Type == "ListInternal" {
		response := msg.Data.(string)
		params := slack.PostMessageParameters{AsUser: true}
		API.PostMessage("cd-phoenix", response, params)
	} else {
		fmt.Println("*****", msg)
		// Ignore other events..
		//fmt.Printf("Unexpected %s: %+v\n", msg.Type, msg.Data)
	}
}
