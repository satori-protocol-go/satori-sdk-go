package event

type SatoriChannelEvent interface {
	EventInterface
}

type satoriChannelEventImpl struct {
	eventImpl
}

func NewSatoriChannelEvent() (SatoriChannelEvent, error) {
	return &satoriChannelEventImpl{}, nil
}
