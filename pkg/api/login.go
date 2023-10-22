package api

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/login"
)

type SatoriLoginApi interface {
	LoginGet() (*login.Login, error)
}

type satoriLoginApiImpl struct {
	cli client.ApiTemplate
}

func newSatoriLoginApi(cli client.ApiTemplate) (SatoriLoginApi, error) {
	return &satoriLoginApiImpl{cli: cli}, nil
}

func (api *satoriLoginApiImpl) LoginGet() (*login.Login, error) {
	var result *login.Login
	err := api.cli.PostForResult("/login.get", &result)
	return result, err
}
