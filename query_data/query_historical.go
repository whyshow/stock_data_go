package query_data

import (
	"../databases"
)

func QueryHistoricalData() int {
	db := databases.MysqlINIT()
	defer db.Close()
	sql := "select * from stock_history"
	result, _ := db.Query(sql)
	var sum = 0
	for result.Next() {
		sum = sum + 1
	}
	return sum
}
