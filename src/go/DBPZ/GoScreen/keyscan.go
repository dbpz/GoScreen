package GoScreen

const (
	EVENT_KEY_DOWN = iota
	EVENT_KEY_UP
	EVENT_KEY_INPUT
)

type keyEvent{
	event_type int
	event_key  int
	is_repeated bool
}

type mouseEvent{
	event_type int
	event_key int
	coordinateX, coordinateY int
}

// a keying device can be a keyboard scanner or a simple tty stdin receiver.
interface keyingDevice{
	receiveNextKeyEvent() keyEvent
}

interface mousingDevice{
	receiveNextMouseEvent() mouseEvent
}
