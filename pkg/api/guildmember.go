package api

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guildmember"
)

type SatoriGuildMemberApi interface {
	GuildMemberGet(guild_id, user_id string) (*guildmember.GuildMember, error)
	GuildMemberList(guild_id, next string) (*guildmember.GuildMemberList, error)
	GuildMemberKick(guild_id, user_id string, permanent bool) error
	GuildMemberApprove(message_id string, approve bool, comment string) error
}

type satoriGuildMemberApiImpl struct {
	cli client.ApiTemplate
}

func newSatoriGuildMemberApi(cli client.ApiTemplate) (SatoriGuildMemberApi, error) {
	return &satoriGuildMemberApiImpl{
		cli: cli,
	}, nil

}

func (api *satoriGuildMemberApiImpl) GuildMemberGet(guild_id, user_id string) (*guildmember.GuildMember, error) {
	var result *guildmember.GuildMember
	err := api.cli.PostByRequestForResult("/guild.member.get", map[string]string{
		"guild_id": guild_id,
		"user_id":  user_id,
	}, &result)
	return result, err
}
func (api *satoriGuildMemberApiImpl) GuildMemberList(guild_id, next string) (*guildmember.GuildMemberList, error) {
	var result *guildmember.GuildMemberList
	err := api.cli.PostByRequestForResult("/guild.member.list", map[string]string{
		"guild_id": guild_id,
		"next":     next,
	}, &result)
	return result, err
}
func (api *satoriGuildMemberApiImpl) GuildMemberKick(guild_id, user_id string, permanent bool) error {
	return api.cli.PostByRequest("/guild.member.kick", map[string]interface{}{
		"guild_id":  guild_id,
		"user_id":   user_id,
		"permanent": permanent,
	})
}
func (api *satoriGuildMemberApiImpl) GuildMemberApprove(message_id string, approve bool, comment string) error {
	return api.cli.PostByRequest("/guild.member.approve", map[string]interface{}{
		"message_id": message_id,
		"approve":    approve,
		"comment":    comment,
	})
}
