package kafka

import "github.com/Shopify/sarama"

// BaseConsumer ...
type BaseConsumer struct {
	broker            *sarama.Broker
	startingOffset    int64
	clientName        string
	consumerGroupName string
	topicName         string

	MetaData         *sarama.MetadataResponse
	AvailableOffsets *sarama.OffsetResponse
}

// NewBaseConsumer ...
func NewBaseConsumer(clientName, consumerGroupName, topicName string) (*BaseConsumer, error) {
	baseConsumer := &BaseConsumer{
		clientName:        clientName,
		consumerGroupName: consumerGroupName,
		topicName:         topicName,
	}

	err := baseConsumer.broker.Open(nil)
	if err != nil {
		return nil, err
	}

	return baseConsumer, nil
}

// InitBroker ...
func (bc *BaseConsumer) InitBroker(brokerHost string) {
	bc.broker = sarama.NewBroker(brokerHost)
}

// Destroy ...
func (bc *BaseConsumer) Destroy() {
	bc.broker.Close()
}

// FetchMetaData ...
func (bc *BaseConsumer) FetchMetaData() error {
	request := sarama.MetadataRequest{Topics: []string{}}
	response, err := bc.broker.GetMetadata(bc.clientName, &request)

	bc.MetaData = response
	return err
}

// FetchOffsets ...
func (bc *BaseConsumer) FetchOffsets() error {
	offsetsRequest := sarama.OffsetRequest{}
	offsetsRequest.AddBlock(bc.topicName, 0, sarama.EarliestOffset, 255)
	offsetsResponse, err := bc.broker.GetAvailableOffsets(bc.clientName, &offsetsRequest)
	if err != nil {
		return err
	}

	if offsetsResponse.Blocks[bc.topicName][0].Err != sarama.NoError {
		return offsetsResponse.Blocks[bc.topicName][0].Err
	}

	return nil
}

// FetchOffset ...
func (bc *BaseConsumer) FetchOffset() error {
	offsetFetchRequest := sarama.OffsetFetchRequest{ConsumerGroup: bc.consumerGroupName}
	offsetFetchRequest.AddPartition(bc.topicName, 0)
	offsetFetchResponse, err := bc.broker.FetchOffset(bc.clientName, &offsetFetchRequest)
	if err != nil {
		return err
	}

	if offsetFetchResponse.Blocks[bc.topicName][0].Err != sarama.NoError {
		return err
	}

	bc.startingOffset = offsetFetchResponse.Blocks[bc.topicName][0].Offset + 1

	return nil
}
