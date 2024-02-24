package client

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/satori-protocol-go/satori-sdk-go/pkg/config"
	"github.com/satori-protocol-go/satori-sdk-go/pkg/constant"
	log "github.com/sirupsen/logrus"
)

type WebsocketEventChannel struct {
	mu          sync.Mutex
	Conn        *websocket.Conn
	addr        string
	accessToken string
	sequence    int64
}

func NewWebsocketEventChannel(conf config.SatoriEventConfig) (EventTemplate, error) {
	if strings.HasSuffix("/", conf.Addr) {
		conf.Addr = strings.TrimSuffix(conf.Addr, "/")
	}
	version := conf.Version
	if version == "" {
		version = "v1"
	}
	result := &WebsocketEventChannel{
		addr:        fmt.Sprintf("%s/%s", conf.Addr, version),
		accessToken: conf.AccessToken,
	}
	return result, nil
}
func (cli *WebsocketEventChannel) SetSequence(sequence int64) {
	cli.sequence = sequence
}
func (cli *WebsocketEventChannel) sendIDENTIFY() {
	defer cli.mu.Unlock()
	cli.mu.Lock()
	log.Info("send IDENTIFY")
	err := cli.Conn.WriteJSON(map[string]interface{}{
		"op": constant.SIGN_NUM_IDENTIFY,
		"body": map[string]interface{}{
			"token":    cli.accessToken,
			"sequence": cli.sequence,
		},
	})
	if err != nil {
		log.Errorf("IDENTIFY发送失败:%v", err)
	}
}

func (cli *WebsocketEventChannel) startHeartbeat() {
	for {
		cli.sendPING()
		time.Sleep(10 * time.Second)
	}
}

func (cli *WebsocketEventChannel) sendPING() {
	defer cli.mu.Unlock()
	cli.mu.Lock()
	log.Info("send PING")
	err := cli.Conn.WriteJSON(map[string]interface{}{
		"op": constant.SIGN_NUM_PING,
	})
	if err != nil {
		log.Errorf("PING发送失败:%v", err)
	}
}

func (cli *WebsocketEventChannel) StartListen(ctx context.Context, callback func(message []byte) error) error {
	url := cli.addr + "/events"
	if cli.accessToken != "" {
		url = fmt.Sprintf("%s?access_token=%s", url, cli.accessToken)
	}
	// url := cli.endpoint + "?access_token=" + cli.accessToken
	log.Infof("开始建立连接...地址:%v", url)
	var err error
	cli.Conn, _, err = websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		log.Errorf("建立连接出现错误...,错误信息%v", err)
		return err
	}
	log.Info("建立连接成功！")
	go cli.sendIDENTIFY()
	// loop to send ping
	go cli.startHeartbeat()
	defer cli.Conn.Close()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, message, err := cli.Conn.ReadMessage()
			if err != nil {
				log.Errorf("websocket发生错误:%v,将在3s后重试...", err)
				time.Sleep(3 * time.Second)
				for i := 1; i < 5; i++ {
					log.Infof("尝试重连....,第%v次", i)
					cli.Conn, _, err = websocket.DefaultDialer.DialContext(ctx, url, nil)
					if err != nil {
						internal := 3 * i
						log.Errorf("第%v次重连失败，将在%vs后重试", i, internal)
						// for j := 0; j < i; j++ {
						time.Sleep(time.Duration(internal) * time.Second)
						// }
					}
					// 重连后需要重新认证
					cli.sendIDENTIFY()
				}
				if err != nil {
					log.Errorf("重连失败...,%v", err)
					return err
				}
				continue
			}
			// recover from panic
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("处理回调函数发生错误...%v", err)
				}
			}()
			// 处理消息
			err = callback(message)
			if err != nil {
				log.Errorf("处理消息出现错误...%v", err)
			}
		}
	}
}
