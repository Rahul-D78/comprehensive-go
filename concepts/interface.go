package concepts

import (
	"fmt"
)

func NewNetInfClientProvider() *NetInfClient {
	return &NetInfClient{}
}

type NetworkInterface interface {
	// get network interface info - interface name, type, speed
	GetNetworkInterfaceInfo(serverAddr string) (string, error)
	// get network health using snap protocol - it might be healthy/unhealthy
	GetNetworkHealth(serverAddr string) (string, error)
}

type NetInfClient struct {
}

func (n *NetInfClient) GetNetworkInterfaceInfo(serverAddr string) (string, error) {
	if serverAddr == "" {
		return "", fmt.Errorf("server address is empty")
	}
	// implementation to get network interface info
	return fmt.Sprintf("Network interface info for server %s", serverAddr), nil
}

func (n *NetInfClient) GetNetworkHealth(serverAddr string) (string, error) {
	if serverAddr == "" {
		return "", fmt.Errorf("server address is empty")
	}
	// implementation to get network health
	return fmt.Sprintf("Network health for server %s is healthy", serverAddr), nil
}
