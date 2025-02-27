# TemporalRuleSetAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDTemporalrulesetGet**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetGet) | **Get** /v1/account/{accountID}/temporalruleset | Get Temporal Rule Set List |
| [**V1AccountAccountIDTemporalrulesetPost**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetPost) | **Post** /v1/account/{accountID}/temporalruleset | Create Temporal Rule Set |
| [**V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete) | **Delete** /v1/account/{accountID}/temporalruleset/{temporalRuleSetID} | Delete Temporal Rule Set |
| [**V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet) | **Get** /v1/account/{accountID}/temporalruleset/{temporalRuleSetID} | Get Temporal Rule Set Details |
| [**V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut**](TemporalRuleSetAPI.md#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut) | **Put** /v1/account/{accountID}/temporalruleset/{temporalRuleSetID} | Update Temporal Rule Set |



## V1AccountAccountIDTemporalrulesetGet

> ServiceDocsTemporalRuleSetGetAll V1AccountAccountIDTemporalrulesetGet(accountID, startKey, pageSize)

Get Temporal Rule Set List

Access the temporal rule set list in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleSetAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleSetAPI apiInstance = new TemporalRuleSetAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsTemporalRuleSetGetAll result = apiInstance.V1AccountAccountIDTemporalrulesetGet(accountID, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleSetAPI#V1AccountAccountIDTemporalrulesetGet");
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

[**ServiceDocsTemporalRuleSetGetAll**](ServiceDocsTemporalRuleSetGetAll.md)

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


## V1AccountAccountIDTemporalrulesetPost

> ServiceDocsTemporalRuleSetGetSingle V1AccountAccountIDTemporalrulesetPost(accountID, temporalruleset)

Create Temporal Rule Set

Develop a new temporal rule set for an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleSetAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleSetAPI apiInstance = new TemporalRuleSetAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alphanumeric
        ServiceVOIPTemporalRuleSetAddEditData temporalruleset = ; // ServiceVOIPTemporalRuleSetAddEditData | payload fields
        try {
            ServiceDocsTemporalRuleSetGetSingle result = apiInstance.V1AccountAccountIDTemporalrulesetPost(accountID, temporalruleset);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleSetAPI#V1AccountAccountIDTemporalrulesetPost");
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
| **temporalruleset** | [**ServiceVOIPTemporalRuleSetAddEditData**](ServiceVOIPTemporalRuleSetAddEditData.md)| payload fields | |

### Return type

[**ServiceDocsTemporalRuleSetGetSingle**](ServiceDocsTemporalRuleSetGetSingle.md)

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


## V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete

> ServiceDocsTemporalRuleSetGetSingle V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete(accountID, temporalRuleSetID)

Delete Temporal Rule Set

Delete the temporal rule set from an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleSetAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleSetAPI apiInstance = new TemporalRuleSetAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string temporalRuleSetID = temporalRuleSetID_example; // string | temporal rule set ID, 32 alpha numeric
        try {
            ServiceDocsTemporalRuleSetGetSingle result = apiInstance.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete(accountID, temporalRuleSetID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleSetAPI#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDDelete");
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
| **temporalRuleSetID** | **string**| temporal rule set ID, 32 alpha numeric | |

### Return type

[**ServiceDocsTemporalRuleSetGetSingle**](ServiceDocsTemporalRuleSetGetSingle.md)

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


## V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet

> ServiceDocsTemporalRuleSetGetSingle V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet(accountID, temporalRuleSetID)

Get Temporal Rule Set Details

Acquire details about a temporal rule set in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleSetAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleSetAPI apiInstance = new TemporalRuleSetAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string temporalRuleSetID = temporalRuleSetID_example; // string | Temporal Ruleset ID, 32 alpha numeric
        try {
            ServiceDocsTemporalRuleSetGetSingle result = apiInstance.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet(accountID, temporalRuleSetID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleSetAPI#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDGet");
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
| **temporalRuleSetID** | **string**| Temporal Ruleset ID, 32 alpha numeric | |

### Return type

[**ServiceDocsTemporalRuleSetGetSingle**](ServiceDocsTemporalRuleSetGetSingle.md)

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


## V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut

> ServiceDocsTemporalRuleSetGetSingle V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut(accountID, temporalRuleSetID, reqBody)

Update Temporal Rule Set

Efficiently adjust the temporal rule set in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.TemporalRuleSetAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        TemporalRuleSetAPI apiInstance = new TemporalRuleSetAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string temporalRuleSetID = temporalRuleSetID_example; // string | Temporal Ruleset ID, 32 alpha numeric
        ServiceVOIPTemporalRuleSetAddEditData reqBody = ; // ServiceVOIPTemporalRuleSetAddEditData | payload fields
        try {
            ServiceDocsTemporalRuleSetGetSingle result = apiInstance.V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut(accountID, temporalRuleSetID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TemporalRuleSetAPI#V1AccountAccountIDTemporalrulesetTemporalRuleSetIDPut");
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
| **temporalRuleSetID** | **string**| Temporal Ruleset ID, 32 alpha numeric | |
| **reqBody** | [**ServiceVOIPTemporalRuleSetAddEditData**](ServiceVOIPTemporalRuleSetAddEditData.md)| payload fields | |

### Return type

[**ServiceDocsTemporalRuleSetGetSingle**](ServiceDocsTemporalRuleSetGetSingle.md)

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

