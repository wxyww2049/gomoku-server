package service

import (
	"encoding/json"
	"gomoku-server/dao"
	"gomoku-server/data/dto"
	"gomoku-server/data/po"
	"log"
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

func GetMsg(msg *dto.Message, nty any) {
	msgByte, _ := json.Marshal(msg.Data)
	err := json.Unmarshal(msgByte, &nty)
	if err != nil {
		log.Println("解析数据失败")
	}
}
func Setup() {
	userDao = dao.UserDao{Tx: dao.DB}
	maPlayer = make(map[string]*po.Player, 20)
	maRoom = make(map[string]*po.Room, 20)
	players = make([]*po.Player, 0)
	rooms = make([]*po.Room, 0)
	ExUserService = UserService{}
}
