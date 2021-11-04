package app_heroestaking

import (
	"github.com/robfig/cron/v3"
	"log"
	"runtime"
	"spaceclan1/spaceclan-data-gatherer/controllers"
)

func StartApplication() {
	initScheduler()
	runtime.Goexit()
}

func initScheduler() {
	c := cron.New()
	_, err := c.AddFunc("*/5 * * * *", controllers.HeroestakingController.FetchPoolIncreaseTransactions)
	if err != nil {
		log.Fatal(err)

	}
	c.Start()
}
