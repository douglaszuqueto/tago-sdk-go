package pubsub

import (
	"log"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/core/mqtt"
)

func (c *ps) connectMQTT() (err error) {
	c.mqtt.Lock()
	defer c.mqtt.Unlock()
	if c.mqtt.client != nil {
		return nil
	}

	c.mqtt.client = mqtt.NewClient(
		"iotformakers-sdk",
		"mqtt.tago.io",
		"admin",
		c.token,
	)

	// c.mqtt.ctx, c.mqtt.cancel = context.WithCancel(context.Background())
	if err := c.mqtt.client.Connect(); err != nil {
		log.Println("tago-sdk: Could not connect to MQTT", err)
		return err
	}
	log.Println("tago-sdk: Connected to MQTT")
	return nil
}

func (c *ps) closeMQTT() error {
	c.mqtt.Lock()
	defer c.mqtt.Unlock()
	if c.mqtt.client == nil {
		return nil
	}
	log.Println("tago-sdk: Disconnecting from MQTT...")
	c.mqtt.cancel()
	c.mqtt.client.Disconnect()
	c.mqtt.client = nil
	return nil
}
