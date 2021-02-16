package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/thamthee/exp-kafka/configs"
	"github.com/thamthee/exp-kafka/foundation/kafka"
)

func main() {
	log := log.New(os.Stdout, "main:", log.LstdFlags | log.Lshortfile)

	if err := run(log); err != nil {
		log.Fatalln("server err: ", err)
	}
}

func run(log *log.Logger) error {
	// ====================================================
	// Configuration
	file, err := os.Open("./configs/config.yml")
	if err != nil {
		return errors.Wrap(err, "read file config")
	}

	conf, err := configs.New(file, "yml")
	if err != nil {
		return errors.Wrap(err, "parse from file")
	}

	r, err := kafka.Reader(conf.Kafka.Brokers, conf.Kafka.ClientID, "demo")
	if err != nil {
		return errors.Wrap(err, "opening kafka reader")
	}

	defer r.Close()

	log.Println("start consuming ...")
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, m.Value)
	}
}
