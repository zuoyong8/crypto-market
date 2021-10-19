package router

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"

	"crypto-market/handler/binance"
	"crypto-market/handler/currency"
	"crypto-market/handler/indicator"
	"crypto-market/middleware"
)

//路由配置
func RouterEngine(zapLogger *zap.Logger) *gin.Engine {

	engine := gin.New()
	if app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("crypto-market-api"),
		newrelic.ConfigLicense("b96c746e9a168a6969315a21804178af1e00NRAL"),
	); err != nil {
		os.Exit(1)
	} else {
		engine.Use(nrgin.Middleware(app))
	}

	//gin中间件使用
	engine.Use(middleware.Cors())
	engine.Use(middleware.Secure())
	engine.Use(middleware.Language())
	engine.Use(middleware.Ginzap(zapLogger, time.RFC3339, true))
	engine.Use(middleware.RecoveryWithZap(zapLogger, true))

	apiGroup := engine.Group("api")
	apiV1Group := apiGroup.Group("v1")

	apiV3Group := apiGroup.Group("v3")
	//币安行情
	binanceApiGroup(apiV3Group)
	//
	currencyGroup(apiV1Group)

	//
	indicatorGroup(apiV1Group)

	return engine
}

//币安行情
func binanceApiGroup(rg *gin.RouterGroup) {
	//币种配置信息
	rg.GET("/exchangeInfo", binance.GetExchangeInfo)
	//24h价格变化
	rg.GET("/ticker/24hr", binance.GetTicker24hr)
}

//
func currencyGroup(rg *gin.RouterGroup) {
	//获取汇率
	rg.GET("/currency/rate", currency.GetCurrencyRate)
}

//
func indicatorGroup(rg *gin.RouterGroup) {
	//获取汇率
	rg.GET("/indicator/list", indicator.GetIndicators)
}
