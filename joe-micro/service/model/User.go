package model

import (
	"github.com/wangmhgo/microservice-project/joe-micro/lib/orm"
)

type  User struct {
	Uid   string     `gorm:"Column:uid"`
	Username  string   `gorm:"Column:username"`
	Avatar  string     `gorm:"Column:avatar"`
    UserType   int      `gorm:"Column:userType"`
}



func GetByID(uid string) ( User, error) {
	var user User
	err := orm.GetDB().Where("uid = ?", uid).Find(&user).Error
	return user, err
}



