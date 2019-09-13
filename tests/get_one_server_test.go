package tests

import (
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/servers"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetServerSetup(test_framework *testing.T) {
	URL = "servers/" + SERVERID
}

func TestGetServerShortSuccess(test_framework *testing.T) {
	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(server), &servers.ShortServer{}).
		Return(1, nil)

	servers.ID = SERVERID
	servers.Full = false
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestGetServerLongSuccess(test_framework *testing.T) {
	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(server), &servers.LongServer{}).
		Return(1, nil)

	servers.ID = SERVERID
	servers.Full = true
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestGetServerClientFailure(test_framework *testing.T) {
	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), testutil.TestError)

	servers.ID = SERVERID
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError("get server")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(server), &servers.ShortServer{}).
		Return(1, errors.New(testutil.PrinterUnmarshalErrorMsg))

	servers.ID = SERVERID
	servers.Full = false
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertErrorCode(test_framework, err, ctlerrors.Errormap[testutil.PrinterUnmarshalErrorMsg])
}