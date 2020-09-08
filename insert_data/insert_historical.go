package insert_data

import (
	"../databases"
	"../model"
	"../read_file"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func InsertHistoricalData(path string) (bool, int, error) {
	var sum = 0
	// connect databases
	db := databases.MysqlINIT()
	// 定义数据结构体
	d := model.CompanyData{}
	//读取 csv 数据
	content, _ := read_file.ReadCsv(path)
	if content == nil {
		fmt.Println("由于限制下载，系统暂停 1小时")
		time.Sleep(1 * time.Hour)
		return false, 0, nil
	}
	//获取文件名 原来做stock code
	fileName := read_file.GetFileName(path)
	// 遍历 行放进结构体
	for _, row := range content {
		// 过滤 第一行数据
		if row[0] == "日期" {
			continue
		}
		// 加工数据
		d.TIME = row[0]
		d.CODE = fileName
		d.NAME = row[2]
		d.PRICE, _ = strconv.ParseFloat(row[3], 64)
		d.HIGH, _ = strconv.ParseFloat(row[4], 64)
		d.LOW, _ = strconv.ParseFloat(row[5], 64)
		d.OPEN, _ = strconv.ParseFloat(row[6], 64)
		d.YESTCLOSE, _ = strconv.ParseFloat(row[7], 64)
		d.UPDOWN, _ = strconv.ParseFloat(row[8], 64)
		d.PERCENT, _ = strconv.ParseFloat(row[9], 64)
		d.HS, _ = strconv.ParseFloat(row[10], 64)
		d.VOLUME, _ = strconv.ParseFloat(row[11], 64)
		d.TURNOVER, _ = strconv.ParseFloat(row[12], 64)
		d.TCAP, _ = strconv.ParseFloat(row[13], 64)
		d.MCAP, _ = strconv.ParseFloat(row[14], 64)
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
func InsertLSData(db *sql.DB, data model.CompanyData) (bool, error) {
	//使用DB结构体实例方法Prepare预处理插入,Prepare会返回一个stmt对象
	stmt, err := db.Prepare("insert into `stock_history`(stock_code,stock_name,stock_price,stock_open,stock_high,stock_hs," +
		"stock_low,stock_percent,stock_mcap,stock_tcap,stock_turnover,stock_updown,stock_volume,stock_yestclose,stock_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("预处理失败:", err)
		return false, err
	}
	//使用Stmt对象执行预处理参数
	_, err = stmt.Exec(data.CODE, data.NAME, data.PRICE, data.OPEN, data.HIGH, data.HS, data.LOW, data.PERCENT, data.MCAP, data.TCAP, data.TURNOVER, data.UPDOWN, data.VOLUME, data.YESTCLOSE, data.TIME)
	if err == nil {
		return true, nil
	}
	return false, err
}
