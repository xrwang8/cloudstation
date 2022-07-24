package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	version      bool
	ossProvider  string
	aliAccessKey string
	aliSecretKey string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cloud-station-cli",
	Short: "cloud-station-cli 文件中转服务",
	Long:  `cloud-station-cli ...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("0.0.1")
			return nil
		}
		return errors.New("no flags find")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&ossProvider, "oss_provider", "p", "aliyun", "the oss provider [aliyun/qcloud]")
	RootCmd.PersistentFlags().StringVarP(&aliAccessKey, "ali_access_id", "i", "", "the ali oss access id")
	RootCmd.PersistentFlags().StringVarP(&aliSecretKey, "ali_secret_key", "k", "", "the ali oss access key")
	RootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "the cloud-station-cli version")
}
