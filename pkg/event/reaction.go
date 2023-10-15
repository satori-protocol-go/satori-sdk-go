package event

type SatoriReactionEvent interface {
	EventInterface
	ListenReactionAdded(callback EventHandlerCallback)
	ListenReactionRemoved(callback EventHandlerCallback)
}

type satoriReactionEventImpl struct {
	eventImpl
}

func NewSatoriReactionEvent() (SatoriReactionEvent, error) {
	return &satoriReactionEventImpl{}, nil
}

func (impl *satoriReactionEventImpl) ListenReactionAdded(callback EventHandlerCallback) {
	impl.addHandlers("reaction-added", callback)
}

func (impl *satoriReactionEventImpl) ListenReactionRemoved(callback EventHandlerCallback) {
	impl.addHandlers("reaction-deleted", callback)
}
