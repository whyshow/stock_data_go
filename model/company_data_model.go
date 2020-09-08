package model

// 数据库结构体
type CompanyData struct {
	CODE        string  // 股票代码
	WY_CODE     string  // 网易股票代码
	HY_CODE     string  // 行业代码
	NAME        string  //名称
	PRICE       float64 // 价格
	OPEN        float64 // 今开
	FIVE_MINUTE float64 // 5分钟涨跌额
	HIGH        float64 // 最高价
	HS          float64 // 换手率
	LB          float64 // 量比
	LOW         float64 // 最低
	MCAP        float64 // 流通市值
	MFRATIO2    float64 //净利润
	MFRATIO10   float64 //主营收
	MFSUM       float64 // 每股收益
	PE          float64 // 市盈率
	PERCENT     float64 // 涨跌幅
	TCAP        float64 //总市值
	TURNOVER    float64 //成交额
	UPDOWN      float64 // 涨跌额
	VOLUME      float64 // 成交量
	WB          float64 //委比
	YESTCLOSE   float64 //昨收
	ZF          float64 //振幅
	TIME        string  //日期
}
