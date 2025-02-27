# \AccountAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountidChildrenGet**](AccountAPI.md#V1AccountAccountidChildrenGet) | **Get** /v1/account/{accountid}/children | Get Sub Account List
[**V1AccountAccountidDelete**](AccountAPI.md#V1AccountAccountidDelete) | **Delete** /v1/account/{accountid} | Delete Account
[**V1AccountAccountidDnsrecordGet**](AccountAPI.md#V1AccountAccountidDnsrecordGet) | **Get** /v1/account/{accountid}/dnsrecord | 
[**V1AccountAccountidDnsrecordPost**](AccountAPI.md#V1AccountAccountidDnsrecordPost) | **Post** /v1/account/{accountid}/dnsrecord | 
[**V1AccountAccountidDnsrecordPut**](AccountAPI.md#V1AccountAccountidDnsrecordPut) | **Put** /v1/account/{accountid}/dnsrecord | 
[**V1AccountAccountidGet**](AccountAPI.md#V1AccountAccountidGet) | **Get** /v1/account/{accountid} | Get Account Details
[**V1AccountAccountidLimitGet**](AccountAPI.md#V1AccountAccountidLimitGet) | **Get** /v1/account/{accountid}/limit | Get Account Limits
[**V1AccountAccountidLimitPut**](AccountAPI.md#V1AccountAccountidLimitPut) | **Put** /v1/account/{accountid}/limit | Set Account Limits
[**V1AccountAccountidPost**](AccountAPI.md#V1AccountAccountidPost) | **Post** /v1/account/{accountid} | Create Sub Account
[**V1AccountAccountidPut**](AccountAPI.md#V1AccountAccountidPut) | **Put** /v1/account/{accountid} | Update Account
[**V1AccountApikeyGet**](AccountAPI.md#V1AccountApikeyGet) | **Get** /v1/account/apikey | 
[**V1AccountGet**](AccountAPI.md#V1AccountGet) | **Get** /v1/account | Get Account List
[**V1AccountPost**](AccountAPI.md#V1AccountPost) | **Post** /v1/account | Create Account



## V1AccountAccountidChildrenGet

> ServiceDocsAccountGetAll V1AccountAccountidChildrenGet(ctx, accountid).StartKey(startKey).PageSize(pageSize).Execute()

Get Sub Account List



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
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidChildrenGet(context.Background(), accountid).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidChildrenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidChildrenGet`: ServiceDocsAccountGetAll
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidChildrenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidChildrenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsAccountGetAll**](ServiceDocsAccountGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDelete

> ServiceDocsAccountGetSingle V1AccountAccountidDelete(ctx, accountid).Execute()

Delete Account



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidDelete(context.Background(), accountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDelete`: ServiceDocsAccountGetSingle
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDnsrecordGet

> ServiceDocsAccountGetSingle V1AccountAccountidDnsrecordGet(ctx, accountid).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidDnsrecordGet(context.Background(), accountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidDnsrecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDnsrecordGet`: ServiceDocsAccountGetSingle
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidDnsrecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDnsrecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDnsrecordPost

> ServiceDocsAccountGetSingle V1AccountAccountidDnsrecordPost(ctx, accountid).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidDnsrecordPost(context.Background(), accountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidDnsrecordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDnsrecordPost`: ServiceDocsAccountGetSingle
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidDnsrecordPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDnsrecordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidDnsrecordPut

> ServiceDocsAccountGetSingle V1AccountAccountidDnsrecordPut(ctx, accountid).Dnsrecord(dnsrecord).Execute()





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
	dnsrecord := *openapiclient.NewServiceUpdateRecordTypeForAccount("RecordType_example") // ServiceUpdateRecordTypeForAccount | record type fields with value SRV, CNAME

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidDnsrecordPut(context.Background(), accountid).Dnsrecord(dnsrecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidDnsrecordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidDnsrecordPut`: ServiceDocsAccountGetSingle
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidDnsrecordPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidDnsrecordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsrecord** | [**ServiceUpdateRecordTypeForAccount**](ServiceUpdateRecordTypeForAccount.md) | record type fields with value SRV, CNAME | 

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidGet

> ServiceDocsAccountGetSingle V1AccountAccountidGet(ctx, accountid).Execute()

Get Account Details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidGet(context.Background(), accountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidGet`: ServiceDocsAccountGetSingle
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidLimitGet

> ServiceDocsAccountLimit V1AccountAccountidLimitGet(ctx, accountid).Execute()

Get Account Limits



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidLimitGet(context.Background(), accountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidLimitGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidLimitGet`: ServiceDocsAccountLimit
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidLimitGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidLimitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsAccountLimit**](ServiceDocsAccountLimit.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidLimitPut

> ServiceDocsAccountLimit V1AccountAccountidLimitPut(ctx, accountid).Limit(limit).Execute()

Set Account Limits



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
	limit := *openapiclient.NewServiceVOIPAccountLimit2(int32(123), int32(123), int32(123)) // ServiceVOIPAccountLimit2 | account fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidLimitPut(context.Background(), accountid).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidLimitPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidLimitPut`: ServiceDocsAccountLimit
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidLimitPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidLimitPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | [**ServiceVOIPAccountLimit2**](ServiceVOIPAccountLimit2.md) | account fields | 

### Return type

[**ServiceDocsAccountLimit**](ServiceDocsAccountLimit.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidPost

> ServiceDocsAccountGetSingle V1AccountAccountidPost(ctx, accountid).Account(account).Execute()

Create Sub Account



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
	account := *openapiclient.NewServiceVOIPAccountAddData("Name_example", "Timezone_example") // ServiceVOIPAccountAddData | account fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidPost(context.Background(), accountid).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidPost`: ServiceDocsAccountGetSingle
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**ServiceVOIPAccountAddData**](ServiceVOIPAccountAddData.md) | account fields | 

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidPut

> ServiceDocsAccountGetSingle V1AccountAccountidPut(ctx, accountid).Account(account).Execute()

Update Account



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
	account := *openapiclient.NewServiceVOIPAccountEditData("Name_example", "Timezone_example") // ServiceVOIPAccountEditData | account fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountAccountidPut(context.Background(), accountid).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountAccountidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidPut`: ServiceDocsAccountGetSingle
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountAccountidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**ServiceVOIPAccountEditData**](ServiceVOIPAccountEditData.md) | account fields | 

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountApikeyGet

> ServiceDocsAccountAPIKey V1AccountApikeyGet(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountApikeyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountApikeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountApikeyGet`: ServiceDocsAccountAPIKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountApikeyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountApikeyGetRequest struct via the builder pattern


### Return type

[**ServiceDocsAccountAPIKey**](ServiceDocsAccountAPIKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountGet

> ServiceDocsAccountGetAll V1AccountGet(ctx).StartKey(startKey).PageSize(pageSize).Execute()

Get Account List



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
	startKey := "startKey_example" // string | start_key for pagination that was returned as next_start_key from your previous call (optional)
	pageSize := int32(56) // int32 | number of records to return, range 1 to 50 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountGet(context.Background()).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountGet`: ServiceDocsAccountGetAll
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsAccountGetAll**](ServiceDocsAccountGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountPost

> ServiceDocsAccountGetSingle V1AccountPost(ctx).Account(account).Execute()

Create Account



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
	account := *openapiclient.NewServiceVOIPAccountAddData("Name_example", "Timezone_example") // ServiceVOIPAccountAddData | account fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.V1AccountPost(context.Background()).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.V1AccountPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountPost`: ServiceDocsAccountGetSingle
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.V1AccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | [**ServiceVOIPAccountAddData**](ServiceVOIPAccountAddData.md) | account fields | 

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

