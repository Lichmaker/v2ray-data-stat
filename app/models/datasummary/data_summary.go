package datasummary

import "time"

type DataSummary struct {
	ID           uint64 `gorm:"primaryKey"`
	UplinkByte   uint64 `gorm:"column:uplink_byte;type:bigint unsigned"`
	DownlinkByte uint64 `gorm:"column:downlink_byte;type:bigint unsigned"`
	Date         string `gorm:"column:date;type:date;uniqueIndex:date_username"`
	Username     string `gorm:"column:username;type:varchar(64);not null;uniqueIndex:date_username;index"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;index"`
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (DataSummary) TableName() string {
	return "data_summary"
}
