package api

import (
	"github.com/satori-protocol-go/satori-model-go/pkg/message"
	"github.com/satori-protocol-go/satori-sdk-go/pkg/client"
)

type SatoriMessageApi interface {
	MessageCreate(channel_id, content string) ([]message.Message, error)
	MessageGet(channel_id, message_id string) (*message.Message, error)
	MessageDelete(channel_id, message_id string) error
	MessageUpdate(channel_id, message_id, content string) error
	MessageList(channel_id, next string) (*message.MessageList, error)
}

type satoriMessageApiImpl struct {
	cli client.ApiTemplate
}

func newSatoriMessageApi(cli client.ApiTemplate) (SatoriMessageApi, error) {
	return &satoriMessageApiImpl{cli}, nil
}

func (api *satoriMessageApiImpl) MessageCreate(channel_id, content string) ([]message.Message, error) {
	var result []message.Message
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

func (api *satoriMessageApiImpl) MessageGet(channel_id, message_id string) (*message.Message, error) {
	var result *message.Message
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
func (api *satoriMessageApiImpl) MessageList(channel_id, next string) (*message.MessageList, error) {
	var result *message.MessageList
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
