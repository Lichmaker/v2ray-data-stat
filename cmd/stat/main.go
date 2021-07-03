package main

import (
	"v2raydatastat/bootstrap"
	"v2raydatastat/config"
	"v2raydatastat/pkg/stat"

	"github.com/robfig/cron/v3"
)

func main() {
	config.Initialize()
	bootstrap.SetupDB()
	bootstrap.SetupGRPCConnect()

	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", stat.Handle)
	// c.AddFunc("*/5 * * * * *", func() { fmt.Println(123) })
	c.Start()
	defer c.Stop()
	select {}
}
