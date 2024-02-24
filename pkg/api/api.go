package api

import (
	"github.com/satori-protocol-go/satori-sdk-go/pkg/client"
	"github.com/satori-protocol-go/satori-sdk-go/pkg/config"
)

type SatoriApi interface {
	SatoriChannelApi
	SatoriGuildApi
	SatoriGuildMemberApi
	SatoriGuildRoleApi
	SatoriLoginApi
	SatoriMessageApi
	SatoriReactionApi
	SatoriUserApi
}

type SatoriApiImpl struct {
	SatoriChannelApi
	SatoriGuildApi
	SatoriGuildMemberApi
	SatoriGuildRoleApi
	SatoriLoginApi
	SatoriMessageApi
	SatoriReactionApi
	SatoriUserApi
}

func NewSatorApiByConfig(conf config.SatoriConfig) (SatoriApi, error) {
	client, err := client.NewApiTemplate(conf.Api)
	if err != nil {
		return nil, err
	}

	channelApi, err := newSatoriChannelApi(client)
	if err != nil {
		return nil, err
	}

	guildApi, err := newSatoriGuildApi(client)
	if err != nil {
		return nil, err
	}

	guildMemberApi, err := newSatoriGuildMemberApi(client)
	if err != nil {
		return nil, err
	}

	guildRoleApi, err := newSatoriGuildRoleApi(client)
	if err != nil {
		return nil, err
	}
	loginApi, err := newSatoriLoginApi(client)
	if err != nil {
		return nil, err
	}
	messageApi, err := newSatoriMessageApi(client)
	if err != nil {
		return nil, err
	}
	reactionApi, err := newSatoriReactionApi(client)
	if err != nil {
		return nil, err
	}
	userApi, err := newSatoriUserApi(client)
	if err != nil {
		return nil, err
	}
	return &SatoriApiImpl{
		channelApi,
		guildApi,
		guildMemberApi,
		guildRoleApi,
		loginApi,
		messageApi,
		reactionApi,
		userApi,
	}, nil
}
