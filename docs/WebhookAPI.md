# \WebhookAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1WebhookAccountAccountIDGet**](WebhookAPI.md#V1WebhookAccountAccountIDGet) | **Get** /v1/webhook/account/{accountID} | Get Webhook List
[**V1WebhookAccountAccountIDPost**](WebhookAPI.md#V1WebhookAccountAccountIDPost) | **Post** /v1/webhook/account/{accountID} | Create Webhook
[**V1WebhookAccountAccountIDWebhookIDDelete**](WebhookAPI.md#V1WebhookAccountAccountIDWebhookIDDelete) | **Delete** /v1/webhook/account/{accountID}/{webhookID} | Delete Webhook
[**V1WebhookAccountAccountIDWebhookIDGet**](WebhookAPI.md#V1WebhookAccountAccountIDWebhookIDGet) | **Get** /v1/webhook/account/{accountID}/{webhookID} | Get Webhook Details
[**V1WebhookAccountAccountIDWebhookIDPut**](WebhookAPI.md#V1WebhookAccountAccountIDWebhookIDPut) | **Put** /v1/webhook/account/{accountID}/{webhookID} | Update Webhook



## V1WebhookAccountAccountIDGet

> ServiceDocsWebhookGetAll V1WebhookAccountAccountIDGet(ctx, accountID).PageSize(pageSize).CurrentPage(currentPage).Execute()

Get Webhook List



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
	pageSize := int32(56) // int32 | number of records to return, range 1 to 50 (optional)
	currentPage := int32(56) // int32 | Current Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.V1WebhookAccountAccountIDGet(context.Background(), accountID).PageSize(pageSize).CurrentPage(currentPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.V1WebhookAccountAccountIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WebhookAccountAccountIDGet`: ServiceDocsWebhookGetAll
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.V1WebhookAccountAccountIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebhookAccountAccountIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | number of records to return, range 1 to 50 | 
 **currentPage** | **int32** | Current Page | 

### Return type

[**ServiceDocsWebhookGetAll**](ServiceDocsWebhookGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebhookAccountAccountIDPost

> ServiceDocsWebhookGetSingle V1WebhookAccountAccountIDPost(ctx, accountID).Body(body).Execute()

Create Webhook



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
	body := *openapiclient.NewServiceWebhookAdd("CallbackMethod_example", "CallbackUrl_example", "Title_example", "WebhookType_example") // ServiceWebhookAdd | Webhook data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.V1WebhookAccountAccountIDPost(context.Background(), accountID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.V1WebhookAccountAccountIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WebhookAccountAccountIDPost`: ServiceDocsWebhookGetSingle
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.V1WebhookAccountAccountIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebhookAccountAccountIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ServiceWebhookAdd**](ServiceWebhookAdd.md) | Webhook data | 

### Return type

[**ServiceDocsWebhookGetSingle**](ServiceDocsWebhookGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebhookAccountAccountIDWebhookIDDelete

> ServiceDocsWebhookDelete V1WebhookAccountAccountIDWebhookIDDelete(ctx, accountID, webhookID).Execute()

Delete Webhook



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
	webhookID := int32(56) // int32 | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.V1WebhookAccountAccountIDWebhookIDDelete(context.Background(), accountID, webhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.V1WebhookAccountAccountIDWebhookIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WebhookAccountAccountIDWebhookIDDelete`: ServiceDocsWebhookDelete
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.V1WebhookAccountAccountIDWebhookIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID | 
**webhookID** | **int32** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebhookAccountAccountIDWebhookIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsWebhookDelete**](ServiceDocsWebhookDelete.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebhookAccountAccountIDWebhookIDGet

> ServiceDocsWebhookGetSingle V1WebhookAccountAccountIDWebhookIDGet(ctx, accountID, webhookID).Execute()

Get Webhook Details



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
	webhookID := int32(56) // int32 | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.V1WebhookAccountAccountIDWebhookIDGet(context.Background(), accountID, webhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.V1WebhookAccountAccountIDWebhookIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WebhookAccountAccountIDWebhookIDGet`: ServiceDocsWebhookGetSingle
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.V1WebhookAccountAccountIDWebhookIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID | 
**webhookID** | **int32** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebhookAccountAccountIDWebhookIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsWebhookGetSingle**](ServiceDocsWebhookGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebhookAccountAccountIDWebhookIDPut

> ServiceDocsWebhookGetSingle V1WebhookAccountAccountIDWebhookIDPut(ctx, accountID, webhookID).Body(body).Execute()

Update Webhook



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
	webhookID := "webhookID_example" // string | Webhook ID
	body := *openapiclient.NewServiceWebhookEdit("CallbackMethod_example", "CallbackUrl_example", "Title_example", "WebhookType_example") // ServiceWebhookEdit | Updated webhook data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.V1WebhookAccountAccountIDWebhookIDPut(context.Background(), accountID, webhookID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.V1WebhookAccountAccountIDWebhookIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WebhookAccountAccountIDWebhookIDPut`: ServiceDocsWebhookGetSingle
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.V1WebhookAccountAccountIDWebhookIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID | 
**webhookID** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebhookAccountAccountIDWebhookIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ServiceWebhookEdit**](ServiceWebhookEdit.md) | Updated webhook data | 

### Return type

[**ServiceDocsWebhookGetSingle**](ServiceDocsWebhookGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

