package channel

import (
	"errors"

	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriChannelApi interface {
	ChannelGet(channel_id string) (*Channel, error)
	ChannelList(guild_id string, next string) (*ChannelList, error)
	ChannelCreate(guild_id string, data Channel) (*Channel, error)
	ChannelUpdate(channel_id string, data Channel) error
	ChannelDelete(channel_id string) error
	UserChannelCreate(user_id string) (*Channel, error)
}

type satoriChannelApiImpl struct {
	cli client.ApiTemplate
}

func NewSatoriChannelApi(cli client.ApiTemplate) (SatoriChannelApi, error) {
	return &satoriChannelApiImpl{
		cli: cli,
	}, nil
}

func (api *satoriChannelApiImpl) ChannelGet(channel_id string) (*Channel, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriChannelApiImpl) ChannelList(guild_id string, next string) (*ChannelList, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriChannelApiImpl) ChannelCreate(guild_id string, data Channel) (*Channel, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriChannelApiImpl) ChannelUpdate(channel_id string, data Channel) error {
	// todo
	return errors.New("unSupport")
}
func (api *satoriChannelApiImpl) ChannelDelete(channel_id string) error {
	// todo
	return errors.New("unSupport")
}
func (api *satoriChannelApiImpl) UserChannelCreate(user_id string) (*Channel, error) {
	// todo
	return nil, errors.New("unSupport")
}
