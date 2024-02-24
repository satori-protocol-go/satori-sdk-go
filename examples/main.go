package main

import (
	"context"
	"fmt"

	"github.com/satori-protocol-go/satori-sdk-go/pkg/api"
	"github.com/satori-protocol-go/satori-sdk-go/pkg/config"
	"github.com/satori-protocol-go/satori-sdk-go/pkg/event"
)

func main() {
	conf := config.SatoriConfig{
		Api: config.SatoriApiConfig{
			Type:     "http",
			Platform: "red",
			SelfId:   "11123456789",
			Endpoint: "http://127.0.0.1:5140", //https://
		},
		Event: config.SatoriEventConfig{
			Type: "websocket",
			Addr: "ws://127.0.0.1:5140", // wss://
		},
		// Event: config.SatoriEventConfig{
		// 	Type: "webhook",
		// 	Addr: "0.0.0.0:8080", // 客户端服务的地址，127.0.0.1:8080 限制访问
		// },
	}
	satoriApi, err := api.NewSatorApiByConfig(conf)
	if err != nil {
		panic(err)
	}
	guildList, err := satoriApi.GuildList("")
	if err != nil {
		fmt.Printf("err:%v", err)
	} else {
		fmt.Printf("guildList is :%v,next is :%s", guildList.Data, guildList.Next)
	}

	satoriEvent, err := event.NewSatorEventByConfig(conf)
	if err != nil {
		panic(err)
	}
	satoriEvent.ListenMessageCreated(func(e event.Event) error {
		fmt.Printf("message.created %v", e)
		return nil
	})
	ctx := context.Background()
	err = satoriEvent.StartWithCtx(ctx)
	if err != nil {
		panic(err)
	}

}
