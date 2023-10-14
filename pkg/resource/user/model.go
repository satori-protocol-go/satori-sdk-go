package user

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	IsBot  bool   `json:"is_bot,omitempty"`
}

type UserList struct {
	Data []User `json:"data"`
	Next string `json:"next,omitempty"`
}
