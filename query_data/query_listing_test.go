package query_data

import (
	"../databases"
	"fmt"
	"testing"
)

// 读取所有上市公司代码
type tStockCode struct {
	CODE    string // 股票代码
	WY_CODE string // 网易股票代码
	HY_CODE string // 行业代码
}

func TestQueryListingAllData(t *testing.T) {
	db := databases.MysqlINIT()
	defer db.Close()
	// 结构体
	stockCode := tStockCode{}
	// 数组结构体
	stockCodeArry := []tStockCode{}

	sql := "select stock_code,stock_wy_code,stock_hy_code from stock_listing "
	if rows, err := db.Query(sql); err == nil {
		for rows.Next() {
			rows.Scan(&stockCode.CODE, &stockCode.WY_CODE, &stockCode.HY_CODE)
			stockCodeArry = append(stockCodeArry, stockCode)
		}
		for _, v := range stockCodeArry {
			fmt.Println(v)
		}
	}

}
