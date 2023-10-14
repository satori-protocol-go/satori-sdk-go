package message

import (
	"errors"

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
	// todo
	return &satoriMessageApiImpl{cli}, nil
}

func (api *satoriMessageApiImpl) MessageCreate(channel_id, content string) ([]Message, error) {
	// todo
	return nil, errors.New("unSupport")
}

func (api *satoriMessageApiImpl) MessageGet(channel_id, message_id string) (*Message, error) {
	// todo
	return nil, errors.New("unSupport")

}
func (api *satoriMessageApiImpl) MessageDelete(channel_id, message_id string) error {
	// todo
	return errors.New("unSupport")

}
func (api *satoriMessageApiImpl) MessageUpdate(channel_id, message_id, content string) error {
	// todo
	return errors.New("unSupport")

}
func (api *satoriMessageApiImpl) MessageList(channel_id, next string) (*MessageList, error) {
	// todo
	return nil, errors.New("unSupport")

}
