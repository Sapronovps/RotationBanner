package kafka

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type Producer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer(brokersList, topic string, retryMax int) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = retryMax

	// Адреса брокеров Kafka
	brokers := []string{brokersList}

	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		return nil, err
	}

	return &Producer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *Producer) SendMessage(data, eventType string) error {
	_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
		Headers: []sarama.RecordHeader{
			{Key: []byte("type"), Value: []byte(eventType)},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) SendCustomMessage(err error, message, eventType string, logger *zap.Logger) {
	errSend := error(nil)
	if err != nil {
		errSend = p.SendMessage("Ошибка создания слота: "+err.Error(), "addSlot")
	} else {
		errSend = p.SendMessage(message, eventType)
	}

	if errSend != nil {
		logger.Error("Ошибка отправки события в kafka" + errSend.Error())
	}
}
