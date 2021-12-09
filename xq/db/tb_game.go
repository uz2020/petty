package db

import "time"

type TbGame struct {
	Id        int64     `json:"id" form:"id" gorm:"column:id;primaryKey;autoIncrement"`
	TableId   string    `json:"table_id" form:"table_id" gorm:"column:table_id;size:200;not null;uniqueIndex:uix_table_id"`
	GameId    string    `json:"game_id" form:"game_id" gorm:"column:game_id;size:200;not null;uniqueIndex:uix_game_id"`
	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;index:idx_country_a"`
}

func (*TbGame) TableName() string {
	return "game"
}
