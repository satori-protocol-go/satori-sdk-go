package user

type User struct {
}

type UserList struct {
	Data []User `json:"data"`
	Next string `json:"next,omitempty"`
}
