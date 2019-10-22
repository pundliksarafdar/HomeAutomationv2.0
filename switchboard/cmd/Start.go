package main

import (
	"agent/router"
	"agent/schedular"
)

func main() {
	schedular.StartSchedule()
	router.StartWebServer()
}
