package event

type SatoriGuildRoleEvent interface {
	EventInterface
	ListenGuildRoleCreated(callback EventHandlerCallback)
	ListenGuildRoleUpdated(callback EventHandlerCallback)
	ListenGuildRoleDeleted(callback EventHandlerCallback)
}

type satoriGuildRoleEventImpl struct {
	eventImpl
}

func NewSatoriGuildRoleEvent() (SatoriGuildRoleEvent, error) {
	return &satoriGuildRoleEventImpl{}, nil
}

func (impl *satoriGuildRoleEventImpl) ListenGuildRoleCreated(callback EventHandlerCallback) {
	impl.addHandlers("guild-role-created", callback)
}

func (impl *satoriGuildRoleEventImpl) ListenGuildRoleUpdated(callback EventHandlerCallback) {
	impl.addHandlers("guild-role-updated", callback)
}

func (impl *satoriGuildRoleEventImpl) ListenGuildRoleDeleted(callback EventHandlerCallback) {
	impl.addHandlers("guild-role-deleted", callback)
}
