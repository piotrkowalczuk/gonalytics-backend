package kafka

// ActionsConsumer ...
type ActionsConsumer struct {
	BaseConsumer
}

// NewActionsConsumer ...
func NewActionsConsumer(url, clientName, consumerGroupName, topicName string) (*ActionsConsumer, error) {
	baseConsumer, err := NewBaseConsumer(url, clientName, consumerGroupName, topicName)

	if err != nil {
		return nil, err
	}

	actionsConsumer := &ActionsConsumer{
		BaseConsumer: *baseConsumer,
	}

	return actionsConsumer, nil
}
