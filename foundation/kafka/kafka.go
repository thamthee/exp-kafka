package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

func Writer(brokers []string, clientID, topic string) (*kafka.Writer, error) {
	dialer := &kafka.Dialer{
		ClientID: clientID,
		Timeout:  5 * time.Second,
	}

	config := kafka.WriterConfig{
		Brokers:          brokers,
		Topic:            topic,
		Dialer:           dialer,
		Balancer:         &kafka.LeastBytes{},
		WriteTimeout:     2 * time.Second,
		ReadTimeout:      2 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	return kafka.NewWriter(config), nil
}

func Reader(brokers []string, clientID, topic string) (*kafka.Reader, error) {
	dialer := &kafka.Dialer{
		ClientID: clientID,
		Timeout:  5 * time.Second,
	}

	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         clientID,
		Topic:           topic,
		Dialer:          dialer,
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         1 * time.Second,
		ReadLagInterval: -1,
	}

	return kafka.NewReader(config), nil
}
