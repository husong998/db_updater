package app

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Parser struct {
}

func (o *Parser) Parse(fileName string) (result []Item, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	result = make([]Item, 0, len(records)-1)

	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 3 {
			err = fmt.Errorf("records %v with less than 3 fields", i)
		}
		var (
			id, stock uint64
			price     float64
		)
		id, err = strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			return
		}
		price, err = strconv.ParseFloat(record[1], 64)
		if err != nil {
			return
		}
		stock, err = strconv.ParseUint(record[2], 10, 32)
		if err != nil {
			return
		}

		result = append(result, Item{
			ID:    uint(id),
			Price: price,
			Stock: uint(stock),
		})
	}
	return
}
