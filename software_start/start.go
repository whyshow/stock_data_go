package software_start

import (
	"../get_stock_data"
	"../insert_data"
	"../query_data"
	"../update_data"
	"fmt"
	"time"
)

func Start() {

	// 判断行业数据表中是否有数据，如果没有则下载数据
	if sum := query_data.QueryHyData(); sum == 0 {
		// 获取行业板块信息
		var url2 = "http://quotes.money.163.com/hs/realtimedata/service/plate.php?host=/hs/realtimedata/service/plate.php&page=0&query=TYPE:HANGYE&fields=RN,NAME,STOCK_COUNT,PE,LB,HSL,PERCENT,TURNOVER,VOLUME,PLATE_ID,TYPE_CODE,PRICE,UPNUM,DOWNNUM,MAXPERCENTSTOCK,MINPERCENTSTOCK&sort=PERCENT&order=desc&count=50&type=query&callback=callback_1932607065&req=01636"
		_, b := get_stock_data.GetIndustryData(url2)
		insert_data.InsertHYData(b)
	}
	// 判断上市公司数据表中是否有数据，如果没有则下载数据
	if sum, _, _ := query_data.QueryListingAllData(); sum == 0 {

	}
	// 判断历史数据表中是否有数据，如果没有则下载数据
	if sum := query_data.QueryHistoricalData(); sum == 0 {
		// 读取所有股票代码
		_, listing, _ := query_data.QueryListingAllData()
		// 遍历所有股票代码
		for _, v := range listing {
			// 根据股票代码下载并保存数据
			if b, path, _ := get_stock_data.DownloadHistoricalData(v.CODE); b {
				// 根据保存数据的路径和名称导入数据
				for true {
					if b, sum, _ := insert_data.InsertHistoricalData(path); b {
						fmt.Println("代码", v.CODE, "共插入", sum, "条数据")
						break
					}
				}
			}
		}
	}

	//设置时间周期
	myTicker := time.NewTicker(time.Second)
	for {
		//当前时间
		nowTime := <-myTicker.C
		// 非周六周日 和规定的时间才能启动
		if nowTime.Weekday().String() != "Saturday" && nowTime.Weekday().String() != "Sunday" && nowTime.Hour() == 23 && nowTime.Minute() == 01 && nowTime.Second() == 30 {
			fmt.Println("时间到了,任务启动", nowTime)
			/** 更新上市公司股票数据 **/
			// 从数据库读取股票代码
			if _, result, err := query_data.QueryListingAllData(); err == nil {
				for _, v := range result {
					if json, err := get_stock_data.GetNowStockData(v.WY_CODE); err == nil {
						if err := update_data.UpdataNowStockData(json); err == nil {
							fmt.Println("股票", v.WY_CODE, "更新成功")
						} else {
							fmt.Println("股票", v.WY_CODE, "更新失败 ", err)
						}
					} else {
						fmt.Println("请求数据错误", err)
					}

					time.Sleep(1 * time.Second / 5)
				}

			} else {
				fmt.Println(err)
			}
			fmt.Println("时间到了,任务完成", nowTime)
		} else {
			fmt.Println("时间未到", nowTime)
		}
	}
}

func InsertData(path string) (int, bool) {
	b, sum, _ := insert_data.InsertHistoricalData(path)
	if b {
		return sum, true
	}
	return 0, false
}
