package db

import "time"

type TbMakeFriend struct {
	Id         int64     `json:"id" form:"id" gorm:"column:id;primaryKey;autoIncrement"`
	FromUserId string    `json:"from_user_id" form:"from_user_id" gorm:"column:from_user_id;size:200;not null;uniqueIndex:uix_from_user_id"`
	ToUserId   string    `json:"to_user_id" form:"to_user_id" gorm:"column:to_user_id;size:200;not null;uniqueIndex:uix_to_user_id"`
	CreatedAt  time.Time `json:"created_at" form:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;index:idx_country_a"`
}

func (*TbMakeFriend) TableName() string {
	return "make_friend"
}
