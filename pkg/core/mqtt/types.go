package mqtt

// Message Message
type Message struct {
	Value []byte
}

func (m *Message) String() string {
	return string(m.Value)
}

// Payload Payload
type Payload struct {
	Topic   string
	Message Message
}
