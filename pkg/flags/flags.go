package flags

import (
	"flag"
	"log"
)

type ParsedFlags struct {
	StationsPath string
	OWMApiKey    string
	Kafka        string
}

func Parse() ParsedFlags {
	var p ParsedFlags
	flag.StringVar(&p.OWMApiKey, "apikey", "", "open weather map API Key (Get one at https://openweathermap.org/appid)")
	flag.StringVar(&p.StationsPath, "stations", "stations.txt", "path to stations file (one station id per line)")
	flag.StringVar(&p.Kafka, "kafka", "localhost:9092", "kafka adresses, comma seperated")
	flag.Parse()
	validate(p)
	return p
}

func validate(p ParsedFlags) {
	if p.OWMApiKey == "" {
		log.Fatal("Please provide a")
	}
}
