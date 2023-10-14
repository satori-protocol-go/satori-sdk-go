package user

import (
	"errors"

	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriUserApi interface {
	UserGet(user_id string) (*User, error)
	FriendList(next string) (*UserList, error)
	FriendApprove(message_id string, approve bool, comment string) error
}

type satoriUserApiImpl struct {
	cli client.ApiTemplate
}

func NewSatoriUserApi(cli client.ApiTemplate) (SatoriUserApi, error) {
	// todo
	return &satoriUserApiImpl{cli}, nil
}
func (api *satoriUserApiImpl) UserGet(user_id string) (*User, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriUserApiImpl) FriendList(next string) (*UserList, error) {
	// todo
	return nil, errors.New("unSupport")
}
func (api *satoriUserApiImpl) FriendApprove(message_id string, approve bool, comment string) error {
	// todo
	return errors.New("unSupport")
}
