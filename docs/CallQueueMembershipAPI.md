# \

All URIs are relative to *http://api.beta.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDQueuemembershipPost**](CallQueueMembershipAPI.md#V1AccountAccountIDQueuemembershipPost) | **Post** /v1/account/{accountID}/queuemembership | Grant Queue Membership to User
[**V1AccountAccountIDQueuemembershipRecipientIDDisablePost**](CallQueueMembershipAPI.md#V1AccountAccountIDQueuemembershipRecipientIDDisablePost) | **Post** /v1/account/{accountID}/queuemembership/{recipientID}/disable | Disable Queue Membership
[**V1AccountAccountIDQueuemembershipRecipientIDEnablePost**](CallQueueMembershipAPI.md#V1AccountAccountIDQueuemembershipRecipientIDEnablePost) | **Post** /v1/account/{accountID}/queuemembership/{recipientID}/enable | Enable Queue Membership



## V1AccountAccountIDQueuemembershipPost

> ServiceDocsQueueMembershipOutput V1AccountAccountIDQueuemembershipPost(ctx, accountID).ReqBody(reqBody).Execute()

Grant Queue Membership to User



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
	reqBody := *openapiclient.NewServiceVOIPQueueMembershipAddData("RecipientId_example") // ServiceVOIPQueueMembershipAddData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipPost(context.Background(), accountID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDQueuemembershipPost`: ServiceDocsQueueMembershipOutput
	fmt.Fprintf(os.Stdout, "Response from `CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDQueuemembershipPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reqBody** | [**ServiceVOIPQueueMembershipAddData**](ServiceVOIPQueueMembershipAddData.md) | payload fields | 

### Return type

[**ServiceDocsQueueMembershipOutput**](ServiceDocsQueueMembershipOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDQueuemembershipRecipientIDDisablePost

> ServiceAPIResponse V1AccountAccountIDQueuemembershipRecipientIDDisablePost(ctx, accountID, recipientID).Execute()

Disable Queue Membership



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipRecipientIDDisablePost(context.Background(), accountID, recipientID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipRecipientIDDisablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDQueuemembershipRecipientIDDisablePost`: ServiceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipRecipientIDDisablePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**recipientID** | **string** | Recipient ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDQueuemembershipRecipientIDDisablePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceAPIResponse**](ServiceAPIResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDQueuemembershipRecipientIDEnablePost

> ServiceAPIResponse V1AccountAccountIDQueuemembershipRecipientIDEnablePost(ctx, accountID, recipientID).ReqBody(reqBody).Execute()

Enable Queue Membership



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
	reqBody := *openapiclient.NewServiceVOIPCallQueueEnableMembershipData(false) // ServiceVOIPCallQueueEnableMembershipData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipRecipientIDEnablePost(context.Background(), accountID, recipientID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipRecipientIDEnablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDQueuemembershipRecipientIDEnablePost`: ServiceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `CallQueueMembershipAPI.V1AccountAccountIDQueuemembershipRecipientIDEnablePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**recipientID** | **string** | Recipient ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDQueuemembershipRecipientIDEnablePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPCallQueueEnableMembershipData**](ServiceVOIPCallQueueEnableMembershipData.md) | payload fields | 

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

