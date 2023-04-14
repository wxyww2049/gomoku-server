package po

type Player struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Role  string `json:"role"`
	Color int8   `json:"color"`
	Ready bool   `json:"ready"`
}
