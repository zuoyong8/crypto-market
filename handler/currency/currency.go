package currency

import (
	"crypto-market/handler/util"
	"crypto-market/model"
	"crypto-market/service/currency"
	"fmt"

	"github.com/gin-gonic/gin"
)

//获取汇率
func GetCurrencyRate(c *gin.Context) {
	var req model.CurrencyRateReq

	if err := c.ShouldBindQuery(&req); err != nil {
		util.ResponseErrorJson(c, 1000001)
		return
	}

	//
	if req.From == "" {
		req.From = "CNY"
	}
	if req.To == "" {
		req.To = "USD"
	}

	//
	util.ResponseJson(c, "", gin.H{
		"currency": fmt.Sprintf("%v-%v", req.From, req.To),
		"rate":     currency.GetCurrencyRate(&req),
	})

}
