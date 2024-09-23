# Taiwan stock

> Query Taiwan company with specific condition in stock market for investigation

- Api link: [api](https://openapi.twse.com.tw/)
    - Fields:
        - /exchangeReport/STOCK_DAY_AVG_ALL
        - /exchangeReport/STOCK_DAY_ALL

## Steps

1. crawl `https://www.twse.com.tw/exchangeReport/STOCK_DAY_ALL` and sort based on the 漲跌價差 
    - get the stockNo
2. Use `https://www.twse.com.tw/exchangeReport/STOCK_DAY?date=20240901&stockNo=2330` to get the value for each stockNo of the month

## Ref
- calcualte kd value:
    - [link](https://medium.com/%E5%8F%B0%E8%82%A1etf%E8%B3%87%E6%96%99%E7%A7%91%E5%AD%B8-%E7%A8%8B%E5%BC%8F%E9%A1%9E/%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80-%E8%87%AA%E5%BB%BAkd%E5%80%BC-819d6fd707c8)
- twstock repo
    - [link](https://github.com/mlouielu/twstock/tree/master)
