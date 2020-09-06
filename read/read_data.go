package read

import (
	"encoding/csv"
	"github.com/axgle/mahonia"
	"log"
	"os"
	"path"
	"strings"
)

func ReadCsv(fileName string) [][]string {
	// 针对小文件，一次性读取所有的文件。注意，r要重新赋值，因为readall是读取剩下的
	fs1, _ := os.Open(fileName)
	defer fs1.Close()
	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	r1 := csv.NewReader(decoder.NewReader(fs1))
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	return content
}

func GetFileName(filepath string) string {
	filenameWithSuffix := path.Base(filepath)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}
