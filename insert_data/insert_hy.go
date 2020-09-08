package insert_data

import (
	"../databases"
	"../get_stock_data"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tidwall/gjson"
)

type HyData struct {
	hy_code           string
	hy_name           string
	hy_company_amount int64
}

//插入行业数据
func InsertHYData(json string) (bool, error) {
	// 处理JSON数据
	// 定义数据结构体
	hy_data := []HyData{}
	hy := HyData{}
	result := gjson.GetMany(json, "list.#.NAME", "list.#.PLATE_ID", "list.#.STOCK_COUNT")
	name := result[0].Array()
	code := result[1].Array()
	amount := result[2].Array()
	for i, res := range name {
		hy.hy_name = res.String()
		hy.hy_code = code[i].String()
		hy.hy_company_amount = amount[i].Int()
		hy_data = append(hy_data, hy)
	}
	// connect databases
	db := databases.MysqlINIT()
	defer db.Close()
	for _, v := range hy_data {
		hy.hy_code = v.hy_code
		hy.hy_name = v.hy_name
		hy.hy_company_amount = v.hy_company_amount
		//使用DB结构体实例方法Prepare预处理插入,Prepare会返回一个stmt对象
		stmt, err := db.Prepare("insert into `stock_hy`(hy_code,hy_name,hy_company_amount) values(?,?,?)")
		if err != nil {
			fmt.Println("预处理失败:", err)
			return false, err
		}
		//使用Stmt对象执行预处理参数
		_, err = stmt.Exec(hy.hy_code, hy.hy_name, hy.hy_company_amount)
		if err != nil {
			fmt.Println(err)
			return false, nil
		}
		// 根据行业代码查询行业内的股票并插入数据库
		json := get_stock_data.GetHyAllCompanyData(hy.hy_code)
		InsertListingData(json, hy.hy_code)
	}
	return true, nil
}

//
