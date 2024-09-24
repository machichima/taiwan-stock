package main

import (
	"encoding/json"
	"fmt"
	"os"

	tools "github.com/machichima/taiwan-stock/internal"
)

func main() {

	stockInfoList := []tools.StockInfo{}

    // Read from json file if exists
    // else crawl the data
	reader, err := os.Open("stockInfo.json")
	if err != nil {
		if os.IsNotExist(err) {
            
            // crawl for data
			stockInfoList, err = tools.GetAllStockInfoMonth("9")
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

    fmt.Println(stockInfoList)

}
