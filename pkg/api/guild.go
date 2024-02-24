package api

import (
	"github.com/satori-protocol-go/satori-model-go/pkg/guild"
	"github.com/satori-protocol-go/satori-sdk-go/pkg/client"
)

type SatoriGuildApi interface {
	GuildGet(guild_id string) (*guild.Guild, error)
	GuildList(next string) (*guild.GuildList, error)
	HandleGuildRequest(message_id string, approve bool, comment string) error
}

type satoriGuildApiImpl struct {
	cli client.ApiTemplate
}

func newSatoriGuildApi(cli client.ApiTemplate) (SatoriGuildApi, error) {
	return &satoriGuildApiImpl{cli: cli}, nil
}

func (api *satoriGuildApiImpl) GuildGet(guild_id string) (*guild.Guild, error) {
	var result *guild.Guild
	err := api.cli.PostByRequestForResult("/guild.get", map[string]string{
		"guild_id": guild_id,
	}, &result)
	return result, err
}

func (api *satoriGuildApiImpl) GuildList(next string) (*guild.GuildList, error) {
	var result *guild.GuildList
	err := api.cli.PostByRequestForResult("/guild.list", map[string]string{
		"next": next,
	}, &result)
	return result, err
}

func (api *satoriGuildApiImpl) HandleGuildRequest(message_id string, approve bool, comment string) error {
	return api.cli.PostByRequest("/guild.approve", map[string]interface{}{
		"message_id": message_id,
		"approve":    approve,
		"comment":    comment,
	})
}
