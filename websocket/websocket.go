package websocket

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
	m      *melody.Melody
	logger *logrus.Logger
	idMap  sync.Map
)

func SendToRoomPlayer(room *po.Room) {
	type reMsg struct {
		Info *po.Player `json:"info"`
		Room *po.Room   `json:"room"`
	}
	msg := &dto.Message{
		Code: constants.UpdateRoomAndPlayer,
	}
	if room == nil {
		return
	}

	if room.Owner != nil {
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
func InitMelody() *melody.Melody {
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

func Connect(s *melody.Session) {
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

func DisConnect(s *melody.Session) {
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

func Receive(s *melody.Session, msgByte []byte) { //收到消息侯的分发
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
	}
}

func Broadcast(msg *dto.Message) { //向所有对话广播消息
	msgByte, _ := json.Marshal(msg)
	if err := m.Broadcast(msgByte); err != nil {
		logger.Error(err)
	}
}

func send2Pid(id string, msg *dto.Message) {
	ts, ok := idMap.Load(id)
	if !ok {
		panic("找不到该id")
		return
	}
	s, ook := ts.(*melody.Session)
	if !ook {
		logger.Error("ts不是melody.session")
	}
	send(s, msg)
	return
}
