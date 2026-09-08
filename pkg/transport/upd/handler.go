package udp

type Handler interface {
	HandleMsg(data []byte)
}

func HandleFunc(call func(data []byte)) Handler {
	return handleFunc(call)
}

type handleFunc func(data []byte)

func (h handleFunc) HandleMsg(data []byte) {
	h(data)
}
