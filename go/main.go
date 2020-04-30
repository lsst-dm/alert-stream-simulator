package main

import (
	"bytes"
	"context"
	"flag"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/lsst-dm/alert-stream-simulator/schema-avro"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/gzip"
	"github.com/segmentio/kafka-go/lz4"
	"github.com/segmentio/kafka-go/snappy"
	"github.com/segmentio/kafka-go/zstd"
)

func main() {
	broker := flag.String("broker", "localhost:9092", "bootstrap Kafka broker to connect with")
	topic := flag.String("topic", "alertsd", "topic to write to")
	file := flag.String("file", "", "file to publish")
	workers := flag.Int("workers", 8, "number of worker goroutines")
	compression := flag.String("compression", "null", "compression codec to use (null | lz4 | gzip | snappy | zstd)")
	flag.Parse()
	var codec kafka.CompressionCodec
	switch *compression {
	case "null":
		codec = nil
	case "lz4":
		codec = lz4.NewCompressionCodec()
	case "gzip":
		codec = gzip.NewCompressionCodec()
	case "snappy":
		codec = snappy.NewCompressionCodec()
	case "zstd":
		codec = zstd.NewCompressionCodec()
	}
	err := run(*broker, *topic, *file, *workers, codec)
	if err != nil {
		log.Fatal(err)
	}
}

func run(brokerURL, topic, filePath string, workers int, codec kafka.CompressionCodec) error {
	ch := make(chan *schema.Alert, 100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		err := readFile(ctx, filePath, ch)
		if err != nil {
			panic(err)
		}
	}()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          []string{brokerURL},
		Topic:            topic,
		BatchSize:        1,
		BatchBytes:       10_000_000,
		CompressionCodec: codec,
	})
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		go func() {
			publishAlerts(ctx, ch, w)
			wg.Done()
		}()
		wg.Add(1)
	}
	start := time.Now()
	wg.Wait()
	end := time.Now()
	log.Printf("done in %v (%.2f/s)", end.Sub(start), 1000.0/(end.Sub(start).Seconds()))
	return nil
}

func publishAlerts(ctx context.Context, ch chan *schema.Alert, w *kafka.Writer) {
	buf := bytes.NewBuffer(nil)
	for alert := range ch {
		err := alert.Serialize(buf)
		if err != nil {
			panic(err)
		}
		msg := kafka.Message{
			Value: buf.Bytes(),
		}
		err = w.WriteMessages(ctx, msg)
		if err != nil {
			panic(err)
		}
		buf.Reset()
	}
}

func readFile(ctx context.Context, fp string, dst chan *schema.Alert) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	reader, err := schema.NewAlertReader(f)
	if err != nil {
		return err
	}

	defer close(dst)
	for {
		alert, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		select {
		case <-ctx.Done():
			return nil
		case dst <- alert:
		}
	}
}
