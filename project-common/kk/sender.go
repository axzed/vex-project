package kk

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

// LogData 日志数据
type LogData struct {
	// kafka Topic
	Topic string
	//json数据
	Data []byte
}

// KafkaWriter kafka writer
type KafkaWriter struct {
	w    *kafka.Writer
	data chan LogData
}

// GetWriter 获取kafka writer
func GetWriter(addr string) *KafkaWriter {
	w := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Balancer: &kafka.LeastBytes{},
	}
	k := &KafkaWriter{
		w:    w,
		data: make(chan LogData, 100),
	}
	go k.sendKafka()
	return k
}

// Send 发送数据
func (w *KafkaWriter) Send(data LogData) {
	// 将数据放入writer缓冲
	w.data <- data
}

// Close 关闭
func (w *KafkaWriter) Close() {
	if w.w != nil {
		w.w.Close()
	}
}

// sendKafka 发送到kafka
func (w *KafkaWriter) sendKafka() {
	for {
		select {
		case data := <-w.data:
			messages := []kafka.Message{
				{
					Topic: data.Topic,
					Key:   []byte("logMsg"),
					Value: data.Data,
				},
			}
			var err error
			const retries = 3
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			// retry sending the message 3 times
			for i := 0; i < retries; i++ {
				// attempt to create topic prior to publishing the message
				err = w.w.WriteMessages(ctx, messages...)
				if err == nil {
					break
				}
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					continue
				}
				if err != nil {
					log.Printf("kafka send writemessage err %s \n", err.Error())
				}
			}
		}
	}

}
