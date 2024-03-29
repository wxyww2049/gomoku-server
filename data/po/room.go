package po

type Room struct {
	Name   string      `json:"name"`
	Id     string      `json:"id"`     //roomId
	Msg    []PlayerMsg `json:"msg"`    //用户谈天记录
	Steps  []Chess     `json:"steps"`  //用户下棋记录
	Owner  *Player     `json:"owner"`  //拥有者
	Player *Player     `json:"player"` //非拥有者玩家
	Status int         `json:"status"` //0表示等待中，1表示人全了还未开始，2表示正在比赛
	Color  int         `json:"color"`  //1表示黑棋，0表示白棋
	Winner int         `json:"winner"` //1表示黑棋赢，-1表示白棋赢
}
