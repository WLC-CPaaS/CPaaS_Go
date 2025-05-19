# \MenuAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDMenuGet**](MenuAPI.md#V1AccountAccountIDMenuGet) | **Get** /v1/account/{accountID}/menu | Get Menu List
[**V1AccountAccountIDMenuMenuIDDelete**](MenuAPI.md#V1AccountAccountIDMenuMenuIDDelete) | **Delete** /v1/account/{accountID}/menu/{menuID} | Delete Menu
[**V1AccountAccountIDMenuMenuIDGet**](MenuAPI.md#V1AccountAccountIDMenuMenuIDGet) | **Get** /v1/account/{accountID}/menu/{menuID} | Get Menu Details
[**V1AccountAccountIDMenuMenuIDPut**](MenuAPI.md#V1AccountAccountIDMenuMenuIDPut) | **Put** /v1/account/{accountID}/menu/{menuID} | Update Menu
[**V1AccountAccountIDMenuPost**](MenuAPI.md#V1AccountAccountIDMenuPost) | **Post** /v1/account/{accountID}/menu | Create Menu



## V1AccountAccountIDMenuGet

> MenuOutputList v1accountaccountidmenuget(ctx, accountID).StartKey(startKey).PageSize(pageSize).Execute()

Get Menu List



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
	resp, r, err := apiClient.MenuAPI.V1AccountAccountIDMenuGet(context.Background(), accountID).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuAPI.V1AccountAccountIDMenuGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMenuGet`: MenuOutputList
	fmt.Fprintf(os.Stdout, "Response from `MenuAPI.V1AccountAccountIDMenuGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMenuGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**MenuOutputList**](MenuOutputList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDMenuMenuIDDelete

> MenuOutputDetail v1accountaccountidmenumenuiddelete(ctx, accountID, menuID).Execute()

Delete Menu



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
	menuID := "menuID_example" // string | Menu ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MenuAPI.V1AccountAccountIDMenuMenuIDDelete(context.Background(), accountID, menuID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuAPI.V1AccountAccountIDMenuMenuIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMenuMenuIDDelete`: MenuOutputDetail
	fmt.Fprintf(os.Stdout, "Response from `MenuAPI.V1AccountAccountIDMenuMenuIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**menuID** | **string** | Menu ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMenuMenuIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MenuOutputDetail**](MenuOutputDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDMenuMenuIDGet

> MenuOutputDetail v1accountaccountidmenumenuidget(ctx, accountID, menuID).Execute()

Get Menu Details



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
	menuID := "menuID_example" // string | Menu ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MenuAPI.V1AccountAccountIDMenuMenuIDGet(context.Background(), accountID, menuID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuAPI.V1AccountAccountIDMenuMenuIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMenuMenuIDGet`: MenuOutputDetail
	fmt.Fprintf(os.Stdout, "Response from `MenuAPI.V1AccountAccountIDMenuMenuIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**menuID** | **string** | Menu ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMenuMenuIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MenuOutputDetail**](MenuOutputDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDMenuMenuIDPut

> MenuOutputDetail v1accountaccountidmenumenuidput(ctx, accountID, menuID).ReqBody(reqBody).Execute()

Update Menu



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
	menuID := "menuID_example" // string | Menu ID, 32 alpha numeric
	reqBody := *openapiclient.NewMenuInputData("Name_example") // MenuInputData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MenuAPI.V1AccountAccountIDMenuMenuIDPut(context.Background(), accountID, menuID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuAPI.V1AccountAccountIDMenuMenuIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMenuMenuIDPut`: MenuOutputDetail
	fmt.Fprintf(os.Stdout, "Response from `MenuAPI.V1AccountAccountIDMenuMenuIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**menuID** | **string** | Menu ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMenuMenuIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**MenuInputData**](MenuInputData.md) | payload fields | 

### Return type

[**MenuOutputDetail**](MenuOutputDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDMenuPost

> MenuOutputDetail v1accountaccountidmenupost(ctx, accountID).Menu(menu).Execute()

Create Menu



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
	menu := *openapiclient.NewMenuInputData("Name_example") // MenuInputData | Menu data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MenuAPI.V1AccountAccountIDMenuPost(context.Background(), accountID).Menu(menu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuAPI.V1AccountAccountIDMenuPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMenuPost`: MenuOutputDetail
	fmt.Fprintf(os.Stdout, "Response from `MenuAPI.V1AccountAccountIDMenuPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alphanumeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMenuPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **menu** | [**MenuInputData**](MenuInputData.md) | Menu data | 

### Return type

[**MenuOutputDetail**](MenuOutputDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

