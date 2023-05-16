package service

/**
服务层的包装
*/
import (
	"encoding/json"
	"gomoku-server/dao"
	"gomoku-server/data/dto"
	"gomoku-server/data/po"
	"log"
)

var (
	userDao  dao.UserDao           //数据库操作，弃用
	maPlayer map[string]*po.Player //实现由pid到玩家指针的映射
	players  []*po.Player          //玩家列表，弃用
	rooms    []*po.Room            //房间列表
	maRoom   map[string]*po.Room   //实现由rid到房间指针的映射
)

/*
*
实例化不同的服务类并导出
*/

var (
	ExUserService   UserService
	ExPlayerService PlayerService
	ExRoomService   RoomService
)

func GetMsg(msg *dto.Message, nty any) { //用于数据解析，讲msg解析为所需格式
	msgByte, _ := json.Marshal(msg.Data)
	err := json.Unmarshal(msgByte, &nty)
	if err != nil {
		log.Println("解析数据失败")
	}
}

/*
*
初始化服务层
*/
func Setup() {
	userDao = dao.UserDao{Tx: dao.DB}
	maPlayer = make(map[string]*po.Player, 20)
	maRoom = make(map[string]*po.Room, 20)
	players = make([]*po.Player, 0)
	rooms = make([]*po.Room, 0)
	ExUserService = UserService{}
}
