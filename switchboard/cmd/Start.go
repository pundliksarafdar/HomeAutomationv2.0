package main

import (
	"agent/router"
	"agent/schedular"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Starting Scheduler......")
	schedular.StartSchedule()
	fmt.Printf("Starting server......")
	router.StartWebServer()
	log.Print("hiiiiiiiiiiiiiiiiii")
}
