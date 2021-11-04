package main

import "spaceclan1/spaceclan-api/app_cacher"

// application to aggregate actions by date,month and cache it in memcached or redis (have not considered yet)
func main() {
	app_cacher.StartApplication()
}
