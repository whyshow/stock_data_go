package get_stock_data

import (
	"../chrome_link"
)

// 根据行业代码获取当前行业下所有公司的股票最新数据
func GetHyAllCompanyData(hy_code string) string {
	url := "http://quotes.money.163.com/hs/service/diyrank.php?host=http%3A%2F%2Fquotes.money.163.com%2Fhs%2Fservice%2Fdiyrank.php&page=" +
		"0&query=PLATE_IDS%3A" + hy_code + "&fields=NO%2CSYMBOL%2CNAME%2CPRICE%2CPERCENT%2CUPDOWN%2CFIVE_MINUTE%2COPEN%2CYESTCLOSE%2CHIG" +
		"H%2CLOW%2CVOLUME%2CTURNOVER%2CHS%2CLB%2CWB%2CZF%2CPE%2CMCAP%2CTCAP%2CMFSUM%2CMFRATIO.MFRATIO2%2CMFRATIO.MFRATIO10%2CSNAME%2CCO" +
		"DE%2CANNOUNMT%2CUVSNEWS&sort=PERCENT&order=desc&count=500&type=query"

	// 通过chrome 浏览器请求 HTTP页面
	body, _ := chrome_link.ChromeHttpGet(url)
	//转 String 类型
	b := string(body)
	return b
}
