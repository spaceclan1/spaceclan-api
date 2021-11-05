package app_web_api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"spaceclan1/spaceclan-api/controllers"
)

var (
	router = gin.New()
)

func StartApplication() {
	log.SetLevel(log.InfoLevel)
	log.Info("web api started")
	router.Use(gin.Recovery())
	mapUrls()
	router.Run(":8080")
}

func mapUrls() {

	router.GET("/vip2/deposits", controllers.Vip2Controller.GetDeposits)
	//router.GET("/vip2/month/{:account}", controllers.Vip2Controller.GetMonthlyRewards)
	//router.GET("/endpoint2", controllers.Controller2.Method2)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": http.StatusText(http.StatusNotFound), "message": "Page not found", "status": http.StatusNotFound})
	})
}
