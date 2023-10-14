package guild

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriGuildApi interface {
	GuildGet(guild_id string) (*Guild, error)
	GuildList(next string) (*GuildList, error)
	HandleGuildRequest(message_id string, approve bool, comment string) error
}

type satoriGuildApiImpl struct {
	cli client.ApiTemplate
}

func NewSatoriGuildApi(cli client.ApiTemplate) (SatoriGuildApi, error) {
	return &satoriGuildApiImpl{cli: cli}, nil
}

func (api *satoriGuildApiImpl) GuildGet(guild_id string) (*Guild, error) {
	var result *Guild
	err := api.cli.PostByRequestForResult("/guild.get", map[string]string{
		"guild_id": guild_id,
	}, &result)
	return result, err
}

func (api *satoriGuildApiImpl) GuildList(next string) (*GuildList, error) {
	var result *GuildList
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
