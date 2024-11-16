package packer

import (
	"fmt"
	"log"

	"github.com/devcastops/client_control/config"
	"github.com/go-openapi/strfmt"
	packer "github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2023-01-01/client/packer_service"
	sdkconfig "github.com/hashicorp/hcp-sdk-go/config"
	"github.com/hashicorp/hcp-sdk-go/httpclient"
)

func GetPackerImageGCP(channel string, config config.Config) (string, error) {
	fmt.Printf("Getting image: %s\n", channel)
	hcpConfig, err := sdkconfig.NewHCPConfig(
		sdkconfig.FromEnv(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Construct HTTP client config
	httpclientConfig := httpclient.Config{
		HCPConfig: hcpConfig,
	}

	// Initialize SDK http client
	cl, err := httpclient.New(httpclientConfig)
	if err != nil {
		log.Fatal(err)
	}
	packerClient := packer.New(cl, strfmt.Default)

	channelParams := packer.NewPackerServiceGetChannelParams()
	channelParams.LocationOrganizationID = config.Packer.OrganizationID
	channelParams.LocationProjectID = config.Packer.ProjectID
	channelParams.BucketName = config.Packer.BucketName
	channelParams.ChannelName = channel
	c, _ := packerClient.PackerServiceGetChannel(channelParams, cl.DefaultAuthentication)

	return c.Payload.Channel.Version.Builds[0].Labels["self_link"], nil

}
