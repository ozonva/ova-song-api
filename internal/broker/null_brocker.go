package broker

func NewNullBroker() Broker {
	return &nullBroker{}
}

type nullBroker struct{}

func (n *nullBroker) SendEvent(_ SongEvent) error { return nil }

func (n *nullBroker) Close() error { return nil }
