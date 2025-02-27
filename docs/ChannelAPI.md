# \ChannelAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDChannelChannelIDGet**](ChannelAPI.md#V1AccountAccountIDChannelChannelIDGet) | **Get** /v1/account/{accountID}/channel/{channelID} | Get Channel Details
[**V1AccountAccountIDChannelChannelIDPost**](ChannelAPI.md#V1AccountAccountIDChannelChannelIDPost) | **Post** /v1/account/{accountID}/channel/{channelID} | Associate Action to Channel
[**V1AccountAccountIDChannelChannelIDPut**](ChannelAPI.md#V1AccountAccountIDChannelChannelIDPut) | **Put** /v1/account/{accountID}/channel/{channelID} | Associate Metaflow to Channel
[**V1AccountAccountIDChannelGet**](ChannelAPI.md#V1AccountAccountIDChannelGet) | **Get** /v1/account/{accountID}/channel | Get Account Channel List
[**V1AccountAccountIDDeviceDeviceIDChannelGet**](ChannelAPI.md#V1AccountAccountIDDeviceDeviceIDChannelGet) | **Get** /v1/account/{accountID}/device/{deviceID}/channel | Get Device Channel List
[**V1AccountAccountIDUserUserIDChannelGet**](ChannelAPI.md#V1AccountAccountIDUserUserIDChannelGet) | **Get** /v1/account/{accountID}/user/{userID}/channel | Get User Channel List



## V1AccountAccountIDChannelChannelIDGet

> ServiceDocsChannelGetSingle V1AccountAccountIDChannelChannelIDGet(ctx, accountID, channelID).Execute()

Get Channel Details



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
	channelID := "channelID_example" // string | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.V1AccountAccountIDChannelChannelIDGet(context.Background(), accountID, channelID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.V1AccountAccountIDChannelChannelIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDChannelChannelIDGet`: ServiceDocsChannelGetSingle
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.V1AccountAccountIDChannelChannelIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**channelID** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDChannelChannelIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsChannelGetSingle**](ServiceDocsChannelGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDChannelChannelIDPost

> ServiceAPIResponse V1AccountAccountIDChannelChannelIDPost(ctx, accountID, channelID).ReqBody(reqBody).Execute()

Associate Action to Channel



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
	channelID := "channelID_example" // string | Channel ID
	reqBody := *openapiclient.NewServiceChannelRunActionData("Action_example") // ServiceChannelRunActionData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.V1AccountAccountIDChannelChannelIDPost(context.Background(), accountID, channelID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.V1AccountAccountIDChannelChannelIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDChannelChannelIDPost`: ServiceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.V1AccountAccountIDChannelChannelIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**channelID** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDChannelChannelIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceChannelRunActionData**](ServiceChannelRunActionData.md) | payload fields | 

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


## V1AccountAccountIDChannelChannelIDPut

> ServiceAPIResponse V1AccountAccountIDChannelChannelIDPut(ctx, accountID, channelID).ReqBody(reqBody).Execute()

Associate Metaflow to Channel



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
	channelID := "channelID_example" // string | Channel ID
	reqBody := *openapiclient.NewServiceChannelRunMetaflowData("Module_example") // ServiceChannelRunMetaflowData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.V1AccountAccountIDChannelChannelIDPut(context.Background(), accountID, channelID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.V1AccountAccountIDChannelChannelIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDChannelChannelIDPut`: ServiceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.V1AccountAccountIDChannelChannelIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**channelID** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDChannelChannelIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceChannelRunMetaflowData**](ServiceChannelRunMetaflowData.md) | payload fields | 

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


## V1AccountAccountIDChannelGet

> ServiceDocsChannelGet V1AccountAccountIDChannelGet(ctx, accountID).Execute()

Get Account Channel List



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
	resp, r, err := apiClient.ChannelAPI.V1AccountAccountIDChannelGet(context.Background(), accountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.V1AccountAccountIDChannelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDChannelGet`: ServiceDocsChannelGet
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.V1AccountAccountIDChannelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDChannelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsChannelGet**](ServiceDocsChannelGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDDeviceDeviceIDChannelGet

> ServiceDocsChannelGet V1AccountAccountIDDeviceDeviceIDChannelGet(ctx, accountID, deviceID).Execute()

Get Device Channel List



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
	deviceID := "deviceID_example" // string | Device ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.V1AccountAccountIDDeviceDeviceIDChannelGet(context.Background(), accountID, deviceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.V1AccountAccountIDDeviceDeviceIDChannelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDDeviceDeviceIDChannelGet`: ServiceDocsChannelGet
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.V1AccountAccountIDDeviceDeviceIDChannelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**deviceID** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDDeviceDeviceIDChannelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsChannelGet**](ServiceDocsChannelGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDUserUserIDChannelGet

> ServiceDocsChannelGet V1AccountAccountIDUserUserIDChannelGet(ctx, accountID, userID).Execute()

Get User Channel List



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.V1AccountAccountIDUserUserIDChannelGet(context.Background(), accountID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.V1AccountAccountIDUserUserIDChannelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDUserUserIDChannelGet`: ServiceDocsChannelGet
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.V1AccountAccountIDUserUserIDChannelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**userID** | **string** | User ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDUserUserIDChannelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsChannelGet**](ServiceDocsChannelGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

