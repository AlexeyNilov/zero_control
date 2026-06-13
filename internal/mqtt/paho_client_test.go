package mqtt

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

func TestPahoReconnectHandlerResubscribesAfterInitialConnect(t *testing.T) {
	var logs bytes.Buffer
	client := NewPahoClient("tcp://localhost:1883", "zero-control-bot", log.New(&logs, "", 0))
	pahoClient := &stubPahoClient{}
	subscription := mqttSubscription{
		topic:   NotifyTopic,
		qos:     notifyQoS,
		handler: func(paho.Client, paho.Message) {},
	}

	reconnect := client.reconnectHandler(subscription)
	reconnect(pahoClient)

	if pahoClient.subscribeCount != 0 {
		t.Fatalf("subscribeCount = %d, want no subscription on initial connect", pahoClient.subscribeCount)
	}

	reconnect(pahoClient)

	if pahoClient.subscribeCount != 1 {
		t.Fatalf("subscribeCount = %d, want one subscription after reconnect", pahoClient.subscribeCount)
	}

	if pahoClient.topic != NotifyTopic {
		t.Fatalf("topic = %q, want %q", pahoClient.topic, NotifyTopic)
	}

	if pahoClient.qos != notifyQoS {
		t.Fatalf("qos = %d, want %d", pahoClient.qos, notifyQoS)
	}

	if !strings.Contains(logs.String(), "mqtt resubscribed topic=zero-control/notify") {
		t.Fatalf("logs = %q, want resubscribe log", logs.String())
	}
}

type stubPahoClient struct {
	subscribeCount int
	topic          string
	qos            byte
	handler        paho.MessageHandler
}

func (s *stubPahoClient) IsConnected() bool {
	return true
}

func (s *stubPahoClient) IsConnectionOpen() bool {
	return true
}

func (s *stubPahoClient) Connect() paho.Token {
	return completedToken{}
}

func (s *stubPahoClient) Disconnect(uint) {}

func (s *stubPahoClient) Publish(string, byte, bool, interface{}) paho.Token {
	return completedToken{}
}

func (s *stubPahoClient) Subscribe(topic string, qos byte, handler paho.MessageHandler) paho.Token {
	s.subscribeCount++
	s.topic = topic
	s.qos = qos
	s.handler = handler
	return completedToken{}
}

func (s *stubPahoClient) SubscribeMultiple(map[string]byte, paho.MessageHandler) paho.Token {
	return completedToken{}
}

func (s *stubPahoClient) Unsubscribe(...string) paho.Token {
	return completedToken{}
}

func (s *stubPahoClient) AddRoute(string, paho.MessageHandler) {}

func (s *stubPahoClient) OptionsReader() paho.ClientOptionsReader {
	return paho.ClientOptionsReader{}
}

type completedToken struct{}

func (completedToken) Wait() bool {
	return true
}

func (completedToken) WaitTimeout(time.Duration) bool {
	return true
}

func (completedToken) Done() <-chan struct{} {
	done := make(chan struct{})
	close(done)
	return done
}

func (completedToken) Error() error {
	return nil
}
