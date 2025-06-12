# \

All URIs are relative to *http://API_HOSTNAME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDProvisionFilenameGet**](ProvisionAPI.md#V1AccountAccountIDProvisionFilenameGet) | **Get** /v1/account/{accountID}/provision/{filename} | 



## V1AccountAccountIDProvisionFilenameGet

> *os.File V1AccountAccountIDProvisionFilenameGet(ctx, accountID, filename).Execute()





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
	filename := "filename_example" // string | Name of config file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisionAPI.V1AccountAccountIDProvisionFilenameGet(context.Background(), accountID, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisionAPI.V1AccountAccountIDProvisionFilenameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDProvisionFilenameGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProvisionAPI.V1AccountAccountIDProvisionFilenameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Account ID, 32 alpha numeric | 
**filename** | **string** | Name of config file | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountAccountIDProvisionFilenameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

