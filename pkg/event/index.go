package event

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/dezhishen/satori-sdk-go/pkg/client"
	"github.com/dezhishen/satori-sdk-go/pkg/config"
	"github.com/dezhishen/satori-sdk-go/pkg/constant"
	"github.com/dezhishen/satori-sdk-go/pkg/resource/login"
	"github.com/tidwall/gjson"
)

type SatoriEvent interface {
	SatoriChannelEvent
	SatoriGuildEvent
	SatoriGuildMemberEvent
	SatoriGuildRoleEvent
	SatoriLoginEvent
	SatoriMessageEvent
	SatoriUserEvent
	StartWithCtx(ctx context.Context) error
}

type SatoriEventImpl struct {
	eventImpl
	SatoriChannelEvent
	SatoriGuildEvent
	SatoriGuildMemberEvent
	SatoriGuildRoleEvent
	SatoriLoginEvent
	SatoriMessageEvent
	SatoriUserEvent
	template   client.EventTemplate
	allHandler map[string][]EventHandlerCallback
}

func (s *SatoriEventImpl) StartWithCtx(ctx context.Context) error {
	return s.template.StartListen(ctx, func(message []byte) error {
		op := gjson.GetBytes(message, "op").Uint()
		switch constant.SignNum(op) {
		case constant.SIGN_NUM_EVENT:
			var e Event
			body := gjson.GetBytes(message, "body").Raw
			err := json.Unmarshal([]byte(body), &e)
			if err != nil {
				return fmt.Errorf("handle listen decoder err :%v,raw:%v", err, message)
			}
			log.Infof("recive EVENT: %s", e.Type)
			s.template.SetSequence(e.Id)
			handler, ok := s.allHandler[e.Type]
			if ok {
				for _, callback := range handler {
					go runCallbakc(e, callback)
				}
			}
			return nil
		case constant.SIGN_NUM_PONG:
			log.Infof("recive PONG")
			return nil
		case constant.SIGN_NUM_READY:
			var logins []login.Login
			_ = json.Unmarshal([]byte(gjson.GetBytes(message, "body").Raw), &logins)
			log.Infof("recive READY:%v", logins)
			return nil
		default:
			return fmt.Errorf("unsupport signNum %d", op)
		}
	})
}

func runCallbakc(e Event, callback EventHandlerCallback) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("处理回调函数发生错误...%v", err)
		}
	}()
	callback(e)
}

// eventHandlers implements SatoriEvent.
func (s *SatoriEventImpl) eventHandlers() []EventHandler {
	return s.handlers
}

func (impl *SatoriEventImpl) addHandlers(handlers []EventHandler) {
	if len(handlers) == 0 {
		return
	}
	for _, handler := range handlers {
		impl.handlers = append(impl.handlers, handler)
		if len(impl.allHandler[handler.EventType]) == 0 {
			impl.allHandler[handler.EventType] = []EventHandlerCallback{handler.EventHandler}
		} else {
			impl.allHandler[handler.EventType] = append(impl.allHandler[handler.EventType], handler.EventHandler)
		}
	}
}

func NewSatorEventByConfig(conf config.SatoriConfig) (SatoriEvent, error) {
	template, err := client.NewEventTemplate(conf.Event)
	if err != nil {
		return nil, err
	}
	result := &SatoriEventImpl{
		template:   template,
		allHandler: make(map[string][]EventHandlerCallback),
	}
	// channel
	channelEvent, err := NewSatoriChannelEvent()
	if err != nil {
		return nil, err
	}
	result.addHandlers(channelEvent.eventHandlers())
	result.SatoriChannelEvent = channelEvent
	// guild
	guildEvent, err := NewSatoriGuildEvent()
	if err != nil {
		return nil, err
	}
	result.addHandlers(guildEvent.eventHandlers())
	result.SatoriGuildEvent = guildEvent
	// guild member
	guildMemberEvent, err := NewSatoriGuildMemberEvent()
	if err != nil {
		return nil, err
	}
	result.addHandlers(guildMemberEvent.eventHandlers())
	result.SatoriGuildMemberEvent = guildMemberEvent
	// guild role
	guildRoleEvent, err := NewSatoriGuildRoleEvent()
	if err != nil {
		return nil, err
	}
	result.addHandlers(guildRoleEvent.eventHandlers())
	result.SatoriGuildRoleEvent = guildRoleEvent
	// login
	loginEvent, err := NewSatoriLoginEvent()
	if err != nil {
		return nil, err
	}
	result.addHandlers(loginEvent.eventHandlers())
	result.SatoriLoginEvent = loginEvent
	// message
	messageEvent, err := NewSatoriMessageEvent()
	if err != nil {
		return nil, err
	}
	result.addHandlers(messageEvent.eventHandlers())
	result.SatoriMessageEvent = messageEvent
	// user
	userEvent, err := NewSatoriUserEvent()
	if err != nil {
		return nil, err
	}
	result.addHandlers(userEvent.eventHandlers())
	result.SatoriUserEvent = userEvent
	return result, nil
}
