package service

import "gomoku-server/dao"

var (
	userDao dao.UserDao
)

var (
	ExUserService UserService
)

func Setup() {
	userDao = dao.UserDao{Tx: dao.DB}
	ExUserService = UserService{}
}
