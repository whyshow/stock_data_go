package insert_data

import (
	"../databases"
	"../model"
	"fmt"
	"github.com/tidwall/gjson"
)

// 插入公司股票信息数据

// 传入 json 类型数据
// 向 stock_listing 表中循环插入数据
func InsertListingData(json string, hy_code string) (bool, error) {
	// connect databases
	db := databases.MysqlINIT()
	defer db.Close()
	//结构体
	compant_data := model.CompanyData{}
	//c_data := []CompanyData{}
	// 处理json 数据
	result := gjson.GetMany(json, "list.#.CODE", "list.#.FIVE_MINUTE", "list.#.HIGH", "list.#.HS", "list.#.LB", "list.#.LOW",
		"list.#.MCAP", "list.#.MFRATIO.MFRATIO2", "list.#.MFRATIO.MFRATIO10", "list.#.MFSUM", "list.#.NAME", "list.#.OPEN", "list.#.PE", "list.#.PERCENT",
		"list.#.PRICE", "list.#.SYMBOL", "list.#.TCAP", "list.#.TURNOVER", "list.#.UPDOWN", "list.#.VOLUME", "list.#.WB", "list.#.YESTCLOSE",
		"list.#.ZF")
	value5 := gjson.Get(json, "list.#")
	CODE := result[0].Array()
	//FIVE_MINUTE := result[1].Array()
	HIGH := result[2].Array()
	//HS := result[3].Array()
	//LB := result[4].Array()
	LOW := result[5].Array()
	//MCAP := result[6].Array()
	//MFRATIO2 := result[7].Array()
	//MFRATIO10 := result[8].Array()
	//MFSUM := result[9].Array()
	NAME := result[10].Array()
	OPEN := result[11].Array()
	//PE := result[12].Array()
	PERCENT := result[13].Array()
	PRICE := result[14].Array()
	SYMBOL := result[15].Array()
	//TCAP := result[16].Array()
	TURNOVER := result[17].Array()
	UPDOWN := result[18].Array()
	VOLUME := result[19].Array()
	WB := result[20].Array()
	YESTCLOSE := result[21].Array()
	//ZF := result[22].Array()
	var i int64
	for i = 0; i < value5.Int(); i++ {
		compant_data.NAME = NAME[i].String()
		compant_data.WY_CODE = CODE[i].String()
		compant_data.CODE = SYMBOL[i].String()
		compant_data.HY_CODE = hy_code
		//compant_data.FIVE_MINUTE = FIVE_MINUTE[i].Float()
		compant_data.HIGH = HIGH[i].Float()
		//compant_data.HS = HS[i].Float()
		compant_data.LOW = LOW[i].Float()
		//compant_data.LB = LB[i].Float()
		//compant_data.MCAP = MCAP[i].Float()
		//compant_data.MFRATIO2 = MFRATIO2[i].Float()
		//compant_data.MFRATIO10 = MFRATIO10[i].Float()
		//compant_data.MFSUM = MFSUM[i].Float()
		compant_data.OPEN = OPEN[i].Float()
		//compant_data.PE = PE[i].Float()
		compant_data.PERCENT = PERCENT[i].Float()
		compant_data.PRICE = PRICE[i].Float()
		//compant_data.TCAP = TCAP[i].Float()
		compant_data.TURNOVER = TURNOVER[i].Float()
		compant_data.UPDOWN = UPDOWN[i].Float()
		compant_data.VOLUME = VOLUME[i].Float()
		compant_data.WB = WB[i].Float()
		compant_data.YESTCLOSE = YESTCLOSE[i].Float()
		//compant_data.ZF = ZF[i].Float()
		//c_data = append(c_data, compant_data)
		//使用DB结构体实例方法Prepare预处理插入,Prepare会返回一个stmt对象
		stmt, err := db.Prepare("insert into `stock_listing`(stock_code,stock_wy_code,stock_hy_code,stock_name,stock_price,stock_open,stock_five_minute,stock_high," +
			"stock_lb,stock_low,stock_mcap,stock_mfratio2,stock_mfratio10,stock_mfsum,stock_pe,stock_percent,stock_tcap,stock_turnover,stock_updown,stock_volume,stock_wb," +
			"stock_yestclose,stock_zf) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			fmt.Println("预处理失败:", err)
			return false, err
		}
		//使用Stmt对象执行预处理参数
		_, err = stmt.Exec(compant_data.CODE, compant_data.WY_CODE, compant_data.HY_CODE, compant_data.NAME, compant_data.PRICE, compant_data.OPEN, compant_data.FIVE_MINUTE, compant_data.HIGH,
			compant_data.LB, compant_data.LOW, compant_data.MCAP, compant_data.MFRATIO2, compant_data.MFRATIO10, compant_data.MFSUM, compant_data.PE, compant_data.PERCENT, compant_data.TCAP,
			compant_data.TURNOVER, compant_data.UPDOWN, compant_data.VOLUME, compant_data.WB, compant_data.YESTCLOSE, compant_data.ZF)
		if err != nil {
			fmt.Println(err)
		}
	}
	return false, nil
}
