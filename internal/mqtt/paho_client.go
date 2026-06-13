package mqtt

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type PahoClient struct {
	brokerURL string
	clientID  string
	logger    *log.Logger
}

type mqttSubscription struct {
	topic   string
	qos     byte
	handler paho.MessageHandler
}

func NewPahoClient(brokerURL, clientID string, logger *log.Logger) *PahoClient {
	return &PahoClient{
		brokerURL: brokerURL,
		clientID:  clientID,
		logger:    logger,
	}
}

func (c *PahoClient) Subscribe(ctx context.Context, topic string, qos byte, handler func([]byte)) error {
	subscription := mqttSubscription{
		topic: topic,
		qos:   qos,
		handler: func(_ paho.Client, message paho.Message) {
			handler(message.Payload())
		},
	}

	client := paho.NewClient(c.options(subscription))
	if err := waitFor(client.Connect()); err != nil {
		return fmt.Errorf("connect to mqtt broker: %w", err)
	}

	if err := subscribe(client, subscription); err != nil {
		client.Disconnect(250)
		return fmt.Errorf("subscribe to topic %s: %w", topic, err)
	}

	c.logger.Printf("mqtt subscribed topic=%s", topic)
	go disconnectWhenDone(ctx, client)
	return nil
}

func (c *PahoClient) options(subscription mqttSubscription) *paho.ClientOptions {
	return paho.NewClientOptions().
		AddBroker(c.brokerURL).
		SetClientID(c.clientID).
		SetAutoReconnect(true).
		SetConnectTimeout(10 * time.Second).
		SetOnConnectHandler(c.reconnectHandler(subscription)).
		SetConnectionLostHandler(func(_ paho.Client, err error) {
			c.logger.Printf("mqtt connection lost: %v", err)
		})
}

func (c *PahoClient) reconnectHandler(subscription mqttSubscription) paho.OnConnectHandler {
	var mu sync.Mutex
	initialConnectSeen := false

	return func(client paho.Client) {
		mu.Lock()
		if !initialConnectSeen {
			initialConnectSeen = true
			mu.Unlock()
			return
		}
		mu.Unlock()

		if err := subscribe(client, subscription); err != nil {
			c.logger.Printf("mqtt resubscribe failed topic=%s error=%v", subscription.topic, err)
			return
		}
		c.logger.Printf("mqtt resubscribed topic=%s", subscription.topic)
	}
}

func subscribe(client paho.Client, subscription mqttSubscription) error {
	return waitFor(client.Subscribe(subscription.topic, subscription.qos, subscription.handler))
}

func waitFor(token paho.Token) error {
	token.Wait()
	return token.Error()
}

func disconnectWhenDone(ctx context.Context, client paho.Client) {
	<-ctx.Done()
	client.Disconnect(250)
}
