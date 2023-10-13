package api

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
	"github.com/dezhishen/satori-sdk-go/pkg/config"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/channel"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guild"
)

type SatoriApi interface {
	channel.SatoriChannelApi
	guild.SatoriGuildApi
}

type SatoriApiImpl struct {
	channel.SatoriChannelApi
	guild.SatoriGuildApi
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
	return &SatoriApiImpl{
		channelApi,
		guildApi,
	}, nil
}
