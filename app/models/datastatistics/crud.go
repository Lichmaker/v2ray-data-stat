package datastatistics

import (
	"v2raydatastat/pkg/model"
)

func (datastatistics *DataStatistics) Create() (err error) {
	result := model.DB.Create(&datastatistics)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
