package mqtt

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client client
type Client interface {
	Connect() error
	Disconnect()

	IsConnected() bool

	Subscribe(topic string, handler Handler) mqtt.Token
	Unsubscribe(topic string) mqtt.Token
	Publish(topic string, msg []byte) mqtt.Token
}

type defaultClient struct {
	id       string
	host     string
	user     string
	password string

	client      mqtt.Client
	isConnected bool

	sync.Mutex

	s sync.RWMutex

	subscriptions map[string]mqtt.MessageHandler
}

// NewClient NewClient
func NewClient(id string, host, user, password string) Client {
	d := &defaultClient{
		id:       id,
		host:     host,
		user:     user,
		password: password,

		subscriptions: make(map[string]mqtt.MessageHandler),
	}

	return d
}

func (c *defaultClient) connectionLostHandler(client mqtt.Client, err error) {
	log.Println("[MQTT] Conexão perdida:")

	// c.s.Lock()
	// c.isConnected = false
	// c.s.Unlock()
}

func (c *defaultClient) onConnectHandler(client mqtt.Client) {
	log.Println("[MQTT] Conectado:")

	// c.s.Lock()
	// c.isConnected = true
	// c.s.Unlock()

	c.Lock()
	defer c.Unlock()

	for topic, handler := range c.subscriptions {
		log.Println("mqtt: re-subscribing to topic:", topic)
		c.subscribe(topic, handler)
	}
}

func (c *defaultClient) connect(clientID string, uri *url.URL) mqtt.Client {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))

	opts.SetClientID(clientID)
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)

	opts.SetKeepAlive(5 * time.Second)
	opts.SetPingTimeout(5 * time.Second)
	opts.SetConnectTimeout(10 * time.Second)
	opts.SetWriteTimeout(10 * time.Second)
	// opts.SetMaxReconnectInterval(5 * time.Second)

	opts.SetConnectionLostHandler(c.connectionLostHandler)
	opts.SetOnConnectHandler(c.onConnectHandler)

	opts.SetUsername(c.user)
	opts.SetPassword(c.password)

	c.client = mqtt.NewClient(opts)

	c.connectLoop()

	return c.client
}

// Connect Connect
func (c *defaultClient) Connect() error {
	host := fmt.Sprintf("tcp://%s:%s", c.host, "1883")
	uri, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}

	c.client = c.connect(c.id, uri)

	return nil
}

func (c *defaultClient) connectLoop() {
	for {
		if token := c.client.Connect(); token.Wait() && token.Error() != nil {
			log.Println("[MQTT] tentando reconectar-se:", token.Error())
			time.Sleep(1000 * time.Millisecond)

			return
		}

		c.Lock()
		c.isConnected = true
		c.Unlock()

		break
	}
}

// Disconnect Disconnect
func (c *defaultClient) Disconnect() {
	log.Println("[MQTT] Fechando conexão:")

	c.Lock()
	defer c.Unlock()

	if c.isConnected {
		c.client.Disconnect(250)
	}
	log.Println("[MQTT] Conexão fechada!")
}

// IsConnected IsConnected
func (c *defaultClient) IsConnected() bool {
	c.Lock()
	defer c.Unlock()

	return c.isConnected
}

func (c *defaultClient) publish(topic string, msg []byte) mqtt.Token {
	return c.client.Publish(topic, 0, false, msg)
}

func (c *defaultClient) subscribe(topic string, handler mqtt.MessageHandler) mqtt.Token {
	c.Lock()
	defer c.Unlock()

	c.subscriptions[topic] = handler
	return c.client.Subscribe(topic, 0, handler)
}

func (c *defaultClient) unsubscribe(topic string) mqtt.Token {
	c.Lock()
	defer c.Unlock()

	delete(c.subscriptions, topic)
	return c.client.Unsubscribe(topic)
}
