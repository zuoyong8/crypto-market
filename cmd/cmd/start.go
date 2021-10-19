package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"crypto-market/common/xlog"
	"crypto-market/config"
	"crypto-market/initializr"
	"crypto-market/router"
)

var (
	configPath string
	logPath    string
	port       int
	StartCmd   = &cobra.Command{
		Use:     "run",
		Short:   "start run server",
		Example: "server run -port 8888 -config /etc/config/config.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

//初始化
func init() {
	StartCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config/config.yml", "config file path")
	StartCmd.PersistentFlags().StringVarP(&logPath, "log", "l", "./log", "log file path")
}

//显示终端信息
func usage() {
	usageStr := `
	-------------------------------------------------------------------------------------------------
					*        *******   *        *   *  *
					*           *      *  *     *   * *
					*           *      *    *   *   *
					*           *      *      * *   * *
					*******  ********  *        *   *   *

				   welcome to use crypto market command
					    copyright @zorro
	-------------------------------------------------------------------------------------------------
	`
	fmt.Printf("%s\n", usageStr)
}

//初始化各种服务
func setup() {
	//加载配置文件
	if err := config.ReadConfig("./config/config.toml"); err != nil {
		panic("read config error")
	}
	//初始化redis
	initializr.InitRedis()
}

//运行
func run() error {
	//初始化日志
	zapLog := xlog.Init("server", logPath)

	//获取路由
	engine := router.RouterEngine(zapLog)
	//启动
	engine.Run(fmt.Sprintf(":%v", port))
	return nil
}
