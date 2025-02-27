# WebhookAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1WebhookAccountAccountIDGet**](WebhookAPI.md#V1WebhookAccountAccountIDGet) | **Get** /v1/webhook/account/{accountID} | Get Webhook List |
| [**V1WebhookAccountAccountIDPost**](WebhookAPI.md#V1WebhookAccountAccountIDPost) | **Post** /v1/webhook/account/{accountID} | Create Webhook |
| [**V1WebhookAccountAccountIDWebhookIDDelete**](WebhookAPI.md#V1WebhookAccountAccountIDWebhookIDDelete) | **Delete** /v1/webhook/account/{accountID}/{webhookID} | Delete Webhook |
| [**V1WebhookAccountAccountIDWebhookIDGet**](WebhookAPI.md#V1WebhookAccountAccountIDWebhookIDGet) | **Get** /v1/webhook/account/{accountID}/{webhookID} | Get Webhook Details |
| [**V1WebhookAccountAccountIDWebhookIDPut**](WebhookAPI.md#V1WebhookAccountAccountIDWebhookIDPut) | **Put** /v1/webhook/account/{accountID}/{webhookID} | Update Webhook |



## V1WebhookAccountAccountIDGet

> ServiceDocsWebhookGetAll V1WebhookAccountAccountIDGet(accountID, pageSize, currentPage)

Get Webhook List

Retrieve the webhook list in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.WebhookAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        WebhookAPI apiInstance = new WebhookAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        int32 currentPage = 56; // int32 | Current Page
        try {
            ServiceDocsWebhookGetAll result = apiInstance.V1WebhookAccountAccountIDGet(accountID, pageSize, currentPage);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling WebhookAPI#V1WebhookAccountAccountIDGet");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountID** | **string**| Account ID | |
| **pageSize** | **int32**| number of records to return, range 1 to 50 | [optional] |
| **currentPage** | **int32**| Current Page | [optional] |

### Return type

[**ServiceDocsWebhookGetAll**](ServiceDocsWebhookGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **404** | Not Found |  -  |
| **500** | Internal Server Error |  -  |


## V1WebhookAccountAccountIDPost

> ServiceDocsWebhookGetSingle V1WebhookAccountAccountIDPost(accountID, body)

Create Webhook

Create a webhook for a specific account ID.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.WebhookAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        WebhookAPI apiInstance = new WebhookAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID
        ServiceWebhookAdd body = ; // ServiceWebhookAdd | Webhook data
        try {
            ServiceDocsWebhookGetSingle result = apiInstance.V1WebhookAccountAccountIDPost(accountID, body);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling WebhookAPI#V1WebhookAccountAccountIDPost");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountID** | **string**| Account ID | |
| **body** | [**ServiceWebhookAdd**](ServiceWebhookAdd.md)| Webhook data | |

### Return type

[**ServiceDocsWebhookGetSingle**](ServiceDocsWebhookGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **500** | Internal Server Error |  -  |


## V1WebhookAccountAccountIDWebhookIDDelete

> ServiceDocsWebhookDelete V1WebhookAccountAccountIDWebhookIDDelete(accountID, webhookID)

Delete Webhook

Remove a webhook identified by its ID for a particular account ID.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.WebhookAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        WebhookAPI apiInstance = new WebhookAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID
        int32 webhookID = 56; // int32 | Webhook ID
        try {
            ServiceDocsWebhookDelete result = apiInstance.V1WebhookAccountAccountIDWebhookIDDelete(accountID, webhookID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling WebhookAPI#V1WebhookAccountAccountIDWebhookIDDelete");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountID** | **string**| Account ID | |
| **webhookID** | **int32**| Webhook ID | |

### Return type

[**ServiceDocsWebhookDelete**](ServiceDocsWebhookDelete.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **404** | Not Found |  -  |
| **500** | Internal Server Error |  -  |


## V1WebhookAccountAccountIDWebhookIDGet

> ServiceDocsWebhookGetSingle V1WebhookAccountAccountIDWebhookIDGet(accountID, webhookID)

Get Webhook Details

Access details about a single webhook ID for an individual account ID.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.WebhookAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        WebhookAPI apiInstance = new WebhookAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID
        int32 webhookID = 56; // int32 | Webhook ID
        try {
            ServiceDocsWebhookGetSingle result = apiInstance.V1WebhookAccountAccountIDWebhookIDGet(accountID, webhookID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling WebhookAPI#V1WebhookAccountAccountIDWebhookIDGet");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountID** | **string**| Account ID | |
| **webhookID** | **int32**| Webhook ID | |

### Return type

[**ServiceDocsWebhookGetSingle**](ServiceDocsWebhookGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **404** | Not Found |  -  |
| **500** | Internal Server Error |  -  |


## V1WebhookAccountAccountIDWebhookIDPut

> ServiceDocsWebhookGetSingle V1WebhookAccountAccountIDWebhookIDPut(accountID, webhookID, body)

Update Webhook

Update a webhook identified by its ID for a distinct account ID.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.WebhookAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        WebhookAPI apiInstance = new WebhookAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID
        string webhookID = webhookID_example; // string | Webhook ID
        ServiceWebhookEdit body = ; // ServiceWebhookEdit | Updated webhook data
        try {
            ServiceDocsWebhookGetSingle result = apiInstance.V1WebhookAccountAccountIDWebhookIDPut(accountID, webhookID, body);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling WebhookAPI#V1WebhookAccountAccountIDWebhookIDPut");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountID** | **string**| Account ID | |
| **webhookID** | **string**| Webhook ID | |
| **body** | [**ServiceWebhookEdit**](ServiceWebhookEdit.md)| Updated webhook data | |

### Return type

[**ServiceDocsWebhookGetSingle**](ServiceDocsWebhookGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **404** | Not Found |  -  |
| **500** | Internal Server Error |  -  |

