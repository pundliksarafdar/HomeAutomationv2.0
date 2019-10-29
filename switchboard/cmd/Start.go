package main

import (
	"agent/router"
	"agent/schedular"
	"fmt"
)

func main() {
	fmt.Printf("Starting Scheduler......")
	schedular.StartSchedule()
	fmt.Printf("Starting server......")
	router.StartWebServer()
}
