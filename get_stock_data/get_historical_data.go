package get_stock_data

import (
	"../chrome_link"
	"fmt"
	"io/ioutil"
	"os"
)

// 下载个股历史交易数据
// 保持为 scv文件
// 返回是否true，保持路径，错误信息
func DownloadHistoricalData(code string) (bool, string, error) {
	var url = "http://quotes.money.163.com/service/chddata.html?code=1" + code +
		"&start=20180109&end=20200904&fields=TCLOSE;HIGH;LOW;TOPEN;LCLOSE;CHG;PCHG;TURNOVER;VOTURNOVER;VATURNOVER;TCAP;MCAP"
	body, err := chrome_link.ChromeHttpGet(url)
	if err != nil {
		return false, "", err
	}
	// 获取目录
	dir, _ := os.Getwd()
	rootPath := dir + "/csv/"
	// 创建文件夹
	if err := os.Mkdir(rootPath, 0666); err != nil {
		fmt.Println("文件夹创建失败")
	}
	f, err := os.Create(rootPath + code + ".csv")
	if err != nil {
		return false, "", err
	}
	defer f.Close()
	if err := ioutil.WriteFile(rootPath+code+".csv", body, 0666); err != nil {
		return false, "", err
	}
	return true, rootPath + code + ".csv", nil
}
