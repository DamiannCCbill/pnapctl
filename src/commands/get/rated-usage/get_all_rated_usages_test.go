package rated_usage

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/models/tables"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllRatedUsages_FullTable(test_framework *testing.T) {
	responseList := billingmodels.GenerateRatedUsageRecordSdkList()
	queryParams := billingmodels.GenerateGetRatedUsageQueryParams()
	setQueryParams(queryParams)

	Full = true

	var recordTables []interface{}

	for _, record := range responseList {
		ratedUsageRecord, _ := tables.RatedUsageRecordFromSdk(record, commandName)
		recordTables = append(recordTables, ratedUsageRecord)
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(queryParams).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables, "get rated-usage").
		Return(nil)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

// Currently the short table is an empty struct.
func TestGetAllRatedUsages_ShortTable(test_framework *testing.T) {
	responseList := billingmodels.GenerateRatedUsageRecordSdkList()
	queryParams := billingmodels.GenerateGetRatedUsageQueryParams()
	setQueryParams(queryParams)

	Full = false

	var recordTables []interface{}

	for _, record := range responseList {
		ratedUsageRecord, _ := tables.ShortRatedUsageRecordFromSdk(record, commandName)
		recordTables = append(recordTables, ratedUsageRecord)
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(queryParams).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables, "get rated-usage").
		Return(nil)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllRatedUsages_KeycloakFailure(test_framework *testing.T) {
	queryParams := billingmodels.GenerateGetRatedUsageQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(queryParams).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllRatedUsages_PrinterFailure(test_framework *testing.T) {
	responseList := billingmodels.GenerateRatedUsageRecordSdkList()
	queryParams := billingmodels.GenerateGetRatedUsageQueryParams()
	setQueryParams(queryParams)

	var recordTables []interface{}

	for _, record := range responseList {
		ratedUsageRecord, _ := tables.ShortRatedUsageRecordFromSdk(record, commandName)
		recordTables = append(recordTables, ratedUsageRecord)
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(queryParams).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables, "get rated-usage").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func TestGetAllRatedUsages_ServerError(test_framework *testing.T) {
	queryParams := billingmodels.GenerateGetRatedUsageQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(queryParams).
		Return(nil, WithResponse(500, nil), nil)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	expectedMessage := "Command 'get rated-usage' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetAllRatedUsages_InvalidParams(test_framework *testing.T) {
	queryParams := billingmodels.GenerateGetRatedUsageQueryParams()
	queryParams.FromYearMonth = "0000/00"
	setQueryParams(queryParams)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.Equal(test_framework, fmt.Sprintf("'FromYearMonth' (%s) is not in the valid format (YYYY-MM)", FromYearMonth), err.Error())
}

func setQueryParams(queryparmas billingmodels.RatedUsageGetQueryParams) {
	FromYearMonth = queryparmas.FromYearMonth
	ToYearMonth = queryparmas.ToYearMonth
	ProductCategory = string(*queryparmas.ProductCategory)
}
