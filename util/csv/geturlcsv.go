package csv

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"io"
	"finrgo/utilcsv"
	"os"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	utf8reader := DecodeWindows1251reader(resp.Body)
	reader := csv.NewReader(utf8reader)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecodeWindows1251reader(ba io.Reader) io.Reader {
	dec := charmap.Windows1251.NewDecoder()
	return dec.Reader(ba)
}

func DecodeWindows1251(ba []byte) string {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.Bytes(ba)
	return string(out)
}

//http://www.moex.com/ru/derivatives/contracts.aspx

/*call
Полный код контракта:	Si-12.17M211217CA60000
Код в торговой системе:	Si60000BL7
Дата начала обращения:	23.11.2016
Последний день обращения:	21.12.2017
Дата исполнения:	21.12.2017
http://www.moex.com/ru/derivatives/contractresults-exp.aspx?day1=20161123&day2=20171221&code=Si-12.17M211217CA60000
*/

/*
Полный код контракта:	Si-3.18M180118PA60000
Код в торговой системе:	Si60000BM8
Дата начала обращения:	13.10.2017
Последний день обращения:	18.01.2018
Дата исполнения:	18.01.2018
http://www.moex.com/ru/derivatives/contractresults-exp.aspx?day1=20171013&day2=20180118&code=Si-3.18M180118PA60000
 */
func getConvertCSV() [][]string {
	//url := "http://www.moex.com/ru/derivatives/optionsdesk-utilcsv.aspx?code=Si-12.17&sid=1&sub=&marg=1&delivery=30-11-17"
	url := "http://www.moex.com/ru/derivatives/contractsumresults-exp.aspx?day1=20170608&day2=20171220&codeb=SBRF-12.17M201217CA%20xxxxx"
	data, err := readCSVFromUrl(url)
	if err != nil {
		fmt.Println("CSV read error")
		panic(err)
	}
	return data
}

func DumpCsvData(data [][]string, isHead bool) {
	for idx, row := range data {
		if idx == 0 {
			if isHead {
				fmt.Println(row)
			}
		} else {
			fmt.Println(row)
		}
		//if idx == 6 && isHead {
		//	break
		//}
	}
}

func getConvertFile() {
	dat, err := ioutil.ReadFile("1.utilcsv")
	if err != nil {
		panic(err)
	}
	fmt.Print(DecodeWindows1251(dat[:222]))
}

func main() {
	fileName := "resultSBER.csv"
	dataCsv := getConvertCSV()
	DumpCsvData(dataCsv, true)

	//os.Exit(0)
	intStart :=0
	if _, err := os.Stat(fileName); err == nil {
		intStart =1
	}

	csv.SaveCsv(fileName, dataCsv[intStart:])
}
