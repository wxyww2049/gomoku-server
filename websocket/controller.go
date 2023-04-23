package websocket

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gomoku-server/constants"
	"gomoku-server/data/dto"
	"gomoku-server/service"
	"gopkg.in/olahol/melody.v1"
)

func GetAllRooms() {
	msg := &dto.Message{
		Code: constants.GetAllPlayers,
	}
	msg.Data = service.ExRoomService.GetAllRoom()
	Broadcast(msg)
}
func GetAllPlayers() {
	msg := &dto.Message{
		Code: constants.GetAllPlayers,
	}
	msg.Data = service.ExPlayerService.GetAllPlayers()
	Broadcast(msg)
}
func CreateRoom(s *melody.Session, msg *dto.Message) {
	pId, ok := s.Get("Id")
	if !ok {
		fmt.Println("创建房间失败，找不到该用户")
		return
	}
	rid := uuid.NewV4().String()
	msg.Data = service.ExRoomService.CreateNewRoom(pId.(string), rid)
	send(s, msg)
	GetAllRooms()
}
func EnterRoom(s *melody.Session, msg *dto.Message) {
	pId, ok := s.Get("Id")
	if !ok {
		fmt.Println("加入房间失败，找不到该用户")
		return
	}
	rid := msg.Data.(string)
	r := service.ExRoomService.EnterRoom(pId.(string), rid)
	msg.Data = r.Player.Name + "加入了房间"
	send2Pid(r.Owner.Id, msg)
	send2Pid(r.Player.Id, msg)
	GetAllRooms()
}
