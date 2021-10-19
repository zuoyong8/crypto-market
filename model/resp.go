package model

//
type ExchangeInfoResp struct {
	Timezone        string      `json:"timezone"`
	ServerTime      int64       `json:"serverTime"`
	RateLimits      interface{} `json:"rateLimits"`
	ExchangeFilters interface{} `json:"exchangeFilters"`
	Symbols         interface{} `json:"symbols"`
}
