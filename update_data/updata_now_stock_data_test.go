package update_data

import (
	"../databases"
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

// 根据股票代码 更新最新的 股票数据
type CompanyData struct {
	CODE        string  // 股票代码
	PRICE       float64 // 价格
	OPEN        float64 // 今开
	FIVE_MINUTE float64 // 5分钟涨跌额
	HIGH        float64 // 最高价
	LOW         float64 // 最低
	PERCENT     float64 // 涨跌幅
	UPDOWN      float64 // 涨跌额
	VOLUME      float64 // 成交量
	YESTCLOSE   float64 //昨收
	TIME        string  //更新时间
}

func TestUpdataNowStockData(t *testing.T) {
	db := databases.MysqlINIT()
	defer db.Close()
	arr := CompanyData{}
	json := "{\"1000002\":{\"code\": \"1000002\", \"percent\": 0.012825, \"high\": 29.05, \"askvol3\": 21200, \"askvol2\": 150213, \"askvol5\": 70500, \"askvol4\": 37500, \"price\": 28.43, \"open\": 27.95, \"bid5\": 28.39, \"bid4\": 28.4, \"bid3\": 28.41, \"bid2\": 28.42, \"bid1\": 28.43, \"low\": 27.95, \"updown\": 0.36, \"type\": \"SZ\", \"bidvol1\": 227139, \"status\": 0, \"bidvol3\": 26500, \"bidvol2\": 16100, \"symbol\": \"000002\", \"update\": \"2020/09/07 15:59:58\", \"bidvol5\": 23500, \"bidvol4\": 168600, \"volume\": 143229421, \"askvol1\": 396832, \"ask5\": 28.48, \"ask4\": 28.47, \"ask1\": 28.44, \"name\": \"\\u4e07  \\u79d1\\uff21\", \"ask3\": 28.46, \"ask2\": 28.45, \"arrow\": \"\\u2191\", \"time\": \"2020/09/07 15:59:54\", \"yestclose\": 28.07, \"turnover\": 4103510527.79} }"
	// 解析JSON数据
	results := gjson.GetMany(json, "*.percent", "*.high", "*.price", "*.open", "*.low", "*.updown", "*.symbol",
		"*.volume", "*.time", "*.yestclose")
	arr.PERCENT = results[0].Float()
	arr.HIGH = results[1].Float()
	arr.OPEN = results[3].Float()
	arr.LOW = results[4].Float()
	arr.PRICE = results[2].Float()
	arr.UPDOWN = results[5].Float()
	arr.CODE = results[6].String()
	arr.VOLUME = results[7].Float()
	arr.TIME = results[8].String()
	arr.YESTCLOSE = results[9].Float()
	// 更新到数据库中
	sql := fmt.Sprintf("UPDATE stock_listing SET stock_price=?,stock_open=?,stock_high=?,stock_percent=?,stock_low=?,stock_updown=?,stock_volume=?," +
		"stock_yestclose=?,stock_time=? WHERE stock_code=?")
	if _, err := db.Exec(sql, arr.PRICE, arr.OPEN, arr.HIGH, arr.PERCENT, arr.LOW, arr.UPDOWN, arr.VOLUME, arr.YESTCLOSE, arr.TIME, arr.CODE); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("更新成功")
	}

}
