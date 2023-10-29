package api

import (
	"github.com/dezhishen/satori-model-go/pkg/user"
	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriUserApi interface {
	UserGet(user_id string) (*user.User, error)
	FriendList(next string) (*user.UserList, error)
	FriendApprove(message_id string, approve bool, comment string) error
}

type satoriUserApiImpl struct {
	cli client.ApiTemplate
}

func newSatoriUserApi(cli client.ApiTemplate) (SatoriUserApi, error) {
	return &satoriUserApiImpl{cli}, nil
}
func (api *satoriUserApiImpl) UserGet(user_id string) (*user.User, error) {
	var result *user.User
	err := api.cli.PostByRequestForResult("/user.get", map[string]string{"user_id": user_id}, &result)
	return result, err
}
func (api *satoriUserApiImpl) FriendList(next string) (*user.UserList, error) {
	var result *user.UserList
	err := api.cli.PostByRequestForResult("/user.get", map[string]string{"next": next}, &result)
	return result, err
}
func (api *satoriUserApiImpl) FriendApprove(message_id string, approve bool, comment string) error {
	return api.cli.PostByRequest("/user.get",
		map[string]interface{}{
			"message_id": message_id,
			"approve":    approve,
			"comment":    comment,
		})
}
