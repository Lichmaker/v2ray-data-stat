package stat

import (
	"errors"
	"log"
	"strconv"
	"time"
	"v2raydatastat/app/models/datastatistics"
	"v2raydatastat/app/models/datasummary"
	"v2raydatastat/pkg/config"
	"v2raydatastat/pkg/v2rayrequest"

	"gorm.io/gorm"
)

var model datasummary.DataSummary

func Handle() {
	reset, _ := strconv.ParseBool(config.GetString("v2ray.reset"))
	data, e := v2rayrequest.GetDataStat(reset)
	if e != nil {
		log.Println(e)
		return
	}
	log.Println(data)

	var err error

	currentDate := time.Now().Local().Format("2006-01-02")
	for _, value := range data {
		if value.DownlinkByte == 0 && value.UplinkByte == 0 {
			// 没数据，直接忽略
			continue
		}

		statModel := datastatistics.DataStatistics{
			UplinkByte:   uint64(value.UplinkByte),
			DownlinkByte: uint64(value.DownlinkByte),
			Username:     value.Username,
		}
		statModel.Create()

		model, err = datasummary.GetByNameAndDate(value.Username, currentDate)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新记录
			model = datasummary.DataSummary{
				Username: value.Username,
				Date:     currentDate,
			}
			model.Create()
		}

		model.Increase(value.UplinkByte, value.DownlinkByte)
	}
}
