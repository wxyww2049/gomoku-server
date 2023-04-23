package service

import (
	"gomoku-server/data/po"
)

type PlayerService struct{}

func (p PlayerService) Connect(id string, name string) bool {
	_, ok := maPlayer[id]
	if ok {
		return false
	}

	player := po.Player{
		Id:    id,
		Name:  name,
		Color: -1,
	}
	maPlayer[id] = &player
	players = append(players, &player)
	return true
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
