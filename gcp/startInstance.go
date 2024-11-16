package gcp

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
)

func (client *Client) CreateStartInstance(
	Zone,
	StackType,
	MachineType,
	Subnetwork,
	ServiceAccount string,
	DiskSize int,
	Scopes,
	Tags []string,
	Lables,
	Metadata map[string]string,
) *StartInstance {
	start := &StartInstance{
		Client:         client,
		Zone:           Zone,
		StackType:      StackType,
		MachineType:    fmt.Sprintf("zones/%s/machineTypes/%s", Zone, MachineType),
		Subnetwork:     fmt.Sprintf("projects/devcastops/regions/%s/subnetworks/%s", Zone[:len(Zone)-2], Subnetwork),
		ServiceAccount: ServiceAccount,
		Scopes:         Scopes,
		DiskSize:       int64(DiskSize),
		Lables:         Lables,
		Metadata:       Metadata,
		Tags:           Tags,
	}
	return start
}

func (opts *StartInstance) StartInstance(image, name string) error {
	fmt.Printf("Starting GCP instance with image: %s\n", image)
	ctx := context.Background()

	c, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	trueobj := true

	metadata := []*computepb.Items{}
	for k, v := range opts.Metadata {
		metadata = append(metadata, &computepb.Items{Key: &k, Value: &v})
	}

	opt := computepb.InsertInstanceRequest{
		Project: opts.Client.Project,
		Zone:    opts.Zone,
		InstanceResource: &computepb.Instance{
			Name:        &name,
			MachineType: &opts.MachineType,
			NetworkInterfaces: []*computepb.NetworkInterface{
				{
					StackType:  &opts.StackType,
					Subnetwork: &opts.Subnetwork,
				},
			},
			ServiceAccounts: []*computepb.ServiceAccount{{Email: &opts.ServiceAccount, Scopes: opts.Scopes}},
			Tags:            &computepb.Tags{Items: opts.Tags},
			Disks: []*computepb.AttachedDisk{
				{
					AutoDelete: &trueobj,
					Boot:       &trueobj,
					DeviceName: &name,
					InitializeParams: &computepb.AttachedDiskInitializeParams{
						DiskSizeGb:  &opts.DiskSize,
						SourceImage: &image,
					},
					DiskSizeGb: &opts.DiskSize,
				},
			},
			ShieldedInstanceConfig: &computepb.ShieldedInstanceConfig{EnableIntegrityMonitoring: &trueobj, EnableVtpm: &trueobj},
			Labels:                 opts.Lables,
			Metadata:               &computepb.Metadata{Items: metadata},
		},
	}
	res, err := c.Insert(ctx, &opt)
	if err != nil {
		return err
	}
	res.Wait(ctx)
	fmt.Println("Created")
	return nil
}
