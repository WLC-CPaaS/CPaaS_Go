# \

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDCallqueueGet**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueGet) | **Get** /v1/account/{accountID}/callqueue | Get Call Queues
[**V1AccountAccountIDCallqueuePost**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueuePost) | **Post** /v1/account/{accountID}/callqueue | Create Call Queue
[**V1AccountAccountIDCallqueueQueueIDDelete**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueQueueIDDelete) | **Delete** /v1/account/{accountID}/callqueue/{queueID} | Delete Call Queue
[**V1AccountAccountIDCallqueueQueueIDGet**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueQueueIDGet) | **Get** /v1/account/{accountID}/callqueue/{queueID} | Get Call Queue Details
[**V1AccountAccountIDCallqueueQueueIDPut**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueQueueIDPut) | **Put** /v1/account/{accountID}/callqueue/{queueID} | Update Call Queue
[**V1AccountAccountIDCallqueueQueueIDStatusGet**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueQueueIDStatusGet) | **Get** /v1/account/{accountID}/callqueue/{queueID}/status | Get Call Queue Status
[**V1AccountAccountIDQueuerolesGet**](CallQueueManagementAPI.md#V1AccountAccountIDQueuerolesGet) | **Get** /v1/account/{accountID}/queueroles | Get Queue Roles of Account
[**V1AccountAccountIDQueuerolesQueueIDPost**](CallQueueManagementAPI.md#V1AccountAccountIDQueuerolesQueueIDPost) | **Post** /v1/account/{accountID}/queueroles/{queueID} | Assign Queue Role to Call Queue



## V1AccountAccountIDCallqueueGet

> ServiceDocsCallQueueGetAll V1AccountAccountIDCallqueueGet(ctx, accountID).Execute()

Get Call Queues



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
	resp, r, err := apiClient.CallQueueManagementAPI.V1AccountAccountIDCallqueueGet(context.Background(), accountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueManagementAPI.V1AccountAccountIDCallqueueGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallqueueGet`: ServiceDocsCallQueueGetAll
	fmt.Fprintf(os.Stdout, "Response from `CallQueueManagementAPI.V1AccountAccountIDCallqueueGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallqueueGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsCallQueueGetAll**](ServiceDocsCallQueueGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallqueuePost

> ServiceDocsCallQueueGetSingle V1AccountAccountIDCallqueuePost(ctx, accountID).ReqBody(reqBody).Execute()

Create Call Queue



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
	reqBody := *openapiclient.NewServiceVOIPCallQueueAddEditData("Name_example") // ServiceVOIPCallQueueAddEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueManagementAPI.V1AccountAccountIDCallqueuePost(context.Background(), accountID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueManagementAPI.V1AccountAccountIDCallqueuePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallqueuePost`: ServiceDocsCallQueueGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallQueueManagementAPI.V1AccountAccountIDCallqueuePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallqueuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reqBody** | [**ServiceVOIPCallQueueAddEditData**](ServiceVOIPCallQueueAddEditData.md) | payload fields | 

### Return type

[**ServiceDocsCallQueueGetSingle**](ServiceDocsCallQueueGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallqueueQueueIDDelete

> ServiceDocsCallQueueGetSingle V1AccountAccountIDCallqueueQueueIDDelete(ctx, accountID, queueID).Execute()

Delete Call Queue



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
	queueID := "queueID_example" // string | Queue ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDDelete(context.Background(), accountID, queueID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallqueueQueueIDDelete`: ServiceDocsCallQueueGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**queueID** | **string** | Queue ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallqueueQueueIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCallQueueGetSingle**](ServiceDocsCallQueueGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallqueueQueueIDGet

> ServiceDocsCallQueueGetSingle V1AccountAccountIDCallqueueQueueIDGet(ctx, accountID, queueID).Execute()

Get Call Queue Details



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
	queueID := "queueID_example" // string | Queue ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDGet(context.Background(), accountID, queueID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallqueueQueueIDGet`: ServiceDocsCallQueueGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**queueID** | **string** | Queue ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallqueueQueueIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCallQueueGetSingle**](ServiceDocsCallQueueGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallqueueQueueIDPut

> ServiceDocsCallQueueGetSingle V1AccountAccountIDCallqueueQueueIDPut(ctx, accountID, queueID).ReqBody(reqBody).Execute()

Update Call Queue



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
	queueID := "queueID_example" // string | Queue ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPCallQueueAddEditData("Name_example") // ServiceVOIPCallQueueAddEditData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDPut(context.Background(), accountID, queueID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallqueueQueueIDPut`: ServiceDocsCallQueueGetSingle
	fmt.Fprintf(os.Stdout, "Response from `CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**queueID** | **string** | Queue ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallqueueQueueIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPCallQueueAddEditData**](ServiceVOIPCallQueueAddEditData.md) | payload fields | 

### Return type

[**ServiceDocsCallQueueGetSingle**](ServiceDocsCallQueueGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDCallqueueQueueIDStatusGet

> ServiceDocsCallQueueGetSingleStatus V1AccountAccountIDCallqueueQueueIDStatusGet(ctx, accountID, queueID).Execute()

Get Call Queue Status



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
	queueID := "queueID_example" // string | Queue ID, 32 alpha numeric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDStatusGet(context.Background(), accountID, queueID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDCallqueueQueueIDStatusGet`: ServiceDocsCallQueueGetSingleStatus
	fmt.Fprintf(os.Stdout, "Response from `CallQueueManagementAPI.V1AccountAccountIDCallqueueQueueIDStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**queueID** | **string** | Queue ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDCallqueueQueueIDStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceDocsCallQueueGetSingleStatus**](ServiceDocsCallQueueGetSingleStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDQueuerolesGet

> ServiceDocsCallQueueGetRoles V1AccountAccountIDQueuerolesGet(ctx, accountID).Execute()

Get Queue Roles of Account



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
	resp, r, err := apiClient.CallQueueManagementAPI.V1AccountAccountIDQueuerolesGet(context.Background(), accountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueManagementAPI.V1AccountAccountIDQueuerolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDQueuerolesGet`: ServiceDocsCallQueueGetRoles
	fmt.Fprintf(os.Stdout, "Response from `CallQueueManagementAPI.V1AccountAccountIDQueuerolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDQueuerolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocsCallQueueGetRoles**](ServiceDocsCallQueueGetRoles.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountAccountIDQueuerolesQueueIDPost

> ServiceAPIResponse V1AccountAccountIDQueuerolesQueueIDPost(ctx, accountID, queueID).ReqBody(reqBody).Execute()

Assign Queue Role to Call Queue



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
	queueID := "queueID_example" // string | Queue ID, 32 alpha numeric
	reqBody := *openapiclient.NewServiceVOIPCallQueueRoleAssignData("Action_example", []map[string][]string{map[string][]string{"key": []string{"Inner_example"}}}, false) // ServiceVOIPCallQueueRoleAssignData | payload fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallQueueManagementAPI.V1AccountAccountIDQueuerolesQueueIDPost(context.Background(), accountID, queueID).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallQueueManagementAPI.V1AccountAccountIDQueuerolesQueueIDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDQueuerolesQueueIDPost`: ServiceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `CallQueueManagementAPI.V1AccountAccountIDQueuerolesQueueIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**queueID** | **string** | Queue ID, 32 alpha numeric | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDQueuerolesQueueIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reqBody** | [**ServiceVOIPCallQueueRoleAssignData**](ServiceVOIPCallQueueRoleAssignData.md) | payload fields | 

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

