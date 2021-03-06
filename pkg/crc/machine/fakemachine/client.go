package fakemachine

import (
	"errors"

	"github.com/code-ready/crc/pkg/crc/machine"
	"github.com/code-ready/crc/pkg/crc/network"
)

func NewClient() *Client {
	return &Client{}
}

type Client struct {
	Failing bool
}

func (c *Client) Delete(deleteConfig machine.DeleteConfig) (machine.DeleteResult, error) {
	return machine.DeleteResult{}, errors.New("not implemented")
}

func (c *Client) GetConsoleURL(consoleConfig machine.ConsoleConfig) (machine.ConsoleResult, error) {
	return machine.ConsoleResult{}, errors.New("not implemented")
}

func (c *Client) GetProxyConfig(machineName string) (*network.ProxyConfig, error) {
	return nil, errors.New("not implemented")
}

func (c *Client) IP(ipConfig machine.IPConfig) (machine.IPResult, error) {
	return machine.IPResult{}, errors.New("not implemented")
}

func (c *Client) PowerOff(powerOff machine.PowerOffConfig) (machine.PowerOffResult, error) {
	return machine.PowerOffResult{}, errors.New("not implemented")
}

func (c *Client) Start(startConfig machine.StartConfig) (machine.StartResult, error) {
	return machine.StartResult{}, errors.New("not implemented")
}

func (c *Client) Stop(stopConfig machine.StopConfig) (machine.StopResult, error) {
	return machine.StopResult{}, errors.New("not implemented")
}

func (c *Client) Status(statusConfig machine.ClusterStatusConfig) (machine.ClusterStatusResult, error) {
	if c.Failing {
		return machine.ClusterStatusResult{
			Success: false,
			Error:   "broken",
		}, errors.New("broken")
	}
	return machine.ClusterStatusResult{
		Name:             "crc",
		CrcStatus:        "Running",
		OpenshiftStatus:  "Running",
		OpenshiftVersion: "4.5.1",
		DiskUse:          10_000_000_000,
		DiskSize:         20_000_000_000,
		Success:          true,
	}, nil
}

func (c *Client) Exists(name string) (bool, error) {
	return true, nil
}
