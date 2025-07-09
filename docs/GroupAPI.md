# \

All URIs are relative to *http://api.beta.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDGroupGet**](GroupAPI.md#V1AccountAccountIDGroupGet) | **Get** /v1/account/{accountID}/group | Get Group List
[**V1AccountAccountIDGroupGroupIDDelete**](GroupAPI.md#V1AccountAccountIDGroupGroupIDDelete) | **Delete** /v1/account/{accountID}/group/{groupID} | Delete Group
[**V1AccountAccountIDGroupGroupIDGet**](GroupAPI.md#V1AccountAccountIDGroupGroupIDGet) | **Get** /v1/account/{accountID}/group/{groupID} | Get Group Details
[**V1AccountAccountIDGroupGroupIDPut**](GroupAPI.md#V1AccountAccountIDGroupGroupIDPut) | **Put** /v1/account/{accountID}/group/{groupID} | Update Group
[**V1AccountAccountIDGroupPost**](GroupAPI.md#V1AccountAccountIDGroupPost) | **Post** /v1/account/{accountID}/group | Create Group



## V1AccountAccountIDGroupGet

> ServiceDocGroupGetAll V1AccountAccountIDGroupGet(ctx, accountID).StartKey(startKey).PageSize(pageSize).Execute()

Get Group List



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
	resp, r, err := apiClient.GroupAPI.V1AccountAccountIDGroupGet(context.Background(), accountID).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.V1AccountAccountIDGroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDGroupGet`: ServiceDocGroupGetAll
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.V1AccountAccountIDGroupGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDGroupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocGroupGetAll**](ServiceDocGroupGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDGroupGroupIDDelete

> ServiceDocGroupGetSingle V1AccountAccountIDGroupGroupIDDelete(ctx, accountID, groupID).Execute()

Delete Group



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
	groupID := "groupID_example" // string | group ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.V1AccountAccountIDGroupGroupIDDelete(context.Background(), accountID, groupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.V1AccountAccountIDGroupGroupIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDGroupGroupIDDelete`: ServiceDocGroupGetSingle
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.V1AccountAccountIDGroupGroupIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**groupID** | **string** | group ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDGroupGroupIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocGroupGetSingle**](ServiceDocGroupGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDGroupGroupIDGet

> ServiceDocGroupGetSingle V1AccountAccountIDGroupGroupIDGet(ctx, accountID, groupID).Execute()

Get Group Details



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
	groupID := "groupID_example" // string | Group ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.V1AccountAccountIDGroupGroupIDGet(context.Background(), accountID, groupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.V1AccountAccountIDGroupGroupIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDGroupGroupIDGet`: ServiceDocGroupGetSingle
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.V1AccountAccountIDGroupGroupIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**groupID** | **string** | Group ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDGroupGroupIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocGroupGetSingle**](ServiceDocGroupGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDGroupGroupIDPut

> ServiceDocGroupGetSingle V1AccountAccountIDGroupGroupIDPut(ctx, accountID, groupID).ReqBody(reqBody).Execute()

Update Group



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
	groupID := "groupID_example" // string | Group ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPGroupAddEdit2(map[string]ServiceEndpoint{"key": *openapiclient.NewServiceEndpoint()}, "Name_example") // ServiceVOIPGroupAddEdit2 | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.V1AccountAccountIDGroupGroupIDPut(context.Background(), accountID, groupID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.V1AccountAccountIDGroupGroupIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDGroupGroupIDPut`: ServiceDocGroupGetSingle
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.V1AccountAccountIDGroupGroupIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**groupID** | **string** | Group ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDGroupGroupIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPGroupAddEdit2**](ServiceVOIPGroupAddEdit2.md) | payload fields | 

### Return type

[**ServiceDocGroupGetSingle**](ServiceDocGroupGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDGroupPost

> ServiceDocGroupGetSingle V1AccountAccountIDGroupPost(ctx, accountID).Group(group).Execute()

Create Group



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
	accountID := "accountID_example" // string | Account ID
	group := *openapiclient.NewServiceVOIPGroupAddEdit2(map[string]ServiceEndpoint{"key": *openapiclient.NewServiceEndpoint()}, "Name_example") // ServiceVOIPGroupAddEdit2 | group fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.V1AccountAccountIDGroupPost(context.Background(), accountID).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.V1AccountAccountIDGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDGroupPost`: ServiceDocGroupGetSingle
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.V1AccountAccountIDGroupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**ServiceVOIPGroupAddEdit2**](ServiceVOIPGroupAddEdit2.md) | group fields | 

### Return type

[**ServiceDocGroupGetSingle**](ServiceDocGroupGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

