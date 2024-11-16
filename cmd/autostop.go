package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/devcastops/client_control/config"
	"github.com/devcastops/client_control/gcp"
	"github.com/spf13/cobra"
)

var autoStopCmd = &cobra.Command{
	Use:   "autostop",
	Short: "Run to autoStop the instance",
	RunE:  autoStop,
}

func init() {

	autoStopCmd.PersistentFlags().StringP("provider", "p", "GCP", "cloud provider to use (GCP)")
	autoStopCmd.PersistentFlags().StringP("name", "n", "test", "name to give the instance")
	autoStopCmd.PersistentFlags().StringP("autostoptime", "t", "00:00", "time to stop the server")
	autoStopCmd.PersistentFlags().StringP("autostopdate", "d", "", "date to stop the server")
	autoStopCmd.PersistentFlags().StringP("config", "c", "config.json", "location of config file")
	rootCmd.AddCommand(autoStopCmd)
}

func autoStop(cmd *cobra.Command, args []string) error {
	params, err := getAutoStopParams(cmd)
	if err != nil {
		return err
	}
	config, err := config.Load(params.Config)
	if err != nil {
		return err
	}
	switch params.provider {
	case "GCP":
		err = autoStopGCP(params, config)
		if err != nil {
			return err
		}
	}

	return nil

}

func autoStopGCP(params AutoStopParams, config config.Config) error {
	killdt := params.DateTime
	nowdt := time.Now()
	if params.DateSet {
		killdt = killdt.AddDate(nowdt.Year(), 0, 0)
	} else {
		killdt = killdt.AddDate(nowdt.Year(), int(nowdt.Month())-1, nowdt.Day()-1)
	}
	if killdt.Before(nowdt) {
		killdt = killdt.AddDate(0, 0, 1)
	}
	fmt.Printf("Now:	%s\nNeeded:	%s\n", nowdt, killdt)
	getInstance := gcp.CreateClient(config.GCP.Project).CreateGetInstance(config.GCP.Compute.Zone)
	for killdt.After(nowdt) {
		time.Sleep(time.Second)
		nowdt = time.Now()
		fmt.Printf("Now:	%s\nNeeded:	%s\n", nowdt, killdt)
		_, err := getInstance.GetInstance(params.Name)
		if err != nil {
			if strings.Contains(err.Error(), "Error 404") {
				return nil
			}
			return err
		}
	}
	return gcp.CreateClient(config.GCP.Project).CreateStopInstance(config.GCP.Compute.Zone).StopInstance(params.Name)
}
