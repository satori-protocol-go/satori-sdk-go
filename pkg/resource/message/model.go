package message

import (
	"github.com/dezhishen/satori-sdk-go/pkg/resource/channel"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guild"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guildmember"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/user"
)

type Message struct {
	Id       string                   `json:"id"`
	Content  string                   `json:"content"`
	Channel  *channel.Channel         `json:"channel,omitempty"`
	Guild    *guild.Guild             `json:"guild,omitempty"`
	Member   *guildmember.GuildMember `json:"member,omitempty"`
	User     *user.User               `json:"user,omitempty"`
	CreateAt int64                    `json:"create_at,omitempty"`
	UpdateAt int64                    `json:"update_at,omitempty"`
}

type MessageList struct {
	Data []Message `json:"data"`
	Next string    `json:"next,omitempty"`
}

type MessageContent struct {
	Text     string
	Elements []MessageElement
}

type MessageElement interface {
	Tag() string
}

func (m *Message) Decode(content *MessageContent) error {
	raw, err := messageContent2RawMessage(content)
	if err != nil {
		return err
	}
	m.Content = raw
	return nil
}

func (m *Message) Encode() (*MessageContent, error) {
	return rawMessage2MessageContent(m.Content)
}
