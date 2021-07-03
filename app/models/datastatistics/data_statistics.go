package datastatistics

import "time"

type DataStatistics struct {
	ID           uint64    `gorm:"primaryKey"`
	UplinkByte   uint64    `gorm:"column:uplink_byte;type:bigint unsigned"`
	DownlinkByte uint64    `gorm:"column:downlink_byte;type:bigint unsigned"`
	Username     string    `gorm:"column:username;type:varchar(64);not null;index"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;index"`
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (DataStatistics) TableName() string {
	return "data_statistics"
}
