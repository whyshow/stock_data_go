package get_stock_data

// 获取股市行业板块数据
// 插入到数据库中
import (
	"../chrome_link"
)

func GetIndustryData(url string) (bool, string) {
	// 通过chrome 浏览器请求 HTTP页面
	body, err := chrome_link.ChromeHttpGet(url)
	if err != nil {
		return false, ""
	} //转 String 类型
	b := string(body)
	// 计算长度
	l := len(b)
	// 截取字符生产 json 数据
	b = b[20 : l-1]
	//fmt.Println(b)
	// 将行业信息数据插入数据库中
	return false, b
}
