# \CallflowAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDCallflowCallflowIDDelete**](CallflowAPI.md#V1AccountAccountIDCallflowCallflowIDDelete) | **Delete** /v1/account/{accountID}/callflow/{callflowID} | Delete Call Group
[**V1AccountAccountIDCallflowCallflowIDGet**](CallflowAPI.md#V1AccountAccountIDCallflowCallflowIDGet) | **Get** /v1/account/{accountID}/callflow/{callflowID} | Get Call Group Details
[**V1AccountAccountIDCallflowCallflowIDPut**](CallflowAPI.md#V1AccountAccountIDCallflowCallflowIDPut) | **Put** /v1/account/{accountID}/callflow/{callflowID} | Update Call Group
[**V1AccountAccountIDCallflowGet**](CallflowAPI.md#V1AccountAccountIDCallflowGet) | **Get** /v1/account/{accountID}/callflow | Get Callflow List
[**V1AccountAccountIDCallflowPost**](CallflowAPI.md#V1AccountAccountIDCallflowPost) | **Post** /v1/account/{accountID}/callflow | Create Call Group



## V1AccountAccountIDCallflowCallflowIDDelete

> ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDDelete(ctx, accountID, callflowID).Execute()

Delete Call Group



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
	callflowID := "callflowID_example" // string | callflow ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallflowAPI.V1AccountAccountIDCallflowCallflowIDDelete(context.Background(), accountID, callflowID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallflowAPI.V1AccountAccountIDCallflowCallflowIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallflowCallflowIDDelete`: ServiceDocsCallflowGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallflowAPI.V1AccountAccountIDCallflowCallflowIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**callflowID** | **string** | callflow ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallflowCallflowIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCallflowGetSingle**](ServiceDocsCallflowGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallflowCallflowIDGet

> ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDGet(ctx, accountID, callflowID).Execute()

Get Call Group Details



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
	callflowID := "callflowID_example" // string | Callflow ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallflowAPI.V1AccountAccountIDCallflowCallflowIDGet(context.Background(), accountID, callflowID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallflowAPI.V1AccountAccountIDCallflowCallflowIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallflowCallflowIDGet`: ServiceDocsCallflowGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallflowAPI.V1AccountAccountIDCallflowCallflowIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**callflowID** | **string** | Callflow ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallflowCallflowIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCallflowGetSingle**](ServiceDocsCallflowGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallflowCallflowIDPut

> ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDPut(ctx, accountID, callflowID).ReqBody(reqBody).Execute()

Update Call Group



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
	callflowID := "callflowID_example" // string | Callflow ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceCallflowAddEditData(*openapiclient.NewServiceCallflowAddEditFlowData("Module_example"), []string{"Numbers_example"}, []string{"Patterns_example"}) // ServiceCallflowAddEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallflowAPI.V1AccountAccountIDCallflowCallflowIDPut(context.Background(), accountID, callflowID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallflowAPI.V1AccountAccountIDCallflowCallflowIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallflowCallflowIDPut`: ServiceDocsCallflowGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallflowAPI.V1AccountAccountIDCallflowCallflowIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**callflowID** | **string** | Callflow ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallflowCallflowIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceCallflowAddEditData**](ServiceCallflowAddEditData.md) | payload fields | 

### Return type

[**ServiceDocsCallflowGetSingle**](ServiceDocsCallflowGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallflowGet

> ServiceDocsCallflowGetAll V1AccountAccountIDCallflowGet(ctx, accountID).StartKey(startKey).PageSize(pageSize).Execute()

Get Callflow List



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
	startKey := "startKey_example" // string | start_key for pagination that was returned as next_start_key from your previous call (optional)
	pageSize := int32(56) // int32 | number of records to return, range 1 to 50 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallflowAPI.V1AccountAccountIDCallflowGet(context.Background(), accountID).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallflowAPI.V1AccountAccountIDCallflowGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallflowGet`: ServiceDocsCallflowGetAll
	fmt.Fprintf(os.Stdout, "Response from `CallflowAPI.V1AccountAccountIDCallflowGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallflowGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsCallflowGetAll**](ServiceDocsCallflowGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallflowPost

> ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowPost(ctx, accountID).Request(request).Execute()

Create Call Group



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
	accountID := "accountID_example" // string | Account ID, 32 alpha-numeric
	request := *openapiclient.NewServiceCallflowAddEditData(*openapiclient.NewServiceCallflowAddEditFlowData("Module_example"), []string{"Numbers_example"}, []string{"Patterns_example"}) // ServiceCallflowAddEditData | Call flow configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallflowAPI.V1AccountAccountIDCallflowPost(context.Background(), accountID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallflowAPI.V1AccountAccountIDCallflowPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallflowPost`: ServiceDocsCallflowGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallflowAPI.V1AccountAccountIDCallflowPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha-numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallflowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ServiceCallflowAddEditData**](ServiceCallflowAddEditData.md) | Call flow configuration | 

### Return type

[**ServiceDocsCallflowGetSingle**](ServiceDocsCallflowGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

