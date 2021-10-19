package service

import (
	"crypto-market/common/xhttp"
	"crypto-market/config"
	"crypto-market/db"
	"log"
)

const (
	prefixUrl = "https://api.binance.com"
)

var (
	binanceApis = map[string]string{
		"ticker24hr":   prefixUrl + "/api/v3/ticker/24hr",
		"exchangeInfo": prefixUrl + "/api/v3/exchangeInfo",
		"klines":       prefixUrl + "/api/v3/klines",
		"depth":        prefixUrl + "/api/v3/depth",
		"recentTrade":  prefixUrl + "/api/v3/trades",
		"time":         prefixUrl + "/api/v3/time",
		"quoteTrades":  prefixUrl + "/api/quote/v1/trades",
		"quoteTickers": prefixUrl + "/api/quote/v1/broker/tickers",
		"aggTrades":    prefixUrl + "/api/v3/aggTrades",
	}
)

//
func SetDataToRedis(method string) {
	bs, err := xhttp.Get(binanceApis[method], nil)
	if err == nil && bs != nil {
		result := string(bs)
		switch config.GetRedisMode() {
		case "cluster":
			err = db.RedisClusterClient().Set(method, result, 0).Err()
		case "single":
			err = db.RedisClient().Set(method, result, 0).Err()
		}
	}
	if err != nil {
		log.Printf("error:%v", err)
	}
}
