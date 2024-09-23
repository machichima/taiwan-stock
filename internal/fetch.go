package tools

import (
	"encoding/json"
	"net/http"
)

const allStockUrl string = "https://www.twse.com.tw/exchangeReport/STOCK_DAY_ALL"
const oneStockDateUrl string = "https://www.twse.com.tw/exchangeReport/STOCK_DAY?date=%s&stockNo=%s"

const retryTimes int = 5 // maximum retry time

type AllStockStruct struct {
	Stat string `json:"stat"`
    Date string `json:"date"`
    Title string `json:"title"`
    Fields []string `json:"fields"`
    Data [][]string `json:"data"`
}


func FetchOneStockMonth(stockID string, month string) error {
    return nil
}


func GetAllStockID() ([]string, error) {
    allStockInfo, err := fetchAllStockInfo()
    if err != nil {
        return []string{}, err
    }

    var stockIDs []string

    for _, infoList := range allStockInfo.Data {
        stockIDs = append(stockIDs, infoList[0])
    }

    return stockIDs, nil
}


func fetchAllStockInfo() (AllStockStruct, error) {
	var res *http.Response
	for i := 0; i < retryTimes; i++ {
		var err error
		res, err = http.Get(allStockUrl)
		if err == nil {
			break
		}
	}
    defer res.Body.Close()

    allStockInfo := new(AllStockStruct)
    if err := json.NewDecoder(res.Body).Decode(allStockInfo); err != nil {
        return AllStockStruct{}, err
    }

    return *allStockInfo, nil
}


