package datasummary

import (
	"v2raydatastat/pkg/model"

	"gorm.io/gorm"
)

func GetByNameAndDate(username string, date string) (DataSummary, error) {
	var dataSummary DataSummary

	if err := model.DB.Where("username = ? AND date = ?",username, date).First(&dataSummary).Error; err != nil {
		return dataSummary, err
	}

	return dataSummary, nil
}

func (datasummary *DataSummary) Increase(uplink int64, downlink int64) error {
	uplinkResult := model.DB.Model(&datasummary).Update("uplink_byte", gorm.Expr("uplink_byte + ?", uplink))
	if uplinkResult.Error != nil {
		return uplinkResult.Error
	} 

	downlinkResult := model.DB.Model(&datasummary).Update("downlink_byte", gorm.Expr("downlink_byte + ?", downlink))
	if downlinkResult.Error != nil {
		return downlinkResult.Error
	} 

	return nil
}

func (datasummary *DataSummary) Create() error {
	result := model.DB.Create(&datasummary)
	if err := result.Error; err != nil {
		return err
	} else {
		return nil
	}
}
