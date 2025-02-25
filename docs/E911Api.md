# \E911API

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1E911Get**](E911API.md#V1E911Get) | **Get** /v1/e911 | Get E911 List
[**V1E911LocationLocationIDActivatePut**](E911API.md#V1E911LocationLocationIDActivatePut) | **Put** /v1/e911/location/{locationID}/activate | Activate E911 Location
[**V1E911LocationLocationIDDelete**](E911API.md#V1E911LocationLocationIDDelete) | **Delete** /v1/e911/location/{locationID} | Delete E911 Location
[**V1E911LocationValidatePut**](E911API.md#V1E911LocationValidatePut) | **Put** /v1/e911/location/validate | Validate a Location
[**V1E911PhoneNumberDelete**](E911API.md#V1E911PhoneNumberDelete) | **Delete** /v1/e911/{phoneNumber} | Delete E911 Phone Number
[**V1E911PhoneNumberLocationActiveGet**](E911API.md#V1E911PhoneNumberLocationActiveGet) | **Get** /v1/e911/{phoneNumber}/location/active | Get Actvie Location for a Phone Number
[**V1E911PhoneNumberLocationGet**](E911API.md#V1E911PhoneNumberLocationGet) | **Get** /v1/e911/{phoneNumber}/location | Get Location List for Phone Number
[**V1E911Post**](E911API.md#V1E911Post) | **Post** /v1/e911 | Create an E911 Location



## V1E911Get

> ServiceDocE911URIsApiOutput V1E911Get(ctx).Execute()

Get E911 List



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
	resp, r, err := apiClient.E911API.V1E911Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E911API.V1E911Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1E911Get`: ServiceDocE911URIsApiOutput
	fmt.Fprintf(os.Stdout, "Response from `E911API.V1E911Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1E911GetRequest struct via the builder pattern


### Return type

[**ServiceDocE911URIsApiOutput**](ServiceDocE911URIsApiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1E911LocationLocationIDActivatePut

> ServiceDocE911ActiveLocationOutput V1E911LocationLocationIDActivatePut(ctx, locationID).Execute()

Activate E911 Location



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
	locationID := "locationID_example" // string | Location ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.E911API.V1E911LocationLocationIDActivatePut(context.Background(), locationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E911API.V1E911LocationLocationIDActivatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1E911LocationLocationIDActivatePut`: ServiceDocE911ActiveLocationOutput
	fmt.Fprintf(os.Stdout, "Response from `E911API.V1E911LocationLocationIDActivatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationID** | **string** | Location ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1E911LocationLocationIDActivatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocE911ActiveLocationOutput**](ServiceDocE911ActiveLocationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1E911LocationLocationIDDelete

> ServiceDocE911RemoveLocationOutput V1E911LocationLocationIDDelete(ctx, locationID).Execute()

Delete E911 Location



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
	locationID := "locationID_example" // string | Location ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.E911API.V1E911LocationLocationIDDelete(context.Background(), locationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E911API.V1E911LocationLocationIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1E911LocationLocationIDDelete`: ServiceDocE911RemoveLocationOutput
	fmt.Fprintf(os.Stdout, "Response from `E911API.V1E911LocationLocationIDDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationID** | **string** | Location ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1E911LocationLocationIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocE911RemoveLocationOutput**](ServiceDocE911RemoveLocationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1E911LocationValidatePut

> ServiceDocE911ValidateLocationOutput V1E911LocationValidatePut(ctx).ReqBody(reqBody).Execute()

Validate a Location



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
	reqBody := *openapiclient.NewServiceE911ValidateLocationInput(*openapiclient.NewServiceE911LocationInput("Address1_example", "Community_example", "PostalCode_example", "State_example")) // ServiceE911ValidateLocationInput | location details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.E911API.V1E911LocationValidatePut(context.Background()).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E911API.V1E911LocationValidatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1E911LocationValidatePut`: ServiceDocE911ValidateLocationOutput
	fmt.Fprintf(os.Stdout, "Response from `E911API.V1E911LocationValidatePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1E911LocationValidatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reqBody** | [**ServiceE911ValidateLocationInput**](ServiceE911ValidateLocationInput.md) | location details | 

### Return type

[**ServiceDocE911ValidateLocationOutput**](ServiceDocE911ValidateLocationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1E911PhoneNumberDelete

> ServiceDocE911RemoveURIApiOutput V1E911PhoneNumberDelete(ctx, phoneNumber).Execute()

Delete E911 Phone Number



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
	phoneNumber := "phoneNumber_example" // string | Phone Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.E911API.V1E911PhoneNumberDelete(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E911API.V1E911PhoneNumberDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1E911PhoneNumberDelete`: ServiceDocE911RemoveURIApiOutput
	fmt.Fprintf(os.Stdout, "Response from `E911API.V1E911PhoneNumberDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | Phone Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1E911PhoneNumberDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocE911RemoveURIApiOutput**](ServiceDocE911RemoveURIApiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1E911PhoneNumberLocationActiveGet

> ServiceDocE911ActiveLocationURIApiOutput V1E911PhoneNumberLocationActiveGet(ctx, phoneNumber).Execute()

Get Actvie Location for a Phone Number



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
	phoneNumber := "phoneNumber_example" // string | Phone Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.E911API.V1E911PhoneNumberLocationActiveGet(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E911API.V1E911PhoneNumberLocationActiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1E911PhoneNumberLocationActiveGet`: ServiceDocE911ActiveLocationURIApiOutput
	fmt.Fprintf(os.Stdout, "Response from `E911API.V1E911PhoneNumberLocationActiveGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | Phone Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1E911PhoneNumberLocationActiveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocE911ActiveLocationURIApiOutput**](ServiceDocE911ActiveLocationURIApiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1E911PhoneNumberLocationGet

> ServiceDocE911LocationsURIApiOutput V1E911PhoneNumberLocationGet(ctx, phoneNumber).Execute()

Get Location List for Phone Number



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
	phoneNumber := "phoneNumber_example" // string | Phone Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.E911API.V1E911PhoneNumberLocationGet(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E911API.V1E911PhoneNumberLocationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1E911PhoneNumberLocationGet`: ServiceDocE911LocationsURIApiOutput
	fmt.Fprintf(os.Stdout, "Response from `E911API.V1E911PhoneNumberLocationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | Phone Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1E911PhoneNumberLocationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDocE911LocationsURIApiOutput**](ServiceDocE911LocationsURIApiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1E911Post

> ServiceDocE911AddLocationOutput V1E911Post(ctx).ReqBody(reqBody).Execute()

Create an E911 Location



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
	reqBody := *openapiclient.NewServiceE911AddLocationInput(*openapiclient.NewServiceE911LocationInput("Address1_example", "Community_example", "PostalCode_example", "State_example"), *openapiclient.NewServiceE911URIInput("CallerName_example", "Uri_example")) // ServiceE911AddLocationInput | location details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.E911API.V1E911Post(context.Background()).ReqBody(reqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `E911API.V1E911Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1E911Post`: ServiceDocE911AddLocationOutput
	fmt.Fprintf(os.Stdout, "Response from `E911API.V1E911Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1E911PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reqBody** | [**ServiceE911AddLocationInput**](ServiceE911AddLocationInput.md) | location details | 

### Return type

[**ServiceDocE911AddLocationOutput**](ServiceDocE911AddLocationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

