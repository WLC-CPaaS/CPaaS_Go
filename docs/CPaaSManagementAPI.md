# \

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MgmtUserGet**](CPaaSManagementAPI.md#V1MgmtUserGet) | **Get** /v1/mgmt/user | Get All CPaaS Users
[**V1MgmtUserPost**](CPaaSManagementAPI.md#V1MgmtUserPost) | **Post** /v1/mgmt/user | Invite CPaaS User
[**V1MgmtUserUserIDDelete**](CPaaSManagementAPI.md#V1MgmtUserUserIDDelete) | **Delete** /v1/mgmt/user/{userID} | Delete CPaaS User
[**V1MgmtUserUserIDGet**](CPaaSManagementAPI.md#V1MgmtUserUserIDGet) | **Get** /v1/mgmt/user/{userID} | Get CPaaS User Details
[**V1MgmtUserUserIDPut**](CPaaSManagementAPI.md#V1MgmtUserUserIDPut) | **Put** /v1/mgmt/user/{userID} | Update CPaaS User Role



## V1MgmtUserGet

> ServiceDocsAdminUserGetAll V1MgmtUserGet(ctx).PageSize(pageSize).StartKey(startKey).Sort(sort).Email(email).Role(role).FirstName(firstName).LastName(lastName).Execute()

Get All CPaaS Users



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
	pageSize := int32(56) // int32 | number of records to return, range 1 to 100 (optional)
	startKey := "startKey_example" // string | unique to fetch next records (optional)
	sort := "sort_example" // string | sorting the records by email(default)/role/first_name/last_name, _A is for ascending and _D is for descending, eg: sort=role_A,email_D (optional)
	email := "email_example" // string | Email (optional)
	role := "role_example" // string | User Role (optional)
	firstName := "firstName_example" // string | First Name (optional)
	lastName := "lastName_example" // string | Last Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CPaaSManagementAPI.V1MgmtUserGet(context.Background()).PageSize(pageSize).StartKey(startKey).Sort(sort).Email(email).Role(role).FirstName(firstName).LastName(lastName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CPaaSManagementAPI.V1MgmtUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MgmtUserGet`: ServiceDocsAdminUserGetAll
	fmt.Fprintf(os.Stdout, "Response from `CPaaSManagementAPI.V1MgmtUserGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MgmtUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | number of records to return, range 1 to 100 | 
 **startKey** | **string** | unique to fetch next records | 
 **sort** | **string** | sorting the records by email(default)/role/first_name/last_name, _A is for ascending and _D is for descending, eg: sort&#x3D;role_A,email_D | 
 **email** | **string** | Email | 
 **role** | **string** | User Role | 
 **firstName** | **string** | First Name | 
 **lastName** | **string** | Last Name | 

### Return type

[**ServiceDocsAdminUserGetAll**](ServiceDocsAdminUserGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MgmtUserPost

> ServiceDocsAdminUserGetSingle V1MgmtUserPost(ctx).ReqBody(reqBody).Execute()

Invite CPaaS User



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
	reqBody := *openapiclient.NewServiceAdminUserAddData("Email_example", "Role_example") // ServiceAdminUserAddData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CPaaSManagementAPI.V1MgmtUserPost(context.Background()).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CPaaSManagementAPI.V1MgmtUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MgmtUserPost`: ServiceDocsAdminUserGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CPaaSManagementAPI.V1MgmtUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MgmtUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reqBody** | [**ServiceAdminUserAddData**](ServiceAdminUserAddData.md) | payload fields | 

### Return type

[**ServiceDocsAdminUserGetSingle**](ServiceDocsAdminUserGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MgmtUserUserIDDelete

> ServiceDocsAdminUserDelete V1MgmtUserUserIDDelete(ctx, userID).Execute()

Delete CPaaS User



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
	userID := "userID_example" // string | User ID, numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CPaaSManagementAPI.V1MgmtUserUserIDDelete(context.Background(), userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CPaaSManagementAPI.V1MgmtUserUserIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MgmtUserUserIDDelete`: ServiceDocsAdminUserDelete
	fmt.Fprintf(os.Stdout, "Response from `CPaaSManagementAPI.V1MgmtUserUserIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User ID, numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MgmtUserUserIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsAdminUserDelete**](ServiceDocsAdminUserDelete.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MgmtUserUserIDGet

> ServiceDocsAdminUserGetSingle V1MgmtUserUserIDGet(ctx, userID).Execute()

Get CPaaS User Details



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
	userID := "userID_example" // string | User ID, numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CPaaSManagementAPI.V1MgmtUserUserIDGet(context.Background(), userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CPaaSManagementAPI.V1MgmtUserUserIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MgmtUserUserIDGet`: ServiceDocsAdminUserGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CPaaSManagementAPI.V1MgmtUserUserIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User ID, numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MgmtUserUserIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsAdminUserGetSingle**](ServiceDocsAdminUserGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MgmtUserUserIDPut

> ServiceDocsAdminUserGetSingle V1MgmtUserUserIDPut(ctx, userID).ReqBody(reqBody).Execute()

Update CPaaS User Role



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
	userID := "userID_example" // string | User ID, numeric
	reqBody := *openapiclient.NewServiceAdminUserEditData("Role_example") // ServiceAdminUserEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CPaaSManagementAPI.V1MgmtUserUserIDPut(context.Background(), userID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CPaaSManagementAPI.V1MgmtUserUserIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MgmtUserUserIDPut`: ServiceDocsAdminUserGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CPaaSManagementAPI.V1MgmtUserUserIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User ID, numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MgmtUserUserIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reqBody** | [**ServiceAdminUserEditData**](ServiceAdminUserEditData.md) | payload fields | 

### Return type

[**ServiceDocsAdminUserGetSingle**](ServiceDocsAdminUserGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

