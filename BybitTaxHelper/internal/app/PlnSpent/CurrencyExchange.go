package plnspent

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Struktury zgodne z JSON-em zwracanym przez API
type ExchangeRatesSeries struct {
	Table    string `json:"table"`
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Rates    []Rate `json:"rates"`
}

type Rate struct {
	No            string  `json:"no"`
	EffectiveDate string  `json:"effectiveDate"`
	Mid           float64 `json:"mid"`
}

func getEuroToPlnRate(date time.Time) (float64, string) {
	date = adjustToLastBusinessDay(date)
	res, formattedDate := getRateFromApi(date)

	rate := getRateFromResponse(res)
	return rate, formattedDate
}

// NBP provides rate only for business days
func adjustToLastBusinessDay(date time.Time) time.Time {
	switch date.Weekday() {
	case time.Saturday:
		date = date.AddDate(0, 0, -1)
	case time.Sunday:
		date = date.AddDate(0, 0, -2)
	}

	return date
}

func getRateFromApi(date time.Time) (*http.Response, string) {
	formattedDate := date.Format("2006-01-02")
	url := fmt.Sprintf("https://api.nbp.pl/api/exchangerates/rates/a/eur/%s/?format=json", formattedDate)
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Set("User-Agent", "plnspent-client/1.0")
	res, _ := client.Do(req)

	return res, formattedDate
}

func getRateFromResponse(res *http.Response) float64 {
	body, _ := io.ReadAll(res.Body)

	var rates ExchangeRatesSeries
	if err := json.Unmarshal(body, &rates); err != nil {
		log.Fatal("Coudnt Unmarshal response")
	}
	if len(rates.Rates) == 0 {
		log.Fatal("Coudnt get rate from response")
	}

	return rates.Rates[0].Mid
}
