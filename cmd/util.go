package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type StartParams struct {
	Config         string
	provider       string
	Name           string
	node_pool      string
	machine_type   string
	packer_channel string
}
type StopParams struct {
	Config   string
	provider string
	Name     string
}
type AutoStopParams struct {
	Config   string
	provider string
	Name     string
	DateTime time.Time
	DateSet  bool
}

func getStartParams(cmd *cobra.Command) (StartParams, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	Config, err := getFlagInput(os.Stdout, scanner, cmd, "config")
	if err != nil {
		return StartParams{}, err
	}
	fmt.Printf("config path set to %s\n", Config)

	provider, err := getFlagInput(os.Stdout, scanner, cmd, "provider")
	if err != nil {
		return StartParams{}, err
	}
	fmt.Printf("file location set to %s\n", provider)

	name, err := getFlagInput(os.Stdout, scanner, cmd, "name")
	if err != nil {
		return StartParams{}, err
	}
	fmt.Printf("file name set to %s\n", name)

	node_pool, err := getFlagDefault(cmd, "node_pool", name)
	if err != nil {
		return StartParams{}, err
	}
	fmt.Printf("file version set to %s\n", node_pool)
	machine_type, err := getFlagInput(os.Stdout, scanner, cmd, "machine_type")
	if err != nil {
		return StartParams{}, err
	}
	fmt.Printf("file output set to %s\n", machine_type)

	packer_channel, err := getFlagInput(os.Stdout, scanner, cmd, "packer_channel")
	if err != nil {
		return StartParams{}, err
	}
	fmt.Printf("file output set to %s\n", packer_channel)

	params := StartParams{
		Config:         Config,
		provider:       provider,
		Name:           name,
		node_pool:      node_pool,
		machine_type:   machine_type,
		packer_channel: packer_channel,
	}

	return params, nil
}

func getStopParams(cmd *cobra.Command) (StopParams, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	Config, err := getFlagInput(os.Stdout, scanner, cmd, "config")
	if err != nil {
		return StopParams{}, err
	}
	fmt.Printf("config path set to %s\n", Config)
	fmt.Printf("Language set to %s\n", Config)
	provider, err := getFlagInput(os.Stdout, scanner, cmd, "provider")
	if err != nil {
		return StopParams{}, err
	}
	fmt.Printf("file location set to %s\n", provider)

	name, err := getFlagInput(os.Stdout, scanner, cmd, "name")
	if err != nil {
		return StopParams{}, err
	}
	fmt.Printf("file name set to %s\n", name)

	params := StopParams{
		Config:   Config,
		provider: provider,
		Name:     name,
	}

	return params, nil
}

func getAutoStopParams(cmd *cobra.Command) (AutoStopParams, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	Config, err := getFlagInput(os.Stdout, scanner, cmd, "config")
	if err != nil {
		return AutoStopParams{}, err
	}
	fmt.Printf("config path set to %s\n", Config)

	provider, err := getFlagInput(os.Stdout, scanner, cmd, "provider")
	if err != nil {
		return AutoStopParams{}, err
	}
	fmt.Printf("file location set to %s\n", provider)

	name, err := getFlagInput(os.Stdout, scanner, cmd, "name")
	if err != nil {
		return AutoStopParams{}, err
	}
	fmt.Printf("file name set to %s\n", name)

	// Time calc

	autostoptime, err := getFlagInput(os.Stdout, scanner, cmd, "autostoptime")
	if err != nil {
		return AutoStopParams{}, err
	}
	fmt.Printf("file autostoptime set to %s\n", autostoptime)

	autostopdate, err := cmd.Flags().GetString("autostopdate")
	if err != nil {
		return AutoStopParams{}, err
	}
	var t time.Time
	var dateset bool
	if autostopdate == "" {
		dateset = false
		t, err = time.Parse("15:04", autostoptime)
		if err != nil {
			return AutoStopParams{}, err
		}
	} else {
		dateset = true
		t, err = time.Parse("2/1 15:04", fmt.Sprintf("%s %s", autostopdate, autostoptime)) //(day/month hour/minute)
		if err != nil {
			return AutoStopParams{}, err
		}
	}

	params := AutoStopParams{
		Config:   Config,
		provider: provider,
		Name:     name,
		DateTime: t,
		DateSet:  dateset,
	}

	return params, nil
}

func getFlagInput(out io.Writer, scanner *bufio.Scanner, cmd *cobra.Command, flagName string) (string, error) {
	flag, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return "", err
	}
	if flag != "" {
		return flag, nil
	}
	fmt.Fprintf(out, "what %s is your server: ", flagName)

	scanner.Scan()

	return scanner.Text(), nil
}

func getFlagDefault(cmd *cobra.Command, flagName string, defaultValue string) (string, error) {
	flag, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return "", err
	}
	if flag != "" {
		return flag, nil
	}
	return defaultValue, nil
}
