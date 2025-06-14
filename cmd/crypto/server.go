package crypto

import (
	"fmt"
	"github.com/lanyulei/toolkit/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"openops/pkg/crypto"
	"os"
)

/*
  @Author : lanyulei
  @Desc :
*/

var (
	configYml string
	secret    string
	value     string
	StartCmd  = &cobra.Command{
		Use:          "crypto",
		Short:        "symmetric encryption",
		Example:      "openops crypto -k xxx -v xxx",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func setup() {
	// 加载配置文件
	viper.SetConfigFile(configYml) // 指定配置文件
	err := viper.ReadInConfig()    // 读取配置信息
	if err != nil {                // 读取配置信息失败
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	// 日志配置
	logger.Setup(
		viper.GetString(`log.level`),
		viper.GetString(`log.path`),
		viper.GetInt(`log.maxSize`),
		viper.GetBool(`log.localtime`),
		viper.GetBool(`log.compress`),
		viper.GetBool(`log.console`),
		nil,
	)
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "specify the profile to start the service")
	_ = StartCmd.MarkPersistentFlagRequired("config")

	StartCmd.PersistentFlags().StringVarP(&secret, "secret", "s", "", "secret required for symmetric encryption")

	StartCmd.PersistentFlags().StringVarP(&value, "value", "v", "", "the string to be encrypted")
	_ = StartCmd.MarkPersistentFlagRequired("value")
}

func run() (err error) {
	var (
		result string
	)

	if secret == "" {
		secret = os.Getenv("OPENOPS_CRYPTO_AES_KEY")
	}

	if secret == "" {
		logger.Fatalf("secret is required")
	}

	result, err = crypto.AesEncryptCBC([]byte(secret), []byte(value))
	if err != nil {
		return
	}

	fmt.Printf("After encryption: %s\n", result)
	return
}
