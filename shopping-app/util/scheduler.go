package util

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func StartScheduler() {

	c := cron.New()

	c.AddFunc("@every 1m", func() {
		fmt.Println("Running")
	})

	c.Start()
}
