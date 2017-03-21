package dolores_slack

import (
	"encoding/json"

	"github.com/nlopes/slack"
)

var (
	SlackAttachmentDefaultColor = "#36a64f"
)

func SampleAttachment() slack.Attachment {
	attachment := slack.Attachment{
		Color:    SlackAttachmentDefaultColor,
		Fallback: "Required plain-text summary of the attachment.",

		AuthorName:    "Bobby Tables",
		AuthorSubname: "BobT",
		AuthorLink:    "http://flickr.com/bobby/",
		AuthorIcon:    "http://flickr.com/icons/bobby.jpg",

		Title:     "Slack API Documentation",
		TitleLink: "https://api.slack.com/",
		Text:      "Optional text that appears within the attachment",
		Pretext:   "Optional text that appears above the attachment block",

		ImageURL: "http://my-website.com/path/to/image.jpg",
		ThumbURL: "http://example.com/path/to/thumb.png",

		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Priority",
				Value: "High",
				Short: false,
			},
		},

		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "Oye",
				Text:  "oye oye",
				Style: "b",
				Type:  "button",
				Value: "oye",
			},
		},

		MarkdownIn: []string{},

		Footer:     "Slack API",
		FooterIcon: "https://platform.slack-edge.com/img/default_application_icon.png",
		Ts:         json.Number("123456789"),
	}
	return attachment
}
