package dolores_slack

import "github.com/nlopes/slack"

func IsDbAdmin(emailID string) bool {
	for _, dbAdminEmailID := range DbAdminEmailIds {
		if emailID == dbAdminEmailID {
			return true
		}

	}
	return false
}

func IsAdmin(emailID string) bool {
	for _, adminEmailID := range DoloresAdminEmailIds {
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
