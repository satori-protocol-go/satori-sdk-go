package login

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriLoginApi interface {
	LoginGet() (*Login, error)
}

type satoriLoginApiImpl struct {
	cli client.ApiTemplate
}

func NewSatoriLoginApi(cli client.ApiTemplate) (SatoriLoginApi, error) {
	return &satoriLoginApiImpl{cli: cli}, nil
}

func (api *satoriLoginApiImpl) LoginGet() (*Login, error) {
	var result *Login
	err := api.cli.PostForResult("/login.get", &result)
	return result, err
}
