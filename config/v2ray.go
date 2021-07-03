package config

import (
	"v2raydatastat/pkg/config"
)

func init() {
	config.Add("v2ray", config.StrMap{
		"reset": config.Env("STAT_RESET", "TRUE"),
	})
}
