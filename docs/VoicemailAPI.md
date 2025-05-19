# \VoicemailAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDVoicemailGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailGet) | **Get** /v1/account/{accountID}/voicemail | Get Voicemail Box List
[**V1AccountAccountIDVoicemailPost**](VoicemailAPI.md#V1AccountAccountIDVoicemailPost) | **Post** /v1/account/{accountID}/voicemail | Create Voicemail Box
[**V1AccountAccountIDVoicemailVoicemailIDDelete**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDDelete) | **Delete** /v1/account/{accountID}/voicemail/{voicemailID} | Delete Voicemail Box
[**V1AccountAccountIDVoicemailVoicemailIDGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDGet) | **Get** /v1/account/{accountID}/voicemail/{voicemailID} | Get Voicemail Box Details
[**V1AccountAccountIDVoicemailVoicemailIDMessageGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageGet) | **Get** /v1/account/{accountID}/voicemail/{voicemailID}/message | Get Voicemail Message List
[**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete) | **Delete** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID} | Delete Voicemail Message
[**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet) | **Get** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID} | Get Voicemail Message Details
[**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut) | **Put** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID} | Update Voicemail Message
[**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet) | **Get** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID}/raw | Get Voicemail Message File
[**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost) | **Post** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID}/raw | Add Voicemail Message File
[**V1AccountAccountIDVoicemailVoicemailIDMessagePost**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessagePost) | **Post** /v1/account/{accountID}/voicemail/{voicemailID}/message | Create Voicemail Message
[**V1AccountAccountIDVoicemailVoicemailIDPut**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDPut) | **Put** /v1/account/{accountID}/voicemail/{voicemailID} | Update Voicemail Box



## V1AccountAccountIDVoicemailGet

> ServiceDocsVoicemailGetAll v1accountaccountidvoicemailget(ctx, accountID).StartKey(startKey).PageSize(pageSize).Execute()

Get Voicemail Box List



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
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailGet(context.Background(), accountID).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailGet`: ServiceDocsVoicemailGetAll
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsVoicemailGetAll**](ServiceDocsVoicemailGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailPost

> ServiceDocsVoicemailGetSingle v1accountaccountidvoicemailpost(ctx, accountID).Voicemail(voicemail).Execute()

Create Voicemail Box



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
	accountID := "accountID_example" // string | account ID, 32 alphanumeric
	voicemail := *openapiclient.NewServiceVOIPVoicemailAddEditData("Mailbox_example", "Name_example") // ServiceVOIPVoicemailAddEditData | voicemail payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailPost(context.Background(), accountID).Voicemail(voicemail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailPost`: ServiceDocsVoicemailGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | account ID, 32 alphanumeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voicemail** | [**ServiceVOIPVoicemailAddEditData**](ServiceVOIPVoicemailAddEditData.md) | voicemail payload fields | 

### Return type

[**ServiceDocsVoicemailGetSingle**](ServiceDocsVoicemailGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDDelete

> ServiceDocsVoicemailGetSingle v1accountaccountidvoicemailvoicemailiddelete(ctx, accountID, voicemailID).Execute()

Delete Voicemail Box



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
	voicemailID := "voicemailID_example" // string | Voicemail ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDDelete(context.Background(), accountID, voicemailID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDDelete`: ServiceDocsVoicemailGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**voicemailID** | **string** | Voicemail ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsVoicemailGetSingle**](ServiceDocsVoicemailGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDGet

> ServiceDocsVoicemailGetSingle v1accountaccountidvoicemailvoicemailidget(ctx, accountID, voicemailID).Execute()

Get Voicemail Box Details



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
	voicemailID := "voicemailID_example" // string | Voicemail ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDGet(context.Background(), accountID, voicemailID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDGet`: ServiceDocsVoicemailGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**voicemailID** | **string** | Voicemail ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsVoicemailGetSingle**](ServiceDocsVoicemailGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDMessageGet

> ServiceDocsVoicemailMessageGetAll v1accountaccountidvoicemailvoicemailidmessageget(ctx, accountID, voicemailID).StartKey(startKey).PageSize(pageSize).Execute()

Get Voicemail Message List



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
	voicemailID := "voicemailID_example" // string | voicemail ID, 32 alpha numeric
	startKey := "startKey_example" // string | start_key for pagination that was returned as next_start_key from your previous call (optional)
	pageSize := int32(56) // int32 | number of records to return, range 1 to 50 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageGet(context.Background(), accountID, voicemailID).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDMessageGet`: ServiceDocsVoicemailMessageGetAll
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**voicemailID** | **string** | voicemail ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDMessageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsVoicemailMessageGetAll**](ServiceDocsVoicemailMessageGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete

> ServiceDocsVoicemailMessageGetSingle v1accountaccountidvoicemailvoicemailidmessagemessageiddelete(ctx, accountID, voicemailID, messageID).Execute()

Delete Voicemail Message



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
	voicemailID := "voicemailID_example" // string | Voicemail ID, 32 alpha numeric
	messageID := "messageID_example" // string | message ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete(context.Background(), accountID, voicemailID, messageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete`: ServiceDocsVoicemailMessageGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**voicemailID** | **string** | Voicemail ID, 32 alpha numeric | 
**messageID** | **string** | message ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceDocsVoicemailMessageGetSingle**](ServiceDocsVoicemailMessageGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet

> ServiceDocsVoicemailMessageGetSingle v1accountaccountidvoicemailvoicemailidmessagemessageidget(ctx, accountID, voicemailID, messageID).Execute()

Get Voicemail Message Details



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
	voicemailID := "voicemailID_example" // string | Voicemail ID, 32 alpha numeric
	messageID := "messageID_example" // string | Message ID, 39 (yyyymm-<32 id>)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet(context.Background(), accountID, voicemailID, messageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet`: ServiceDocsVoicemailMessageGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**voicemailID** | **string** | Voicemail ID, 32 alpha numeric | 
**messageID** | **string** | Message ID, 39 (yyyymm-&lt;32 id&gt;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceDocsVoicemailMessageGetSingle**](ServiceDocsVoicemailMessageGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut

> ServiceDocsVoicemailMessageGetSingle v1accountaccountidvoicemailvoicemailidmessagemessageidput(ctx, accountID, voicemailID, messageID).ReqBody(reqBody).Execute()

Update Voicemail Message



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
	voicemailID := "voicemailID_example" // string | Voicemail ID, 32 alpha numeric
	messageID := "messageID_example" // string | Message ID, 39 (yyyymm-<32 id>)
	reqBody := *openapiclient.NewServiceVOIPVoicemailMessageChange() // ServiceVOIPVoicemailMessageChange | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut(context.Background(), accountID, voicemailID, messageID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut`: ServiceDocsVoicemailMessageGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**voicemailID** | **string** | Voicemail ID, 32 alpha numeric | 
**messageID** | **string** | Message ID, 39 (yyyymm-&lt;32 id&gt;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reqBody** | [**ServiceVOIPVoicemailMessageChange**](ServiceVOIPVoicemailMessageChange.md) | payload fields | 

### Return type

[**ServiceDocsVoicemailMessageGetSingle**](ServiceDocsVoicemailMessageGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet

> *os.File v1accountaccountidvoicemailvoicemailidmessagemessageidrawget(ctx, accountID, voicemailID, messageID).Execute()

Get Voicemail Message File



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
	accountID := "accountID_example" // string | Account ID, unique 32-character alphanumeric identifier
	voicemailID := "voicemailID_example" // string | Voicemail Box ID, unique 32-character alphanumeric identifier
	messageID := "messageID_example" // string | Message ID, unique 32-character alphanumeric identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet(context.Background(), accountID, voicemailID, messageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, unique 32-character alphanumeric identifier | 
**voicemailID** | **string** | Voicemail Box ID, unique 32-character alphanumeric identifier | 
**messageID** | **string** | Message ID, unique 32-character alphanumeric identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost

> map[string]interface{} v1accountaccountidvoicemailvoicemailidmessagemessageidrawpost(ctx, accountID, voicemailID, messageID).File(file).Execute()

Add Voicemail Message File



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
	accountID := "accountID_example" // string | Account ID, 32 alphanumeric characters
	voicemailID := "voicemailID_example" // string | Voicemail ID, 32 alphanumeric characters
	messageID := "messageID_example" // string | Message ID, 32 alphanumeric characters
	file := os.NewFile(1234, "some_file") // *os.File | Audio file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost(context.Background(), accountID, voicemailID, messageID).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alphanumeric characters | 
**voicemailID** | **string** | Voicemail ID, 32 alphanumeric characters | 
**messageID** | **string** | Message ID, 32 alphanumeric characters | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **file** | ***os.File** | Audio file to upload | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDMessagePost

> ServiceDocsVoicemailMessageGetSingle v1accountaccountidvoicemailvoicemailidmessagepost(ctx, accountID, voicemailID).Message(message).Execute()

Create Voicemail Message



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
	accountID := "accountID_example" // string | account ID, 32 alphanumeric
	voicemailID := "voicemailID_example" // string | voicemail ID, 32 alphanumeric
	message := *openapiclient.NewServiceVOIPVoicemailMessageAddData("Folder_example") // ServiceVOIPVoicemailMessageAddData | voicemail message payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessagePost(context.Background(), accountID, voicemailID).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDMessagePost`: ServiceDocsVoicemailMessageGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDMessagePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | account ID, 32 alphanumeric | 
**voicemailID** | **string** | voicemail ID, 32 alphanumeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDMessagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | [**ServiceVOIPVoicemailMessageAddData**](ServiceVOIPVoicemailMessageAddData.md) | voicemail message payload fields | 

### Return type

[**ServiceDocsVoicemailMessageGetSingle**](ServiceDocsVoicemailMessageGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDVoicemailVoicemailIDPut

> ServiceDocsVoicemailGetSingle v1accountaccountidvoicemailvoicemailidput(ctx, accountID, voicemailID).ReqBody(reqBody).Execute()

Update Voicemail Box



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
	voicemailID := "voicemailID_example" // string | Voicemail ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPVoicemailAddEditData("Mailbox_example", "Name_example") // ServiceVOIPVoicemailAddEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDPut(context.Background(), accountID, voicemailID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDVoicemailVoicemailIDPut`: ServiceDocsVoicemailGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.V1AccountAccountIDVoicemailVoicemailIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**voicemailID** | **string** | Voicemail ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDVoicemailVoicemailIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPVoicemailAddEditData**](ServiceVOIPVoicemailAddEditData.md) | payload fields | 

### Return type

[**ServiceDocsVoicemailGetSingle**](ServiceDocsVoicemailGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

