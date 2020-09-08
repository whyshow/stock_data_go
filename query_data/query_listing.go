package query_data

import (
	"../databases"
)

// 读取所有上市公司代码
type StockCode struct {
	CODE    string // 股票代码
	WY_CODE string // 网易股票代码
	HY_CODE string // 行业代码
}

func QueryListingAllData() (int, []StockCode, error) {
	db := databases.MysqlINIT()
	defer db.Close()
	// 结构体
	stockCode := StockCode{}
	// 数组结构体
	stockCodeArry := []StockCode{}

	var sum = 0
	sql := "select stock_code,stock_wy_code,stock_hy_code from stock_listing "
	if rows, err := db.Query(sql); err == nil {
		for rows.Next() {
			sum = sum + 1
			rows.Scan(&stockCode.CODE, &stockCode.WY_CODE, &stockCode.HY_CODE)
			stockCodeArry = append(stockCodeArry, stockCode)
		}
	} else {
		return sum, stockCodeArry, err
	}
	return sum, stockCodeArry, nil
}
