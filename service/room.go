package service

import (
	"gomoku-server/data/po"
	"log"
)

type RoomService struct{}

func (R *RoomService) CheckNumOfRoom(room *po.Room) int {
	var ret = 0
	if room.Owner != nil {
		ret += 1
	}
	if room.Player != nil {
		ret += 1
	}
	return ret
}
func (R *RoomService) GetRoomById(id string) *po.Room {
	r, ok := maRoom[id]
	if ok == false {
		return nil
	}
	return r
}
func (R *RoomService) CreateNewRoom(playerId string, id string) *po.Room {
	r := &po.Room{
		Id:     id,
		Owner:  maPlayer[playerId],
		Status: 0,
	}
	maRoom[id] = r
	rooms = append(rooms, r)
	return r
}
func (R *RoomService) GetAllRoom() []*po.Room {
	return rooms
}
func (R *RoomService) DelRoom(id string) {
	_, ok := maRoom[id]
	if ok == false {
		log.Println("删除房间失败，ID:", id)
		return
	}
	delete(maRoom, id)
	for index, value := range rooms {
		if value.Id == id {
			rooms = append(rooms[:index], rooms[index+1:]...)
			break
		}
	}

}
func (R *RoomService) EnterRoom(playerId string, roomId string) *po.Room {

	r := maRoom[roomId]
	r.Player = maPlayer[playerId]
	r.Status = 1
	return r
}
func (R *RoomService) ExitRoom(playerId string, roomId string) bool {
	r := maRoom[roomId]
	p := maPlayer[playerId]
	if r.Owner == p {
		if r.Player != nil {
			r.Owner = r.Player
			r.Player = nil
			return true
		} else {
			r.Owner = nil
			R.DelRoom(roomId)
			return true
		}
	} else if r.Player == p {
		r.Player = nil
		return true
	} else {
		return false
	}
}
