package csv

import "testing"

func TestSaveCsv(t *testing.T) {
	var data = [][]string{{"Line3", "Hello Readers of"}, {"Line2", "golangcode.com"}}
	SaveCsv(data)
}