package service

import (
	"encoding/json"
	"fmt"
	"gomoku-server/data/dto"
	"gomoku-server/data/po"
	"gomoku-server/pkg/ren"
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
func (R *RoomService) CreateNewRoom(playerId string, id string, msg *dto.Message) *po.Room {

	pa := maPlayer[playerId]
	r := &po.Room{
		Id:     id,
		Owner:  pa,
		Status: 0,
	}
	type roomMsg struct {
		Name string `json:"name"`
	}
	pa.Color = 0
	msgbyte, _ := json.Marshal(msg.Data)
	var tDetail roomMsg
	err := json.Unmarshal(msgbyte, &tDetail)

	if err != nil {
		fmt.Println("创建房间失败，解析数据失败")
		return nil
	}

	if tDetail.Name == "" {
		r.Name = pa.Name + "的房间"
	} else {
		r.Name = tDetail.Name
	}
	r.Winner = 0
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
	pa := maPlayer[playerId]
	pa.Color = 1 - r.Owner.Color
	if r.Player != nil {
		log.Println("房间已满")
		return nil
	}
	r.Player = pa
	r.Status = 1
	return r
}
func (R *RoomService) ExitRoom(playerId string, roomId string) bool {
	r := maRoom[roomId]
	p := maPlayer[playerId]
	log.Println("p is", *p)
	if r.Owner.Id == p.Id {
		if r.Player != nil {
			r.Owner = r.Player
			r.Player = nil
			r.Status = 0
			return true
		} else {
			r.Owner = nil
			log.Println("开始删除房间", roomId)
			ExRoomService.DelRoom(roomId)
			return true
		}
	} else if r.Player.Id == p.Id {
		r.Player = nil
		r.Status = 0
		return true
	} else {
		log.Println("没有找到对应房间")
		return false
	}
}
func (R *RoomService) StartGame(roomId string) (*po.Room, bool) {
	r := maRoom[roomId]
	if r.Status == 1 {
		r.Status = 2
		r.Steps = nil
		r.Color = 1 //黑棋先手
		if r.Winner != 0 {
			var tmp = r.Owner.Color
			r.Owner.Color = r.Player.Color
			r.Player.Color = tmp
		}
		r.Winner = 0
		return r, true
	} else {
		return r, false
	}
}

func ConvertToBoard(steps *[]po.Chess) ([15][15]int, int, int) {
	var board [15][15]int
	var x = 0
	var y = 0
	for _, value := range *steps {
		x = value.I
		y = value.J
		if value.Color == 0 {
			board[value.I][value.J] = -1
		} else {
			board[value.I][value.J] = 1
		}
	}
	return board, x, y
}
func (R *RoomService) AddNewChess(roomId string, msg *dto.Message) (*po.Room, int, *string) { //int 3为添加成功 0为出现错误 -1游戏结束白色方赢，1为游戏结束黑色方赢
	r := maRoom[roomId]
	var tchess po.Chess

	GetMsg(msg, &tchess)

	if r.Color != tchess.Color {
		remsg := "不是你的回合"
		return r, 0, &remsg
	}

	r.Steps = append(r.Steps, tchess)

	board, x, y := ConvertToBoard(&r.Steps)
	fb := ren.CheckForbid(board, x, y)
	if fb != 0 {
		remsg := "禁手"
		switch fb {
		case 1:
			remsg = "三三禁手"
			break
		case 2:
			remsg = "四四禁手"
			break
		case 3:
			remsg = "长连禁手"
			break
		default:
			break
		}
		r.Steps = append(r.Steps[:len(r.Steps)-1])
		return r, 0, &remsg
	}
	res := ren.CheckWin(board, x, y)
	if res != 0 {
		r.Status = 1
		r.Winner = res
		return r, res, nil
	}

	r.Color = 1 - r.Color
	return r, 3, nil
}
