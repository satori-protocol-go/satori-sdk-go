package guildmember

type SatoriGuildMemberApi interface {
	GuildMemberGet(guild_id, user_id string) (GuildMember, error)
	GuildMemberList(guild_id, next string) (GuildMemberList, error)
	GuildMemberKick(guild_id, user_id string, permanent bool) error
	GuildMemberApprove(message_id string, approve bool, comment string) error
}
