package main

import "cron-job/service"

func main() {
	app := service.NewApp()
	app.Run()
}
