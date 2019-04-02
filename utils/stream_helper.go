package utils

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
)

//StreamHelper is an interface that allows interacting with streams
//go:generate counterfeiter -o ../fakes/fake_stream_helper.go . StreamHelper
type StreamHelper interface {
	StartStreamPullPush(context.Context, resources.DataStream, chan string) (err error)
}

type restToKafkaHelper struct {
	logger *log.Logger
	ctx    context.Context
}

//NewStreamHelper creates a stream helper
func NewStreamHelper(c context.Context, l *log.Logger) StreamHelper {
	return &restToKafkaHelper{
		logger: l,
		ctx:    c,
	}
}

//StartStreamPullPush starts a go routine to fetch data from source and send it to destination
func (r *restToKafkaHelper) StartStreamPullPush(ctx context.Context, stream resources.DataStream, c chan string) (err error) {
	r.logger.Printf("entering-start-stream-pull-push")
	defer r.logger.Printf("exiting-start-stream-pull-push")
	var resp *http.Response
	r.logger.Printf("received-stream-%#v", stream)
	stop := false
	go func() {
		<-c
		stop = true
	}()
	go func() {
		r.logger.Printf("entered-go-routine")
		resp, err = http.Get(stream.DataStreamSourceURL)
		if err != nil {
			r.logger.Printf("failed-getting-stream")
			err = fmt.Errorf("failed-getting-stream")
			return
		}
		r.logger.Printf("succeeded-getting-stream")
		reader := bufio.NewReader(resp.Body)
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = 5
		config.Producer.Return.Successes = true
		brokers := []string{stream.DataStreamDestinationURL}
		producer, err := sarama.NewSyncProducer(brokers, config)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := producer.Close(); err != nil {
				panic(err)
			}
		}()
		for !stop {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				r.logger.Printf("failed-reading-stream-%#v", err)
			}
			msg := &sarama.ProducerMessage{
				Topic: stream.DataStreamTopic,
				Value: sarama.StringEncoder(line),
			}
			_, _, err = producer.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}()
	return err
}
