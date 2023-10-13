package channel

type Channel struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type ChannelList struct {
	Data []Channel `json:"data"`
	Next string    `json:"next,omitempty"`
}
