package service

import (
	"gomoku-server/data/dto"
	"gomoku-server/data/po"
)

type PlayerService struct{}

func (p PlayerService) Connect(id string, name string) (*po.Player, bool) {
	_, ok := maPlayer[id]
	if ok {
		return nil, false
	}

	player := po.Player{
		Id:    id,
		Name:  name,
		Color: -1,
	}
	maPlayer[id] = &player
	players = append(players, &player)
	return &player, true
}
func (p PlayerService) DisConnect(id string) bool {
	_, ok := maPlayer[id]
	if !ok {
		return false
	}
	delete(maPlayer, id)
	for index, value := range players {
		if value.Id == id {
			players = append(players[:index], players[index+1:]...)
			break
		}
	}
	return true
}
func (p PlayerService) GetAllPlayers() []*po.Player {
	return players
}
func (p PlayerService) Rename(id string, msg *dto.Message) (*po.Player, bool) {
	player, ok := maPlayer[id]
	if !ok {
		return nil, false
	}
	type renameMsg struct {
		Name string `json:"name"`
	}
	var tname renameMsg
	GetMsg(msg, &tname)
	player.Name = tname.Name

	return player, true
}
