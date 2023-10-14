package login

import (
	"errors"

	"github.com/dezhishen/satori-sdk-go/pkg/client"
)

type SatoriLoginApi interface {
	LoginGet() (*Login, error)
}

type satoriLoginApiImpl struct {
	cli client.ApiTemplate
}

func NewSatoriLoginApi(cli client.ApiTemplate) (SatoriLoginApi, error) {
	// todo
	return &satoriLoginApiImpl{cli: cli}, nil
}

func (api *satoriLoginApiImpl) LoginGet() (*Login, error) {
	// todo
	return nil, errors.New("unSupport")
}
