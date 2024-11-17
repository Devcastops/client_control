package cmd

import (
	"fmt"

	"github.com/devcastops/client_control/cloudflare"
	"github.com/devcastops/client_control/common"
	"github.com/devcastops/client_control/config"
	"github.com/devcastops/client_control/gcp"
	"github.com/devcastops/client_control/packer"
	"github.com/devcastops/client_control/webhook"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run to start the instance",
	RunE:  start,
}

func init() {

	startCmd.PersistentFlags().StringP("provider", "p", "GCP", "cloud provider to use (GCP)")
	startCmd.PersistentFlags().StringP("name", "n", "test", "name to give the instance")
	startCmd.PersistentFlags().StringP("node_pool", "l", "", "The node_pool to deploy the client in")
	startCmd.PersistentFlags().StringP("machine_type", "m", "e2-standard-2", "The machine_type to use")
	startCmd.PersistentFlags().StringP("packer_channel", "i", "live", "packer channel to pull from")
	startCmd.PersistentFlags().StringP("config", "c", "config.json", "location of config file")
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) error {
	params, err := getStartParams(cmd)
	if err != nil {
		return err
	}
	config, err := config.Load(params.Config)
	if err != nil {
		return err
	}
	var info *common.Instance
	switch params.provider {
	case "GCP":
		info, err = startGCP(params, config)
		if err != nil {
			return err
		}
	}
	err = cloudflare.UpdateDNS(config, info.Name, info.Ip)
	if err != nil {
		return err
	}
	err = webhook.SendMessage(config.Webhook.Url, fmt.Sprintf("Started instance:\nName: %s\nNode pool: %s\nIP: %s", params.Name, params.node_pool, info.Ip))

	if err != nil {
		return err
	}

	return nil

}

func startGCP(params StartParams, config config.Config) (*common.Instance, error) {
	image, err := packer.GetPackerImageGCP(params.packer_channel, config)
	if err != nil {
		return &common.Instance{}, err
	}

	client := gcp.CreateClient(config.GCP.Project)

	err = client.CreateStartInstance(
		config.GCP.Compute.Zone,
		"IPV4_IPV6",
		params.machine_type,
		config.GCP.Compute.Subnetwork,
		config.GCP.Compute.ServiceAccount,
		config.GCP.Compute.DiskSize,
		[]string{
			"https://www.googleapis.com/auth/devstorage.read_only",
			"https://www.googleapis.com/auth/logging.write",
			"https://www.googleapis.com/auth/monitoring.write",
			"https://www.googleapis.com/auth/servicecontrol",
			"https://www.googleapis.com/auth/service.management.readonly",
			"https://www.googleapis.com/auth/trace.append",
		},
		[]string{
			"nomadclient",
			params.node_pool,
		},
		map[string]string{
			"nodepool": params.node_pool,
		},
		map[string]string{
			"startup-script": fmt.Sprintf(`
#! /bin/bash 
echo '{"client":{"node_pool":"%s"}}' | jq -rM '.'>/etc/nomad.d/node_pool.hcl
echo '{"client":{"servers": %s, "artifact":{"decompression_file_count_limit":0}}}' | jq -rM '.'>/etc/nomad.d/client_build.hcl
systemctl restart nomad
      `, params.node_pool, config.Nomad.ServerIPs),
			"enable-oslogin": "TRUE",
		},
	).StartInstance(image, params.Name)
	if err != nil {
		return &common.Instance{}, err
	}

	return client.CreateGetInstance(config.GCP.Compute.Zone).GetInstance(params.Name)

}
