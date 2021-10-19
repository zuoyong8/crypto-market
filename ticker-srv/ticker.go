package main

import (
	"crypto-market/config"
	"crypto-market/initializr"
	"crypto-market/ticker-srv/service"

	"github.com/jasonlvhit/gocron"
)

//
func main() {
	//加载配置文件
	if err := config.ReadConfig("./config/config.toml"); err != nil {
		panic("read config error")
	}
	//初始化redis
	initializr.InitRedis()
	// defer initializr.CloseRedis()

	//24h价格变化
	gocron.Every(1).Minutes().Do(service.SetDataToRedis, "ticker24hr")

	//币种配置信息
	gocron.Every(2).Minutes().Do(service.SetDataToRedis, "exchangeInfo")

	//深度
	// gocron.Every(34).Seconds().Do(service.SetDataToRedis, "depth")

	//归集成交历史(近期)
	// gocron.Every(2).Minutes().Do(service.SetDataToRedis, "aggTrades")

	//
	// gocron.Every(43).Seconds().Do(service.SetDataToRedis, "time")

	//
	// gocron.Every(43).Seconds().Do(service.SetDataToRedis, "quoteTrades")

	//
	// gocron.Every(43).Seconds().Do(service.SetDataToRedis, "quoteTickers")

	// gocron.Every(34).Seconds().Do(service.SetDataToRedis, "klines")

	<-gocron.Start()

}
