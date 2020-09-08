package get_stock_data

import (
	"../chrome_link"
	"fmt"
)

// 根据股票代码查询最新的数据
// 返回 json 数据
func GetNowStockData(wy_code string) (string, error) {
	url := "http://api.money.126.net/data/feed/" + wy_code + ",money.api"
	// 模拟谷歌浏览器请求数据
	body, err := chrome_link.ChromeHttpGet(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//转 String 类型
	b := string(body)
	// 计算长度
	l := len(b)
	// 截取字符生产 json 数据
	b = b[21 : l-2]
	return b, err
}
