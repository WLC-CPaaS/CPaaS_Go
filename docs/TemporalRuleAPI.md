# \

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDTemporalruleGet**](TemporalRuleAPI.md#V1AccountAccountIDTemporalruleGet) | **Get** /v1/account/{accountID}/temporalrule | Get Temporal Rule List
[**V1AccountAccountIDTemporalrulePost**](TemporalRuleAPI.md#V1AccountAccountIDTemporalrulePost) | **Post** /v1/account/{accountID}/temporalrule | Create Temporal Rule
[**V1AccountAccountIDTemporalruleTemporalRuleIDDelete**](TemporalRuleAPI.md#V1AccountAccountIDTemporalruleTemporalRuleIDDelete) | **Delete** /v1/account/{accountID}/temporalrule/{temporalRuleID} | Delete Temporal Rule
[**V1AccountAccountIDTemporalruleTemporalRuleIDGet**](TemporalRuleAPI.md#V1AccountAccountIDTemporalruleTemporalRuleIDGet) | **Get** /v1/account/{accountID}/temporalrule/{temporalRuleID} | Get Temporal Rule Details
[**V1AccountAccountIDTemporalruleTemporalRuleIDPut**](TemporalRuleAPI.md#V1AccountAccountIDTemporalruleTemporalRuleIDPut) | **Put** /v1/account/{accountID}/temporalrule/{temporalRuleID} | Update Temporal Rule



## V1AccountAccountIDTemporalruleGet

> ServiceDocsTemporalRuleGetAll V1AccountAccountIDTemporalruleGet(ctx, accountID).StartKey(startKey).PageSize(pageSize).Execute()

Get Temporal Rule List



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
	resp, r, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalruleGet(context.Background(), accountID).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleAPI.V1AccountAccountIDTemporalruleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalruleGet`: ServiceDocsTemporalRuleGetAll
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleAPI.V1AccountAccountIDTemporalruleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalruleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsTemporalRuleGetAll**](ServiceDocsTemporalRuleGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDTemporalrulePost

> ServiceDocsTemporalRuleGetSingle V1AccountAccountIDTemporalrulePost(ctx, accountID).Temporalrule(temporalrule).Execute()

Create Temporal Rule



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
	accountID := "accountID_example" // string | Account ID, 32 alphanumeric
	temporalrule := *openapiclient.NewServiceVOIPTemporalRuleAddEdit2("Cycle_example", "Name_example") // ServiceVOIPTemporalRuleAddEdit2 | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalrulePost(context.Background(), accountID).Temporalrule(temporalrule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleAPI.V1AccountAccountIDTemporalrulePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalrulePost`: ServiceDocsTemporalRuleGetSingle
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleAPI.V1AccountAccountIDTemporalrulePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alphanumeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalrulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **temporalrule** | [**ServiceVOIPTemporalRuleAddEdit2**](ServiceVOIPTemporalRuleAddEdit2.md) | payload fields | 

### Return type

[**ServiceDocsTemporalRuleGetSingle**](ServiceDocsTemporalRuleGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDTemporalruleTemporalRuleIDDelete

> ServiceDocsTemporalRuleGetSingle V1AccountAccountIDTemporalruleTemporalRuleIDDelete(ctx, accountID, temporalRuleID).Execute()

Delete Temporal Rule



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
	temporalRuleID := "temporalRuleID_example" // string | temporal rule ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDDelete(context.Background(), accountID, temporalRuleID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalruleTemporalRuleIDDelete`: ServiceDocsTemporalRuleGetSingle
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**temporalRuleID** | **string** | temporal rule ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalruleTemporalRuleIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsTemporalRuleGetSingle**](ServiceDocsTemporalRuleGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDTemporalruleTemporalRuleIDGet

> ServiceDocsTemporalRuleGetSingle V1AccountAccountIDTemporalruleTemporalRuleIDGet(ctx, accountID, temporalRuleID).Execute()

Get Temporal Rule Details



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
	temporalRuleID := "temporalRuleID_example" // string | Temporal Rule ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDGet(context.Background(), accountID, temporalRuleID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalruleTemporalRuleIDGet`: ServiceDocsTemporalRuleGetSingle
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**temporalRuleID** | **string** | Temporal Rule ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalruleTemporalRuleIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsTemporalRuleGetSingle**](ServiceDocsTemporalRuleGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDTemporalruleTemporalRuleIDPut

> ServiceDocsTemporalRuleGetSingle V1AccountAccountIDTemporalruleTemporalRuleIDPut(ctx, accountID, temporalRuleID).ReqBody(reqBody).Execute()

Update Temporal Rule



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
	temporalRuleID := "temporalRuleID_example" // string | Temporal Rule ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPTemporalRuleAddEdit2("Cycle_example", "Name_example") // ServiceVOIPTemporalRuleAddEdit2 | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDPut(context.Background(), accountID, temporalRuleID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalruleTemporalRuleIDPut`: ServiceDocsTemporalRuleGetSingle
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleAPI.V1AccountAccountIDTemporalruleTemporalRuleIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**temporalRuleID** | **string** | Temporal Rule ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalruleTemporalRuleIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPTemporalRuleAddEdit2**](ServiceVOIPTemporalRuleAddEdit2.md) | payload fields | 

### Return type

[**ServiceDocsTemporalRuleGetSingle**](ServiceDocsTemporalRuleGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

