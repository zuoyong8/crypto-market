package config

import "github.com/spf13/viper"

//读取配置文件
func ReadConfig(path string) error {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

//
func GetRedisMode() string {
	return viper.Sub("application").GetString("redis_mode")
}
