package message

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriMessageApi interface {
	MessageCreate(channel_id, content string) ([]Message, error)
	MessageGet(channel_id, message_id string) (*Message, error)
	MessageDelete(channel_id, message_id string) error
	MessageUpdate(channel_id, message_id, content string) error
	MessageList(channel_id, next string) (*MessageList, error)
}

type satoriMessageApiImpl struct {
	cli client.ApiTemplate
}

func NewSatoriMessageApi(cli client.ApiTemplate) (SatoriMessageApi, error) {
	return &satoriMessageApiImpl{cli}, nil
}

func (api *satoriMessageApiImpl) MessageCreate(channel_id, content string) ([]Message, error) {
	var result []Message
	err := api.cli.PostByRequestForResult(
		"/message.create",
		map[string]string{
			"channel_id": channel_id,
			"content":    content,
		},
		result,
	)
	return result, err
}

func (api *satoriMessageApiImpl) MessageGet(channel_id, message_id string) (*Message, error) {
	var result *Message
	err := api.cli.PostByRequestForResult(
		"/message.get",
		map[string]string{
			"channel_id": channel_id,
			"message_id": message_id,
		},
		result,
	)
	return result, err

}
func (api *satoriMessageApiImpl) MessageDelete(channel_id, message_id string) error {
	return api.cli.PostByRequest("/message.delete", map[string]string{
		"channel_id": channel_id,
		"message_id": message_id,
	})

}
func (api *satoriMessageApiImpl) MessageUpdate(channel_id, message_id, content string) error {
	return api.cli.PostByRequest("/message.update", map[string]string{
		"channel_id": channel_id,
		"message_id": message_id,
		"content":    content,
	})

}
func (api *satoriMessageApiImpl) MessageList(channel_id, next string) (*MessageList, error) {
	var result *MessageList
	err := api.cli.PostByRequestForResult(
		"/message.list",
		map[string]string{
			"channel_id": channel_id,
			"next":       next,
		},
		result,
	)
	return result, err
}
