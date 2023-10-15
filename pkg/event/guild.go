package event

type SatoriGuildEvent interface {
	EventInterface
	ListenGuildAdded(callback EventHandlerCallback)
	ListenGuildUpdated(callback EventHandlerCallback)
	ListenGuildRemoved(callback EventHandlerCallback)
	ListenGuildRequest(callback EventHandlerCallback)
}

type satoriGuildEventImpl struct {
	eventImpl
}

func NewSatoriGuildEvent() (SatoriGuildEvent, error) {
	return &satoriGuildEventImpl{}, nil
}

func (impl *satoriGuildEventImpl) ListenGuildAdded(callback EventHandlerCallback) {
	impl.addHandlers("guild-added", callback)
}

func (impl *satoriGuildEventImpl) ListenGuildUpdated(callback EventHandlerCallback) {
	impl.addHandlers("guild-updated", callback)
}

func (impl *satoriGuildEventImpl) ListenGuildRemoved(callback EventHandlerCallback) {
	impl.addHandlers("guild-removed", callback)
}

func (impl *satoriGuildEventImpl) ListenGuildRequest(callback EventHandlerCallback) {
	impl.addHandlers("guild-request", callback)
}
