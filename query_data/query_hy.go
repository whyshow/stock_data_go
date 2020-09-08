package query_data

import (
	"../databases"
)

func QueryHyData() int {
	db := databases.MysqlINIT()
	defer db.Close()
	sql := "select * from stock_hy"
	result, _ := db.Query(sql)
	var sum = 0
	for result.Next() {
		result.Scan()
		sum = sum + 1
	}
	return sum
}
