/*
White Label Communications CPaas API Documentation

Testing MetaflowAPIService

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

func Test_openapi_MetaflowAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MetaflowAPIService V1AccountAccountIDDeviceDeviceIDMetaflowDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var deviceID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowDelete(context.Background(), accountID, deviceID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaflowAPIService V1AccountAccountIDDeviceDeviceIDMetaflowGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var deviceID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowGet(context.Background(), accountID, deviceID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaflowAPIService V1AccountAccountIDDeviceDeviceIDMetaflowPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var deviceID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowPost(context.Background(), accountID, deviceID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaflowAPIService V1AccountAccountIDMetaflowDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDMetaflowDelete(context.Background(), accountID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaflowAPIService V1AccountAccountIDMetaflowGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDMetaflowGet(context.Background(), accountID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaflowAPIService V1AccountAccountIDMetaflowPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDMetaflowPost(context.Background(), accountID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaflowAPIService V1AccountAccountIDUserUserIDMetaflowDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var userID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowDelete(context.Background(), accountID, userID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaflowAPIService V1AccountAccountIDUserUserIDMetaflowGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var userID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowGet(context.Background(), accountID, userID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaflowAPIService V1AccountAccountIDUserUserIDMetaflowPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var userID string

		resp, httpRes, err := apiClient.MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowPost(context.Background(), accountID, userID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
