package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"openiam/common/router"
	"openiam/pkg/config"
	"openiam/pkg/tools"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
	"github.com/lanyulei/toolkit/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "start API server",
		Example:      "openiam server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	// 加载配置文件
	config.Setup(configYml)

	// 日志配置
	logger.Setup(
		viper.GetString(`log.level`),
		viper.GetString(`log.path`),
		viper.GetInt(`log.maxSize`),
		viper.GetBool(`log.localTime`),
		viper.GetBool(`log.compress`),
		viper.GetBool(`log.console`),
		nil,
	)

	// 数据库配置
	db.Setup(
		viper.GetString("db.type"),
		viper.GetString("db.dsn"),
		viper.GetInt("db.maxIdleConn"),
		viper.GetInt("db.maxOpenConn"),
		viper.GetInt("db.connMaxLifetime"),
	)
}

func run() (err error) {
	if viper.GetString("server.mode") == config.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}

	// 路由引擎实例
	r := gin.Default()
	router.Setup(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")),
		Handler: r,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func(cancel context.CancelFunc) {
		cancel()

		// 关闭 redis 连接
		redis.StopChRedis()
	}(cancel)

	go func() {
		// 服务连接
		if viper.GetBool("ssl.enable") {
			if err := srv.ListenAndServeTLS(viper.GetString("ssl.pem"), viper.GetString("ssl.key")); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Fatal("listen: ", err)
			}
		}
	}()

	fmt.Println("\nServer run at:")
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", viper.GetInt("server.port"))
	fmt.Printf("-  Network: http://%s:%d/ \r\n", tools.GetLocalHost(), viper.GetInt("server.port"))
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n\n", time.Now().Format("2006-01-02 15:04:05.000"))
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	<-quit

	fmt.Printf("%s Shutdown Server ... \r\n", time.Now().Format("2006-01-02 15:04:05"))

	if err = srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}

	return nil
}
