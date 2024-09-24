package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const allStockUrl string = "https://www.twse.com.tw/exchangeReport/STOCK_DAY_ALL"
const oneStockDateUrl string = "https://www.twse.com.tw/exchangeReport/STOCK_DAY?date=%s&stockNo=%s"

const retryTimes int = 5 // maximum retry time

func GetAllStockInfoMonth(month string) ([]StockInfo, error) {
	stockIdTitles, err := GetAllStockIDTitle()
	if err != nil {
		return []StockInfo{}, err
	}

	var allStockInfo []StockInfo

	startTime := time.Now()
	for _, idTitle := range stockIdTitles {
		// TODO: add oneStockInfo to the list of StockInfo
		oneStockInfo, err := FetchOneStockMonth(idTitle.ID, month)
		if err != nil {
			return []StockInfo{}, nil
		}

		fmt.Println(idTitle.ID)

		closingPrices := []ClosingPriceDate{}
		for _, data := range oneStockInfo.Data {
			fmt.Println(data)
			floatPrice, err := strconv.ParseFloat(data[6], 32)
			if err != nil {
				if errors.Is(err, strconv.ErrSyntax) {
					floatPrice = 0.0
				} else {
					return []StockInfo{}, err
				}
			}

			closingPrices = append(
				closingPrices, ClosingPriceDate{
					Date:         data[0],
					ClosingPrice: float32(floatPrice),
				})
		}

		// TODO: Check the closingPrices contain the correct info

		allStockInfo = append(
			allStockInfo,
			StockInfo{
				ID:            idTitle.ID,
				Title:         idTitle.Title,
				ClosingPrices: closingPrices,
			},
		)

		time.Sleep(100 * time.Millisecond) // delay one second to prevent being blocked
	}

	fmt.Println(time.Since(startTime))

	// Save result to json
	file, err := json.MarshalIndent(allStockInfo, "", "  ")
	if err := os.WriteFile("stockInfo.json", file, os.ModePerm); err != nil {
		return []StockInfo{}, err
	}

	return allStockInfo, nil

}

func FetchOneStockMonth(stockID string, month string) (StockDataStruct, error) {
	url := fmt.Sprintf(oneStockDateUrl, month, stockID)

	var res *http.Response

	for i := 0; i < retryTimes; i++ {
		var err error
		res, err = http.Get(url)
		if err == nil {
			break
		}
	}
	defer res.Body.Close()

	oneStockInfo := new(StockDataStruct)
	if err := json.NewDecoder(res.Body).Decode(oneStockInfo); err != nil {
		return StockDataStruct{}, err
	}

	return *oneStockInfo, nil
}

func GetAllStockIDTitle() ([]StockIDTitle, error) {
	allStockInfo, err := fetchAllStockInfo()
	if err != nil {
		return []StockIDTitle{}, err
	}

	var stockIDTitleList []StockIDTitle

	for _, infoList := range allStockInfo.Data {
		stockIDTitle := StockIDTitle{
			ID:    infoList[0],
			Title: infoList[1],
		}
		stockIDTitleList = append(stockIDTitleList, stockIDTitle)
	}

	return stockIDTitleList, nil
}

func fetchAllStockInfo() (StockDataStruct, error) {
	var res *http.Response
	for i := 0; i < retryTimes; i++ {
		var err error
		res, err = http.Get(allStockUrl)
		if err == nil {
			break
		}
	}
	defer res.Body.Close()

	allStockInfo := new(StockDataStruct)
	if err := json.NewDecoder(res.Body).Decode(allStockInfo); err != nil {
		return StockDataStruct{}, err
	}

	return *allStockInfo, nil
}
