package grpc

import (
	"github.com/shimingyah/pool"
	"v2raydatastat/pkg/config"
)

var p pool.Pool

func NewPool() {
	var (
		err error
		host = config.GetString("grpc.host")
	)
	p, err = pool.New(host, pool.DefaultOptions)
	if err != nil {
		panic(err)
	}
}

func Get() (pool.Conn, error) {
	return p.Get();
}
