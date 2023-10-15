package event

type SatoriMessageEvent interface {
	EventInterface
	ListenMessageCreated(callback EventHandlerCallback)
	ListenMessageUpdated(callback EventHandlerCallback)
	ListenMessageDelete(callback EventHandlerCallback)
}

type satoriMessageEventImpl struct {
	eventImpl
}

func NewSatoriMessageEvent() (SatoriMessageEvent, error) {
	return &satoriMessageEventImpl{}, nil
}

func (impl *satoriMessageEventImpl) ListenMessageCreated(callback EventHandlerCallback) {
	impl.addHandlers("message-created", callback)
}

func (impl *satoriMessageEventImpl) ListenMessageUpdated(callback EventHandlerCallback) {
	impl.addHandlers("message-updated", callback)
}

func (impl *satoriMessageEventImpl) ListenMessageDelete(callback EventHandlerCallback) {
	impl.addHandlers("message-deleted", callback)
}
