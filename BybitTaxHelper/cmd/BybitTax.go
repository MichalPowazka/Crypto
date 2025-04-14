package main

import (
	csvreader "BybitTaxHelper/internal/app/CsvReader"
	plnspent "BybitTaxHelper/internal/app/PlnSpent"
)

func main() {
	data := csvreader.ReadExcelData()
	plnspent.GetPlnSpent(data)
}
