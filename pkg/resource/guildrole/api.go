package guildrole

import (
	"errors"

	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriGuildRoleApi interface {
	GuildRoleSet(guild_id, user_id, role_id string) error
	GuildRoleList(guild_id, next string) (*GuildRoleList, error)
	GuildRoleCreate(guild_id string, role GuildRole) (*GuildRole, error)
	GuildRoleUpdate(guild_id, role_id string, role GuildRole) error
	GuildRoleDelete(guild_id, role_id string) error
}

type satoriGuildRoleApiImpl struct {
	cli client.ApiTemplate
}

func NewSatoriGuildRoleApi(cli client.ApiTemplate) (SatoriGuildRoleApi, error) {
	return &satoriGuildRoleApiImpl{
		cli: cli,
	}, nil

}

func (api *satoriGuildRoleApiImpl) GuildRoleSet(guild_id, user_id, role_id string) error {
	// todo
	return errors.New("unSupport")
}
func (api *satoriGuildRoleApiImpl) GuildRoleList(guild_id, next string) (*GuildRoleList, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriGuildRoleApiImpl) GuildRoleCreate(guild_id string, role GuildRole) (*GuildRole, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriGuildRoleApiImpl) GuildRoleUpdate(guild_id, role_id string, role GuildRole) error {
	// todo
	return errors.New("unSupport")
}
func (api *satoriGuildRoleApiImpl) GuildRoleDelete(guild_id, role_id string) error {
	return errors.New("unSupport")
}
