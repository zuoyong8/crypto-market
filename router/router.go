package router

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"

	"crypto-market/handler/currency"
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
	currencyGroup(apiV1Group)

	// apiV2Group := engine.Group("v2")
	return engine
}

//
func currencyGroup(rg *gin.RouterGroup) {
	//获取汇率
	rg.GET("/currency/rate", currency.GetCurrencyRate)
}
