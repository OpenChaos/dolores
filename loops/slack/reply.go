package dolores_slack

import (
	"fmt"

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
