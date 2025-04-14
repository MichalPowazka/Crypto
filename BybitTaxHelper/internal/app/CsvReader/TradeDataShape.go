package csvreader

import (
	"reflect"
	"time"
)

type TradeDataShape struct {
	Uid            string    `csv:"Uid"`
	SpotPairs      string    `csv:"Spot Pairs"`
	OrderType      string    `csv:"Order Type"`
	Direction      string    `csv:"Direction"`
	FilledValue    float64   `csv:"Filled Value"`
	FilledPrice    float64   `csv:"Filled Price"`
	FilledQuantity float64   `csv:"Filled Quantity"`
	Fees           float64   `csv:"Fees"`
	TransactionID  string    `csv:"Transaction ID"`
	OrderNo        string    `csv:"Order No."`
	Timestamp      time.Time `csv:"Timestamp (UTC+0)"`
}

func (v TradeDataShape) getStructFieldNames() []string {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var fieldNames []string
	for i := 0; i < val.NumField(); i++ {
		fieldNames = append(fieldNames, val.Type().Field(i).Name)
	}
	return fieldNames
}
