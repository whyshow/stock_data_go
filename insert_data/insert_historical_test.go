package insert_data

import (
	"../databases"
	"../read_file"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

type sData struct {
	CODE        string  // 股票代码
	WY_CODE     string  // 网易股票代码
	HY_CODE     string  // 行业代码
	NAME        string  //名称
	PRICE       float64 // 价格
	OPEN        float64 // 今开
	FIVE_MINUTE float64 // 5分钟涨跌额
	HIGH        float64 // 最高价
	HS          float64 // 换手率
	LB          float64 // 量比
	LOW         float64 // 最低
	MCAP        float64 // 流通市值
	MFRATIO2    float64 //净利润
	MFRATIO10   float64 //主营收
	MFSUM       float64 // 每股收益
	PE          float64 // 市盈率
	PERCENT     float64 // 涨跌幅
	TCAP        float64 //总市值
	TURNOVER    float64 //成交额
	UPDOWN      float64 // 涨跌额
	VOLUME      float64 // 成交量
	WB          float64 //委比
	YESTCLOSE   float64 //昨收
	ZF          float64 //振幅
	TIME        string  //日期
}

func TestInsertHistoricalData(t *testing.T) {
	var sum = 0
	// connect databases
	db := databases.MysqlINIT()
	// 定义数据结构体
	d := sData{}
	//读取 csv 数据
	content, _ := read_file.ReadCsv("../csv/000001.csv")
	//获取文件名 原来做stock code
	fileName := read_file.GetFileName("../csv/000001.csv")
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
		if b, err := InsertLSDatas(db, d); !b {
			fmt.Println(err)
		}
	}
	// 关闭数据库
	db.Close()

}

// 插入数据
func InsertLSDatas(db *sql.DB, data sData) (bool, error) {
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
