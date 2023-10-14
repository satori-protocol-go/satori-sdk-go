package guildmember

import (
	"errors"

	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriGuildMemberApi interface {
	GuildMemberGet(guild_id, user_id string) (*GuildMember, error)
	GuildMemberList(guild_id, next string) (*GuildMemberList, error)
	GuildMemberKick(guild_id, user_id string, permanent bool) error
	GuildMemberApprove(message_id string, approve bool, comment string) error
}

type satoriGuildMemberApiImpl struct {
	cli client.ApiTemplate
}

func NewSatoriGuildMemberApi(cli client.ApiTemplate) (SatoriGuildMemberApi, error) {
	return &satoriGuildMemberApiImpl{
		cli: cli,
	}, nil

}

func (api *satoriGuildMemberApiImpl) GuildMemberGet(guild_id, user_id string) (*GuildMember, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriGuildMemberApiImpl) GuildMemberList(guild_id, next string) (*GuildMemberList, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriGuildMemberApiImpl) GuildMemberKick(guild_id, user_id string, permanent bool) error {
	// todo
	return errors.New("unSupport")
}
func (api *satoriGuildMemberApiImpl) GuildMemberApprove(message_id string, approve bool, comment string) error {
	// todo
	return errors.New("unSupport")
}
