package csvreader

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func dataToArray(file *os.File) [][]string {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 11
	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error while creating data array from the file; ", err)
	}

	return data
}

func arrayToStruct(data *[][]string) []TradeDataShape {
	validateHeaderRow(data)
	trades := []TradeDataShape{}
	for _, row := range *data {
		trade := TradeDataShape{
			Uid:            row[0],
			SpotPairs:      row[1],
			OrderType:      row[2],
			Direction:      row[3],
			FilledValue:    parseFloat(row[4]),
			FilledPrice:    parseFloat(row[5]),
			FilledQuantity: parseFloat(row[6]),
			Fees:           parseFloat(row[7]),
			TransactionID:  row[8],
			OrderNo:        row[9],
			Timestamp:      parseTime(row[10]),
		}
		trades = append(trades, trade)
	}
	return trades
}

func validateHeaderRow(data *[][]string) {
	expectedHeaders := TradeDataShape{}.getStructFieldNames()

	for index := range expectedHeaders {
		expectedHeader := expectedHeaders[index]
		accualHeader := sanitizeHeader((*data)[0][index])

		if !strings.Contains(accualHeader, expectedHeader) {
			log.Fatal("Invalid headers row \nexpected header: ", expectedHeader, " \naccual header: ", accualHeader)
		}
	}
	*data = (*data)[1:]
}

func sanitizeHeader(value string) string {
	value = removeSpace(value)
	value = removePunctuation(value)
	return value
}
