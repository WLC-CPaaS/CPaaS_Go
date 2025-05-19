# \DeviceAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountidDeviceDeviceidDelete**](DeviceAPI.md#V1AccountAccountidDeviceDeviceidDelete) | **Delete** /v1/account/{accountid}/device/{deviceid} | Delete Device
[**V1AccountAccountidDeviceDeviceidGet**](DeviceAPI.md#V1AccountAccountidDeviceDeviceidGet) | **Get** /v1/account/{accountid}/device/{deviceid} | Get Device Details
[**V1AccountAccountidDeviceDeviceidPut**](DeviceAPI.md#V1AccountAccountidDeviceDeviceidPut) | **Put** /v1/account/{accountid}/device/{deviceid} | Update Device
[**V1AccountAccountidDeviceDeviceidRebootPost**](DeviceAPI.md#V1AccountAccountidDeviceDeviceidRebootPost) | **Post** /v1/account/{accountid}/device/{deviceid}/reboot | Reboot Device
[**V1AccountAccountidDeviceGet**](DeviceAPI.md#V1AccountAccountidDeviceGet) | **Get** /v1/account/{accountid}/device | Get Device List
[**V1AccountAccountidDevicePost**](DeviceAPI.md#V1AccountAccountidDevicePost) | **Post** /v1/account/{accountid}/device | Create Device



## V1AccountAccountidDeviceDeviceidDelete

> ServiceDocsDeviceGetSingle v1accountaccountiddevicedeviceiddelete(ctx, accountid, deviceid).Execute()

Delete Device



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
	deviceid := "deviceid_example" // string | Device ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.V1AccountAccountidDeviceDeviceidDelete(context.Background(), accountid, deviceid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.V1AccountAccountidDeviceDeviceidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDeviceDeviceidDelete`: ServiceDocsDeviceGetSingle
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.V1AccountAccountidDeviceDeviceidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**deviceid** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDeviceDeviceidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsDeviceGetSingle**](ServiceDocsDeviceGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDeviceDeviceidGet

> ServiceDocsDeviceGetSingle v1accountaccountiddevicedeviceidget(ctx, accountid, deviceid).Execute()

Get Device Details



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
	deviceid := "deviceid_example" // string | Device ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.V1AccountAccountidDeviceDeviceidGet(context.Background(), accountid, deviceid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.V1AccountAccountidDeviceDeviceidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDeviceDeviceidGet`: ServiceDocsDeviceGetSingle
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.V1AccountAccountidDeviceDeviceidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**deviceid** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDeviceDeviceidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsDeviceGetSingle**](ServiceDocsDeviceGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDeviceDeviceidPut

> ServiceDocsDeviceGetSingle v1accountaccountiddevicedeviceidput(ctx, accountid, deviceid).Device(device).Execute()

Update Device



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
	deviceid := "deviceid_example" // string | Device ID, 32 alpha numeric
	device := *openapiclient.NewServiceVOIPDeviceAddEdit2("Name_example", *openapiclient.NewServiceVOIPDeviceAddEdit3a("InviteFormat_example", "Username_example")) // ServiceVOIPDeviceAddEdit2 | device fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.V1AccountAccountidDeviceDeviceidPut(context.Background(), accountid, deviceid).Device(device).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.V1AccountAccountidDeviceDeviceidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDeviceDeviceidPut`: ServiceDocsDeviceGetSingle
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.V1AccountAccountidDeviceDeviceidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**deviceid** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDeviceDeviceidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**ServiceVOIPDeviceAddEdit2**](ServiceVOIPDeviceAddEdit2.md) | device fields | 

### Return type

[**ServiceDocsDeviceGetSingle**](ServiceDocsDeviceGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDeviceDeviceidRebootPost

> ServiceDocsDeviceReboot v1accountaccountiddevicedeviceidrebootpost(ctx, accountid, deviceid).Execute()

Reboot Device



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
	deviceid := "deviceid_example" // string | Device ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.V1AccountAccountidDeviceDeviceidRebootPost(context.Background(), accountid, deviceid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.V1AccountAccountidDeviceDeviceidRebootPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDeviceDeviceidRebootPost`: ServiceDocsDeviceReboot
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.V1AccountAccountidDeviceDeviceidRebootPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**deviceid** | **string** | Device ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDeviceDeviceidRebootPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsDeviceReboot**](ServiceDocsDeviceReboot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDeviceGet

> ServiceDocsDeviceGetAll v1accountaccountiddeviceget(ctx, accountid).StartKey(startKey).PageSize(pageSize).Execute()

Get Device List



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
	resp, r, err := apiClient.DeviceAPI.V1AccountAccountidDeviceGet(context.Background(), accountid).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.V1AccountAccountidDeviceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDeviceGet`: ServiceDocsDeviceGetAll
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.V1AccountAccountidDeviceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDeviceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsDeviceGetAll**](ServiceDocsDeviceGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDevicePost

> ServiceDocsDeviceGetSingle v1accountaccountiddevicepost(ctx, accountid).Device(device).Execute()

Create Device



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
	device := *openapiclient.NewServiceVOIPDeviceAddEdit2("Name_example", *openapiclient.NewServiceVOIPDeviceAddEdit3a("InviteFormat_example", "Username_example")) // ServiceVOIPDeviceAddEdit2 | device fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.V1AccountAccountidDevicePost(context.Background(), accountid).Device(device).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.V1AccountAccountidDevicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDevicePost`: ServiceDocsDeviceGetSingle
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.V1AccountAccountidDevicePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDevicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **device** | [**ServiceVOIPDeviceAddEdit2**](ServiceVOIPDeviceAddEdit2.md) | device fields | 

### Return type

[**ServiceDocsDeviceGetSingle**](ServiceDocsDeviceGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

