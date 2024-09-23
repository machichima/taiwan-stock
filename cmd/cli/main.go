package main

import (
	"fmt"

	tools "github.com/machichima/taiwan-stock/internal"
)

func main() {
    stockIDs, err := tools.GetAllStockID()
    if err != nil {
        panic(err)
    }

    fmt.Println(stockIDs)
}
