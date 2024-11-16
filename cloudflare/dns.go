package cloudflare

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/devcastops/client_control/config"
)

func UpdateDNS(config config.Config, service, ip string) error {
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		return err
	}
	ctx := context.Background()

	zoneId := config.Cloudflare.ZoneId
	fqdn := fmt.Sprintf("%s.devcastops.com", service)
	dnsRecordID := ""
	records, _, err := api.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneId), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i := range records {
		if records[i].Name == fmt.Sprintf("%s.devcastops.com", service) {
			dnsRecordID = records[i].ID
			break
		}
	}
	if dnsRecordID == "" {
		createRes, err := api.CreateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneId), cloudflare.CreateDNSRecordParams{
			Name:    fqdn,
			Content: ip,
			Type:    "A",
		})
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(createRes)
	} else {
		updateRes, err := api.UpdateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneId), cloudflare.UpdateDNSRecordParams{
			ID:      dnsRecordID,
			Name:    fqdn,
			Content: ip,
			Type:    "A",
		})
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(updateRes)
	}
	return nil
}
