package cmd

import (
	"fmt"
	"os"

	"github.com/devcastops/client_control/config"
	"github.com/devcastops/client_control/webhook"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client_control",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.Load("config.json")
		fmt.Println(config)
		return err
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		config, _ := config.Load("config.json")
		webhook.SendMessage(config.Webhook.Url, err.Error())
		os.Exit(1)
	}
}
