package event

import (
	"github.com/dezhishen/satori-sdk-go/pkg/client"
	"github.com/dezhishen/satori-sdk-go/pkg/config"
)

type SatoriEvent interface {
	SatoriChannelEvent
	SatoriGuildEvent
	SatoriGuildMemberEvent
	SatoriGuildRoleEvent
	SatoriLoginEvent
	SatoriMessageEvent
	SatoriUserEvent
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

func NewSatorApiByConfig(conf config.SatoriConfig) (SatoriEvent, error) {
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
