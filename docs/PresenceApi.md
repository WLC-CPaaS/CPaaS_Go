# \PresenceAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDPresenceExtensionPut**](PresenceAPI.md#V1AccountAccountIDPresenceExtensionPut) | **Put** /v1/account/{accountID}/presence/{extension} | Set/Reset Presence for Extension
[**V1AccountAccountIDPresenceGet**](PresenceAPI.md#V1AccountAccountIDPresenceGet) | **Get** /v1/account/{accountID}/presence | Get Presence Details
[**V1AccountAccountIDUserUserIDPresencePut**](PresenceAPI.md#V1AccountAccountIDUserUserIDPresencePut) | **Put** /v1/account/{accountID}/user/{userID}/presence | Set/Reset Presence for User



## V1AccountAccountIDPresenceExtensionPut

> ServiceAPIResponse V1AccountAccountIDPresenceExtensionPut(ctx, accountID, extension).ReqBody(reqBody).Execute()

Set/Reset Presence for Extension



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountID := "accountID_example" // string | Account ID, 32 alpha numeric
	extension := "extension_example" // string | Extension
	reqBody := *openapiclient.NewServiceVOIPPresenceSetResetEditData("Action_example") // ServiceVOIPPresenceSetResetEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PresenceAPI.V1AccountAccountIDPresenceExtensionPut(context.Background(), accountID, extension).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PresenceAPI.V1AccountAccountIDPresenceExtensionPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDPresenceExtensionPut`: ServiceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `PresenceAPI.V1AccountAccountIDPresenceExtensionPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**extension** | **string** | Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDPresenceExtensionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPPresenceSetResetEditData**](ServiceVOIPPresenceSetResetEditData.md) | payload fields | 

### Return type

[**ServiceAPIResponse**](ServiceAPIResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDPresenceGet

> ServiceDocsPresenceGet V1AccountAccountIDPresenceGet(ctx, accountID).Execute()

Get Presence Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountID := "accountID_example" // string | Account ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PresenceAPI.V1AccountAccountIDPresenceGet(context.Background(), accountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PresenceAPI.V1AccountAccountIDPresenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDPresenceGet`: ServiceDocsPresenceGet
	fmt.Fprintf(os.Stdout, "Response from `PresenceAPI.V1AccountAccountIDPresenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDPresenceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsPresenceGet**](ServiceDocsPresenceGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDUserUserIDPresencePut

> ServiceAPIResponse V1AccountAccountIDUserUserIDPresencePut(ctx, accountID, userID).ReqBody(reqBody).Execute()

Set/Reset Presence for User



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountID := "accountID_example" // string | Account ID, 32 alpha numeric
	userID := "userID_example" // string | User ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPPresenceSetResetEditData("Action_example") // ServiceVOIPPresenceSetResetEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PresenceAPI.V1AccountAccountIDUserUserIDPresencePut(context.Background(), accountID, userID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PresenceAPI.V1AccountAccountIDUserUserIDPresencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDUserUserIDPresencePut`: ServiceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `PresenceAPI.V1AccountAccountIDUserUserIDPresencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**userID** | **string** | User ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDUserUserIDPresencePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPPresenceSetResetEditData**](ServiceVOIPPresenceSetResetEditData.md) | payload fields | 

### Return type

[**ServiceAPIResponse**](ServiceAPIResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

