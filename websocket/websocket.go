package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gomoku-server/data/dto"
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
	if !service.ExPlayerService.Connect(id, "unNamed") {
		log.Print("此id已连接")
	} else {
		log.Print(id, "连接成功")
	}
}
func DisConnect(s *melody.Session) {
	id, ok := s.Get("Id")
	fmt.Println(id, ok)
	if ok == false {
		return
	}
	service.ExPlayerService.DisConnect(id.(string))
	idMap.LoadAndDelete(id)
	fmt.Println(id, "断开连接")
}
func Receive(s *melody.Session, msgByte []byte) { //收到消息侯的分发
	msg := &dto.Message{}
	err := json.Unmarshal(msgByte, msg)
	if err != nil {
		send(s, dto.NewErrMsg(err))
		return
	}
	msgOut, err := json.Marshal(msg.Code)
	fmt.Println(msg)
	s.Write(msgOut)
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
		return
	}
	send(s, msg)
}
