package main

import (
	"spaceclan1/spaceclan-api/app_cacher"
	"spaceclan1/spaceclan-api/datasource"
)

// application to aggregate actions by date,month and cache it in memcached or redis (have not considered yet)
func main() {
	datasource.InitMysql()
	datasource.InitRedis()
	app_cacher.StartApplication()
}
