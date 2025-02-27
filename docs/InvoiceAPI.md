# \InvoiceAPI

All URIs are relative to *http://api.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceDailyreportGet**](InvoiceAPI.md#V1FinanceDailyreportGet) | **Get** /v1/finance/dailyreport | Get Daily Account Summary
[**V1FinanceInvoiceGet**](InvoiceAPI.md#V1FinanceInvoiceGet) | **Get** /v1/finance/invoice | Get Invoice Summary
[**V1FinanceMonthlysummaryGet**](InvoiceAPI.md#V1FinanceMonthlysummaryGet) | **Get** /v1/finance/monthlysummary | Get Monthly Summary



## V1FinanceDailyreportGet

> ServiceDocDailyAccountSummaryOutput V1FinanceDailyreportGet(ctx).Year(year).Month(month).ComponentId(componentId).Execute()

Get Daily Account Summary



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
	year := "year_example" // string | Year of the century
	month := "month_example" // string | Month of the year
	componentId := "componentId_example" // string | Component id of the routes

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.V1FinanceDailyreportGet(context.Background()).Year(year).Month(month).ComponentId(componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.V1FinanceDailyreportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceDailyreportGet`: ServiceDocDailyAccountSummaryOutput
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.V1FinanceDailyreportGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceDailyreportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **string** | Year of the century | 
 **month** | **string** | Month of the year | 
 **componentId** | **string** | Component id of the routes | 

### Return type

[**ServiceDocDailyAccountSummaryOutput**](ServiceDocDailyAccountSummaryOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1FinanceInvoiceGet

> ServiceDocInvoiceOutput V1FinanceInvoiceGet(ctx).Year(year).Month(month).Execute()

Get Invoice Summary



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
	year := "year_example" // string | Invoice year
	month := "month_example" // string | Invoice Month

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.V1FinanceInvoiceGet(context.Background()).Year(year).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.V1FinanceInvoiceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceInvoiceGet`: ServiceDocInvoiceOutput
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.V1FinanceInvoiceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceInvoiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **string** | Invoice year | 
 **month** | **string** | Invoice Month | 

### Return type

[**ServiceDocInvoiceOutput**](ServiceDocInvoiceOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1FinanceMonthlysummaryGet

> ServiceDocMonthlySummaryOutput V1FinanceMonthlysummaryGet(ctx).Year(year).Month(month).Execute()

Get Monthly Summary



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
	year := "year_example" // string | Invoice year
	month := "month_example" // string | Invoice month

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.V1FinanceMonthlysummaryGet(context.Background()).Year(year).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.V1FinanceMonthlysummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceMonthlysummaryGet`: ServiceDocMonthlySummaryOutput
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.V1FinanceMonthlysummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceMonthlysummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **string** | Invoice year | 
 **month** | **string** | Invoice month | 

### Return type

[**ServiceDocMonthlySummaryOutput**](ServiceDocMonthlySummaryOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

