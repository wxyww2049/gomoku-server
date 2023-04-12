package po

type User struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
	Status   string `json:"status"`
}
