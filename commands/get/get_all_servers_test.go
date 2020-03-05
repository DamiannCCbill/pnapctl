package get

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/commands/get/servers"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func getAllServersSetup() {
	URL = "servers/"
}

func TestGetAllServersUnmarshallingError(test_framework *testing.T) {
	getAllServersSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, ioutil.NopCloser(bytes.NewBuffer([]byte{0, 5}))), nil)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingErrorBody, "get servers", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllServersShortSuccess(test_framework *testing.T) {
	getAllServersSetup()

	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	shortServer := generators.ConvertLongToShortServers(serverlist)
	PrepareMockPrinter(test_framework).
		PrintOutput(&shortServer, false, "get servers").
		Return(nil)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllServersLongSuccess(test_framework *testing.T) {
	getAllServersSetup()

	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(&serverlist, false, "get servers").
		Return(nil)

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllServersClientFailure(test_framework *testing.T) {
	getAllServersSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestError)

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllServersKeycloakFailure(test_framework *testing.T) {
	getAllServersSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestKeycloakError)

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllServersPrinterFailure(test_framework *testing.T) {
	getAllServersSetup()

	// generate servers
	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(&serverlist, false, "get servers").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}