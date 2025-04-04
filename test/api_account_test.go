/*
White Label Communications CPaas API Documentation

Testing AccountAPIService

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

func Test_openapi_AccountAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountAPIService V1AccountAccountidChildrenGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidChildrenGet(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidDelete(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidDnsrecordGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidDnsrecordGet(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidDnsrecordPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidDnsrecordPost(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidDnsrecordPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidDnsrecordPut(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidGet(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidLimitGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidLimitGet(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidLimitPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidLimitPut(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidPost(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountAccountidPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.AccountAPI.V1AccountAccountidPut(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountApikeyGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccountAPI.V1AccountApikeyGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccountAPI.V1AccountGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService V1AccountPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccountAPI.V1AccountPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
