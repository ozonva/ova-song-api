package broker

const (
	typeCreate      = "Create"
	typeCreateMulti = "CreateMulti"
	typeUpdate      = "Update"
	typeRemove      = "Remove"
)

type SongEvent struct {
	Type string
	Id   int64
}

func NewCreateEvent(id int64) SongEvent {
	return SongEvent{typeCreate, id}
}

func NewCreateMultiEvent(id int64) SongEvent {
	return SongEvent{typeCreateMulti, id}
}

func NewUpdateEvent(id int64) SongEvent {
	return SongEvent{typeUpdate, id}
}

func NewRemoveEvent(id int64) SongEvent {
	return SongEvent{typeRemove, id}
}
