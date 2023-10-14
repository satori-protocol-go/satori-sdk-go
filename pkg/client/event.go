package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/dezhishen/satori-sdk-go/pkg/config"
)

type EventChannel interface {
	StartListen(context.Context, func(message []byte) error) error
	// callback func(message []byte) error
}

func NewEventChannel(conf config.SatoriEventConfig) (EventChannel, error) {
	if conf.Type == "" {
		return nil, errors.New("type is empty")
	}
	switch conf.Type {
	// websocket
	case "websocket":
		return NewWebsocketEventChannel(conf)
	// webhook
	case "webhook":
		return NewWebhookEventChannel(conf)
	// 反向websocket
	default:
		return nil, fmt.Errorf(
			"type %s not support,just support [%s,%s]",
			conf.Type,
			"websocket",
			"webhook",
		)
	}

}
