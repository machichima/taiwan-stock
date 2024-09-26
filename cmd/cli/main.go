package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	tools "github.com/machichima/taiwan-stock/internal"
)

func main() {

	rsvDuration := 9
	month := 9

	stockInfoList := []tools.StockInfo{}

	// Read from json file if exists
	// else crawl the data
	reader, err := os.Open("stockInfo.json")
	if err != nil {
		if os.IsNotExist(err) {

			// crawl for data
			stockInfoList, err = tools.GetAllStockInfoMonth(strconv.Itoa(month))
			if err != nil {
				panic(err)
			}

		} else {
			panic(err)
		}
	} else {
		// read from the json file
		if err := json.NewDecoder(reader).Decode(&stockInfoList); err != nil {
			panic(err)
		}
	}

	// check the number of stock is enough
	if len(stockInfoList[0].ClosingPrices) < tools.EnsureFetchDays {
		// fetch one more month
		stockInfoListNew, err := tools.GetAllStockInfoMonth(strconv.Itoa(month - 1))
		if err != nil {
			panic(err)
		}
		stockInfoList = append(stockInfoList, stockInfoListNew...)

	}
     
	// Save result to json
	file, err := json.MarshalIndent(stockInfoList, "", "  ")
	if err := os.WriteFile("stockInfo.json", file, os.ModePerm); err != nil {
		panic(err)
	}


	// rsv := tools.CalRSVOneStock(stockInfoList[0], rsvDuration)
	// fmt.Println(rsv)

	kInit, dInit := 50, 50
	k, d := tools.CalKDOneStock(float32(kInit), float32(dInit), stockInfoList[0], rsvDuration)

	fmt.Println(k, d)
}
