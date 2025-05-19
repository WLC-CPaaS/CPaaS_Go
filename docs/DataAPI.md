# \

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DataCallDailySummaryGet**](DataAPI.md#V1DataCallDailySummaryGet) | **Get** /v1/data/call_daily_summary | Get Call Daily Summary List
[**V1DataCallDetailGet**](DataAPI.md#V1DataCallDetailGet) | **Get** /v1/data/call_detail | Get Call Detail List
[**V1DataCallMonthlySummaryGet**](DataAPI.md#V1DataCallMonthlySummaryGet) | **Get** /v1/data/call_monthly_summary | Get Call Detail List
[**V1DataEndpointListGet**](DataAPI.md#V1DataEndpointListGet) | **Get** /v1/data/endpoint_list | Get Endpoint List
[**V1DataEventDailySummaryGet**](DataAPI.md#V1DataEventDailySummaryGet) | **Get** /v1/data/event_daily_summary | Get Event Daily Summary List
[**V1DataEventDetailGet**](DataAPI.md#V1DataEventDetailGet) | **Get** /v1/data/event_detail | Get Event Details
[**V1DataEventMonthlySummaryGet**](DataAPI.md#V1DataEventMonthlySummaryGet) | **Get** /v1/data/event_monthly_summary | Get Event Monthly Summary List
[**V1DataFeatureDailySummaryGet**](DataAPI.md#V1DataFeatureDailySummaryGet) | **Get** /v1/data/feature_daily_summary | Get Feature Daily Summary List
[**V1DataFeatureMonthlySummaryGet**](DataAPI.md#V1DataFeatureMonthlySummaryGet) | **Get** /v1/data/feature_monthly_summary | Get Feature Monthly Summary List



## V1DataCallDailySummaryGet

> ServiceDocsCallDailySummary V1DataCallDailySummaryGet(ctx).AccountId(accountId).CallType(callType).EndDate(endDate).PageSize(pageSize).StartDate(startDate).StartKey(startKey).Execute()

Get Call Daily Summary List



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
	accountId := "accountId_example" // string |  (optional)
	callType := "callType_example" // string |  (optional)
	endDate := "endDate_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startDate := "startDate_example" // string |  (optional)
	startKey := "startKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataCallDailySummaryGet(context.Background()).AccountId(accountId).CallType(callType).EndDate(endDate).PageSize(pageSize).StartDate(startDate).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataCallDailySummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataCallDailySummaryGet`: ServiceDocsCallDailySummary
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataCallDailySummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataCallDailySummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **callType** | **string** |  | 
 **endDate** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startDate** | **string** |  | 
 **startKey** | **string** |  | 

### Return type

[**ServiceDocsCallDailySummary**](ServiceDocsCallDailySummary.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataCallDetailGet

> ServiceDocsCallDetail V1DataCallDetailGet(ctx).Account(account).CallType(callType).CalleeName(calleeName).CalleeNumber(calleeNumber).CallerName(callerName).CallerNumber(callerNumber).EndDate(endDate).PageSize(pageSize).StartDate(startDate).StartKey(startKey).Execute()

Get Call Detail List



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
	account := "account_example" // string |  (optional)
	callType := "callType_example" // string |  (optional)
	calleeName := "calleeName_example" // string |  (optional)
	calleeNumber := "calleeNumber_example" // string |  (optional)
	callerName := "callerName_example" // string |  (optional)
	callerNumber := "callerNumber_example" // string |  (optional)
	endDate := "endDate_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startDate := "startDate_example" // string |  (optional)
	startKey := "startKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataCallDetailGet(context.Background()).Account(account).CallType(callType).CalleeName(calleeName).CalleeNumber(calleeNumber).CallerName(callerName).CallerNumber(callerNumber).EndDate(endDate).PageSize(pageSize).StartDate(startDate).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataCallDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataCallDetailGet`: ServiceDocsCallDetail
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataCallDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataCallDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** |  | 
 **callType** | **string** |  | 
 **calleeName** | **string** |  | 
 **calleeNumber** | **string** |  | 
 **callerName** | **string** |  | 
 **callerNumber** | **string** |  | 
 **endDate** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startDate** | **string** |  | 
 **startKey** | **string** |  | 

### Return type

[**ServiceDocsCallDetail**](ServiceDocsCallDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataCallMonthlySummaryGet

> ServiceDocsCallMonthlySummary V1DataCallMonthlySummaryGet(ctx).Account(account).CallType(callType).EndMonth(endMonth).EndYear(endYear).PageSize(pageSize).StartKey(startKey).StartMonth(startMonth).StartYear(startYear).Execute()

Get Call Detail List



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
	account := "account_example" // string |  (optional)
	callType := "callType_example" // string |  (optional)
	endMonth := int32(56) // int32 |  (optional)
	endYear := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startKey := "startKey_example" // string |  (optional)
	startMonth := int32(56) // int32 |  (optional)
	startYear := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataCallMonthlySummaryGet(context.Background()).Account(account).CallType(callType).EndMonth(endMonth).EndYear(endYear).PageSize(pageSize).StartKey(startKey).StartMonth(startMonth).StartYear(startYear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataCallMonthlySummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataCallMonthlySummaryGet`: ServiceDocsCallMonthlySummary
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataCallMonthlySummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataCallMonthlySummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** |  | 
 **callType** | **string** |  | 
 **endMonth** | **int32** |  | 
 **endYear** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **startKey** | **string** |  | 
 **startMonth** | **int32** |  | 
 **startYear** | **int32** |  | 

### Return type

[**ServiceDocsCallMonthlySummary**](ServiceDocsCallMonthlySummary.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataEndpointListGet

> ServiceDocsEndpointList V1DataEndpointListGet(ctx).EndpointName(endpointName).FeatureName(featureName).PageSize(pageSize).StartKey(startKey).TransactionType(transactionType).Version(version).Execute()

Get Endpoint List



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
	endpointName := "endpointName_example" // string |  (optional)
	featureName := "featureName_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startKey := "startKey_example" // string |  (optional)
	transactionType := "transactionType_example" // string |  (optional)
	version := "version_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataEndpointListGet(context.Background()).EndpointName(endpointName).FeatureName(featureName).PageSize(pageSize).StartKey(startKey).TransactionType(transactionType).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataEndpointListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataEndpointListGet`: ServiceDocsEndpointList
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataEndpointListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataEndpointListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointName** | **string** |  | 
 **featureName** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startKey** | **string** |  | 
 **transactionType** | **string** |  | 
 **version** | **string** |  | 

### Return type

[**ServiceDocsEndpointList**](ServiceDocsEndpointList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataEventDailySummaryGet

> ServiceDocsEventDailySummary V1DataEventDailySummaryGet(ctx).AccountId(accountId).Component(component).EndDate(endDate).PageSize(pageSize).StartDate(startDate).StartKey(startKey).Execute()

Get Event Daily Summary List



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
	accountId := "accountId_example" // string |  (optional)
	component := "component_example" // string |  (optional)
	endDate := "endDate_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startDate := "startDate_example" // string |  (optional)
	startKey := "startKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataEventDailySummaryGet(context.Background()).AccountId(accountId).Component(component).EndDate(endDate).PageSize(pageSize).StartDate(startDate).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataEventDailySummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataEventDailySummaryGet`: ServiceDocsEventDailySummary
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataEventDailySummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataEventDailySummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **component** | **string** |  | 
 **endDate** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startDate** | **string** |  | 
 **startKey** | **string** |  | 

### Return type

[**ServiceDocsEventDailySummary**](ServiceDocsEventDailySummary.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataEventDetailGet

> ServiceDocsEventDetail V1DataEventDetailGet(ctx).AccountId(accountId).Component(component).EndDateTime(endDateTime).EventName(eventName).ExecStatus(execStatus).PageSize(pageSize).StartDateTime(startDateTime).StartKey(startKey).Username(username).Execute()

Get Event Details



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
	accountId := "accountId_example" // string |  (optional)
	component := "component_example" // string |  (optional)
	endDateTime := "endDateTime_example" // string |  (optional)
	eventName := "eventName_example" // string |  (optional)
	execStatus := "execStatus_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startDateTime := "startDateTime_example" // string |  (optional)
	startKey := "startKey_example" // string |  (optional)
	username := "username_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataEventDetailGet(context.Background()).AccountId(accountId).Component(component).EndDateTime(endDateTime).EventName(eventName).ExecStatus(execStatus).PageSize(pageSize).StartDateTime(startDateTime).StartKey(startKey).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataEventDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataEventDetailGet`: ServiceDocsEventDetail
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataEventDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataEventDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **component** | **string** |  | 
 **endDateTime** | **string** |  | 
 **eventName** | **string** |  | 
 **execStatus** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startDateTime** | **string** |  | 
 **startKey** | **string** |  | 
 **username** | **string** |  | 

### Return type

[**ServiceDocsEventDetail**](ServiceDocsEventDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataEventMonthlySummaryGet

> ServiceDocsEventMonthlySummary V1DataEventMonthlySummaryGet(ctx).AccountId(accountId).Component(component).EndMonth(endMonth).EndYear(endYear).PageSize(pageSize).StartKey(startKey).StartMonth(startMonth).StartYear(startYear).Execute()

Get Event Monthly Summary List



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
	accountId := "accountId_example" // string |  (optional)
	component := "component_example" // string |  (optional)
	endMonth := int32(56) // int32 |  (optional)
	endYear := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startKey := "startKey_example" // string |  (optional)
	startMonth := int32(56) // int32 |  (optional)
	startYear := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataEventMonthlySummaryGet(context.Background()).AccountId(accountId).Component(component).EndMonth(endMonth).EndYear(endYear).PageSize(pageSize).StartKey(startKey).StartMonth(startMonth).StartYear(startYear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataEventMonthlySummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataEventMonthlySummaryGet`: ServiceDocsEventMonthlySummary
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataEventMonthlySummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataEventMonthlySummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **component** | **string** |  | 
 **endMonth** | **int32** |  | 
 **endYear** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **startKey** | **string** |  | 
 **startMonth** | **int32** |  | 
 **startYear** | **int32** |  | 

### Return type

[**ServiceDocsEventMonthlySummary**](ServiceDocsEventMonthlySummary.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataFeatureDailySummaryGet

> ServiceDocsFeatureDailySummary V1DataFeatureDailySummaryGet(ctx).EndDate(endDate).FeatureName(featureName).PageSize(pageSize).StartDate(startDate).StartKey(startKey).Execute()

Get Feature Daily Summary List



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
	endDate := "endDate_example" // string |  (optional)
	featureName := "featureName_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startDate := "startDate_example" // string |  (optional)
	startKey := "startKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataFeatureDailySummaryGet(context.Background()).EndDate(endDate).FeatureName(featureName).PageSize(pageSize).StartDate(startDate).StartKey(startKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataFeatureDailySummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataFeatureDailySummaryGet`: ServiceDocsFeatureDailySummary
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataFeatureDailySummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataFeatureDailySummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endDate** | **string** |  | 
 **featureName** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startDate** | **string** |  | 
 **startKey** | **string** |  | 

### Return type

[**ServiceDocsFeatureDailySummary**](ServiceDocsFeatureDailySummary.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DataFeatureMonthlySummaryGet

> ServiceDocsFeatureMonthlySummary V1DataFeatureMonthlySummaryGet(ctx).EndMonth(endMonth).EndYear(endYear).FeatureName(featureName).PageSize(pageSize).StartKey(startKey).StartMonth(startMonth).StartYear(startYear).Execute()

Get Feature Monthly Summary List



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
	endMonth := int32(56) // int32 |  (optional)
	endYear := int32(56) // int32 |  (optional)
	featureName := "featureName_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startKey := "startKey_example" // string |  (optional)
	startMonth := int32(56) // int32 |  (optional)
	startYear := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.V1DataFeatureMonthlySummaryGet(context.Background()).EndMonth(endMonth).EndYear(endYear).FeatureName(featureName).PageSize(pageSize).StartKey(startKey).StartMonth(startMonth).StartYear(startYear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.V1DataFeatureMonthlySummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DataFeatureMonthlySummaryGet`: ServiceDocsFeatureMonthlySummary
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.V1DataFeatureMonthlySummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DataFeatureMonthlySummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endMonth** | **int32** |  | 
 **endYear** | **int32** |  | 
 **featureName** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startKey** | **string** |  | 
 **startMonth** | **int32** |  | 
 **startYear** | **int32** |  | 

### Return type

[**ServiceDocsFeatureMonthlySummary**](ServiceDocsFeatureMonthlySummary.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

