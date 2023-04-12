package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gomoku-server/data/dto"
	"gopkg.in/olahol/melody.v1"
)

var (
	m      *melody.Melody
	logger *logrus.Logger
)

func InitMelody() *melody.Melody {
	m = melody.New()
	m.HandleMessage(Receive)
	return m
}
func send(s *melody.Session, msg *dto.Message) { //想特定的对话s发送消息
	msgByte, _ := json.Marshal(msg)

	if err := s.Write(msgByte); err != nil {
		logger.Error(err)
	}

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
