package cmd

import (
	"errors"
	"fmt"
	"openiam/cmd/migrate"
	"openiam/cmd/server"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "openiam",
	Short:        "openiam",
	SilenceUsage: true,
	Long:         `openiam`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}

func tip() {
	usageStr := `欢迎使用 openiam 管理系统，可以使用 -h 查看命令帮助`
	fmt.Printf("%s\n", usageStr)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
