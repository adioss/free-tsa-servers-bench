package main

import (
	"github.com/adioss/free-tsa-servers-bench/process"
	"log"
	"os"
)

// by default, re-start process every minutes
const defaultCronScheduling string = "0h1m0s"

func initLog() {
	file, err := os.OpenFile("/mnt/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Print("Logging correctly initialized")
}

func main() {
	initLog()
	scheduling := os.Getenv("CRON_SCHEDULING")
	if scheduling == "" {
		scheduling = defaultCronScheduling
	}
	tsaServer := os.Getenv("TSA_SERVER")
	if tsaServer != "" {
		log.Print("Process will run with this scheduling: " + scheduling)
		process.Loop(tsaServer)
	} else {
		log.Print("Process will run with this scheduling: " + scheduling)
		process.Schedule(scheduling)
	}

}
