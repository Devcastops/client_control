package gcp

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
)

func (client *Client) CreateStopInstance(Zone string) *StopInstance {
	start := &StopInstance{
		Client: client,
		Zone:   Zone,
	}
	return start
}

func (opts *StopInstance) StopInstance(name string) error {
	fmt.Printf("Stopping: %s\n", name)
	ctx := context.Background()

	c, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.Delete(ctx, &computepb.DeleteInstanceRequest{
		Project:  opts.Client.Project,
		Zone:     opts.Zone,
		Instance: name,
	})
	if err != nil {
		return err
	}
	fmt.Println("Deleted")
	return nil
}
