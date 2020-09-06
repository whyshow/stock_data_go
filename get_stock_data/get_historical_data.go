package get_stock_data

import (
	"../chrome_link"
	"io/ioutil"
	"os"
)

// 下载个股历史交易数据
// 保持为 scv文件
func DownloadHistoricalData(code string) (bool, error) {
	var uri = "http://quotes.money.163.com/service/chddata.html?code=1" + code + "&start=20100812&end=20200904&fields=TCLOSE;HIGH;LOW;TOPEN;LCLOSE;CHG;PCHG"
	body := chrome_link.ChromeHttpGet(uri)
	f, _ := os.Create("cvsfile/" + code + ".csv")
	defer f.Close()
	if err := ioutil.WriteFile("cvsfile/"+code+".csv", body, 0666); err != nil {
		return false, err
	}
	return true, nil
}
