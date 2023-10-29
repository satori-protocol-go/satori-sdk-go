package api

import (
	"github.com/dezhishen/satori-model-go/pkg/channel"
	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriChannelApi interface {
	ChannelGet(channel_id string) (*channel.Channel, error)
	ChannelList(guild_id string, next string) (*channel.ChannelList, error)
	ChannelCreate(guild_id string, data channel.Channel) (*channel.Channel, error)
	ChannelUpdate(channel_id string, data channel.Channel) error
	ChannelDelete(channel_id string) error
	UserChannelCreate(user_id string) (*channel.Channel, error)
}

type satoriChannelApiImpl struct {
	cli client.ApiTemplate
}

func newSatoriChannelApi(cli client.ApiTemplate) (SatoriChannelApi, error) {
	return &satoriChannelApiImpl{
		cli: cli,
	}, nil
}

func (api *satoriChannelApiImpl) ChannelGet(channel_id string) (*channel.Channel, error) {
	var result *channel.Channel
	err := api.cli.PostByRequestForResult("/channel.get", map[string]string{
		"channel_id": channel_id,
	}, &result)
	return result, err
}

func (api *satoriChannelApiImpl) ChannelList(guild_id string, next string) (*channel.ChannelList, error) {
	var result *channel.ChannelList
	err := api.cli.PostByRequestForResult("/channel.get", map[string]string{
		"guild_id": guild_id,
		"next":     next,
	}, &result)
	return result, err
}

func (api *satoriChannelApiImpl) ChannelCreate(guild_id string, data channel.Channel) (*channel.Channel, error) {
	var result *channel.Channel
	err := api.cli.PostByRequestForResult("/channel.create", map[string]interface{}{
		"guild_id": guild_id,
		"data":     data,
	}, &result)
	return result, err
}

func (api *satoriChannelApiImpl) ChannelUpdate(channel_id string, data channel.Channel) error {
	return api.cli.PostByRequest("/channel.update", map[string]interface{}{
		"channel_id": channel_id,
		"data":       data,
	})
}

func (api *satoriChannelApiImpl) ChannelDelete(channel_id string) error {
	return api.cli.PostByRequest("/channel.delete", map[string]interface{}{
		"channel_id": channel_id,
	})
}

func (api *satoriChannelApiImpl) UserChannelCreate(user_id string) (*channel.Channel, error) {
	var result *channel.Channel
	err := api.cli.PostByRequestForResult("/user.channel.create", map[string]interface{}{
		"user_id": user_id,
	}, &result)
	return result, err
}
