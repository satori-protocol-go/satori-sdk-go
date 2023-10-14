package client

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"strings"
// 	"time"

// 	"github.com/dezhishen/satori-sdk-go/pkg/config"
// 	"github.com/gorilla/websocket"
// 	log "github.com/sirupsen/logrus"
// )

// type WebsocketApiChannel struct {
// 	Ctx             context.Context
// 	ready           chan bool
// 	Conn            *websocket.Conn
// 	endpoint        string
// 	accessToken     string
// 	resultCallbacks map[string]map[string]func(data []byte) error
// 	timeout         time.Duration
// }

// type actionReq struct {
// 	Action string      `json:"action"`
// 	Params interface{} `json:"params"`
// 	Echo   string      `json:"echo"`
// }

// type actionResult struct {
// 	Data    interface{} `json:"data"`
// 	Status  string      `json:"status"`
// 	Retcode int64       `json:"retcode"`
// 	Msg     string      `json:"msg"`
// 	Wording string      `json:"wording"`
// 	Echo    string      `json:"echo"`
// }

// func NewWebsocketApiTemplate(conf config.SatoriApiConfig) (ApiTemplate, error) {
// 	if strings.HasSuffix("/", conf.Endpoint) {
// 		conf.Endpoint = strings.TrimSuffix(conf.Endpoint, "/")
// 	}
// 	if conf.Timeout == 0 {
// 		// 默认10s
// 		conf.Timeout = DEFAULT_TIMEOUT
// 	}
// 	result := &WebsocketApiChannel{
// 		endpoint:        conf.Endpoint,
// 		ready:           make(chan bool),
// 		Ctx:             context.Background(),
// 		timeout:         time.Duration(conf.Timeout) * time.Millisecond,
// 		accessToken:     conf.AccessToken,
// 		resultCallbacks: make(map[string]map[string]func(data []byte) error),
// 	}
// 	go func() {
// 		err := result.listenApiResult()
// 		if err != nil {
// 			log.Errorf("listen api result error: %s", err)
// 			panic(err)
// 		}
// 	}()
// 	if <-result.ready {
// 		return result, nil
// 	} else {
// 		return nil, fmt.Errorf("websocket api channel init failed")
// 	}
// }

// func (cli *WebsocketApiChannel) Post(action string) error {
// 	echo := fmt.Sprintf("%s/%d", action, time.Now().UnixNano())
// 	reqData := actionReq{
// 		Action: action,
// 		Echo:   echo,
// 	}
// 	err := cli.Conn.WriteJSON(reqData)
// 	if err != nil {
// 		cli.Conn.Close()
// 		cli.Conn = nil
// 	}
// 	return err
// }

// func (cli *WebsocketApiChannel) PostForResult(action string, result interface{}) error {
// 	reqId := fmt.Sprintf("%d", time.Now().UnixNano())
// 	echo := fmt.Sprintf("%s/%s", action, reqId)
// 	reqData := actionReq{
// 		Action: action,
// 		Echo:   echo,
// 	}
// 	// 注册回调
// 	channel := make(chan byte)
// 	cli.registerCallbackWithChannel(action, reqId, channel, &result)
// 	defer func() {
// 		cli.unregisterCallback(action, reqId)
// 		close(channel)
// 	}()
// 	err := cli.Conn.WriteJSON(reqData)
// 	if err != nil {
// 		cli.Conn.Close()
// 		cli.Conn = nil
// 	}
// 	// 等待回调
// 	select {
// 	case <-channel:
// 		return nil
// 	case <-time.After(cli.timeout):
// 		return fmt.Errorf("request timeout")
// 	}
// }

// func (cli *WebsocketApiChannel) PostByRequest(action string, params interface{}) error {
// 	reqData := actionReq{
// 		Action: action,
// 		Params: params,
// 	}
// 	err := cli.Conn.WriteJSON(reqData)
// 	if err != nil {
// 		cli.Conn.Close()
// 		cli.Conn = nil
// 	}
// 	return err
// }

// func (cli *WebsocketApiChannel) PostByRequestForResult(action string, params interface{}, result interface{}) error {
// 	reqId := fmt.Sprintf("%d", time.Now().UnixNano())
// 	echo := fmt.Sprintf("%s/%s", action, reqId)
// 	reqData := actionReq{
// 		Action: action,
// 		Params: params,
// 		Echo:   echo,
// 	}
// 	// 注册回调
// 	channel := make(chan byte)
// 	cli.registerCallbackWithChannel(action, reqId, channel, &result)
// 	defer func() {
// 		cli.unregisterCallback(action, reqId)
// 		close(channel)
// 	}()
// 	err := cli.Conn.WriteJSON(reqData)
// 	if err != nil {
// 		cli.Conn.Close()
// 		cli.Conn = nil
// 	}
// 	select {
// 	case <-channel:
// 		return nil
// 	case <-time.After(cli.timeout):
// 		return fmt.Errorf("request timeout")
// 	}
// }

// func (cli *WebsocketApiChannel) listenApiResult() error {
// 	url := fmt.Sprintf("%s/api", cli.endpoint)
// 	if cli.accessToken != "" {
// 		url = fmt.Sprintf("%s?access_token=%s", url, cli.accessToken)
// 	}
// 	// url := cli.endpoint + "?access_token=" + cli.accessToken
// 	log.Infof("开始建立连接...地址:%v", url)
// 	var err error
// 	cli.Conn, _, err = websocket.DefaultDialer.DialContext(cli.Ctx, url, nil)
// 	if err != nil {
// 		log.Infof("建立连接出现错误...,错误信息%v", err)
// 		return err
// 	}
// 	log.Info("建立连接成功！")
// 	cli.ready <- true
// 	defer cli.Conn.Close()
// 	for {
// 		select {
// 		case <-cli.Ctx.Done():
// 			return nil
// 		default:
// 			_, message, err := cli.Conn.ReadMessage()
// 			if err != nil {
// 				log.Errorf("websocket发生错误:%v,将在3s后重试...", err)
// 				time.Sleep(3 * time.Second)
// 				for i := 1; i < 5; i++ {
// 					log.Infof("尝试重连....,第%v次", i)
// 					cli.Conn, _, err = websocket.DefaultDialer.DialContext(cli.Ctx, url, nil)
// 					if err != nil {
// 						internal := 3 * i
// 						log.Errorf("第%v次重连失败，将在%vs后重试", i, internal)
// 						// for j := 0; j < i; j++ {
// 						time.Sleep(time.Duration(internal) * time.Second)
// 						// }
// 					}
// 				}
// 				if err != nil {
// 					log.Errorf("重连失败...,%v", err)
// 					return err
// 				}
// 				continue
// 			}
// 			// recover from panic
// 			defer func() {
// 				if err := recover(); err != nil {
// 					log.Errorf("处理回调函数发生错误...%v", err)
// 				}
// 			}()
// 			var result actionResult
// 			err = json.Unmarshal(message, &result)
// 			if err != nil {
// 				log.Errorf("解析回调数据出现错误...%v", err)
// 				continue
// 			}
// 			cli.handlerCallback(result.Echo, message)
// 			// 将消息结果放入通道
// 		}
// 	}
// }

// func (cli *WebsocketApiChannel) Listen(ctx context.Context, callback func(message []byte) error) error {
// 	url := fmt.Sprintf("%s/event", cli.endpoint)
// 	if cli.accessToken != "" {
// 		url = fmt.Sprintf("%s?access_token=%s", url, cli.accessToken)
// 	}
// 	// url := cli.endpoint + "?access_token=" + cli.accessToken
// 	log.Infof("开始建立连接...地址:%v", url)
// 	c, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
// 	if err != nil {
// 		log.Infof("建立连接出现错误...,错误信息%v", err)
// 		return err
// 	}
// 	log.Info("建立连接成功！")
// 	defer c.Close()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil
// 		default:
// 			_, message, err := c.ReadMessage()
// 			if err != nil {
// 				log.Errorf("websocket发生错误:%v,将在3s后重试...", err)
// 				time.Sleep(3 * time.Second)
// 				for i := 1; i < 5; i++ {
// 					log.Infof("尝试重连....,第%v次", i)
// 					c, _, err = websocket.DefaultDialer.DialContext(ctx, url, nil)
// 					if err != nil {
// 						internal := 3 * i
// 						log.Errorf("第%v次重连失败，将在%vs后重试", i, internal)
// 						// for j := 0; j < i; j++ {
// 						time.Sleep(time.Duration(internal) * time.Second)
// 						// }
// 					}
// 				}
// 				if err != nil {
// 					log.Errorf("重连失败...,%v", err)
// 					return err
// 				}
// 				continue
// 			}
// 			// recover from panic
// 			defer func() {
// 				if err := recover(); err != nil {
// 					log.Errorf("处理回调函数发生错误...%v", err)
// 				}
// 			}()
// 			callback(message)
// 		}
// 	}
// }

// func (cli *WebsocketApiChannel) registerCallbackWithChannel(
// 	action string,
// 	reqId string,
// 	channel chan byte,
// 	result interface{},
// ) {
// 	if _, ok := cli.resultCallbacks[action]; !ok {
// 		cli.resultCallbacks[action] = make(map[string]func(data []byte) error)
// 	}
// 	cli.resultCallbacks[action][reqId] =
// 		func(data []byte) error {
// 			json.Unmarshal(data, &result)
// 			channel <- 1
// 			return nil
// 		}
// }

// func (cli *WebsocketApiChannel) unregisterCallback(
// 	action string,
// 	reqId string,
// ) {
// 	if _, ok := cli.resultCallbacks[action]; ok {
// 		delete(cli.resultCallbacks[action], reqId)
// 	}
// }

// func (cli *WebsocketApiChannel) handlerCallback(
// 	echo string,
// 	message []byte,
// ) {
// 	if echo == "" {
// 		return
// 	}
// 	action := strings.Split(echo, "/")[0]
// 	reqId := strings.Split(echo, "/")[1]
// 	if _, ok := cli.resultCallbacks[action]; ok {
// 		if _, ok := cli.resultCallbacks[action][reqId]; ok {
// 			cli.resultCallbacks[action][reqId](message)
// 		}
// 	}
// }
