package mqtt

import (
	"context"
	"fmt"
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type PahoClient struct {
	brokerURL string
	clientID  string
	logger    *log.Logger
}

func NewPahoClient(brokerURL, clientID string, logger *log.Logger) *PahoClient {
	return &PahoClient{
		brokerURL: brokerURL,
		clientID:  clientID,
		logger:    logger,
	}
}

func (c *PahoClient) Subscribe(ctx context.Context, topic string, qos byte, handler func([]byte)) error {
	client := paho.NewClient(c.options())
	if err := waitFor(client.Connect()); err != nil {
		return fmt.Errorf("connect to mqtt broker: %w", err)
	}

	if err := waitFor(client.Subscribe(topic, qos, func(_ paho.Client, message paho.Message) {
		handler(message.Payload())
	})); err != nil {
		client.Disconnect(250)
		return fmt.Errorf("subscribe to topic %s: %w", topic, err)
	}

	c.logger.Printf("mqtt subscribed topic=%s", topic)
	go disconnectWhenDone(ctx, client)
	return nil
}

func (c *PahoClient) options() *paho.ClientOptions {
	return paho.NewClientOptions().
		AddBroker(c.brokerURL).
		SetClientID(c.clientID).
		SetAutoReconnect(true).
		SetConnectTimeout(10 * time.Second).
		SetConnectionLostHandler(func(_ paho.Client, err error) {
			c.logger.Printf("mqtt connection lost: %v", err)
		})
}

func waitFor(token paho.Token) error {
	token.Wait()
	return token.Error()
}

func disconnectWhenDone(ctx context.Context, client paho.Client) {
	<-ctx.Done()
	client.Disconnect(250)
}
