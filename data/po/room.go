package po

type Room struct {
	Id     int         `json:"id"`
	Msg    []PlayerMsg `json:"msg"`
	Steps  []Chess     `json:"steps"`
	Owner  Player      `json:"owner"`
	Player Player      `json:"player"`
	Status bool        `json:"status"`
}
