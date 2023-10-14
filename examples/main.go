package main

import (
	"fmt"

	"github.com/dezhishen/satori-sdk-go/pkg/api"
	"github.com/dezhishen/satori-sdk-go/pkg/config"
)

func main() {
	conf := config.SatoriConfig{
		Api: config.SatoriApiConfig{
			Type:        "http",
			Endpoint:    "http://127.0.0.1:6379",
			AccessToken: "adawdawd",
		},
		Event: config.SatoriEventConfig{
			Type: "webhook",
			Addr: "0.0.0.0:8080",
		},
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
	// ...
	// todo event examples
}
