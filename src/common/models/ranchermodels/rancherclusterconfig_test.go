package ranchermodels

import (
	"fmt"
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
)

func TestRancherClusterConfigToSdk(test_framework *testing.T) {
	rancherClusterConfig := GeneratecliRancherClusterConfig()
	sdkRancherClusterConfig := *rancherClusterConfig.ToSdk()

	assertEqualRancherClusterConfig(test_framework, rancherClusterConfig, sdkRancherClusterConfig)
}

func TestRancherClusterConfigFromSdk(test_framework *testing.T) {
	sdkRancherClusterConfig := GeneratesdkRancherClusterConfig()
	rancherClusterConfig := *RancherClusterConfigFromSdk(&sdkRancherClusterConfig)

	assertEqualRancherClusterConfig(test_framework, rancherClusterConfig, sdkRancherClusterConfig)
}

func TestRancherClusterConfigToTableString_nilConfig(test_framework *testing.T) {
	result := RancherClusterConfigToTableString(nil)
	assert.Equal(test_framework, "", result)
}

func TestNodePoolsToTableStrings_withClusterConfig(test_framework *testing.T) {
	sdkModel := GeneratesdkRancherClusterConfig()

	result := RancherClusterConfigToTableString(&sdkModel)

	assert.Equal(test_framework, result, generateClusterConfigResultString(&sdkModel))
}

func generateClusterConfigResultString(config *ranchersdk.RancherClusterConfig) string {
	return fmt.Sprintf("Token: %s, Domain: %s", *config.Token, *config.ClusterDomain)
}

func assertEqualRancherClusterConfig(test_framework *testing.T, cliRancherClusterConfig RancherClusterConfig, sdkRancherClusterConfig ranchersdk.RancherClusterConfig) {
	assert.Equal(test_framework, cliRancherClusterConfig.Token, sdkRancherClusterConfig.Token)
	assert.Equal(test_framework, cliRancherClusterConfig.TlsSan, sdkRancherClusterConfig.TlsSan)
	assert.Equal(test_framework, cliRancherClusterConfig.EtcdSnapshotScheduleCron, sdkRancherClusterConfig.EtcdSnapshotScheduleCron)
	assert.Equal(test_framework, cliRancherClusterConfig.EtcdSnapshotRetention, sdkRancherClusterConfig.EtcdSnapshotRetention)
	assert.Equal(test_framework, cliRancherClusterConfig.NodeTaint, sdkRancherClusterConfig.NodeTaint)
	assert.Equal(test_framework, cliRancherClusterConfig.ClusterDomain, sdkRancherClusterConfig.ClusterDomain)

	if !assertNilEquality(test_framework, "Certificates", cliRancherClusterConfig.Certificates, sdkRancherClusterConfig.Certificates) {
		assertEqualRancherClusterCertificates(test_framework, *cliRancherClusterConfig.Certificates, *sdkRancherClusterConfig.Certificates)
	}
}
