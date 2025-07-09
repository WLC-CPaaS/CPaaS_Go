# \

All URIs are relative to *http://api.beta.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountidPhonenumberGet**](PhoneNumberAPI.md#V1AccountAccountidPhonenumberGet) | **Get** /v1/account/{accountid}/phonenumber | Get Assigned Numbers List
[**V1AccountPhonenumberAssignPost**](PhoneNumberAPI.md#V1AccountPhonenumberAssignPost) | **Post** /v1/account/phonenumber/assign | Assign Number
[**V1AccountPhonenumberDisconnectPost**](PhoneNumberAPI.md#V1AccountPhonenumberDisconnectPost) | **Post** /v1/account/phonenumber/disconnect | Disconnect Number
[**V1AccountPhonenumberGet**](PhoneNumberAPI.md#V1AccountPhonenumberGet) | **Get** /v1/account/phonenumber | Get Unassigned Numbers List
[**V1AccountPhonenumberPost**](PhoneNumberAPI.md#V1AccountPhonenumberPost) | **Post** /v1/account/phonenumber | Purchase Number
[**V1AccountPhonenumberUnassignPost**](PhoneNumberAPI.md#V1AccountPhonenumberUnassignPost) | **Post** /v1/account/phonenumber/unassign | Unassign Number
[**V1PhonenumberSearchGet**](PhoneNumberAPI.md#V1PhonenumberSearchGet) | **Get** /v1/phonenumber/search | Search New Numbers



## V1AccountAccountidPhonenumberGet

> ServiceDocsAccountPhonenumberGetAll V1AccountAccountidPhonenumberGet(ctx, accountid).StartKey(startKey).PageSize(pageSize).Execute()

Get Assigned Numbers List



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
	startKey := "startKey_example" // string | Start key for pagination, obtained from previous responses (optional)
	pageSize := int32(56) // int32 | Number of records to return per page (range: 1 to 50) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberAPI.V1AccountAccountidPhonenumberGet(context.Background(), accountid).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberAPI.V1AccountAccountidPhonenumberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidPhonenumberGet`: ServiceDocsAccountPhonenumberGetAll
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberAPI.V1AccountAccountidPhonenumberGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidPhonenumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | Start key for pagination, obtained from previous responses | 
 **pageSize** | **int32** | Number of records to return per page (range: 1 to 50) | 

### Return type

[**ServiceDocsAccountPhonenumberGetAll**](ServiceDocsAccountPhonenumberGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPhonenumberAssignPost

> ServiceAPIResponseStatusCodeOnly V1AccountPhonenumberAssignPost(ctx).Payload(payload).Execute()

Assign Number



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
	payload := *openapiclient.NewServiceDocsPhonenumberAssignPayload() // ServiceDocsPhonenumberAssignPayload | assignment payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberAPI.V1AccountPhonenumberAssignPost(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberAPI.V1AccountPhonenumberAssignPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountPhonenumberAssignPost`: ServiceAPIResponseStatusCodeOnly
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberAPI.V1AccountPhonenumberAssignPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPhonenumberAssignPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ServiceDocsPhonenumberAssignPayload**](ServiceDocsPhonenumberAssignPayload.md) | assignment payload | 

### Return type

[**ServiceAPIResponseStatusCodeOnly**](ServiceAPIResponseStatusCodeOnly.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPhonenumberDisconnectPost

> ServiceAPIResponseStatusCodeOnly V1AccountPhonenumberDisconnectPost(ctx).Payload(payload).Execute()

Disconnect Number



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
	payload := *openapiclient.NewServiceDocsPhonenumberUnassignPayload() // ServiceDocsPhonenumberUnassignPayload | disconnect payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberAPI.V1AccountPhonenumberDisconnectPost(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberAPI.V1AccountPhonenumberDisconnectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountPhonenumberDisconnectPost`: ServiceAPIResponseStatusCodeOnly
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberAPI.V1AccountPhonenumberDisconnectPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPhonenumberDisconnectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ServiceDocsPhonenumberUnassignPayload**](ServiceDocsPhonenumberUnassignPayload.md) | disconnect payload | 

### Return type

[**ServiceAPIResponseStatusCodeOnly**](ServiceAPIResponseStatusCodeOnly.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPhonenumberGet

> ServiceDocsAccountPhonenumberGetAll V1AccountPhonenumberGet(ctx).StartKey(startKey).PageSize(pageSize).Execute()

Get Unassigned Numbers List



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
	startKey := "startKey_example" // string | Start key for pagination, obtained from previous responses (optional)
	pageSize := int32(56) // int32 | Number of records to return per page (range: 1 to 50) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberAPI.V1AccountPhonenumberGet(context.Background()).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberAPI.V1AccountPhonenumberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountPhonenumberGet`: ServiceDocsAccountPhonenumberGetAll
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberAPI.V1AccountPhonenumberGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPhonenumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startKey** | **string** | Start key for pagination, obtained from previous responses | 
 **pageSize** | **int32** | Number of records to return per page (range: 1 to 50) | 

### Return type

[**ServiceDocsAccountPhonenumberGetAll**](ServiceDocsAccountPhonenumberGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPhonenumberPost

> ServiceDocsOrderPhonenumber V1AccountPhonenumberPost(ctx).Phonenumber(phonenumber).Execute()

Purchase Number



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
	phonenumber := []string{"Property_example"} // []string | phonenumber fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberAPI.V1AccountPhonenumberPost(context.Background()).Phonenumber(phonenumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberAPI.V1AccountPhonenumberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountPhonenumberPost`: ServiceDocsOrderPhonenumber
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberAPI.V1AccountPhonenumberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPhonenumberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phonenumber** | **[]string** | phonenumber fields | 

### Return type

[**ServiceDocsOrderPhonenumber**](ServiceDocsOrderPhonenumber.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPhonenumberUnassignPost

> ServiceAPIResponseStatusCodeOnly V1AccountPhonenumberUnassignPost(ctx).Payload(payload).Execute()

Unassign Number



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
	payload := *openapiclient.NewServiceDocsPhonenumberUnassignPayload() // ServiceDocsPhonenumberUnassignPayload | unassign payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberAPI.V1AccountPhonenumberUnassignPost(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberAPI.V1AccountPhonenumberUnassignPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountPhonenumberUnassignPost`: ServiceAPIResponseStatusCodeOnly
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberAPI.V1AccountPhonenumberUnassignPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPhonenumberUnassignPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ServiceDocsPhonenumberUnassignPayload**](ServiceDocsPhonenumberUnassignPayload.md) | unassign payload | 

### Return type

[**ServiceAPIResponseStatusCodeOnly**](ServiceAPIResponseStatusCodeOnly.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PhonenumberSearchGet

> ServiceDocsPhonenumberSearchGetAll V1PhonenumberSearchGet(ctx).AreaCode(areaCode).Quantity(quantity).Execute()

Search New Numbers



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
	areaCode := "areaCode_example" // string | Area code (exactly 3 numeric characters) example: 610 or 484
	quantity := int32(56) // int32 | Number of records to return (range: 1 to 100, defaults to 100 if not provided) (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberAPI.V1PhonenumberSearchGet(context.Background()).AreaCode(areaCode).Quantity(quantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberAPI.V1PhonenumberSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PhonenumberSearchGet`: ServiceDocsPhonenumberSearchGetAll
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberAPI.V1PhonenumberSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PhonenumberSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **areaCode** | **string** | Area code (exactly 3 numeric characters) example: 610 or 484 | 
 **quantity** | **int32** | Number of records to return (range: 1 to 100, defaults to 100 if not provided) | [default to 100]

### Return type

[**ServiceDocsPhonenumberSearchGetAll**](ServiceDocsPhonenumberSearchGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

