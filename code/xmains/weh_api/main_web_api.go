package main

import (
	"spaceclan1/spaceclan-api/app_web_api"
	"spaceclan1/spaceclan-api/datasource"
)

// application to aggregate actions by date,month and cache it in memcached or redis (have not considered yet)
func main() {
	datasource.InitRedis()
	app_web_api.StartApplication()
}
