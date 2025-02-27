# CallflowAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDCallflowCallflowIDDelete**](CallflowAPI.md#V1AccountAccountIDCallflowCallflowIDDelete) | **Delete** /v1/account/{accountID}/callflow/{callflowID} | Delete Call Group |
| [**V1AccountAccountIDCallflowCallflowIDGet**](CallflowAPI.md#V1AccountAccountIDCallflowCallflowIDGet) | **Get** /v1/account/{accountID}/callflow/{callflowID} | Get Call Group Details |
| [**V1AccountAccountIDCallflowCallflowIDPut**](CallflowAPI.md#V1AccountAccountIDCallflowCallflowIDPut) | **Put** /v1/account/{accountID}/callflow/{callflowID} | Update Call Group |
| [**V1AccountAccountIDCallflowGet**](CallflowAPI.md#V1AccountAccountIDCallflowGet) | **Get** /v1/account/{accountID}/callflow | Get Callflow List |
| [**V1AccountAccountIDCallflowPost**](CallflowAPI.md#V1AccountAccountIDCallflowPost) | **Post** /v1/account/{accountID}/callflow | Create Call Group |



## V1AccountAccountIDCallflowCallflowIDDelete

> ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDDelete(accountID, callflowID)

Delete Call Group

Delete a callflow in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallflowAPI apiInstance = new CallflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string callflowID = callflowID_example; // string | callflow ID, 32 alpha numeric
        try {
            ServiceDocsCallflowGetSingle result = apiInstance.V1AccountAccountIDCallflowCallflowIDDelete(accountID, callflowID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallflowAPI#V1AccountAccountIDCallflowCallflowIDDelete");
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
| **callflowID** | **string**| callflow ID, 32 alpha numeric | |

### Return type

[**ServiceDocsCallflowGetSingle**](ServiceDocsCallflowGetSingle.md)

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


## V1AccountAccountIDCallflowCallflowIDGet

> ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDGet(accountID, callflowID)

Get Call Group Details

Get the details for a single callflow in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallflowAPI apiInstance = new CallflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string callflowID = callflowID_example; // string | Callflow ID, 32 alpha numeric
        try {
            ServiceDocsCallflowGetSingle result = apiInstance.V1AccountAccountIDCallflowCallflowIDGet(accountID, callflowID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallflowAPI#V1AccountAccountIDCallflowCallflowIDGet");
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
| **callflowID** | **string**| Callflow ID, 32 alpha numeric | |

### Return type

[**ServiceDocsCallflowGetSingle**](ServiceDocsCallflowGetSingle.md)

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


## V1AccountAccountIDCallflowCallflowIDPut

> ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDPut(accountID, callflowID, reqBody)

Update Call Group

Update the details for a single callflow in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallflowAPI apiInstance = new CallflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string callflowID = callflowID_example; // string | Callflow ID, 32 alpha numeric
        ServiceCallflowAddEditData reqBody = ; // ServiceCallflowAddEditData | payload fields
        try {
            ServiceDocsCallflowGetSingle result = apiInstance.V1AccountAccountIDCallflowCallflowIDPut(accountID, callflowID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallflowAPI#V1AccountAccountIDCallflowCallflowIDPut");
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
| **callflowID** | **string**| Callflow ID, 32 alpha numeric | |
| **reqBody** | [**ServiceCallflowAddEditData**](ServiceCallflowAddEditData.md)| payload fields | |

### Return type

[**ServiceDocsCallflowGetSingle**](ServiceDocsCallflowGetSingle.md)

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


## V1AccountAccountIDCallflowGet

> ServiceDocsCallflowGetAll V1AccountAccountIDCallflowGet(accountID, startKey, pageSize)

Get Callflow List

Permit a user to view the callflow details in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallflowAPI apiInstance = new CallflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsCallflowGetAll result = apiInstance.V1AccountAccountIDCallflowGet(accountID, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallflowAPI#V1AccountAccountIDCallflowGet");
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

[**ServiceDocsCallflowGetAll**](ServiceDocsCallflowGetAll.md)

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


## V1AccountAccountIDCallflowPost

> ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowPost(accountID, request)

Create Call Group

Create instructions for routing a call to a user or system.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallflowAPI apiInstance = new CallflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha-numeric
        ServiceCallflowAddEditData request = ; // ServiceCallflowAddEditData | Call flow configuration
        try {
            ServiceDocsCallflowGetSingle result = apiInstance.V1AccountAccountIDCallflowPost(accountID, request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallflowAPI#V1AccountAccountIDCallflowPost");
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
| **accountID** | **string**| Account ID, 32 alpha-numeric | |
| **request** | [**ServiceCallflowAddEditData**](ServiceCallflowAddEditData.md)| Call flow configuration | |

### Return type

[**ServiceDocsCallflowGetSingle**](ServiceDocsCallflowGetSingle.md)

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

