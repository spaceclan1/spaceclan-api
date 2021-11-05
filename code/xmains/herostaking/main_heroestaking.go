package main

import (
	"spaceclan1/spaceclan-api/app_heroestaking"
	"spaceclan1/spaceclan-api/datasource"
)

// application to fetch herostaking accounts actions from APIEndpoints and insert into database
func main() {
	datasource.InitMysql()
	app_heroestaking.StartApplication()
}
