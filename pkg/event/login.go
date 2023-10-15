package event

type SatoriLoginEvent interface {
	EventInterface
	ListenLoginAdded(callback EventHandlerCallback)
	ListenLoginUpdated(callback EventHandlerCallback)
	ListenLoginRemoved(callback EventHandlerCallback)
}

type satoriLoginEventImpl struct {
	eventImpl
}

func NewSatoriLoginEvent() (SatoriLoginEvent, error) {
	return &satoriLoginEventImpl{}, nil
}

func (impl *satoriLoginEventImpl) ListenLoginAdded(callback EventHandlerCallback) {
	impl.addHandlers("login-added", callback)
}

func (impl *satoriLoginEventImpl) ListenLoginUpdated(callback EventHandlerCallback) {
	impl.addHandlers("login-updated", callback)
}

func (impl *satoriLoginEventImpl) ListenLoginRemoved(callback EventHandlerCallback) {
	impl.addHandlers("login-removed", callback)
}
