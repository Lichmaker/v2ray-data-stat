package v2rayrequest

import (
	"context"
	"strings"
	"time"
	"v2raydatastat/pkg/grpc"

	"github.com/v2fly/v2ray-core/v4/app/stats/command"
)

var DataStat []struct {
	Username string
	Type     string
	Value    string
}

func GetDataStat(reset bool) ([]struct {
	Username     string
	UplinkByte   int64
	DownlinkByte int64
}, error) {
	var DataStat []struct {
		Username     string
		UplinkByte   int64
		DownlinkByte int64
	}
	usernameMap := make(map[string]struct{})
	uplinkMap := make(map[string]int64)
	downlinkMap := make(map[string]int64)

	connect, err := grpc.Get()
	if err != nil {
		return nil, err
	}
	defer connect.Close()

	client := command.NewStatsServiceClient(connect.Value())
	request := new(command.QueryStatsRequest)
	request.Reset_ = reset
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	respond, err := client.QueryStats(ctx, request)
	if err != nil {
		return nil, err
	}

	for _, value := range respond.Stat {
		slice1 := strings.Split(value.Name, ">>>")

		if slice1[3] == "uplink" {
			uplinkMap[slice1[1]] = value.Value
		} else {
			downlinkMap[slice1[1]] = value.Value
		}
		usernameMap[slice1[1]] = struct{}{}
	}

	for _username := range usernameMap {
		var s struct {
			Username     string
			UplinkByte   int64
			DownlinkByte int64
		}
		if _v, ok := uplinkMap[_username]; ok {
			s.UplinkByte = _v
		} else {
			s.UplinkByte = 0
		}
		if _v, ok := downlinkMap[_username]; ok {
			s.DownlinkByte = _v
		} else {
			s.DownlinkByte = 0
		}
		s.Username = _username
		DataStat = append(DataStat, s)
	}

	return DataStat, nil
}
