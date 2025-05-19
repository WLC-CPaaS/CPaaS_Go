# \

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDLoginrecipientRecipientIDPost**](CallQueueRecipientAPI.md#V1AccountAccountIDLoginrecipientRecipientIDPost) | **Post** /v1/account/{accountID}/loginrecipient/{recipientID} | Login as Recipient
[**V1AccountAccountIDQueuerecipientGet**](CallQueueRecipientAPI.md#V1AccountAccountIDQueuerecipientGet) | **Get** /v1/account/{accountID}/queuerecipient | Change Recipient Status
[**V1AccountAccountIDRecipientRecipientIDStatusPost**](CallQueueRecipientAPI.md#V1AccountAccountIDRecipientRecipientIDStatusPost) | **Post** /v1/account/{accountID}/recipient/{recipientID}/status | Get Recipient List



## V1AccountAccountIDLoginrecipientRecipientIDPost

> ServiceDocsCallQueueResponseShort V1AccountAccountIDLoginrecipientRecipientIDPost(ctx, accountID, recipientID).ReqBody(reqBody).Execute()

Login as Recipient



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
	recipientID := "recipientID_example" // string | Recipient ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPCallQueueRecipientLoginLogoutData("Action_example") // ServiceVOIPCallQueueRecipientLoginLogoutData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueRecipientAPI.V1AccountAccountIDLoginrecipientRecipientIDPost(context.Background(), accountID, recipientID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueRecipientAPI.V1AccountAccountIDLoginrecipientRecipientIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDLoginrecipientRecipientIDPost`: ServiceDocsCallQueueResponseShort
	fmt.Fprintf(os.Stdout, "Response from `CallQueueRecipientAPI.V1AccountAccountIDLoginrecipientRecipientIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**recipientID** | **string** | Recipient ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDLoginrecipientRecipientIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPCallQueueRecipientLoginLogoutData**](ServiceVOIPCallQueueRecipientLoginLogoutData.md) | payload fields | 

### Return type

[**ServiceDocsCallQueueResponseShort**](ServiceDocsCallQueueResponseShort.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDQueuerecipientGet

> ServiceDocsGetQueueRecipients V1AccountAccountIDQueuerecipientGet(ctx, accountID).Execute()

Change Recipient Status



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
	resp, r, err := apiClient.CallQueueRecipientAPI.V1AccountAccountIDQueuerecipientGet(context.Background(), accountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueRecipientAPI.V1AccountAccountIDQueuerecipientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDQueuerecipientGet`: ServiceDocsGetQueueRecipients
	fmt.Fprintf(os.Stdout, "Response from `CallQueueRecipientAPI.V1AccountAccountIDQueuerecipientGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDQueuerecipientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsGetQueueRecipients**](ServiceDocsGetQueueRecipients.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDRecipientRecipientIDStatusPost

> ServiceAPIResponse V1AccountAccountIDRecipientRecipientIDStatusPost(ctx, accountID, recipientID).ReqBody(reqBody).Execute()

Get Recipient List



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
	recipientID := "recipientID_example" // string | Recipient ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPCallQueueRecipientStatusData("Status_example") // ServiceVOIPCallQueueRecipientStatusData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueRecipientAPI.V1AccountAccountIDRecipientRecipientIDStatusPost(context.Background(), accountID, recipientID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueRecipientAPI.V1AccountAccountIDRecipientRecipientIDStatusPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDRecipientRecipientIDStatusPost`: ServiceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `CallQueueRecipientAPI.V1AccountAccountIDRecipientRecipientIDStatusPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**recipientID** | **string** | Recipient ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDRecipientRecipientIDStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPCallQueueRecipientStatusData**](ServiceVOIPCallQueueRecipientStatusData.md) | payload fields | 

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

