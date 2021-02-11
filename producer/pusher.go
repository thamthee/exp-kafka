package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Pusher struct {
	log    *log.Logger
	writer *kafka.Writer
}

func New(log *log.Logger, w *kafka.Writer) Pusher {
	return Pusher{
		log:    log,
		writer: w,
	}
}

func (p Pusher) SendMsg(ctx context.Context, key, value []byte, now time.Time) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  now.UTC(),
	}

	return p.writer.WriteMessages(ctx, msg)
}
