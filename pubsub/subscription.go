package pubsub

import (
	"time"
)

type Subscription struct {
	Service                       string
	Name                          string
	Topic                         string
	Endpoint                      string
	AckDeadline                   time.Duration
	RetainAckedMessages           bool
	RetentionDuration             time.Duration
	EnableOrdering                bool
	ExpirationTTL                 time.Duration
	Filter                        string
	DeliverExactlyOnce            bool
	DeadLetterTopic               string
	MaxDeadLetterDeliveryAttempts int
	RetryMinimumBackoff           *time.Duration
	RetryMaximumBackoff           *time.Duration
}

func (s *Subscription) GetSubscriptionID() string {
	return s.Name
}
