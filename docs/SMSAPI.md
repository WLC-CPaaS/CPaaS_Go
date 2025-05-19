# \SMSAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SmsAccountAccountIDCampaignCampaignIDImportGet**](SMSAPI.md#V1SmsAccountAccountIDCampaignCampaignIDImportGet) | **Get** /v1/sms/account/{accountID}/campaign/{campaignID}/import | 
[**V1SmsAccountAccountIDCampaignCampaignIDImportPost**](SMSAPI.md#V1SmsAccountAccountIDCampaignCampaignIDImportPost) | **Post** /v1/sms/account/{accountID}/campaign/{campaignID}/import | 
[**V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet**](SMSAPI.md#V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet) | **Get** /v1/sms/account/{accountID}/campaign/{campaignID}/phonenumber | 
[**V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut**](SMSAPI.md#V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut) | **Put** /v1/sms/account/{accountID}/campaign/{campaignID}/phonenumber | 
[**V1SmsAccountAccountIDCampaignImportGet**](SMSAPI.md#V1SmsAccountAccountIDCampaignImportGet) | **Get** /v1/sms/account/{accountID}/campaign/import | 



## V1SmsAccountAccountIDCampaignCampaignIDImportGet

> ServiceDocsCampaignImportOutput v1smsaccountaccountidcampaigncampaignidimportget(ctx, accountID, campaignID).Execute()





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
	campaignID := "campaignID_example" // string | Campaign ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDImportGet(context.Background(), accountID, campaignID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDImportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmsAccountAccountIDCampaignCampaignIDImportGet`: ServiceDocsCampaignImportOutput
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDImportGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**campaignID** | **string** | Campaign ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmsAccountAccountIDCampaignCampaignIDImportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCampaignImportOutput**](ServiceDocsCampaignImportOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmsAccountAccountIDCampaignCampaignIDImportPost

> ServiceDocsCampaignImportOutput v1smsaccountaccountidcampaigncampaignidimportpost(ctx, accountID, campaignID).Execute()





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
	campaignID := "campaignID_example" // string | Campaign ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDImportPost(context.Background(), accountID, campaignID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmsAccountAccountIDCampaignCampaignIDImportPost`: ServiceDocsCampaignImportOutput
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDImportPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**campaignID** | **string** | Campaign ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmsAccountAccountIDCampaignCampaignIDImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCampaignImportOutput**](ServiceDocsCampaignImportOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet

> ServiceDocsCampaignPhoneNumberOutput v1smsaccountaccountidcampaigncampaignidphonenumberget(ctx, accountID, campaignID).PageNum(pageNum).PageSize(pageSize).Execute()





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
	campaignID := "campaignID_example" // string | Campaign ID
	pageNum := float32(8.14) // float32 | Page number (optional)
	pageSize := float32(8.14) // float32 | Page size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet(context.Background(), accountID, campaignID).PageNum(pageNum).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet`: ServiceDocsCampaignPhoneNumberOutput
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**campaignID** | **string** | Campaign ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNum** | **float32** | Page number | 
 **pageSize** | **float32** | Page size | 

### Return type

[**ServiceDocsCampaignPhoneNumberOutput**](ServiceDocsCampaignPhoneNumberOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut

> ServiceDocsCampaignTagDetagPhonenumbersOutput v1smsaccountaccountidcampaigncampaignidphonenumberput(ctx, accountID, campaignID).ReqBody(reqBody).Execute()





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
	campaignID := "campaignID_example" // string | Campaign ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceCampaignTagDetagPhonenumbers([]string{"PhoneNumbers_example"}) // ServiceCampaignTagDetagPhonenumbers | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut(context.Background(), accountID, campaignID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut`: ServiceDocsCampaignTagDetagPhonenumbersOutput
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**campaignID** | **string** | Campaign ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceCampaignTagDetagPhonenumbers**](ServiceCampaignTagDetagPhonenumbers.md) | payload fields | 

### Return type

[**ServiceDocsCampaignTagDetagPhonenumbersOutput**](ServiceDocsCampaignTagDetagPhonenumbersOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmsAccountAccountIDCampaignImportGet

> ServiceDocsCampaignImportedGetAllOutput v1smsaccountaccountidcampaignimportget(ctx, accountID).PageNum(pageNum).PageSize(pageSize).Execute()





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
	pageNum := float32(8.14) // float32 | Page number (optional)
	pageSize := float32(8.14) // float32 | Page size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.V1SmsAccountAccountIDCampaignImportGet(context.Background(), accountID).PageNum(pageNum).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.V1SmsAccountAccountIDCampaignImportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmsAccountAccountIDCampaignImportGet`: ServiceDocsCampaignImportedGetAllOutput
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.V1SmsAccountAccountIDCampaignImportGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmsAccountAccountIDCampaignImportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNum** | **float32** | Page number | 
 **pageSize** | **float32** | Page size | 

### Return type

[**ServiceDocsCampaignImportedGetAllOutput**](ServiceDocsCampaignImportedGetAllOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

