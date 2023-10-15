package event

type SatoriGuildMemberEvent interface {
	EventInterface
	ListenGuildMemberAdded(callback EventHandlerCallback)
	ListenGuildMemberUpdated(callback EventHandlerCallback)
	ListenGuildMemberRemoved(callback EventHandlerCallback)
	ListenGuildMemberRequest(callback EventHandlerCallback)
}

type satoriGuildMemberEventImpl struct {
	eventImpl
}

func NewSatoriGuildMemberEvent() (SatoriGuildMemberEvent, error) {
	return &satoriGuildMemberEventImpl{}, nil
}

func (impl *satoriGuildMemberEventImpl) ListenGuildMemberAdded(callback EventHandlerCallback) {
	impl.addHandlers("guild-member-added", callback)
}

func (impl *satoriGuildMemberEventImpl) ListenGuildMemberUpdated(callback EventHandlerCallback) {
	impl.addHandlers("guild-member-updated", callback)
}

func (impl *satoriGuildMemberEventImpl) ListenGuildMemberRemoved(callback EventHandlerCallback) {
	impl.addHandlers("guild-member-removed", callback)
}

func (impl *satoriGuildMemberEventImpl) ListenGuildMemberRequest(callback EventHandlerCallback) {
	impl.addHandlers("guild-member-request", callback)
}
