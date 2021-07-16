package common

const (
	GET  = "GET"
	POST = "POST"
)

var (
	//
	RateApiUrl = map[string]string{
		"it120ApiUrl": "https://api.it120.cc/gooking/forex/rate?fromCode=%v&toCode=%v",
		"sinaJsUrl":   "https://hq.sinajs.cn/rn=%vlist=fx_s%v%v",
	}

	Errors = map[int64]string{
		1000001: "参数错误或请求格式不正确",
	}
)
