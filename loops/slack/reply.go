package dolores_slack

import (
	"fmt"
	"log"

	"github.com/nlopes/slack"
)

func Reply(event *slack.MessageEvent, message string) (err error) {
	userName := event.Msg.User
	user, err := API.GetUserInfo(event.Msg.User)
	if err == nil {
		userName = user.Name
	}
	params := slack.PostMessageParameters{}
	params.Username = BotID
	params.AsUser = true
	params.LinkNames = 1 // so slack linkify channel names and usernames https://api.slack.com/docs/message-formatting
	replyMessage := fmt.Sprintf("@%s: %s", userName, message)
	API.PostMessage(event.Msg.Channel, replyMessage, params)
	return
}

func AddReaction(event *slack.MessageEvent, reaction string) (err error) {
	msgRef := slack.NewRefToMessage(event.Msg.Channel, event.Msg.Timestamp)
	if err = API.AddReaction(reaction, msgRef); err != nil {
		log.Printf("[error] Error adding reaction: %s\n", err)
	}
	return
}

func ReplyInteractive(event *slack.MessageEvent, message string, attachments []slack.Attachment) (err error) {
	userName := event.Msg.User
	user, err := API.GetUserInfo(event.Msg.User)
	if err == nil {
		userName = user.Name
	}
	params := slack.PostMessageParameters{}
	params.Username = BotID
	params.AsUser = true
	params.LinkNames = 1

	params.Attachments = attachments //

	replyMessage := fmt.Sprintf("@%s: %s", userName, message)

	channelID, timestamp, msgErr := API.PostMessage(event.Msg.Channel, replyMessage, params)
	if msgErr != nil {
		log.Printf("[error] Error posting message: %s in %v with timestamp %v\n", msgErr, channelID, timestamp)
		return
	}
	return
}
