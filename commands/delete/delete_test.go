package delete

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	delete "phoenixnap.com/pnap-cli/commands/delete/server"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestDeleteServerSuccess(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(SERVERID).
		Return(generators.GenerateDeleteResult(), WithResponse(200, nil), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerNotFound(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(SERVERID).
		Return(bmcapi.DeleteResult{}, WithResponse(404, nil), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	expectedMessage := "Command 'delete server' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteServerError(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(SERVERID).
		Return(bmcapi.DeleteResult{}, WithResponse(500, nil), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	expectedMessage := "Command 'delete server' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteServerClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(SERVERID).
		Return(bmcapi.DeleteResult{}, nil, testutil.TestError)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteServerKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(SERVERID).
		Return(bmcapi.DeleteResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
