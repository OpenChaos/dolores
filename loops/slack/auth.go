package dolores_slack

import "github.com/nlopes/slack"

func AuthenticatedApi(slack_token string) *slack.Client {
	api := slack.New(slack_token)
	return api
}
