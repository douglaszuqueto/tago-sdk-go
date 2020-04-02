package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Handler Handler
type Handler func(client Client, msg Payload)

// Subscribe Subscribe
func (c *defaultClient) Subscribe(topic string, handler Handler) mqtt.Token {
	return c.subscribe(topic, func(mqtt mqtt.Client, msg mqtt.Message) {
		payload := &Payload{
			Topic: msg.Topic(),
			Message: Message{
				Value: msg.Payload(),
			},
		}

		handler(c, *payload)
	})
}

// Publish Publish
func (c *defaultClient) Publish(topic string, msg []byte) mqtt.Token {
	return c.publish(topic, msg)
}

// Unsubscribe Unsubscribe
func (c *defaultClient) Unsubscribe(topic string) mqtt.Token {
	return c.unsubscribe(topic)
}
