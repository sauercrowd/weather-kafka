package weather

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// GetStationsFromFile parses everything into an int array
func GetStationsFromFile(path string) ([]int, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(dat)
	lines := strings.Split(content, "\n")

	IDs := make([]int, 0)
	for _, line := range lines {
		log.Println("Line: ", line)
		if len(line) == 0 {
			continue
		}
		id, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		IDs = append(IDs, id)
	}
	return IDs, nil
}
