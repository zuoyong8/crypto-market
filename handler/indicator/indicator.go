package indicator

import (
	"crypto-market/handler/util"
	"crypto-market/model"
	"crypto-market/service/indicator"

	"github.com/gin-gonic/gin"
)

//获取说明
func GetIndicators(c *gin.Context) {
	//
	var req model.IndicatorReq

	if err := c.ShouldBindQuery(&req); err != nil {
		util.ResponseErrorJson(c, 1000001)
		return
	}
	//
	util.ResponseJson(c, "", indicator.GetIndicators(req.Name))
}
