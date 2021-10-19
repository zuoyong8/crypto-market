package binance

import (
	"crypto-market/db"
	"crypto-market/model"
	"encoding/json"
	"reflect"
	"unsafe"
)

//
func ExchangeInfo() *model.ExchangeInfoResp {
	result := db.RedisClusterClient().Get("exchangeInfo").Val()
	if result != "" {
		var info model.ExchangeInfoResp
		if err := json.Unmarshal(string2Bytes(result), &info); err == nil {
			return &info
		}
	}
	return nil
}

//
func Ticker24r() interface{} {
	result := db.RedisClusterClient().Get("ticker24hr").Val()
	if result != "" {
		var info interface{}
		if err := json.Unmarshal(string2Bytes(result), &info); err == nil {
			return &info
		}
	}
	return nil

}

//
func string2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
