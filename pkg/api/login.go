package api

import (
	"github.com/satori-protocol-go/satori-model-go/pkg/login"
	"github.com/satori-protocol-go/satori-sdk-go/pkg/client"
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
