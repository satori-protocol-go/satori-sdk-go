package api

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/user"
)

type SatoriReactionApi interface {
	// 向特定消息添加表态。
	ReactionCreate(channel_id, message_id, emoji string) error
	// 从特定消息删除某个用户添加的特定表态。如果没有传入用户 ID 则表示删除自己的表态。
	ReactionDelete(channel_id, message_id, emoji, user_id string) error
	// 从特定消息清除某个特定表态。如果没有传入表态名称则表示清除所有表态。
	ReactionClear(channel_id, message_id, emoji string) error
	// 获取添加特定消息的特定表态的用户列表。返回一个 User 的 分页列表。
	ReactionList(channel_id, message_id, emoji, next string) (*user.UserList, error)
}

type satoriReactionApiImpl struct {
	cli client.ApiTemplate
}

func newSatoriReactionApi(cli client.ApiTemplate) (SatoriReactionApi, error) {
	return &satoriReactionApiImpl{cli}, nil
}

func (api *satoriReactionApiImpl) ReactionCreate(channel_id, message_id, emoji string) error {
	return api.cli.PostByRequest("/reaction.create", map[string]string{
		"channel_id": channel_id,
		"message_id": message_id,
		"emoji":      emoji,
	})
}
func (api *satoriReactionApiImpl) ReactionDelete(channel_id, message_id, emoji, user_id string) error {
	return api.cli.PostByRequest("/reaction.delete", map[string]string{
		"channel_id": channel_id,
		"message_id": message_id,
		"emoji":      emoji,
		"user_id":    user_id,
	})
}
func (api *satoriReactionApiImpl) ReactionClear(channel_id, message_id, emoji string) error {
	return api.cli.PostByRequest("/reaction.clear", map[string]string{
		"channel_id": channel_id,
		"message_id": message_id,
		"emoji":      emoji,
	})
}
func (api *satoriReactionApiImpl) ReactionList(channel_id, message_id, emoji, next string) (*user.UserList, error) {
	var result *user.UserList
	err := api.cli.PostByRequestForResult("/reaction.list", map[string]string{
		"channel_id": channel_id,
		"message_id": message_id,
		"emoji":      emoji,
		"next":       next,
	}, &result)
	return result, err
}
