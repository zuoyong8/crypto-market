package usd

import (
	"crypto-market/common"
	"crypto-market/common/xhttp"
	"encoding/json"
	"fmt"
)

//
type it120Currency struct {
	url string
}

//
type it120Result struct {
	Code int `json:"code"`
	Data struct {
		Rate     float64 `json:"rate"`
		ToCode   float64 `json:"toCode"`
		FromCode float64 `json:"fromCode"`
	} `json:"data"`
	Msg string `json:"msg"`
}

//
func newIt120Currency() currency {
	return &it120Currency{
		url: common.RateApiUrl["it120ApiUrl"],
	}
}

//
func (c *it120Currency) GetCurrencyRate(from, to string) float64 {
	bs, err := xhttp.Get(fmt.Sprintf(c.url, from, to), nil)
	if err != nil || bs == nil {
		return 0.0
	}
	//
	var r it120Result
	err = json.Unmarshal(bs, &r)
	if err == nil {
		return r.Data.Rate
	}
	return 0.0
}
