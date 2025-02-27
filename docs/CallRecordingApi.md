# CallRecordingAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDRecordingGet**](CallRecordingAPI.md#V1AccountAccountIDRecordingGet) | **Get** /v1/account/{accountID}/recording | Get Account Call Recording |
| [**V1AccountAccountIDRecordingRecordingIDDelete**](CallRecordingAPI.md#V1AccountAccountIDRecordingRecordingIDDelete) | **Delete** /v1/account/{accountID}/recording/{recordingID} | Delete Call Recording |
| [**V1AccountAccountIDRecordingRecordingIDGet**](CallRecordingAPI.md#V1AccountAccountIDRecordingRecordingIDGet) | **Get** /v1/account/{accountID}/recording/{recordingID} | Get Call Recording Details |
| [**V1AccountAccountIDUserUserIDRecordingGet**](CallRecordingAPI.md#V1AccountAccountIDUserUserIDRecordingGet) | **Get** /v1/account/{accountID}/user/{userID}/recording | Get User Call Recording |



## V1AccountAccountIDRecordingGet

> ServiceDocsCallRecordingGetAll V1AccountAccountIDRecordingGet(accountID)

Get Account Call Recording

Obtain a list of the call recordings within an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallRecordingAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallRecordingAPI apiInstance = new CallRecordingAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocsCallRecordingGetAll result = apiInstance.V1AccountAccountIDRecordingGet(accountID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallRecordingAPI#V1AccountAccountIDRecordingGet");
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

### Return type

[**ServiceDocsCallRecordingGetAll**](ServiceDocsCallRecordingGetAll.md)

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


## V1AccountAccountIDRecordingRecordingIDDelete

> ServiceDocsCallRecordingGetSingle V1AccountAccountIDRecordingRecordingIDDelete(accountID, recordingID)

Delete Call Recording

Delete a single call recording from an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallRecordingAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallRecordingAPI apiInstance = new CallRecordingAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string recordingID = recordingID_example; // string | Recording ID, 39 (yyyymm-<32 id>)
        try {
            ServiceDocsCallRecordingGetSingle result = apiInstance.V1AccountAccountIDRecordingRecordingIDDelete(accountID, recordingID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallRecordingAPI#V1AccountAccountIDRecordingRecordingIDDelete");
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
| **recordingID** | **string**| Recording ID, 39 (yyyymm-&lt;32 id&gt;) | |

### Return type

[**ServiceDocsCallRecordingGetSingle**](ServiceDocsCallRecordingGetSingle.md)

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


## V1AccountAccountIDRecordingRecordingIDGet

> ServiceDocsCallRecordingGetSingle V1AccountAccountIDRecordingRecordingIDGet(accountID, recordingID)

Get Call Recording Details

Access details for each recorded call in an account (e.g., duration, names and numbers of call participants, etc.).

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallRecordingAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallRecordingAPI apiInstance = new CallRecordingAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string recordingID = recordingID_example; // string | Recording ID, 39 (yyyymm-<32 id>)
        try {
            ServiceDocsCallRecordingGetSingle result = apiInstance.V1AccountAccountIDRecordingRecordingIDGet(accountID, recordingID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallRecordingAPI#V1AccountAccountIDRecordingRecordingIDGet");
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
| **recordingID** | **string**| Recording ID, 39 (yyyymm-&lt;32 id&gt;) | |

### Return type

[**ServiceDocsCallRecordingGetSingle**](ServiceDocsCallRecordingGetSingle.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, audio/mp3, audio/mpeg, audio/mpeg3, audio/x-wav, audio/wav


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |


## V1AccountAccountIDUserUserIDRecordingGet

> ServiceDocsCallRecordingGetAll V1AccountAccountIDUserUserIDRecordingGet(accountID, userID)

Get User Call Recording

Retrieve a list of call recordings for a user within an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallRecordingAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallRecordingAPI apiInstance = new CallRecordingAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string userID = userID_example; // string | User ID, 32 alpha numeric
        try {
            ServiceDocsCallRecordingGetAll result = apiInstance.V1AccountAccountIDUserUserIDRecordingGet(accountID, userID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallRecordingAPI#V1AccountAccountIDUserUserIDRecordingGet");
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
| **userID** | **string**| User ID, 32 alpha numeric | |

### Return type

[**ServiceDocsCallRecordingGetAll**](ServiceDocsCallRecordingGetAll.md)

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

