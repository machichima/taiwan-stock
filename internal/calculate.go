package tools

import (
	"errors"
	"fmt"
	"slices"
)

// calcuate the RSV value for one stock. Read the StockInfo struct
// and return the rsv value for it
//
// Calculate the rsv value for "daysBefore" days ahead of the latest date
func CalRSVOneStock(stockInfo StockInfo, duration int, daysBefore int) float32 {

	lastIndex := len(stockInfo.ClosingPrices) - 1 - daysBefore

	// make sure the data count is larger than the duration
	if len(stockInfo.ClosingPrices) < (duration + daysBefore) {
		// do something to get more data
		fmt.Printf("data is not enough for the calculation, at least %d data needed\n", duration+daysBefore)
		fmt.Printf("only %d data available", len(stockInfo.ClosingPrices))
		panic(errors.New("data not enough"))
	}

	// Run through closing Price in reverse order
	closingPrice := make([]float32, 0)
	for i := lastIndex; i > lastIndex-duration; i-- {
		fmt.Println(stockInfo.ClosingPrices[i])
		closingPrice = append(
			closingPrice,
			stockInfo.ClosingPrices[i].ClosingPrice,
		)
	}

	minVal := slices.Min(closingPrice)
	maxVal := slices.Max(closingPrice)

	rsv := (closingPrice[0] - minVal) / (maxVal - minVal) * 100

	fmt.Println(closingPrice, minVal, maxVal)

	return rsv
}

// calcuate the K and D values for one stock. Read the previous k and d value
// and the rsv value (calcuated by CalRSVOneStock) and return the
// K and D values respectively
func CalKDOneStock(kPrev float32, dPrev float32, stockInfo StockInfo, rsvDuration int) (float32, float32) {

	k, d := kPrev, dPrev
	fmt.Println("init k, d: ", k, d)
	for i := Conf.KdDays - 1; i >= 0; i-- {
		fmt.Println(i, " days before")
		rsv := CalRSVOneStock(stockInfo, rsvDuration, i)
		fmt.Println("rsv: ", rsv)
		k = k*(float32(2)/3) + rsv*(float32(1)/3)
		d = d*(float32(2)/3) + k*(float32(1)/3)
		fmt.Println("k, d: ", k, d)
	}

	return k, d
}
