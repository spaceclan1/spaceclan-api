package app_cacher

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"runtime"
	"spaceclan1/spaceclan-api/controllers"
)

func StartApplication() {
	log.SetLevel(log.InfoLevel)
	log.Info("cache started")
	controllers.CacherController.CacheAndAggregate()
	//controllers.CacherController.CacheAndAggregate()
	//controllers.CacherController.CacheAndAggregate()
	//controllers.CacherController.CacheAndAggregate()
	//controllers.CacherController.CacheAndAggregate()
	//initScheduler()
	runtime.Goexit()
}

func initScheduler() {
	c := cron.New()
	_, err := c.AddFunc("*/1 * * * *", controllers.CacherController.CacheAndAggregate)
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}
