package read_file

import (
	"bufio"
	"encoding/csv"
	"github.com/axgle/mahonia"
	"log"
	"os"
	"path"
	"strings"
)

func ReadCsv(fileName string) ([][]string, bool) {
	// 针对小文件，一次性读取所有的文件。注意，r要重新赋值，因为readall是读取剩下的
	fs, _ := os.Open(fileName)
	//判断下载的数据是否正确
	scanner := bufio.NewScanner(fs)
	lineText := ""
	for scanner.Scan() {
		lineText = scanner.Text()
		lineText = lineText[1:9]
		break
	}
	if lineText == "!DOCTYPE" {
		return nil, false
	}
	defer fs.Close()
	fs1, _ := os.Open(fileName)
	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	r1 := csv.NewReader(decoder.NewReader(fs1))
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	return content, true
}

func GetFileName(filepath string) string {
	filenameWithSuffix := path.Base(filepath)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}
