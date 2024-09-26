package tools

type StockDataStruct struct {
	Stat   string     `json:"stat"`
	Date   string     `json:"date"`
	Title  string     `json:"title"`
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
}

type StockIDTitle struct {
	ID    string
	Title string
}

type StockInfo struct {
	ID            string
	Title         string
	ClosingPrices []ClosingPriceDate
}

type ClosingPriceDate struct {
	Date         string
	ClosingPrice float32
}

type Config struct {
	Month           int
	RsvDuration     int
	KdDays          int
	RetryTimes      int
	AllStockUrl     string
	OneStockDateUrl string
}
