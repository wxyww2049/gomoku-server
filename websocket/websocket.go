package websocket

/**
websocket的基础函数和消息分发
*/

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gomoku-server/constants"
	"gomoku-server/data/dto"
	"gomoku-server/data/po"
	"gomoku-server/service"
	"gopkg.in/olahol/melody.v1"
	"log"
	"sync"
)

var (
	m      *melody.Melody //websocket的基础函数和消息分发
	logger *logrus.Logger //打印日志
	idMap  sync.Map       //用于映射pid和session，idMap[pid]指向id为pid的某个用户
)

func SendError(pid string, msg string) {
	send2Pid(pid, &dto.Message{
		Code: constants.Fail,
		Data: msg,
	})
} //返回错误信息

func SendToRoomPlayer(room *po.Room) { //向房间内的玩家发送房间状态
	log.Println("返回房间信息")
	type reMsg struct {
		Info *po.Player `json:"info"`
		Room *po.Room   `json:"room"`
	}
	msg := &dto.Message{
		Code: constants.UpdateRoomAndPlayer,
	}
	if room == nil {
		log.Println("返回房间信息时，未找到房间")
		return
	}

	if room.Owner != nil {
		log.Println("向房主发送消息")
		msg.Data = &reMsg{
			Info: room.Owner,
			Room: room,
		}
		send2Pid(room.Owner.Id, msg)
	}
	if room.Player != nil {
		msg.Data = &reMsg{
			Info: room.Player,
			Room: room,
		}
		send2Pid(room.Player.Id, msg)
	}
}
func sendInfoMsgToRoom(r *po.Room, sentence string) { //向房间内的玩家发送提示消息
	if r.Owner != nil {
		sendInfoMsg(r.Owner.Id, sentence)
	}
	if r.Player != nil {
		sendInfoMsg(r.Player.Id, sentence)
	}
}
func sendInfoMsg(pid string, rmsg string) { //向特定玩家发送提示消息
	msg := &dto.Message{
		Code: constants.InfoMsg,
		Data: rmsg,
	}
	send2Pid(pid, msg)
}

func InitMelody() *melody.Melody { //初始化melody，设置消息处理函数
	m = melody.New()
	m.HandleMessage(Receive)
	m.HandleConnect(Connect)
	m.HandleDisconnect(DisConnect)
	return m
}

func send(s *melody.Session, msg *dto.Message) { //向特定的对话s发送消息
	msgByte, _ := json.Marshal(msg)
	if err := s.Write(msgByte); err != nil {
		logger.Error(err)
	}
}

func Connect(s *melody.Session) { //用户连接时执行的函数，需要初始化用户信息
	id := uuid.NewV4().String()
	idMap.Store(id, s)
	s.Set("Id", id)
	p, ok := service.ExPlayerService.Connect(id, "unNamed")
	if !ok {
		log.Print("此id已连接")
	} else {
		log.Print(id, "连接成功")
	}
	msg := &dto.Message{
		Code: constants.Connect,
		Data: p,
	}
	send(s, msg)
	GetAllRooms()
}

func DisConnect(s *melody.Session) { //用户断开连接时执行的函数，需要退出目前所在的房间
	id, ok := s.Get("Id")
	fmt.Println(id, ok)
	if ok == false {
		return
	}
	rid, ok := s.Get("rid")
	if ok {
		log.Println("找到房间", rid.(string))
		service.ExRoomService.ExitRoom(id.(string), rid.(string))
		r := service.ExRoomService.GetRoomById(rid.(string))
		SendToRoomPlayer(r)
	}
	service.ExPlayerService.DisConnect(id.(string))
	idMap.LoadAndDelete(id)
	fmt.Println(id, "断开连接")
	GetAllRooms()
}

func Receive(s *melody.Session, msgByte []byte) { //收到消息侯的分发，根据收到消息的code分发给不同的处理函数
	msg := &dto.Message{}
	err := json.Unmarshal(msgByte, msg)
	if err != nil {
		send(s, dto.NewErrMsg(err))
		return
	}
	switch msg.Code {
	case constants.GetAllPlayers:
		GetAllPlayers()
	case constants.GetAllRoom:
		GetAllRooms()
	case constants.CreateRoom:
		CreateRoom(s, msg)
	case constants.PlayerRename:
		RenamePlayer(s, msg)
	case constants.EnterRoom:
		EnterRoom(s, msg)
	case constants.ExitRoom:
		ExitRoom(s, msg)
	case constants.StartGame:
		StartGame(s, msg)
	case constants.AddNewChess:
		AddNewChess(s, msg)
	case constants.AdmitDefeat:
		AdmitDefeat(s, msg)
	case constants.ChatMsg:
		ReciveMsg(s, msg)
	}
}

func Broadcast(msg *dto.Message) { //向所有对话广播消息
	msgByte, _ := json.Marshal(msg)
	if err := m.Broadcast(msgByte); err != nil {
		logger.Error(err)
	}
}

func send2Pid(id string, msg *dto.Message) { //根据pid向玩家发送消息
	ts, ok := idMap.Load(id)
	if !ok {
		panic("找不到该id")
		return
	}
	s, ook := ts.(*melody.Session)
	if !ook {
		logger.Error("ts不是melody.session")
	}
	log.Println("向", id, "发送消息", msg)
	send(s, msg)
	return
}
