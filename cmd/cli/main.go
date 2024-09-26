package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	tools "github.com/machichima/taiwan-stock/internal"
)

func main() {

    if err := tools.ReadConfJson("config.json"); err != nil {
        panic(err)
    }

    // -1 cause cal rsvduration contains the current date
    ensureFetchDays := tools.Conf.KdDays + tools.Conf.RsvDuration - 1

	stockInfoList := []tools.StockInfo{}

	// Read from json file if exists
	// else crawl the data
	reader, err := os.Open("stockInfo.json")
	if err != nil {
		if os.IsNotExist(err) {

			// crawl for data
			stockInfoList, err = tools.GetAllStockInfoMonth(strconv.Itoa(tools.Conf.Month))
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
	if len(stockInfoList[0].ClosingPrices) < ensureFetchDays {
		// fetch one more month
		stockInfoListNew, err := tools.GetAllStockInfoMonth(strconv.Itoa(tools.Conf.Month - 1))
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
	k, d := tools.CalKDOneStock(float32(kInit), float32(dInit), stockInfoList[0], tools.Conf.RsvDuration)

	fmt.Println(k, d)
}
