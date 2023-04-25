package websocket

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gomoku-server/constants"
	"gomoku-server/data/dto"
	"gomoku-server/service"
	"gopkg.in/olahol/melody.v1"
	"log"
)

func GetAllRooms() {
	msg := &dto.Message{
		Code: constants.GetAllRoom,
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
func ExitRoom(s *melody.Session, msg *dto.Message) {
	pId, ok := s.Get("Id")
	if !ok {
		fmt.Println("退出房间失败，找不到该用户")
		return
	}
	rid, ok := s.Get("rid")
	if !ok {
		fmt.Println("退出房间失败，找不到该房间")
		return
	}
	log.Println("开始退出房间")
	service.ExRoomService.ExitRoom(pId.(string), rid.(string))
	SendToRoomPlayer(service.ExRoomService.GetRoomById(rid.(string)))
	send(s, msg)
	GetAllRooms()

}
func CreateRoom(s *melody.Session, msg *dto.Message) {

	pId, ok := s.Get("Id")
	if !ok {
		fmt.Println("创建房间失败，找不到该用户")
		return
	}

	rid := uuid.NewV4().String()
	s.Set("rid", rid)

	r := service.ExRoomService.CreateNewRoom(pId.(string), rid, msg)
	SendToRoomPlayer(r)
	//send(s, msg)
	GetAllRooms()
}
func EnterRoom(s *melody.Session, msg *dto.Message) {
	pId, ok := s.Get("Id")
	if !ok {
		fmt.Println("加入房间失败，找不到该用户")
		return
	}
	rid := msg.Data.(string)
	s.Set("rid", rid)
	r := service.ExRoomService.EnterRoom(pId.(string), rid)
	if r == nil {
		msg.Code = constants.Fail
		msg.Data = "房间已满"
		send(s, msg)
	} else {
		SendToRoomPlayer(r)
		GetAllRooms()
	}
}

func RenamePlayer(s *melody.Session, msg *dto.Message) {
	pId, ok := s.Get("Id")
	if !ok {
		fmt.Println("重命名失败，找不到该用户")
		return
	}
	p, ok := service.ExPlayerService.Rename(pId.(string), msg)
	if !ok {
		fmt.Println("重命名失败")
		return
	}
	msg.Data = p
	send(s, msg)
	GetAllRooms()
}
