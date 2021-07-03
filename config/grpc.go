package config

import (
	"v2raydatastat/pkg/config"
)

func init() {
	config.Add("grpc", config.StrMap{
		"host": config.Env("GRPC_HOST", "127.0.0.1:80"),
	})
}
