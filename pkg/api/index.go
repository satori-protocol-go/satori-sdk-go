package api

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
	"github.com/dezhishen/satori-sdk-go/pkg/config"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/channel"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guild"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guildmember"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guildrole"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/login"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/message"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/reaction"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/user"
)

type SatoriApi interface {
	channel.SatoriChannelApi
	guild.SatoriGuildApi
	guildmember.SatoriGuildMemberApi
	guildrole.SatoriGuildRoleApi
	login.SatoriLoginApi
	message.SatoriMessageApi
	reaction.SatoriReactionApi
	user.SatoriUserApi
}

type SatoriApiImpl struct {
	channel.SatoriChannelApi
	guild.SatoriGuildApi
	guildmember.SatoriGuildMemberApi
	guildrole.SatoriGuildRoleApi
	login.SatoriLoginApi
	message.SatoriMessageApi
	reaction.SatoriReactionApi
	user.SatoriUserApi
}

func NewSatorApiByConfig(conf config.SatoriConfig) (SatoriApi, error) {
	client, err := client.NewApiTemplate(conf.Api)
	if err != nil {
		return nil, err
	}

	channelApi, err := channel.NewSatoriChannelApi(client)
	if err != nil {
		return nil, err
	}

	guildApi, err := guild.NewSatoriGuildApi(client)
	if err != nil {
		return nil, err
	}

	guildMemberApi, err := guildmember.NewSatoriGuildMemberApi(client)
	if err != nil {
		return nil, err
	}

	guildRoleApi, err := guildrole.NewSatoriGuildRoleApi(client)
	if err != nil {
		return nil, err
	}
	loginApi, err := login.NewSatoriLoginApi(client)
	if err != nil {
		return nil, err
	}
	messageApi, err := message.NewSatoriMessageApi(client)
	if err != nil {
		return nil, err
	}
	reactionApi, err := reaction.NewSatoriReactionApi(client)
	if err != nil {
		return nil, err
	}
	userApi, err := user.NewSatoriUserApi(client)
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
