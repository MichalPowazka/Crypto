package csvreader

import (
	"log"
	"os"
)

const dataFilePath = "../assets/data.csv"

func ReadExcelData() []TradeDataShape {
	file := readFile()
	defer file.Close()

	array := dataToArray(file)
	data := arrayToStruct(&array)

	return data
}

func readFile() *os.File {
	file, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal("Error while reading the file ", dataFilePath, " error: ", err)
	}
	return file
}
