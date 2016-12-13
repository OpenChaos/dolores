package main

import (
	dolores_helpers "./corecode"
	dolores_slack "./loops/slack"

	"github.com/jasonlvhit/gocron"
)

func prepareScheduler() {
	scheduler := gocron.NewScheduler()
	<-scheduler.Start()
	// more examples: https://github.com/jasonlvhit/gocron/blob/master/example/example.go#L19
}

func main() {
	config := dolores_helpers.ConfigFromFlags()

	go prepareScheduler() // spawn cron scheduler jobs
	dolores_slack.LoopRTMEvents(config)
}
