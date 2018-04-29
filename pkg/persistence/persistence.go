package persistence

import (
	"log"

	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/sauercrowd/weather-kafka/pkg/weather"
)

//Add an entry
func Add(producer *sarama.SyncProducer, we *weather.WeatherEntry) error {
	// convert to json
	bytes, err := json.Marshal(we)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{Topic: "weather", Value: sarama.StringEncoder(string(bytes))}
	partition, offset, err := (*producer).SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
	return nil
}
