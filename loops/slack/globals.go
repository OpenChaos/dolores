package dolores_slack

import "github.com/nlopes/slack"

var (
	BotID              string
	SlackAdminEmailIds []string
	API                *slack.Client
)
