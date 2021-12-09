package db

import "time"

type TbTable struct {
	Id        int64     `json:"id" form:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `json:"name" form:"name" gorm:"column:name;size:100;not null"`
	TableId   string    `json:"table_id" form:"table_id" gorm:"column:table_id;size:200;not null;uniqueIndex:uix_table_id"`
	UserId    string    `json:"user_id" form:"user_id" gorm:"column:user_id;size:200;not null"`
	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;index:idx_country_a"`
}

func (*TbTable) TableName() string {
	return "game_table"
}
