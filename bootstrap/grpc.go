package bootstrap

import "v2raydatastat/pkg/grpc"



func SetupGRPCConnect() {
	grpc.NewPool()
}