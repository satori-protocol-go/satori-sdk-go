package client

import (
	"errors"

	"github.com/dezhishen/satori-sdk-go/pkg/config"
)

type ApiTemplate interface {
	Post(url string) error
	PostForResult(url string, result interface{}) error
	PostByRequest(url string, req interface{}) error
	PostByRequestForResult(url string, req, result interface{}) error
}

func NewApiTemplate(conf *config.SatoriApiConfig) (ApiTemplate, error) {
	if conf == nil {
		return nil, errors.New("api config not found")
	}
	if conf.Type == "http" {
		return NewHttpApiTemplate(conf)
	}
	if conf.Type == "websocket" {
		return NewWebsocketApiTemplate(conf)
	}
	return nil, errors.New("channel not support")
}
