package csvreader

import (
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func parseTime(value string) time.Time {
	layout := "2006-01-02 15:04:05"
	parsed, err := time.Parse(layout, strings.TrimSpace(value))
	if err != nil {
		log.Fatal("Error parsing time: ", err)
	}
	return parsed
}

func parseFloat(value string) float64 {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal("Error parsing float : ", err)
	}
	return parsed
}

func removeSpace(value string) string {
	return strings.ReplaceAll(value, " ", "")
}

func removePunctuation(value string) string {
	var result []rune
	for _, r := range value {
		if !unicode.IsPunct(r) {
			result = append(result, r)
		}
	}
	return string(result)
}
