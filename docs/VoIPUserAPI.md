# \

All URIs are relative to *http://api.beta.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountidUserGet**](VoIPUserAPI.md#V1AccountAccountidUserGet) | **Get** /v1/account/{accountid}/user | Get User List
[**V1AccountAccountidUserPost**](VoIPUserAPI.md#V1AccountAccountidUserPost) | **Post** /v1/account/{accountid}/user | Create User
[**V1AccountAccountidUserUseridDelete**](VoIPUserAPI.md#V1AccountAccountidUserUseridDelete) | **Delete** /v1/account/{accountid}/user/{userid} | Delete User
[**V1AccountAccountidUserUseridGet**](VoIPUserAPI.md#V1AccountAccountidUserUseridGet) | **Get** /v1/account/{accountid}/user/{userid} | Get User Details
[**V1AccountAccountidUserUseridPut**](VoIPUserAPI.md#V1AccountAccountidUserUseridPut) | **Put** /v1/account/{accountid}/user/{userid} | Update User
[**V1AccountAccountidUserUseridUserauthPost**](VoIPUserAPI.md#V1AccountAccountidUserUseridUserauthPost) | **Post** /v1/account/{accountid}/user/{userid}/userauth | Impersonate a User



## V1AccountAccountidUserGet

> ServiceDocsUserGetAll V1AccountAccountidUserGet(ctx, accountid).StartKey(startKey).PageSize(pageSize).Execute()

Get User List



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
	resp, r, err := apiClient.VoIPUserAPI.V1AccountAccountidUserGet(context.Background(), accountid).StartKey(startKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPUserAPI.V1AccountAccountidUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidUserGet`: ServiceDocsUserGetAll
	fmt.Fprintf(os.Stdout, "Response from `VoIPUserAPI.V1AccountAccountidUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startKey** | **string** | start_key for pagination that was returned as next_start_key from your previous call | 
 **pageSize** | **int32** | number of records to return, range 1 to 50 | 

### Return type

[**ServiceDocsUserGetAll**](ServiceDocsUserGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidUserPost

> ServiceDocsUserGetSingle V1AccountAccountidUserPost(ctx, accountid).User(user).Execute()

Create User



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
	user := *openapiclient.NewServiceVOIPUserAdd2("Email_example", "FirstName_example", "LastName_example") // ServiceVOIPUserAdd2 | user fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPUserAPI.V1AccountAccountidUserPost(context.Background(), accountid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPUserAPI.V1AccountAccountidUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidUserPost`: ServiceDocsUserGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoIPUserAPI.V1AccountAccountidUserPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**ServiceVOIPUserAdd2**](ServiceVOIPUserAdd2.md) | user fields | 

### Return type

[**ServiceDocsUserGetSingle**](ServiceDocsUserGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidUserUseridDelete

> ServiceDocsUserGetSingle V1AccountAccountidUserUseridDelete(ctx, accountid, userid).Execute()

Delete User



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
	userid := "userid_example" // string | User ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPUserAPI.V1AccountAccountidUserUseridDelete(context.Background(), accountid, userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPUserAPI.V1AccountAccountidUserUseridDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidUserUseridDelete`: ServiceDocsUserGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoIPUserAPI.V1AccountAccountidUserUseridDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**userid** | **string** | User ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidUserUseridDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsUserGetSingle**](ServiceDocsUserGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidUserUseridGet

> ServiceDocsUserGetSingle V1AccountAccountidUserUseridGet(ctx, accountid, userid).Execute()

Get User Details



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
	userid := "userid_example" // string | User ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPUserAPI.V1AccountAccountidUserUseridGet(context.Background(), accountid, userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPUserAPI.V1AccountAccountidUserUseridGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidUserUseridGet`: ServiceDocsUserGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoIPUserAPI.V1AccountAccountidUserUseridGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**userid** | **string** | User ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidUserUseridGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsUserGetSingle**](ServiceDocsUserGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidUserUseridPut

> ServiceDocsUserGetSingle V1AccountAccountidUserUseridPut(ctx, accountid, userid).User(user).Execute()

Update User



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
	userid := "userid_example" // string | User ID, 32 alpha numeric
	user := *openapiclient.NewServiceVOIPUserAdd2("Email_example", "FirstName_example", "LastName_example") // ServiceVOIPUserAdd2 | user fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPUserAPI.V1AccountAccountidUserUseridPut(context.Background(), accountid, userid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPUserAPI.V1AccountAccountidUserUseridPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidUserUseridPut`: ServiceDocsUserGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoIPUserAPI.V1AccountAccountidUserUseridPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**userid** | **string** | User ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidUserUseridPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **user** | [**ServiceVOIPUserAdd2**](ServiceVOIPUserAdd2.md) | user fields | 

### Return type

[**ServiceDocsUserGetSingle**](ServiceDocsUserGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountidUserUseridUserauthPost

> ServiceDocsImpersonateUserGetSingle V1AccountAccountidUserUseridUserauthPost(ctx, accountid, userid).User(user).Execute()

Impersonate a User



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
	userid := "userid_example" // string | User ID, 32 alpha numeric
	user := *openapiclient.NewServiceVOIPImpersonateUser("Action_example") // ServiceVOIPImpersonateUser | Payload for impersonate a user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPUserAPI.V1AccountAccountidUserUseridUserauthPost(context.Background(), accountid, userid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPUserAPI.V1AccountAccountidUserUseridUserauthPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountidUserUseridUserauthPost`: ServiceDocsImpersonateUserGetSingle
	fmt.Fprintf(os.Stdout, "Response from `VoIPUserAPI.V1AccountAccountidUserUseridUserauthPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account ID, 32 alpha numeric | 
**userid** | **string** | User ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountidUserUseridUserauthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **user** | [**ServiceVOIPImpersonateUser**](ServiceVOIPImpersonateUser.md) | Payload for impersonate a user | 

### Return type

[**ServiceDocsImpersonateUserGetSingle**](ServiceDocsImpersonateUserGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

