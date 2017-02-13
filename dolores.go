package main

import (
	dolores_corecode "dolores/corecode"
	dolores_slack "dolores/loops/slack"
	dolores_memories "dolores/memories"

	"github.com/jasonlvhit/gocron"
)

func prepareScheduler() {
	scheduler := gocron.NewScheduler()
	scheduler.Every(2).Hours().Do(dolores_memories.GcloudComputeInstances)
	<-scheduler.Start()
	//scheduler.Every(3).Minutes().Do(task)
	// more examples: https://github.com/jasonlvhit/gocron/blob/master/example/example.go#L19
}

func main() {
	config := dolores_corecode.ConfigFromFlags()

	go prepareScheduler() // spawn cron scheduler jobs
	dolores_slack.LoopRTMEvents(config)
}
