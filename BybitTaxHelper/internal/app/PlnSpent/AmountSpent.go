package plnspent

import (
	csvreader "BybitTaxHelper/internal/app/CsvReader"
	"fmt"
	"log"
	"strings"
	"time"
)

const buyOperation = "buy"

// const sellOperation = "sell"

func GetPlnSpent(data []csvreader.TradeDataShape) {
	eurBuyOperations := getEurBuyOperations(&data)
	eurSpendings := calculateOperations(&eurBuyOperations)
	fmt.Println("sum: ", eurSpendings)
}

func getEurBuyOperations(data *[]csvreader.TradeDataShape) []csvreader.TradeDataShape {
	operations := []csvreader.TradeDataShape{}
	for _, single := range *data {
		if single.Direction == buyOperation && strings.Contains(single.SpotPairs, "EUR") {
			operations = append(operations, single)
		}
	}
	return operations
}

// only calculate buys in the future if needed change * -1 rate modifier
func calculateOperations(data *[]csvreader.TradeDataShape) float64 {

	var sum float64 = 0.0
	for _, operation := range *data {
		//Should calculate for previous day then exchange happend
		date := operation.Timestamp.AddDate(0, 0, -1)
		rate, rateDate := getEuroToPlnRate(date)
		sum += rate * operation.FilledValue * -1
		log.Println("Initial date: ", operation.Timestamp, " Final date: ", rateDate, " Rate: ", rate, " FilledValue: ", operation.FilledValue)
		time.Sleep(time.Second)
	}

	return sum
}
