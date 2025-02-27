# \MediaAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDMediaMediaIDFileGet**](MediaAPI.md#V1AccountAccountIDMediaMediaIDFileGet) | **Get** /v1/account/{accountID}/media/{mediaID}/file | Get Media File
[**V1AccountAccountIDMediaMediaIDFilePost**](MediaAPI.md#V1AccountAccountIDMediaMediaIDFilePost) | **Post** /v1/account/{accountID}/media/{mediaID}/file | Add Media File
[**V1AccountAccountidMediaGet**](MediaAPI.md#V1AccountAccountidMediaGet) | **Get** /v1/account/{accountid}/media | Get Media List
[**V1AccountAccountidMediaMediaidDelete**](MediaAPI.md#V1AccountAccountidMediaMediaidDelete) | **Delete** /v1/account/{accountid}/media/{mediaid} | Delete Media
[**V1AccountAccountidMediaMediaidGet**](MediaAPI.md#V1AccountAccountidMediaMediaidGet) | **Get** /v1/account/{accountid}/media/{mediaid} | Get Media Details
[**V1AccountAccountidMediaPost**](MediaAPI.md#V1AccountAccountidMediaPost) | **Post** /v1/account/{accountid}/media | Create Media



## V1AccountAccountIDMediaMediaIDFileGet

> *os.File V1AccountAccountIDMediaMediaIDFileGet(ctx, accountID, mediaID).Execute()

Get Media File



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
	mediaID := "mediaID_example" // string | Media ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.V1AccountAccountIDMediaMediaIDFileGet(context.Background(), accountID, mediaID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.V1AccountAccountIDMediaMediaIDFileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMediaMediaIDFileGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.V1AccountAccountIDMediaMediaIDFileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**mediaID** | **string** | Media ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMediaMediaIDFileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, audio/mp3, audio/mpeg, audio/mpeg3, audio/x-wav, audio/wav, audio/ogg, video/x-flv, video/h264, video/mpeg, video/quicktime, video/mp4, video/webm

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDMediaMediaIDFilePost

> ServiceDocsMediaGetSingle V1AccountAccountIDMediaMediaIDFilePost(ctx, accountID, mediaID).File(file).Execute()

Add Media File



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
	mediaID := "mediaID_example" // string | Media ID, 32 alpha numeric
	file := os.NewFile(1234, "some_file") // *os.File | Media file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.V1AccountAccountIDMediaMediaIDFilePost(context.Background(), accountID, mediaID).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.V1AccountAccountIDMediaMediaIDFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMediaMediaIDFilePost`: ServiceDocsMediaGetSingle
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.V1AccountAccountIDMediaMediaIDFilePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**mediaID** | **string** | Media ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMediaMediaIDFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | Media file | 

### Return type

[**ServiceDocsMediaGetSingle**](ServiceDocsMediaGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidMediaGet

> ServiceDocsMediaGetAll V1AccountAccountidMediaGet(ctx, accountid).StartKey(startKey).PageSize(pageSize).Execute()

Get Media List



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
	accountid := "accountid_example" // string | Account ID, 32 alpha numeric
	startKey := "startKey_example" // string | start_key for pagination that was returned as next_start_key from your previous call (optional)
	pageSize := int32(56) // int32 | number of records to return, range 1 to 50 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.V1AccountAccountidMediaGet(context.Background(), accountid).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.V1AccountAccountidMediaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidMediaGet`: ServiceDocsMediaGetAll
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.V1AccountAccountidMediaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidMediaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsMediaGetAll**](ServiceDocsMediaGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidMediaMediaidDelete

> ServiceDocsMediaGetSingle V1AccountAccountidMediaMediaidDelete(ctx, accountid, mediaid).Execute()

Delete Media



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
	accountid := "accountid_example" // string | Account ID, 32 alpha numeric
	mediaid := "mediaid_example" // string | Device ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.V1AccountAccountidMediaMediaidDelete(context.Background(), accountid, mediaid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.V1AccountAccountidMediaMediaidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidMediaMediaidDelete`: ServiceDocsMediaGetSingle
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.V1AccountAccountidMediaMediaidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**mediaid** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidMediaMediaidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsMediaGetSingle**](ServiceDocsMediaGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidMediaMediaidGet

> ServiceDocsMediaGetSingle V1AccountAccountidMediaMediaidGet(ctx, accountid, mediaid).Execute()

Get Media Details



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
	accountid := "accountid_example" // string | Account ID, 32 alpha numeric
	mediaid := "mediaid_example" // string | Media ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.V1AccountAccountidMediaMediaidGet(context.Background(), accountid, mediaid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.V1AccountAccountidMediaMediaidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidMediaMediaidGet`: ServiceDocsMediaGetSingle
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.V1AccountAccountidMediaMediaidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**mediaid** | **string** | Media ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidMediaMediaidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsMediaGetSingle**](ServiceDocsMediaGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidMediaPost

> ServiceDocsMediaGetSingle V1AccountAccountidMediaPost(ctx, accountid).Media(media).Execute()

Create Media



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
	accountid := "accountid_example" // string | Account ID, 32 alpha numeric
	media := *openapiclient.NewServiceVOIPMediaAddEditData("Name_example") // ServiceVOIPMediaAddEditData | Media creation or update payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.V1AccountAccountidMediaPost(context.Background(), accountid).Media(media).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.V1AccountAccountidMediaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidMediaPost`: ServiceDocsMediaGetSingle
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.V1AccountAccountidMediaPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidMediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **media** | [**ServiceVOIPMediaAddEditData**](ServiceVOIPMediaAddEditData.md) | Media creation or update payload | 

### Return type

[**ServiceDocsMediaGetSingle**](ServiceDocsMediaGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

