package event

import (
	"github.com/satori-protocol-go/satori-model-go/pkg/channel"
	"github.com/satori-protocol-go/satori-model-go/pkg/guild"
	"github.com/satori-protocol-go/satori-model-go/pkg/guildmember"
	"github.com/satori-protocol-go/satori-model-go/pkg/guildrole"
	"github.com/satori-protocol-go/satori-model-go/pkg/login"
	"github.com/satori-protocol-go/satori-model-go/pkg/message"
	"github.com/satori-protocol-go/satori-model-go/pkg/user"
)

type SignInfo struct {
	Op   uint64      `json:"op"`
	Body interface{} `json:"body,omitempty"`
}

type Event struct {
	Id        int64                    `json:"id"`                 //事件 ID
	Type      string                   `json:"type"`               //事件类型
	Platform  string                   `json:"platform"`           //接收者的平台名称
	SelfId    string                   `json:"self_id"`            //接收者的平台账号
	Timestamp int64                    `json:"timestamp"`          //事件的时间戳
	Channel   *channel.Channel         `json:"channel,omitempty"`  //事件所属的频道
	Guild     *guild.Guild             `json:"guild,omitempty"`    //事件所属的群组
	Login     *login.Login             `json:"login,omitempty"`    //事件的登录信息
	Member    *guildmember.GuildMember `json:"member,omitempty"`   //事件的目标成员
	Message   *message.Message         `json:"message,omitempty"`  //事件的消息
	Operator  *user.User               `json:"operator,omitempty"` //事件的操作者
	Role      *guildrole.GuildRole     `json:"role,omitempty"`     //事件的目标角色
	User      *user.User               `json:"user,omitempty"`     //事件的目标用户
}

type EventHandlerCallback func(e Event) error

type EventHandler struct {
	EventType    string
	EventHandler EventHandlerCallback
}

type EventInterface interface {
	eventHandlers() []EventHandler
}

type eventImpl struct {
	handlers []EventHandler
}

func (event *eventImpl) eventHandlers() []EventHandler {
	return event.handlers
}

func (event *eventImpl) addHandlers(eventType string, handler EventHandlerCallback) {
	event.handlers = append(event.handlers, EventHandler{
		EventType:    eventType,
		EventHandler: handler,
	})
}
