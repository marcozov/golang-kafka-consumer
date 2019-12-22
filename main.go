package main

import (
	"fmt"
	"context"
	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("test import")
	r := kafka.NewReader(kafka.ReaderConfig{
	    Brokers:   []string{"localhost:9092"},
	    //Topic:     "topic-A",
	    GroupID:   "go-consumer",
	    Topic:     "restaurant-service-command-channel",
	    //Partition: 1,
	    //MinBytes:  10e3, // 10KB
	    //MaxBytes:  10e6, // 10MB
	})
	//r.SetOffset(42)

	for {
	    m, err := r.ReadMessage(context.Background())
	    if err != nil {
		break
	    }
	    fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}
