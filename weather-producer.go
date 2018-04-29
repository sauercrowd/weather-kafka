package main

import (
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sauercrowd/weather-kafka/pkg/flags"
	"github.com/sauercrowd/weather-kafka/pkg/persistence"
	"github.com/sauercrowd/weather-kafka/pkg/weather"
)

const requestsPerMinute = 60

func main() {
	parsedFlags := flags.Parse()

	IDs, err := weather.GetStationsFromFile(parsedFlags.StationsPath)
	if err != nil {
		log.Println(err)
		return
	}

	producer, err := sarama.NewSyncProducer(strings.Split(parsedFlags.Kafka, ","), nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	waitTimeSeconds := int(float64(len(IDs))/float64(requestsPerMinute)*60 + 1)
	for {
		for _, id := range IDs {
			log.Printf("ID [%d]", id)
			wEntry, err := weather.GetWeather(parsedFlags.OWMApiKey, id)
			if err != nil {
				log.Println(err)
				return
			}
			if err := persistence.Add(&producer, wEntry); err != nil {
				log.Println("Could not add entry: ", err)
			}
		}
		time.Sleep(time.Second * time.Duration(waitTimeSeconds))
	}

}
