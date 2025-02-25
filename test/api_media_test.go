/*
White Label Communications CPaas API Documentation

Testing MediaAPIService

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

func Test_openapi_MediaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MediaAPIService V1AccountAccountIDMediaMediaIDFileGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var mediaID string

		resp, httpRes, err := apiClient.MediaAPI.V1AccountAccountIDMediaMediaIDFileGet(context.Background(), accountID, mediaID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService V1AccountAccountIDMediaMediaIDFilePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountID string
		var mediaID string

		resp, httpRes, err := apiClient.MediaAPI.V1AccountAccountIDMediaMediaIDFilePost(context.Background(), accountID, mediaID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService V1AccountAccountidMediaGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.MediaAPI.V1AccountAccountidMediaGet(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService V1AccountAccountidMediaMediaidDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string
		var mediaid string

		resp, httpRes, err := apiClient.MediaAPI.V1AccountAccountidMediaMediaidDelete(context.Background(), accountid, mediaid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService V1AccountAccountidMediaMediaidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string
		var mediaid string

		resp, httpRes, err := apiClient.MediaAPI.V1AccountAccountidMediaMediaidGet(context.Background(), accountid, mediaid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService V1AccountAccountidMediaPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountid string

		resp, httpRes, err := apiClient.MediaAPI.V1AccountAccountidMediaPost(context.Background(), accountid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
