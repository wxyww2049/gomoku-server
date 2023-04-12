package po

type Room struct {
	Id      int         `json:"id"`
	Msg     []PlayerMsg `json:"msg"`
	Steps   []Chess     `json:"steps"`
	Players []Player    `json:"players"`
	Status  bool        `json:"status"`
}
