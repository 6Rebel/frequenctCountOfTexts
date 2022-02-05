package helper

import (
	"FrequenctCountOfTexts/pojo"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"sort"
)

func EncodeJSONBody(resp http.ResponseWriter, statusCode int, data interface{}) {
	resp.WriteHeader(statusCode)
	err := json.NewEncoder(resp).Encode(data)
	if err != nil {
		logrus.Errorf("Error encoding response %v", err)
	}
}

// SortMapByValue sorts the word frequency map based on their frequencies
func SortMapByValue(wordFrequencyMap map[string]int) []pojo.WordFrequency {

	var wordsFrequencies []pojo.WordFrequency
	for key, value := range wordFrequencyMap {
		wordsFrequencies = append(wordsFrequencies, pojo.WordFrequency{Key: key, Value: value})
	}

	sort.Slice(wordsFrequencies, func(i, j int) bool {
		return wordsFrequencies[i].Value > wordsFrequencies[j].Value
	})

	if len(wordsFrequencies) > 10 {
		return wordsFrequencies[:10]
	}

	return wordsFrequencies
}
