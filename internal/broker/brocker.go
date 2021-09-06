package broker

import (
	"time"

	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type Broker interface {
	SendEvent(event SongEvent) error
	Close() error
}

type broker struct {
	producer sarama.SyncProducer
	topic    string
}

func (b *broker) SendEvent(event SongEvent) error {
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}
	log.Debug().Bytes("Send event", bytes).Msg("Event marshalled")

	msg := &sarama.ProducerMessage{
		Topic:     b.topic,
		Key:       sarama.StringEncoder(b.topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, err = b.producer.SendMessage(msg)
	return err
}

func (b *broker) Close() error {
	return b.producer.Close()
}

func NewKafkaBroker(brokers []string, topic string) (Broker, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &broker{
		producer: producer,
		topic:    topic,
	}, nil
}
