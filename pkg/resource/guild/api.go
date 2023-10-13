package guild

import (
	"errors"

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
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriGuildApiImpl) GuildList(next string) (*GuildList, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriGuildApiImpl) HandleGuildRequest(message_id string, approve bool, comment string) error {
	// todo
	return errors.New("unSupport")
}
