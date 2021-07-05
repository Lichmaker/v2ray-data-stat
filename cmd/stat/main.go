package main

import (
	"log"
	"v2raydatastat/bootstrap"
	"v2raydatastat/config"
	"v2raydatastat/pkg/stat"

	"github.com/robfig/cron/v3"
)

func main() {
	log.SetPrefix("[v2ray-stat]")
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
