package update_data

import (
	"../databases"
	"fmt"
	"github.com/tidwall/gjson"
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

func UpdataNowStockData(json string) error {
	db := databases.MysqlINIT()
	defer db.Close()
	arr := CompanyData{}
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
		return err
	}
	return nil
}
