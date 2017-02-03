package dolores_slack

import "github.com/nlopes/slack"

func IsAdmin(emailID string) bool {
	for _, adminEmailID := range SlackAdminEmailIds {
		if emailID == adminEmailID {
			return true
		}
	}
	return false
}

func SenderEmail(event *slack.MessageEvent) string {
	user, err := API.GetUserInfo(event.Msg.User)
	if err == nil {
		return user.Profile.Email
	}
	return ""
}
