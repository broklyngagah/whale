package rabbit

type Message struct {
	exchange  string
	key       string
	mandatory bool
	immediate bool
	data      []byte
}

func NewMessage(exchange, key string, mandatory, immediate bool, data []byte) *Message {
	return &Message{
		exchange:  exchange,
		key:       key,
		mandatory: mandatory,
		immediate: immediate,
		data:      data,
	}
}
