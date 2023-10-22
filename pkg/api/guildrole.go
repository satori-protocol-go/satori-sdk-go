package api

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guildrole"
)

type SatoriGuildRoleApi interface {
	GuildMemberRoleSet(guild_id, user_id, role_id string) error
	GuildMemberRoleUnset(guild_id, user_id, role_id string) error
	GuildRoleList(guild_id, next string) (*guildrole.GuildRoleList, error)
	GuildRoleCreate(guild_id string, role guildrole.GuildRole) (*guildrole.GuildRole, error)
	GuildRoleUpdate(guild_id, role_id string, role guildrole.GuildRole) error
	GuildRoleDelete(guild_id, role_id string) error
}

type satoriGuildRoleApiImpl struct {
	cli client.ApiTemplate
}

func newSatoriGuildRoleApi(cli client.ApiTemplate) (SatoriGuildRoleApi, error) {
	return &satoriGuildRoleApiImpl{
		cli: cli,
	}, nil

}

func (api *satoriGuildRoleApiImpl) GuildMemberRoleSet(guild_id, user_id, role_id string) error {
	return api.cli.PostByRequest("/guild.member.role.set", map[string]string{
		"guild_id": guild_id,
		"user_id":  user_id,
		"role_id":  role_id,
	})
}
func (api *satoriGuildRoleApiImpl) GuildMemberRoleUnset(guild_id, user_id, role_id string) error {
	return api.cli.PostByRequest("/guild.member.role.unset", map[string]string{
		"guild_id": guild_id,
		"user_id":  user_id,
		"role_id":  role_id,
	})
}

func (api *satoriGuildRoleApiImpl) GuildRoleList(guild_id, next string) (*guildrole.GuildRoleList, error) {
	var result *guildrole.GuildRoleList
	err := api.cli.PostByRequestForResult("/guild.role.list", map[string]string{
		"guild_id": guild_id,
		"next":     next,
	}, &result)
	return result, err
}
func (api *satoriGuildRoleApiImpl) GuildRoleCreate(guild_id string, role guildrole.GuildRole) (*guildrole.GuildRole, error) {
	var result *guildrole.GuildRole
	err := api.cli.PostByRequestForResult("/guild.role.create", map[string]interface{}{
		"guild_id": guild_id,
		"role":     role,
	}, &result)
	return result, err
}
func (api *satoriGuildRoleApiImpl) GuildRoleUpdate(guild_id, role_id string, role guildrole.GuildRole) error {
	return api.cli.PostByRequest("/guild.role.update", map[string]interface{}{
		"guild_id": guild_id,
		"role_id":  role_id,
		"role":     role,
	})
}
func (api *satoriGuildRoleApiImpl) GuildRoleDelete(guild_id, role_id string) error {
	return api.cli.PostByRequest("/guild.role.delete", map[string]string{
		"guild_id": guild_id,
		"role_id":  role_id,
	})
}
