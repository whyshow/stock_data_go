package insert

import (
	"../databases"
	"../get_stock_data"
	"../read"
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tidwall/gjson"
	"strconv"
)

type Data struct {
	stock_code          string
	stock_name          string
	stock_date          string
	stock_closing       float64
	stock_high          float64
	stock_low           float64
	stock_opening       float64
	stock_prior         string
	stock_change_amount float64
	stock_amplitude     float64
}

func InsertData(path string) (bool, int, error) {
	var sum = 0
	// connect databases
	db := databases.MysqlINIT()
	// 定义数据结构体
	d := Data{}
	//读取 csv 数据
	content := read.ReadCsv(path)
	//获取文件名 原来做stock code
	fileName := read.GetFileName(path)
	// 遍历 行放进结构体
	for _, row := range content {
		// 过滤 第一行数据
		if row[0] == "日期" {
			continue
		}
		// 加工数据
		d.stock_date = row[0]
		d.stock_code = fileName
		d.stock_name = row[2]
		d.stock_closing, _ = strconv.ParseFloat(row[3], 64)
		d.stock_high, _ = strconv.ParseFloat(row[4], 64)
		d.stock_low, _ = strconv.ParseFloat(row[5], 64)
		d.stock_opening, _ = strconv.ParseFloat(row[6], 64)
		d.stock_prior = row[7]
		d.stock_change_amount, _ = strconv.ParseFloat(row[8], 64)
		d.stock_amplitude, _ = strconv.ParseFloat(row[9], 64)
		sum++
		// 插入数据
		if b, err := InsertLSData(db, d); !b {
			fmt.Println(err)
			return false, sum, err
		}
	}
	// 关闭数据库
	db.Close()
	return true, sum, nil
}

// 插入数据
func InsertLSData(db *sql.DB, data Data) (bool, error) {
	//使用DB结构体实例方法Prepare预处理插入,Prepare会返回一个stmt对象
	stmt, err := db.Prepare("insert into `stock_history`(stock_code,stock_name,stock_date,stock_closing,stock_high,stock_low,stock_opening,stock_prior,stock_change_amount,stock_amplitude) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("预处理失败:", err)
		return false, err
	}
	//使用Stmt对象执行预处理参数
	_, err = stmt.Exec(data.stock_code, data.stock_name, data.stock_date, data.stock_closing, data.stock_high, data.stock_low, data.stock_opening, data.stock_prior, data.stock_change_amount, data.stock_amplitude)
	if err == nil {
		return true, nil
	}
	return false, err
}

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
		var url = "http://quotes.money.163.com/hs/service/diyrank.php?host=http%3A%2F%2Fquotes.money.163.com%2Fhs%2Fservice%2Fdiyrank.php&page=0&query=PLATE_IDS%3A" + hy.hy_code + "&fields=NO%2CSYMBOL%2CNAME%2CPRICE%2CPERCENT%2CUPDOWN%2CFIVE_MINUTE%2COPEN%2CYESTCLOSE%2CHIGH%2CLOW%2CVOLUME%2CTURNOVER%2CHS%2CLB%2CWB%2CZF%2CPE%2CMCAP%2CTCAP%2CMFSUM%2CMFRATIO.MFRATIO2%2CMFRATIO.MFRATIO10%2CSNAME%2CCODE%2CANNOUNMT%2CUVSNEWS&sort=PERCENT&order=desc&count=500&type=query"
		json := get_stock_data.GetNowData(url)
		InsertNowData(json, hy.hy_code)
	}
	return true, nil
}

//
