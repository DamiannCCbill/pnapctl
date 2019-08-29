package tests

import (
	"bytes"
	"testing"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweron"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

func TestPowerOnSetup(t *testing.T) {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/power-on"
}

func TestPowerOnServerSuccess(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{SERVERID})

	if err != nil {
		test_framework.Errorf("Expected no error. Instead got %s", err.Error())
	}
}

func TestPowerOnServerNotFound(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{SERVERID})

	if err.Error() != "404" {
		test_framework.Errorf("Expected '404 NOT FOUND' error. Instead got %s", err.Error())
	}
}

func TestPowerOnServerError(test_framework *testing.T) {
	bmcErr := ctlerrors.BMCError{
		Message:          "Something went wrong!",
		ValidationErrors: []string{},
	}

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(bmcErr)), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{SERVERID})

	if err.Error() != "500" {
		test_framework.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", err.Error())
	}
}
