package get_stock_data

import (
	"../chrome_link"
	"fmt"
	"testing"
)

func TestGetNowStockData(t *testing.T) {
	wy_code := "1000002"
	url := "http://api.money.126.net/data/feed/" + wy_code + ",money.api"
	// 模拟谷歌浏览器请求数据
	body, _ := chrome_link.ChromeHttpGet(url)
	//转 String 类型
	b := string(body)
	// 计算长度
	l := len(b)
	// 截取字符生产 json 数据
	b = b[21 : l-2]
	fmt.Println(b)
}
