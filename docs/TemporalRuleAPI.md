# TemporalRuleAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDTemporalruleGet**](TemporalRuleAPI.md#V1AccountAccountIDTemporalruleGet) | **Get** /v1/account/{accountID}/temporalrule | Get Temporal Rule List |
| [**V1AccountAccountIDTemporalrulePost**](TemporalRuleAPI.md#V1AccountAccountIDTemporalrulePost) | **Post** /v1/account/{accountID}/temporalrule | Create Temporal Rule |
| [**V1AccountAccountIDTemporalruleTemporalRuleIDDelete**](TemporalRuleAPI.md#V1AccountAccountIDTemporalruleTemporalRuleIDDelete) | **Delete** /v1/account/{accountID}/temporalrule/{temporalRuleID} | Delete Temporal Rule |
| [**V1AccountAccountIDTemporalruleTemporalRuleIDGet**](TemporalRuleAPI.md#V1AccountAccountIDTemporalruleTemporalRuleIDGet) | **Get** /v1/account/{accountID}/temporalrule/{temporalRuleID} | Get Temporal Rule Details |
| [**V1AccountAccountIDTemporalruleTemporalRuleIDPut**](TemporalRuleAPI.md#V1AccountAccountIDTemporalruleTemporalRuleIDPut) | **Put** /v1/account/{accountID}/temporalrule/{temporalRuleID} | Update Temporal Rule |



## V1AccountAccountIDTemporalruleGet

> ServiceDocsTemporalRuleGetAll V1AccountAccountIDTemporalruleGet(accountID, startKey, pageSize)

Get Temporal Rule List

Access all temporal rules for an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleAPI apiInstance = new TemporalRuleAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsTemporalRuleGetAll result = apiInstance.V1AccountAccountIDTemporalruleGet(accountID, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleAPI#V1AccountAccountIDTemporalruleGet");
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
| **accountID** | **string**| Account ID, 32 alpha numeric | |
| **startKey** | **string**| start_key for pagination that was returned as next_start_key from your previous call | [optional] |
| **pageSize** | **int32**| number of records to return, range 1 to 50 | [optional] |

### Return type

[**ServiceDocsTemporalRuleGetAll**](ServiceDocsTemporalRuleGetAll.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |


## V1AccountAccountIDTemporalrulePost

> ServiceDocsTemporalRuleGetSingle V1AccountAccountIDTemporalrulePost(accountID, temporalrule)

Create Temporal Rule

Create temporal rules for an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleAPI apiInstance = new TemporalRuleAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alphanumeric
        ServiceVOIPTemporalRuleAddEdit2 temporalrule = ; // ServiceVOIPTemporalRuleAddEdit2 | payload fields
        try {
            ServiceDocsTemporalRuleGetSingle result = apiInstance.V1AccountAccountIDTemporalrulePost(accountID, temporalrule);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleAPI#V1AccountAccountIDTemporalrulePost");
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
| **accountID** | **string**| Account ID, 32 alphanumeric | |
| **temporalrule** | [**ServiceVOIPTemporalRuleAddEdit2**](ServiceVOIPTemporalRuleAddEdit2.md)| payload fields | |

### Return type

[**ServiceDocsTemporalRuleGetSingle**](ServiceDocsTemporalRuleGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |


## V1AccountAccountIDTemporalruleTemporalRuleIDDelete

> ServiceDocsTemporalRuleGetSingle V1AccountAccountIDTemporalruleTemporalRuleIDDelete(accountID, temporalRuleID)

Delete Temporal Rule

Remove a temporal rule from an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleAPI apiInstance = new TemporalRuleAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string temporalRuleID = temporalRuleID_example; // string | temporal rule ID, 32 alpha numeric
        try {
            ServiceDocsTemporalRuleGetSingle result = apiInstance.V1AccountAccountIDTemporalruleTemporalRuleIDDelete(accountID, temporalRuleID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleAPI#V1AccountAccountIDTemporalruleTemporalRuleIDDelete");
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
| **accountID** | **string**| Account ID, 32 alpha numeric | |
| **temporalRuleID** | **string**| temporal rule ID, 32 alpha numeric | |

### Return type

[**ServiceDocsTemporalRuleGetSingle**](ServiceDocsTemporalRuleGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |


## V1AccountAccountIDTemporalruleTemporalRuleIDGet

> ServiceDocsTemporalRuleGetSingle V1AccountAccountIDTemporalruleTemporalRuleIDGet(accountID, temporalRuleID)

Get Temporal Rule Details

View details about individual time rules.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleAPI apiInstance = new TemporalRuleAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string temporalRuleID = temporalRuleID_example; // string | Temporal Rule ID, 32 alpha numeric
        try {
            ServiceDocsTemporalRuleGetSingle result = apiInstance.V1AccountAccountIDTemporalruleTemporalRuleIDGet(accountID, temporalRuleID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleAPI#V1AccountAccountIDTemporalruleTemporalRuleIDGet");
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
| **accountID** | **string**| Account ID, 32 alpha numeric | |
| **temporalRuleID** | **string**| Temporal Rule ID, 32 alpha numeric | |

### Return type

[**ServiceDocsTemporalRuleGetSingle**](ServiceDocsTemporalRuleGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |


## V1AccountAccountIDTemporalruleTemporalRuleIDPut

> ServiceDocsTemporalRuleGetSingle V1AccountAccountIDTemporalruleTemporalRuleIDPut(accountID, temporalRuleID, reqBody)

Update Temporal Rule

Edit the existing temporal rules in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleAPI apiInstance = new TemporalRuleAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string temporalRuleID = temporalRuleID_example; // string | Temporal Rule ID, 32 alpha numeric
        ServiceVOIPTemporalRuleAddEdit2 reqBody = ; // ServiceVOIPTemporalRuleAddEdit2 | payload fields
        try {
            ServiceDocsTemporalRuleGetSingle result = apiInstance.V1AccountAccountIDTemporalruleTemporalRuleIDPut(accountID, temporalRuleID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleAPI#V1AccountAccountIDTemporalruleTemporalRuleIDPut");
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
| **accountID** | **string**| Account ID, 32 alpha numeric | |
| **temporalRuleID** | **string**| Temporal Rule ID, 32 alpha numeric | |
| **reqBody** | [**ServiceVOIPTemporalRuleAddEdit2**](ServiceVOIPTemporalRuleAddEdit2.md)| payload fields | |

### Return type

[**ServiceDocsTemporalRuleGetSingle**](ServiceDocsTemporalRuleGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |

