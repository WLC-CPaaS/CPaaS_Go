# \

All URIs are relative to *http://api.beta.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDTemporalrulesetGet**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetGet) | **Get** /v1/account/{accountID}/temporalruleset | Get Temporal Rule Set List
[**V1AccountAccountIDTemporalrulesetPost**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetPost) | **Post** /v1/account/{accountID}/temporalruleset | Create Temporal Rule Set
[**V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete) | **Delete** /v1/account/{accountID}/temporalruleset/{temporalRuleSetID} | Delete Temporal Rule Set
[**V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet) | **Get** /v1/account/{accountID}/temporalruleset/{temporalRuleSetID} | Get Temporal Rule Set Details
[**V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut) | **Put** /v1/account/{accountID}/temporalruleset/{temporalRuleSetID} | Update Temporal Rule Set



## V1AccountAccountIDTemporalrulesetGet

> ServiceDocsTemporalRuleSetGetAll V1AccountAccountIDTemporalrulesetGet(ctx, accountID).StartKey(startKey).PageSize(pageSize).Execute()

Get Temporal Rule Set List



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
	resp, r, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetGet(context.Background(), accountID).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalrulesetGet`: ServiceDocsTemporalRuleSetGetAll
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalrulesetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsTemporalRuleSetGetAll**](ServiceDocsTemporalRuleSetGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDTemporalrulesetPost

> ServiceDocsTemporalRuleSetGetSingle V1AccountAccountIDTemporalrulesetPost(ctx, accountID).Temporalruleset(temporalruleset).Execute()

Create Temporal Rule Set



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
	temporalruleset := *openapiclient.NewServiceVOIPTemporalRuleSetAddEditData("Name_example") // ServiceVOIPTemporalRuleSetAddEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetPost(context.Background(), accountID).Temporalruleset(temporalruleset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalrulesetPost`: ServiceDocsTemporalRuleSetGetSingle
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alphanumeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalrulesetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **temporalruleset** | [**ServiceVOIPTemporalRuleSetAddEditData**](ServiceVOIPTemporalRuleSetAddEditData.md) | payload fields | 

### Return type

[**ServiceDocsTemporalRuleSetGetSingle**](ServiceDocsTemporalRuleSetGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete

> ServiceDocsTemporalRuleSetGetSingle V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete(ctx, accountID, temporalRuleSetID).Execute()

Delete Temporal Rule Set



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
	temporalRuleSetID := "temporalRuleSetID_example" // string | temporal rule set ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete(context.Background(), accountID, temporalRuleSetID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete`: ServiceDocsTemporalRuleSetGetSingle
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**temporalRuleSetID** | **string** | temporal rule set ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalrulesetTemporalRuleSetIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsTemporalRuleSetGetSingle**](ServiceDocsTemporalRuleSetGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet

> ServiceDocsTemporalRuleSetGetSingle V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet(ctx, accountID, temporalRuleSetID).Execute()

Get Temporal Rule Set Details



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
	temporalRuleSetID := "temporalRuleSetID_example" // string | Temporal Ruleset ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet(context.Background(), accountID, temporalRuleSetID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet`: ServiceDocsTemporalRuleSetGetSingle
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**temporalRuleSetID** | **string** | Temporal Ruleset ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalrulesetTemporalRuleSetIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsTemporalRuleSetGetSingle**](ServiceDocsTemporalRuleSetGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut

> ServiceDocsTemporalRuleSetGetSingle V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut(ctx, accountID, temporalRuleSetID).ReqBody(reqBody).Execute()

Update Temporal Rule Set



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
	temporalRuleSetID := "temporalRuleSetID_example" // string | Temporal Ruleset ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPTemporalRuleSetAddEditData("Name_example") // ServiceVOIPTemporalRuleSetAddEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut(context.Background(), accountID, temporalRuleSetID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut`: ServiceDocsTemporalRuleSetGetSingle
	fmt.Fprintf(os.Stdout, "Response from `TemporalRuleSetAPI.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**temporalRuleSetID** | **string** | Temporal Ruleset ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDTemporalrulesetTemporalRuleSetIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPTemporalRuleSetAddEditData**](ServiceVOIPTemporalRuleSetAddEditData.md) | payload fields | 

### Return type

[**ServiceDocsTemporalRuleSetGetSingle**](ServiceDocsTemporalRuleSetGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

