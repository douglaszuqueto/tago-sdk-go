package pubsub

import (
	"context"
	"log"
	"sync"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/core/mqtt"
)

// PubSub PubSub
type PubSub interface {
	Sub() (<-chan *mqtt.Payload, error)
	Debug() (<-chan *mqtt.Payload, error)

	Pub(msg []byte) error

	UnsubscribeData() error
	UnsubscribeDebug() error

	Close()
}

type ps struct {
	token  string
	ctx    context.Context
	cancel context.CancelFunc
	mqtt   struct {
		sync.RWMutex
		client mqtt.Client
		ctx    context.Context
		cancel context.CancelFunc
	}

	sync.RWMutex

	dataCh  chan *mqtt.Payload
	debugCh chan *mqtt.Payload
}

var (
	dataTopic  = "edge/data"
	debugTopic = "tago/debug"

	dataChLen  = 1
	debugChLen = 1
)

// New New
func New(token string) PubSub {
	c := &ps{
		token: token,
	}

	if err := c.connectMQTT(); err != nil {
		panic(err)
	}

	return c
}

func (d *ps) Pub(msg []byte) error {
	token := d.mqtt.client.Publish("tago/data/post", msg)

	if token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (d *ps) Sub() (<-chan *mqtt.Payload, error) {
	d.Lock()
	defer d.Unlock()

	if d.dataCh != nil {
		return d.dataCh, nil
	}

	d.dataCh = make(chan *mqtt.Payload, dataChLen)

	token := d.mqtt.client.Subscribe(dataTopic, func(_ mqtt.Client, msg mqtt.Payload) {
		d.RLock()
		defer d.RUnlock()

		d.dataCh <- &msg
	})

	token.Wait()
	err := token.Error()
	if err != nil {
		close(d.dataCh)
		d.dataCh = nil
	}

	return d.dataCh, err
}

func (d *ps) Debug() (<-chan *mqtt.Payload, error) {
	d.Lock()
	defer d.Unlock()

	if d.debugCh != nil {
		return d.debugCh, nil
	}

	d.debugCh = make(chan *mqtt.Payload, debugChLen)

	token := d.mqtt.client.Subscribe(debugTopic, func(_ mqtt.Client, msg mqtt.Payload) {
		d.RLock()
		defer d.RUnlock()

		d.debugCh <- &msg
	})

	token.Wait()
	err := token.Error()
	if err != nil {
		close(d.debugCh)
		d.debugCh = nil
	}

	return d.debugCh, err
}

func (d *ps) UnsubscribeData() error {
	token := d.mqtt.client.Unsubscribe(dataTopic)

	if token.Error() != nil {
		return token.Error()
	}

	d.Lock()
	defer d.Unlock()

	if d.dataCh != nil {
		close(d.dataCh)
		d.dataCh = nil
	}

	return nil
}

func (d *ps) UnsubscribeDebug() error {
	token := d.mqtt.client.Unsubscribe(debugTopic)

	if token.Error() != nil {
		return token.Error()
	}

	d.Lock()
	defer d.Unlock()

	if d.debugCh != nil {
		close(d.debugCh)
		d.debugCh = nil
	}

	return nil
}

func (d *ps) Close() {
	log.Println("Fechando conexões!")

	d.closeMQTT()

	d.Lock()
	defer d.Unlock()

	if d.debugCh != nil {
		close(d.debugCh)
		d.debugCh = nil
	}

	if d.dataCh != nil {
		close(d.dataCh)
		d.dataCh = nil
	}

	log.Println("Conexões fechadas")
}
