package po

type Player struct {
	User
	Role  string `json:"role"`
	Color int8   `json:"color"`
	Ready bool   `json:"ready"`
}
