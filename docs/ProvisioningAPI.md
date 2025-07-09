# \

All URIs are relative to *http://api.beta.cpaaslabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountAccountIDProvisionFilenameGet**](ProvisioningAPI.md#V1AccountAccountIDProvisionFilenameGet) | **Get** /v1/account/{accountID}/provision/{filename} | Get Config File Details
[**V1ApBrandBrandFamilyFamilyGet**](ProvisioningAPI.md#V1ApBrandBrandFamilyFamilyGet) | **Get** /v1/ap/brand/{brand}/family/{family} | Get Family Details
[**V1ApBrandBrandFamilyFamilyModelGet**](ProvisioningAPI.md#V1ApBrandBrandFamilyFamilyModelGet) | **Get** /v1/ap/brand/{brand}/family/{family}/model | Get Model List
[**V1ApBrandBrandFamilyFamilyModelModelGet**](ProvisioningAPI.md#V1ApBrandBrandFamilyFamilyModelModelGet) | **Get** /v1/ap/brand/{brand}/family/{family}/model/{model} | Get Model Details
[**V1ApBrandBrandFamilyFamilyModelModelTemplateGet**](ProvisioningAPI.md#V1ApBrandBrandFamilyFamilyModelModelTemplateGet) | **Get** /v1/ap/brand/{brand}/family/{family}/model/{model}/template | Get Template List
[**V1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGet**](ProvisioningAPI.md#V1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGet) | **Get** /v1/ap/brand/{brand}/family/{family}/model/{model}/template/{template} | Get Template Details
[**V1ApBrandBrandFamilyGet**](ProvisioningAPI.md#V1ApBrandBrandFamilyGet) | **Get** /v1/ap/brand/{brand}/family | Get Family List
[**V1ApBrandBrandGet**](ProvisioningAPI.md#V1ApBrandBrandGet) | **Get** /v1/ap/brand/{brand} | Get Brand Details
[**V1ApBrandGet**](ProvisioningAPI.md#V1ApBrandGet) | **Get** /v1/ap/brand | Get Brand List
[**V1ApConfigfileGeneratePost**](ProvisioningAPI.md#V1ApConfigfileGeneratePost) | **Post** /v1/ap/configfile/generate | Generate Config File



## V1AccountAccountIDProvisionFilenameGet

> *os.File V1AccountAccountIDProvisionFilenameGet(ctx, accountID, filename).Execute()

Get Config File Details



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
	resp, r, err := apiClient.ProvisioningAPI.V1AccountAccountIDProvisionFilenameGet(context.Background(), accountID, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1AccountAccountIDProvisionFilenameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountAccountIDProvisionFilenameGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1AccountAccountIDProvisionFilenameGet`: %v\n", resp)
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


## V1ApBrandBrandFamilyFamilyGet

> ProvisioningDocsDocsFamilyOutputSingle V1ApBrandBrandFamilyFamilyGet(ctx, brand, family).Execute()

Get Family Details



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
	brand := "brand_example" // string | brand
	family := "family_example" // string | family

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApBrandBrandFamilyFamilyGet(context.Background(), brand, family).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApBrandBrandFamilyFamilyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApBrandBrandFamilyFamilyGet`: ProvisioningDocsDocsFamilyOutputSingle
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApBrandBrandFamilyFamilyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brand** | **string** | brand | 
**family** | **string** | family | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApBrandBrandFamilyFamilyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProvisioningDocsDocsFamilyOutputSingle**](ProvisioningDocsDocsFamilyOutputSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApBrandBrandFamilyFamilyModelGet

> ProvisioningDocsDocsModelOutput V1ApBrandBrandFamilyFamilyModelGet(ctx, brand, family).ModelName(modelName).PageSize(pageSize).StartKey(startKey).Status(status).Execute()

Get Model List



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
	brand := "brand_example" // string | brand
	family := "family_example" // string | family
	modelName := "modelName_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startKey := "startKey_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelGet(context.Background(), brand, family).ModelName(modelName).PageSize(pageSize).StartKey(startKey).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApBrandBrandFamilyFamilyModelGet`: ProvisioningDocsDocsModelOutput
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brand** | **string** | brand | 
**family** | **string** | family | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApBrandBrandFamilyFamilyModelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modelName** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startKey** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**ProvisioningDocsDocsModelOutput**](ProvisioningDocsDocsModelOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApBrandBrandFamilyFamilyModelModelGet

> ProvisioningDocsDocsModelOutputSingle V1ApBrandBrandFamilyFamilyModelModelGet(ctx, brand, family, model).Execute()

Get Model Details



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
	brand := "brand_example" // string | brand
	family := "family_example" // string | family
	model := "model_example" // string | model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelGet(context.Background(), brand, family, model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApBrandBrandFamilyFamilyModelModelGet`: ProvisioningDocsDocsModelOutputSingle
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brand** | **string** | brand | 
**family** | **string** | family | 
**model** | **string** | model | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApBrandBrandFamilyFamilyModelModelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProvisioningDocsDocsModelOutputSingle**](ProvisioningDocsDocsModelOutputSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApBrandBrandFamilyFamilyModelModelTemplateGet

> ProvisioningDocsDocsTemplatesOutput V1ApBrandBrandFamilyFamilyModelModelTemplateGet(ctx, brand, family, model).Firmware(firmware).PageSize(pageSize).StartKey(startKey).Status(status).TemplateName(templateName).Execute()

Get Template List



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
	brand := "brand_example" // string | brand
	family := "family_example" // string | family
	model := "model_example" // string | model
	firmware := "firmware_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startKey := "startKey_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	templateName := "templateName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelTemplateGet(context.Background(), brand, family, model).Firmware(firmware).PageSize(pageSize).StartKey(startKey).Status(status).TemplateName(templateName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelTemplateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApBrandBrandFamilyFamilyModelModelTemplateGet`: ProvisioningDocsDocsTemplatesOutput
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelTemplateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brand** | **string** | brand | 
**family** | **string** | family | 
**model** | **string** | model | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApBrandBrandFamilyFamilyModelModelTemplateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **firmware** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startKey** | **string** |  | 
 **status** | **string** |  | 
 **templateName** | **string** |  | 

### Return type

[**ProvisioningDocsDocsTemplatesOutput**](ProvisioningDocsDocsTemplatesOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGet

> ProvisioningDocsDocsTemplateOutputSingle V1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGet(ctx, brand, family, model, template).Execute()

Get Template Details



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
	brand := "brand_example" // string | brand
	family := "family_example" // string | family
	model := "model_example" // string | model
	template := "template_example" // string | template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGet(context.Background(), brand, family, model, template).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGet`: ProvisioningDocsDocsTemplateOutputSingle
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brand** | **string** | brand | 
**family** | **string** | family | 
**model** | **string** | model | 
**template** | **string** | template | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApBrandBrandFamilyFamilyModelModelTemplateTemplateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ProvisioningDocsDocsTemplateOutputSingle**](ProvisioningDocsDocsTemplateOutputSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApBrandBrandFamilyGet

> ProvisioningDocsDocsFamilyOutput V1ApBrandBrandFamilyGet(ctx, brand).FamilyName(familyName).PageSize(pageSize).StartKey(startKey).Status(status).Execute()

Get Family List



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
	brand := "brand_example" // string | brand
	familyName := "familyName_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startKey := "startKey_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApBrandBrandFamilyGet(context.Background(), brand).FamilyName(familyName).PageSize(pageSize).StartKey(startKey).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApBrandBrandFamilyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApBrandBrandFamilyGet`: ProvisioningDocsDocsFamilyOutput
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApBrandBrandFamilyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brand** | **string** | brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApBrandBrandFamilyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **familyName** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startKey** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**ProvisioningDocsDocsFamilyOutput**](ProvisioningDocsDocsFamilyOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApBrandBrandGet

> ProvisioningDocsDocsBrandOutputSingle V1ApBrandBrandGet(ctx, brand).Execute()

Get Brand Details



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
	brand := "brand_example" // string | brand id to retrieve a brand

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApBrandBrandGet(context.Background(), brand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApBrandBrandGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApBrandBrandGet`: ProvisioningDocsDocsBrandOutputSingle
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApBrandBrandGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brand** | **string** | brand id to retrieve a brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApBrandBrandGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningDocsDocsBrandOutputSingle**](ProvisioningDocsDocsBrandOutputSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApBrandGet

> ProvisioningDocsDocsBrandsOutput V1ApBrandGet(ctx).BrandName(brandName).PageSize(pageSize).StartKey(startKey).Status(status).Execute()

Get Brand List



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
	brandName := "brandName_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	startKey := "startKey_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApBrandGet(context.Background()).BrandName(brandName).PageSize(pageSize).StartKey(startKey).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApBrandGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApBrandGet`: ProvisioningDocsDocsBrandsOutput
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApBrandGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ApBrandGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brandName** | **string** |  | 
 **pageSize** | **int32** |  | 
 **startKey** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**ProvisioningDocsDocsBrandsOutput**](ProvisioningDocsDocsBrandsOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApConfigfileGeneratePost

> ProvisioningDocsDocsConfigFileOutput V1ApConfigfileGeneratePost(ctx).Params(params).Execute()

Generate Config File



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
	params := *openapiclient.NewModelsGenerateConfigFileRequest(*openapiclient.NewModelsConfigFileParameter(), "TemplateFileId_example", "TemplateId_example") // ModelsGenerateConfigFileRequest | body params to generate config file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.V1ApConfigfileGeneratePost(context.Background()).Params(params).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.V1ApConfigfileGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApConfigfileGeneratePost`: ProvisioningDocsDocsConfigFileOutput
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.V1ApConfigfileGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ApConfigfileGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | [**ModelsGenerateConfigFileRequest**](ModelsGenerateConfigFileRequest.md) | body params to generate config file | 

### Return type

[**ProvisioningDocsDocsConfigFileOutput**](ProvisioningDocsDocsConfigFileOutput.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

