package kafka

import (
	"consume-bluebird/worker/consume/models"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

type DeleteTopicsRequest struct {
	Version int16
	Topics  []string
	Timeout time.Duration
}

func GetConsumerConfig() *consumergroup.Config {
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = 10 * time.Second
	return config
}

func ConnectConsumer(cfgKafka models.ConfigKafka) (*consumergroup.ConsumerGroup, error) {
	config := cfgKafka.ConsumerConfig
	consumer, consumerErr := consumergroup.JoinConsumerGroup(cfgKafka.ConsumerGroup, cfgKafka.KafkaTopics, cfgKafka.Zookeeper, config)
	if consumerErr != nil {
		return nil, consumerErr
	}
	return consumer, nil
}

