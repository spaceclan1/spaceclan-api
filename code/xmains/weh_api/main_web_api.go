package main

import "spaceclan1/spaceclan-api/app_web_api"

// application to aggregate actions by date,month and cache it in memcached or redis (have not considered yet)
func main() {
	app_web_api.StartApplication()
}
