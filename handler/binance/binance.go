package binance

import (
	"crypto-market/handler/util"
	"crypto-market/service/binance"

	"github.com/gin-gonic/gin"
)

//币种配置信息
func GetExchangeInfo(c *gin.Context) {
	util.ResponseJson(c, "", binance.ExchangeInfo())
}

//24h价格变化
func GetTicker24hr(c *gin.Context) {
	util.ResponseJson(c, "", binance.Ticker24r())
}
