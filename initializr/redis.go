package initializr

import (
	"crypto-market/config"
	"crypto-market/db"

	"github.com/spf13/viper"
)

//
func InitRedis() {
	initRedis()
}

//
func CloseRedis() {
	closeRedis()
}

//关闭redis
func closeRedis() {
	switch config.GetRedisMode() {
	case "single":
		db.CloseRedisClient()
	case "cluster":
		db.CloseRedissClusterClient()
	}
}

//初始化redis
func initRedis() {
	switch config.GetRedisMode() {
	case "single":
		initSingleRedis()
	case "cluster":
		initRedisCluster()
	}
}

//初始化Redis单节点
func initSingleRedis() {
	cfgRedis := viper.Sub("redis")
	if cfgRedis == nil {
		panic("config not found redis")
	}
	db.InitRedis(cfgRedis)
}

//初始化redis集群
func initRedisCluster() {
	cfgRedisCluster := viper.Sub("rediscluster")
	if cfgRedisCluster == nil {
		panic("config not found redisCluster")
	}
	db.InitRedisCluster(cfgRedisCluster)
}
