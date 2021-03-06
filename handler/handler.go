package handler

import (
	"FrequenctCountOfTexts/helper"
	"FrequenctCountOfTexts/pojo"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func TopTenMostUsedWords(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request recieved to fetch Top Ten Most Used Words")

	var textInputRequest pojo.TextInput

	err := json.NewDecoder(r.Body).Decode(&textInputRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	textArray := strings.Split(textInputRequest.Text, " ")

	wordOccurrenceMap := make(map[string]int)

	for _, text := range textArray {
		wordOccurrenceMap[text] += 1
	}

	mostOccurrenceWordsMap := helper.SortMapByValue(wordOccurrenceMap)

	helper.EncodeJSONBody(w, http.StatusOK, mostOccurrenceWordsMap)
}
