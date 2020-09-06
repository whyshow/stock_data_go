package main

import (
	"./databases"
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/axgle/mahonia"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
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

func main() {
	path := "./csv/000001.csv"
	db := databases.MysqlINIT()

	d := Data{}
	content := ReadCsv(path)
	fileName := GetFileName(path)

	for _, row := range content {
		if row[0] == "日期" {
			continue
		}
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
		StructInsert(db, d)

	}
	db.Close()
}

// 插入数据
func StructInsert(db *sql.DB, data Data) {

	//使用DB结构体实例方法Prepare预处理插入,Prepare会返回一个stmt对象
	stmt, err := db.Prepare("insert into `stock_history`(stock_code,stock_name,stock_date,stock_closing,stock_high,stock_low,stock_opening,stock_prior,stock_change_amount,stock_amplitude) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("预处理失败:", err)
		return
	}
	//使用Stmt对象执行预处理参数
	_, err = stmt.Exec(data.stock_code, data.stock_name, data.stock_date, data.stock_closing, data.stock_high, data.stock_low, data.stock_opening, data.stock_prior, data.stock_change_amount, data.stock_amplitude)
	if err != nil {
		fmt.Println("执行预处理失败:", err)
		return
	} else {
		//rows,_ := result.RowsAffected()
		//fmt.Println("执行成功,影响行数",rows,"行" )
	}
}

func ReadCsv(fileName string) [][]string {
	// 针对小文件，一次性读取所有的文件。注意，r要重新赋值，因为readall是读取剩下的
	fs1, _ := os.Open(fileName)
	defer fs1.Close()
	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	r1 := csv.NewReader(decoder.NewReader(fs1))
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	return content
}

func GetFileName(filepath string) string {
	filenameWithSuffix := path.Base(filepath)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}
