/*
White Label Communications CPaas API Documentation

Testing TemporalRuleAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_TemporalRuleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TemporalRuleAPIService V1AccountAccountIDTemporalruleGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string

		resp, httpRes, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalruleGet(context.Background(), accountID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemporalRuleAPIService V1AccountAccountIDTemporalrulePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string

		resp, httpRes, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalrulePost(context.Background(), accountID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemporalRuleAPIService V1AccountAccountIDTemporalruleTemporalRuleIDDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var temporalRuleID string

		resp, httpRes, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDDelete(context.Background(), accountID, temporalRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemporalRuleAPIService V1AccountAccountIDTemporalruleTemporalRuleIDGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var temporalRuleID string

		resp, httpRes, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDGet(context.Background(), accountID, temporalRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemporalRuleAPIService V1AccountAccountIDTemporalruleTemporalRuleIDPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var temporalRuleID string

		resp, httpRes, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDPut(context.Background(), accountID, temporalRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
