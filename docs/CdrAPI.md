# \CdrAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDCdrCdrIDGet**](CdrAPI.md#V1AccountAccountIDCdrCdrIDGet) | **Get** /v1/account/{accountID}/cdr/{cdrID} | 
[**V1AccountAccountIDCdrGet**](CdrAPI.md#V1AccountAccountIDCdrGet) | **Get** /v1/account/{accountID}/cdr | 



## V1AccountAccountIDCdrCdrIDGet

> ServiceDocsCdrGetSingle v1accountaccountidcdrcdridget(ctx, accountID, cdrID).Execute()





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
	cdrID := "cdrID_example" // string | CDR ID, string

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CdrAPI.V1AccountAccountIDCdrCdrIDGet(context.Background(), accountID, cdrID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.V1AccountAccountIDCdrCdrIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCdrCdrIDGet`: ServiceDocsCdrGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CdrAPI.V1AccountAccountIDCdrCdrIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**cdrID** | **string** | CDR ID, string | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCdrCdrIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCdrGetSingle**](ServiceDocsCdrGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCdrGet

> ServiceDocsCdrGetAll v1accountaccountidcdrget(ctx, accountID).PageSize(pageSize).StartKey(startKey).CreatedFrom(createdFrom).CreatedTo(createdTo).Execute()





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
	pageSize := "pageSize_example" // string | Page size (Maximum number of results to display per page)
	startKey := "startKey_example" // string | Start key (Starting offset for displaying results)
	createdFrom := "createdFrom_example" // string | For displaying records which are created on or after this timestamp (Supported timestamp formats: iso 8601, unix time in seconds or milliseconds or microseconds or nanoseconds)
	createdTo := "createdTo_example" // string | For displaying records which are created on or before this timestamp (Supported timestamp formats: iso 8601, unix time in seconds or milliseconds or microseconds or nanoseconds)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CdrAPI.V1AccountAccountIDCdrGet(context.Background(), accountID).PageSize(pageSize).StartKey(startKey).CreatedFrom(createdFrom).CreatedTo(createdTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.V1AccountAccountIDCdrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCdrGet`: ServiceDocsCdrGetAll
	fmt.Fprintf(os.Stdout, "Response from `CdrAPI.V1AccountAccountIDCdrGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCdrGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **string** | Page size (Maximum number of results to display per page) | 
 **startKey** | **string** | Start key (Starting offset for displaying results) | 
 **createdFrom** | **string** | For displaying records which are created on or after this timestamp (Supported timestamp formats: iso 8601, unix time in seconds or milliseconds or microseconds or nanoseconds) | 
 **createdTo** | **string** | For displaying records which are created on or before this timestamp (Supported timestamp formats: iso 8601, unix time in seconds or milliseconds or microseconds or nanoseconds) | 

### Return type

[**ServiceDocsCdrGetAll**](ServiceDocsCdrGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

