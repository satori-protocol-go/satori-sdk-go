package event

type SatoriUserEvent interface {
	EventInterface
	ListenFriendRequest(callback EventHandlerCallback)
}

type satoriUserEventImpl struct {
	eventImpl
}

func NewSatoriUserEvent() (SatoriUserEvent, error) {
	return &satoriUserEventImpl{}, nil
}

func (impl *satoriUserEventImpl) ListenFriendRequest(callback EventHandlerCallback) {
	impl.addHandlers("friend-request", callback)
}
