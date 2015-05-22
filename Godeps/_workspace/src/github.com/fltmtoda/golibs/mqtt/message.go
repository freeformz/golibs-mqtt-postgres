package mqtt

type Message interface {
	Duplicate() bool
	Qos() byte
	Retained() bool
	Topic() string
	MessageID() uint16
	Payload() []byte
}
