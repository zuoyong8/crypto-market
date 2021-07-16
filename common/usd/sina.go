package usd

import (
	"crypto-market/common"
	"crypto-market/common/xhttp"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//
type sinaCurrency struct {
	url string
}

//
func newSinaCurrency() currency {
	return &sinaCurrency{
		url: common.RateApiUrl["sinaJsUrl"],
	}
}

//
func (s *sinaCurrency) GetCurrencyRate(from, to string) float64 {
	milliSecond := time.Now().UnixNano() / 1e6
	header := map[string]string{
		"Content-Type": "text/plain; charset=UTF-8",
	}
	bs, err := xhttp.Get(fmt.Sprintf(s.url, milliSecond, strings.ToLower(from), strings.ToLower(to)), header)
	if err != nil || bs == nil {
		return 0.0
	}

	fmt.Println(string(bs))

	result := string(bs)
	ss := strings.Split(result, ",")
	if len(ss) >= 3 {
		f, _ := strconv.ParseFloat(ss[2], 10)
		return f
	}

	return 0.0
}
