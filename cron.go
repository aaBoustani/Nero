package main

import (
	"log"

	"github.com/robfig/cron"
)

func ResetRemainingCRON() {
	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() {
		log.Println("Resetting remaining Nero to give")
		ResetAllRemaining()
	})
	c.Start()
}
