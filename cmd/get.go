package cmd

import (
	"github.com/devcastops/client_control/config"
	"github.com/devcastops/client_control/gcp"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Run to get the instance",
	RunE:  get,
}

func init() {

	getCmd.PersistentFlags().StringP("provider", "p", "GCP", "cloud provider to use (GCP)")
	getCmd.PersistentFlags().StringP("name", "n", "test", "name to give the instance")
	getCmd.PersistentFlags().StringP("config", "c", "config.json", "location of config file")
	rootCmd.AddCommand(getCmd)
}

func get(cmd *cobra.Command, args []string) error {
	params, err := getStopParams(cmd)
	if err != nil {
		return err
	}
	config, err := config.Load(params.Config)
	if err != nil {
		return err
	}
	switch params.provider {
	case "GCP":
		err = getGCP(params, config)
		if err != nil {
			return err
		}
	}

	return nil

}

func getGCP(params StopParams, config config.Config) error {
	_, err := gcp.CreateClient(config.GCP.Project).CreateGetInstance(config.GCP.Compute.Zone).GetInstance(params.Name)

	return err
}
