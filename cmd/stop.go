package cmd

import (
	"github.com/devcastops/client_control/config"
	"github.com/devcastops/client_control/gcp"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Run to stop the instance",
	RunE:  stop,
}

func init() {

	stopCmd.PersistentFlags().StringP("provider", "p", "GCP", "cloud provider to use (GCP)")
	stopCmd.PersistentFlags().StringP("name", "n", "test", "name to give the instance")
	stopCmd.PersistentFlags().StringP("config", "c", "config.json", "location of config file")
	rootCmd.AddCommand(stopCmd)
}

func stop(cmd *cobra.Command, args []string) error {
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
		err = stopGCP(params, config)
		if err != nil {
			return err
		}
	}

	return nil

}

func stopGCP(params StopParams, config config.Config) error {
	return gcp.CreateClient(config.GCP.Project).CreateStopInstance(config.GCP.Compute.Zone).StopInstance(params.Name)
}
