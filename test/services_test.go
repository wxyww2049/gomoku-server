package test

import (
	"fmt"
	"gomoku-server/constants"
	"gomoku-server/data/po"
	"testing"
)

var (
	ma map[string]po.Player
)

func testMa(id string) {
	m := &po.Player{
		Name: "testName",
		Id:   id,
	}
	ma[id] = *m
}
func getName(id string) (string, bool) {
	m, ok := ma[id]
	if !ok {
		return "", false
	}
	return m.Name, true
}
func TestIt(t *testing.T) {
	ma = make(map[string]po.Player, 20)
	testMa("testId")
	name, ok := getName("testId")
	if !ok {
		println("there is no this key")
	} else {
		println(name)
	}
}

func TestCopy(t *testing.T) {

	myPlayer := &po.Player{
		Id: "player1",
	}

	myroom := &po.Room{
		Player: myPlayer,
	}

	myPlayer.Id = "player2"

	fmt.Println(myroom.Player.Id)
}
func TestMsgCode(t *testing.T) {
	fmt.Println(constants.ExitRoom)
}
