package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type PrivateNetworkConfiguration struct {
	GatewayAddress    *string                `yaml:"gatewayAddress" json:"gatewayAddress"`
	ConfigurationType *string                `yaml:"configurationType" json:"configurationType"`
	PrivateNetworks   []ServerPrivateNetwork `yaml:"privateNetworks" json:"privateNetworks"`
}

func (privateNetConf *PrivateNetworkConfiguration) toSdk() *bmcapisdk.PrivateNetworkConfiguration {
	if privateNetConf == nil {
		return nil
	}

	return &bmcapisdk.PrivateNetworkConfiguration{
		GatewayAddress:    privateNetConf.GatewayAddress,
		ConfigurationType: privateNetConf.ConfigurationType,
		PrivateNetworks:   mapServerPrivateNetworkListToSdk(privateNetConf.PrivateNetworks),
	}
}

func privateNetworkConfigurationFromSdk(privateNetworkConfnfiguration *bmcapisdk.PrivateNetworkConfiguration) *PrivateNetworkConfiguration {
	if privateNetworkConfnfiguration == nil {
		return nil
	}

	return &PrivateNetworkConfiguration{
		GatewayAddress:    privateNetworkConfnfiguration.GatewayAddress,
		ConfigurationType: privateNetworkConfnfiguration.ConfigurationType,
		PrivateNetworks:   serverPrivateNetworkListFromSdk(privateNetworkConfnfiguration.PrivateNetworks),
	}
}
