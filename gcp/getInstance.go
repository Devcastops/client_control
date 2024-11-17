package gcp

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/devcastops/client_control/common"
)

func (client *Client) CreateGetInstance(Zone string) *GetInstance {
	start := &GetInstance{
		Client: client,
		Zone:   Zone,
	}
	return start
}

func (opts *GetInstance) GetInstance(name string) (*common.Instance, error) {
	fmt.Printf("GetInstance: %s\n", name)
	ctx := context.Background()

	c, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return &common.Instance{}, err
	}
	defer c.Close()

	res, err := c.Get(ctx, &computepb.GetInstanceRequest{
		Project:  opts.Client.Project,
		Zone:     opts.Zone,
		Instance: name,
	})
	if err != nil {
		return &common.Instance{}, err
	}
	// fmt.Println(res)
	return &common.Instance{
		Ip:   *res.NetworkInterfaces[0].NetworkIP,
		Name: name,
	}, nil
}
