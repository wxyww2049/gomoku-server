package service

import (
	"gomoku-server/dao"
	"gomoku-server/data/po"
)

var (
	userDao  dao.UserDao
	maPlayer map[string]po.Player
)

var (
	ExUserService   UserService
	ExPlayerService PlayerService
)

func Setup() {
	userDao = dao.UserDao{Tx: dao.DB}
	maPlayer = make(map[string]po.Player, 20)

	ExUserService = UserService{}
}
