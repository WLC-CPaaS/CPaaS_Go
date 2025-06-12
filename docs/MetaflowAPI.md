# \

All URIs are relative to *http://API_HOSTNAME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDDeviceDeviceIDMetaflowDelete**](MetaflowAPI.md#V1AccountAccountIDDeviceDeviceIDMetaflowDelete) | **Delete** /v1/account/{accountID}/device/{deviceID}/metaflow | Delete Device Metaflow
[**V1AccountAccountIDDeviceDeviceIDMetaflowGet**](MetaflowAPI.md#V1AccountAccountIDDeviceDeviceIDMetaflowGet) | **Get** /v1/account/{accountID}/device/{deviceID}/metaflow | Get Device Metaflow List
[**V1AccountAccountIDDeviceDeviceIDMetaflowPost**](MetaflowAPI.md#V1AccountAccountIDDeviceDeviceIDMetaflowPost) | **Post** /v1/account/{accountID}/device/{deviceID}/metaflow | Create Device Metaflow
[**V1AccountAccountIDMetaflowDelete**](MetaflowAPI.md#V1AccountAccountIDMetaflowDelete) | **Delete** /v1/account/{accountID}/metaflow | Delete Account Metaflow
[**V1AccountAccountIDMetaflowGet**](MetaflowAPI.md#V1AccountAccountIDMetaflowGet) | **Get** /v1/account/{accountID}/metaflow | Get Account Metaflow List
[**V1AccountAccountIDMetaflowPost**](MetaflowAPI.md#V1AccountAccountIDMetaflowPost) | **Post** /v1/account/{accountID}/metaflow | Create Account Metaflow
[**V1AccountAccountIDUserUserIDMetaflowDelete**](MetaflowAPI.md#V1AccountAccountIDUserUserIDMetaflowDelete) | **Delete** /v1/account/{accountID}/user/{userID}/metaflow | Delete User Metaflow
[**V1AccountAccountIDUserUserIDMetaflowGet**](MetaflowAPI.md#V1AccountAccountIDUserUserIDMetaflowGet) | **Get** /v1/account/{accountID}/user/{userID}/metaflow | Get User Metaflow List
[**V1AccountAccountIDUserUserIDMetaflowPost**](MetaflowAPI.md#V1AccountAccountIDUserUserIDMetaflowPost) | **Post** /v1/account/{accountID}/user/{userID}/metaflow | Create User Metaflow



## V1AccountAccountIDDeviceDeviceIDMetaflowDelete

> ServiceDocsMetaflowGet V1AccountAccountIDDeviceDeviceIDMetaflowDelete(ctx, accountID, deviceID).Execute()

Delete Device Metaflow



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
	deviceID := "deviceID_example" // string | Device ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowDelete(context.Background(), accountID, deviceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDDeviceDeviceIDMetaflowDelete`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**deviceID** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDDeviceDeviceIDMetaflowDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDDeviceDeviceIDMetaflowGet

> ServiceDocsMetaflowGet V1AccountAccountIDDeviceDeviceIDMetaflowGet(ctx, accountID, deviceID).Execute()

Get Device Metaflow List



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
	deviceID := "deviceID_example" // string | Device ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowGet(context.Background(), accountID, deviceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDDeviceDeviceIDMetaflowGet`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**deviceID** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDDeviceDeviceIDMetaflowGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDDeviceDeviceIDMetaflowPost

> ServiceDocsMetaflowGet V1AccountAccountIDDeviceDeviceIDMetaflowPost(ctx, accountID, deviceID).ReqBody(reqBody).Execute()

Create Device Metaflow



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
	deviceID := "deviceID_example" // string | Device ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPMetaflowAddData() // ServiceVOIPMetaflowAddData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowPost(context.Background(), accountID, deviceID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDDeviceDeviceIDMetaflowPost`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDDeviceDeviceIDMetaflowPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**deviceID** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDDeviceDeviceIDMetaflowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPMetaflowAddData**](ServiceVOIPMetaflowAddData.md) | payload fields | 

### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDMetaflowDelete

> ServiceDocsMetaflowGet V1AccountAccountIDMetaflowDelete(ctx, accountID).Execute()

Delete Account Metaflow



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
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDMetaflowDelete(context.Background(), accountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDMetaflowDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMetaflowDelete`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDMetaflowDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMetaflowDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDMetaflowGet

> ServiceDocsMetaflowGet V1AccountAccountIDMetaflowGet(ctx, accountID).Execute()

Get Account Metaflow List



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
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDMetaflowGet(context.Background(), accountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDMetaflowGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMetaflowGet`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDMetaflowGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMetaflowGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDMetaflowPost

> ServiceDocsMetaflowGet V1AccountAccountIDMetaflowPost(ctx, accountID).Metaflow(metaflow).Execute()

Create Account Metaflow



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
	metaflow := *openapiclient.NewServiceVOIPMetaflowAddData() // ServiceVOIPMetaflowAddData | Metaflow fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDMetaflowPost(context.Background(), accountID).Metaflow(metaflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDMetaflowPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDMetaflowPost`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDMetaflowPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDMetaflowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metaflow** | [**ServiceVOIPMetaflowAddData**](ServiceVOIPMetaflowAddData.md) | Metaflow fields | 

### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDUserUserIDMetaflowDelete

> ServiceDocsMetaflowGet V1AccountAccountIDUserUserIDMetaflowDelete(ctx, accountID, userID).Execute()

Delete User Metaflow



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
	userID := "userID_example" // string | user ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowDelete(context.Background(), accountID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDUserUserIDMetaflowDelete`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**userID** | **string** | user ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDUserUserIDMetaflowDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDUserUserIDMetaflowGet

> ServiceDocsMetaflowGet V1AccountAccountIDUserUserIDMetaflowGet(ctx, accountID, userID).Execute()

Get User Metaflow List



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
	userID := "userID_example" // string | user ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowGet(context.Background(), accountID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDUserUserIDMetaflowGet`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**userID** | **string** | user ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDUserUserIDMetaflowGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDUserUserIDMetaflowPost

> ServiceDocsMetaflowGet V1AccountAccountIDUserUserIDMetaflowPost(ctx, accountID, userID).ReqBody(reqBody).Execute()

Create User Metaflow



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
	userID := "userID_example" // string | user ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPMetaflowAddData() // ServiceVOIPMetaflowAddData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowPost(context.Background(), accountID, userID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDUserUserIDMetaflowPost`: ServiceDocsMetaflowGet
	fmt.Fprintf(os.Stdout, "Response from `MetaflowAPI.V1AccountAccountIDUserUserIDMetaflowPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**userID** | **string** | user ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDUserUserIDMetaflowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPMetaflowAddData**](ServiceVOIPMetaflowAddData.md) | payload fields | 

### Return type

[**ServiceDocsMetaflowGet**](ServiceDocsMetaflowGet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

