package query_data

import (
	"../databases"
	"fmt"
	"testing"
)

func TestQueryHistoricalData(t *testing.T) {
	db := databases.MysqlINIT()
	defer db.Close()
	sql := "select * from stock_history"
	result, _ := db.Query(sql)
	var sum = 0
	for result.Next() {
		sum = sum + 1
	}
	fmt.Println(sum)

}
