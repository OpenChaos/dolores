package main

import (
	dolores_helpers "./helpers"
	dolores_slack "./slack"

	"github.com/jasonlvhit/gocron"
	"github.com/nlopes/slack"
)

func prepareScheduler() {
	scheduler := gocron.NewScheduler()
	<-scheduler.Start()
	// more examples: https://github.com/jasonlvhit/gocron/blob/master/example/example.go#L19
}

func main() {
	config := dolores_helpers.ConfigFromFlags()

	dolores_slack.BotID = config["slack-bot-name"]
	dolores_slack.API = dolores_slack.AuthenticatedApi(config["slack-bot-api-token"])
	if config["slack-debug-mode"] == "true" {
		dolores_slack.API.SetDebug(true)
	} else {
		dolores_slack.API.SetDebug(false)
	}
	go prepareScheduler() // spawn cron scheduler jobs

	rtm := dolores_slack.API.NewRTM()
	go rtm.ManageConnection() // spawn slack bot

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.HelloEvent:
				dolores_slack.HelloEvent(ev)

			case *slack.ConnectedEvent:
				dolores_slack.ConnectedEvent(ev)

			case *slack.MessageEvent:
				dolores_slack.MessageEvent(ev)

			case *slack.PresenceChangeEvent:
				dolores_slack.PresenceChangeEvent(ev)

			case *slack.LatencyReport:
				dolores_slack.LatencyReport(ev)

			case *slack.RTMError:
				dolores_slack.RTMError(ev)

			case *slack.InvalidAuthEvent:
				dolores_slack.InvalidAuthEvent(ev)
				panic("invalid slack api token")

			default:
				dolores_slack.DefaultEvent(msg)
			}
		}
	}
}
