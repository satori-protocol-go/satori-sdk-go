package event

import (
	"github.com/dezhishen/satori-sdk-go/pkg/resource/channel"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guild"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guildmember"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/guildrole"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/login"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/message"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/user"
)

type Event struct {
	Id        int64                    //事件 ID
	Type      string                   //事件类型
	Platform  string                   //接收者的平台名称
	SelfId    string                   //接收者的平台账号
	Timestamp int64                    //事件的时间戳
	Channel   *channel.Channel         //事件所属的频道
	Guild     *guild.Guild             //事件所属的群组
	Login     *login.Login             //事件的登录信息
	Member    *guildmember.GuildMember //事件的目标成员
	Message   *message.Message         //事件的消息
	Operator  *user.User               //事件的操作者
	Role      *guildrole.GuildRole     //事件的目标角色
	User      *user.User               //事件的目标用户
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
