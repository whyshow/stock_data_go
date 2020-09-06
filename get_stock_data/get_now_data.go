package get_stock_data

import (
	"../chrome_link"
)

// 获取 股票最新数据

func GetNowData(url string) string {
	// 通过chrome 浏览器请求 HTTP页面
	body := chrome_link.ChromeHttpGet(url)
	//转 String 类型
	b := string(body)
	return b
}
