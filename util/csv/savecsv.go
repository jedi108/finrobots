package csv

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
)
//
//var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}


func SaveCsv(filename string, datas [][]string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		file, err = os.Create(filename)
		checkError("Cannot create file", err)
	} else {

	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range datas {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
	fmt.Println("save ok")
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

