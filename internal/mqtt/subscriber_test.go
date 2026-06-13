package mqtt

import (
	"bytes"
	"context"
	"errors"
	"log"
	"strings"
	"testing"
)

func TestSubscriberSubscribesToNotifyTopicAndForwardsPayloadToMainChat(t *testing.T) {
	client := &stubClient{}
	notifier := &stubNotifier{}
	subscriber := NewSubscriber(client, notifier, log.New(&bytes.Buffer{}, "", 0))

	err := subscriber.Start(context.Background())
	if err != nil {
		t.Fatalf("Start returned error: %v", err)
	}

	if client.topic != NotifyTopic {
		t.Fatalf("topic = %q, want %q", client.topic, NotifyTopic)
	}

	if client.qos != 1 {
		t.Fatalf("qos = %d, want %d", client.qos, 1)
	}

	client.handler([]byte("garage door opened"))

	if notifier.text != "garage door opened" {
		t.Fatalf("text = %q, want %q", notifier.text, "garage door opened")
	}
}

func TestSubscriberLogsReceivedNotificationWithoutPayloadText(t *testing.T) {
	var logs bytes.Buffer
	client := &stubClient{}
	notifier := &stubNotifier{}
	subscriber := NewSubscriber(client, notifier, log.New(&logs, "", 0))

	err := subscriber.Start(context.Background())
	if err != nil {
		t.Fatalf("Start returned error: %v", err)
	}

	client.handler([]byte("garage door opened"))

	logLine := logs.String()
	if !strings.Contains(logLine, "received mqtt notification bytes=18") {
		t.Fatalf("logs = %q, want brief mqtt receipt log", logLine)
	}

	if strings.Contains(logLine, "garage door opened") {
		t.Fatalf("logs = %q, want payload text omitted", logLine)
	}
}

func TestSubscriberReturnsErrorWhenTopicSubscriptionFails(t *testing.T) {
	client := &stubClient{subscribeErr: errors.New("subscribe failed")}
	notifier := &stubNotifier{}
	subscriber := NewSubscriber(client, notifier, log.New(&bytes.Buffer{}, "", 0))

	err := subscriber.Start(context.Background())
	if err == nil {
		t.Fatal("Start returned nil error, want subscription error")
	}

	if notifier.text != "" {
		t.Fatalf("text = %q, want empty notification", notifier.text)
	}
}

func TestSubscriberIgnoresBlankPayloads(t *testing.T) {
	var logs bytes.Buffer
	client := &stubClient{}
	notifier := &stubNotifier{}
	subscriber := NewSubscriber(client, notifier, log.New(&logs, "", 0))

	err := subscriber.Start(context.Background())
	if err != nil {
		t.Fatalf("Start returned error: %v", err)
	}

	client.handler([]byte(" \n\t "))

	if notifier.text != "" {
		t.Fatalf("text = %q, want empty notification", notifier.text)
	}

	if !strings.Contains(logs.String(), "ignored blank mqtt notification") {
		t.Fatalf("logs = %q, want blank notification log", logs.String())
	}
}

type stubClient struct {
	topic        string
	qos          byte
	handler      func([]byte)
	subscribeErr error
}

type stubNotifier struct {
	text string
}

func (s *stubClient) Subscribe(_ context.Context, topic string, qos byte, handler func([]byte)) error {
	s.topic = topic
	s.qos = qos
	s.handler = handler
	return s.subscribeErr
}

func (s *stubNotifier) NotifyMainChat(_ context.Context, text string) error {
	s.text = text
	return nil
}
