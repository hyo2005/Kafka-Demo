package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// tao context moi
	ctx := context.Background()
	//produce message trong go routine
	go produce(ctx)
	consume(ctx)
}

const (
	topic          = "message-log"
	broker1Address = "localhost:9093"
	broker2Address = "localhost:9094"
	broker3Address = "localhost:9095"
)

func produce(ctx context.Context) {
	// Khoi tao counter
	i := 0

	// Khoi tao writer voi cac brokers va topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
	})

	for {
		// Kafka messgage bao gom key va value. Key se quyet dinh
		// xem message se duoc publish vao partitionn nao, broker nao
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// in ra console key cua message
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// in log ra console da send dc message so bao nhieu
		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {
	// khoi tao reader voi cac broker va topic
	// group id dung de phan biet cac consumer trong
	// consumer group, tranh viec nhan trung message
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
		GroupID: "my-group",
	})
	for {
		// method ReadMessage method se block cho den khi nhan duoc event tiep theo
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// in log ra console sau khi nhan duoc message
		fmt.Println("received: ", string(msg.Value))
	}
}
