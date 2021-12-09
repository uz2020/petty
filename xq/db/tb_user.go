package db

import "time"

type TbUser struct {
	Id        int64     `json:"id" form:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `json:"username" form:"username" gorm:"column:username;size:100;not null"`
	UserId    string    `json:"user_id" form:"user_id" gorm:"column:user_id;size:200;not null;uniqueIndex:uix_user_id"`
	Password  string    `json:"password" form:"password" gorm:"column:password;size:200;not null"`
	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;index:idx_country_a"`
}

func (*TbUser) TableName() string {
	return "user"
}
