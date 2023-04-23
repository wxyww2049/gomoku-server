package service

import (
	"gomoku-server/dao"
	"gomoku-server/data/po"
)

var (
	userDao  dao.UserDao
	maPlayer map[string]*po.Player
	players  []*po.Player
	rooms    []*po.Room
	maRoom   map[string]*po.Room
)

var (
	ExUserService   UserService
	ExPlayerService PlayerService
	ExRoomService   RoomService
)

func Setup() {
	userDao = dao.UserDao{Tx: dao.DB}
	maPlayer = make(map[string]*po.Player, 20)
	players = make([]*po.Player, 0)
	rooms = make([]*po.Room, 0)
	ExUserService = UserService{}
}
