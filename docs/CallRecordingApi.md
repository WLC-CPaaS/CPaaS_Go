# \CallRecordingAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDRecordingGet**](CallRecordingAPI.md#V1AccountAccountIDRecordingGet) | **Get** /v1/account/{accountID}/recording | Get Account Call Recording
[**V1AccountAccountIDRecordingRecordingIDDelete**](CallRecordingAPI.md#V1AccountAccountIDRecordingRecordingIDDelete) | **Delete** /v1/account/{accountID}/recording/{recordingID} | Delete Call Recording
[**V1AccountAccountIDRecordingRecordingIDGet**](CallRecordingAPI.md#V1AccountAccountIDRecordingRecordingIDGet) | **Get** /v1/account/{accountID}/recording/{recordingID} | Get Call Recording Details
[**V1AccountAccountIDUserUserIDRecordingGet**](CallRecordingAPI.md#V1AccountAccountIDUserUserIDRecordingGet) | **Get** /v1/account/{accountID}/user/{userID}/recording | Get User Call Recording



## V1AccountAccountIDRecordingGet

> ServiceDocsCallRecordingGetAll v1accountaccountidrecordingget(ctx, accountID).Execute()

Get Account Call Recording



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
	resp, r, err := apiClient.CallRecordingAPI.V1AccountAccountIDRecordingGet(context.Background(), accountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingAPI.V1AccountAccountIDRecordingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDRecordingGet`: ServiceDocsCallRecordingGetAll
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingAPI.V1AccountAccountIDRecordingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDRecordingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsCallRecordingGetAll**](ServiceDocsCallRecordingGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDRecordingRecordingIDDelete

> ServiceDocsCallRecordingGetSingle v1accountaccountidrecordingrecordingiddelete(ctx, accountID, recordingID).Execute()

Delete Call Recording



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
	recordingID := "recordingID_example" // string | Recording ID, 39 (yyyymm-<32 id>)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingAPI.V1AccountAccountIDRecordingRecordingIDDelete(context.Background(), accountID, recordingID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingAPI.V1AccountAccountIDRecordingRecordingIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDRecordingRecordingIDDelete`: ServiceDocsCallRecordingGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingAPI.V1AccountAccountIDRecordingRecordingIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**recordingID** | **string** | Recording ID, 39 (yyyymm-&lt;32 id&gt;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDRecordingRecordingIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCallRecordingGetSingle**](ServiceDocsCallRecordingGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDRecordingRecordingIDGet

> ServiceDocsCallRecordingGetSingle v1accountaccountidrecordingrecordingidget(ctx, accountID, recordingID).Execute()

Get Call Recording Details



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
	recordingID := "recordingID_example" // string | Recording ID, 39 (yyyymm-<32 id>)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingAPI.V1AccountAccountIDRecordingRecordingIDGet(context.Background(), accountID, recordingID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingAPI.V1AccountAccountIDRecordingRecordingIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDRecordingRecordingIDGet`: ServiceDocsCallRecordingGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingAPI.V1AccountAccountIDRecordingRecordingIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**recordingID** | **string** | Recording ID, 39 (yyyymm-&lt;32 id&gt;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDRecordingRecordingIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCallRecordingGetSingle**](ServiceDocsCallRecordingGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, audio/mp3, audio/mpeg, audio/mpeg3, audio/x-wav, audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDUserUserIDRecordingGet

> ServiceDocsCallRecordingGetAll v1accountaccountiduseruseridrecordingget(ctx, accountID, userID).Execute()

Get User Call Recording



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
	resp, r, err := apiClient.CallRecordingAPI.V1AccountAccountIDUserUserIDRecordingGet(context.Background(), accountID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingAPI.V1AccountAccountIDUserUserIDRecordingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDUserUserIDRecordingGet`: ServiceDocsCallRecordingGetAll
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingAPI.V1AccountAccountIDUserUserIDRecordingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**userID** | **string** | User ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDUserUserIDRecordingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCallRecordingGetAll**](ServiceDocsCallRecordingGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

