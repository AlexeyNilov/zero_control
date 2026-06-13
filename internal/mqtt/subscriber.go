package mqtt

import (
	"context"
	"fmt"
	"log"
	"strings"
)

const NotifyTopic = "zero-control/notify"

const notifyQoS byte = 1

type Client interface {
	Subscribe(context.Context, string, byte, func([]byte)) error
}

type Notifier interface {
	NotifyMainChat(context.Context, string) error
}

type Subscriber struct {
	client   Client
	notifier Notifier
	logger   *log.Logger
}

func NewSubscriber(client Client, notifier Notifier, logger *log.Logger) *Subscriber {
	return &Subscriber{
		client:   client,
		notifier: notifier,
		logger:   logger,
	}
}

func (s *Subscriber) Start(ctx context.Context) error {
	err := s.client.Subscribe(ctx, NotifyTopic, notifyQoS, func(payload []byte) {
		s.forwardNotification(ctx, payload)
	})
	if err != nil {
		return fmt.Errorf("subscribe to mqtt notifications: %w", err)
	}

	return nil
}

func (s *Subscriber) forwardNotification(ctx context.Context, payload []byte) {
	text := strings.TrimSpace(string(payload))
	if text == "" {
		s.logger.Printf("ignored blank mqtt notification")
		return
	}

	s.logger.Printf("received mqtt notification bytes=%d", len(payload))

	if err := s.notifier.NotifyMainChat(ctx, text); err != nil {
		s.logger.Printf("mqtt notification delivery error: %v", err)
	}
}
