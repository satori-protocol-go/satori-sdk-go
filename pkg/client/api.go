package client

import (
	"fmt"

	"github.com/dezhishen/satori-sdk-go/pkg/config"
)

type ApiTemplate interface {
	Post(url string) error
	PostForResult(url string, result interface{}) error
	PostByRequest(url string, req interface{}) error
	PostByRequestForResult(url string, req, result interface{}) error
}

func NewApiTemplate(conf config.SatoriApiConfig) (ApiTemplate, error) {
	if conf.Type == "http" {
		return NewHttpApiTemplate(conf)
	}
	return nil, fmt.Errorf(
		"type %s not support,just support [%s]",
		conf.Type,
		"http",
	)
}
