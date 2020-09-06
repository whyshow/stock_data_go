package chrome_link

import (
	"io/ioutil"
	"net/http"
)

// 模拟 Chrome浏览器GET下载请求
func ChromeHttpGet(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
		return nil
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	client := http.Client{}
	// 发送请求
	response, err := client.Do(request)
	if err != nil {
		panic(err)
		return nil
	}
	defer response.Body.Close()
	// 读取数据成 []byte 类型
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
		return nil
	}
	return body
}
