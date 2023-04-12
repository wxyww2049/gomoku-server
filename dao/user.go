package dao

import (
	"gomoku-server/data/po"
	"gorm.io/gorm"
)

type UserDao struct {
	Tx *gorm.DB
}

func (u UserDao) SaveUser(user *po.User) {
	err := u.Tx.Save(user).Error
	if err != nil {
		panic(err)
	}
}
