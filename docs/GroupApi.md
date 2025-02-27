# GroupAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDGroupGet**](GroupAPI.md#V1AccountAccountIDGroupGet) | **Get** /v1/account/{accountID}/group | Get Group List |
| [**V1AccountAccountIDGroupGroupIDDelete**](GroupAPI.md#V1AccountAccountIDGroupGroupIDDelete) | **Delete** /v1/account/{accountID}/group/{groupID} | Delete Group |
| [**V1AccountAccountIDGroupGroupIDGet**](GroupAPI.md#V1AccountAccountIDGroupGroupIDGet) | **Get** /v1/account/{accountID}/group/{groupID} | Get Group Details |
| [**V1AccountAccountIDGroupGroupIDPut**](GroupAPI.md#V1AccountAccountIDGroupGroupIDPut) | **Put** /v1/account/{accountID}/group/{groupID} | Update Group |
| [**V1AccountAccountIDGroupPost**](GroupAPI.md#V1AccountAccountIDGroupPost) | **Post** /v1/account/{accountID}/group | Create Group |



## V1AccountAccountIDGroupGet

> ServiceDocGroupGetAll V1AccountAccountIDGroupGet(accountID, startKey, pageSize)

Get Group List

Get a list of groups associated with an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.GroupAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        GroupAPI apiInstance = new GroupAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocGroupGetAll result = apiInstance.V1AccountAccountIDGroupGet(accountID, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupAPI#V1AccountAccountIDGroupGet");
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

[**ServiceDocGroupGetAll**](ServiceDocGroupGetAll.md)

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


## V1AccountAccountIDGroupGroupIDDelete

> ServiceDocGroupGetSingle V1AccountAccountIDGroupGroupIDDelete(accountID, groupID)

Delete Group

Delete a call group in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.GroupAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        GroupAPI apiInstance = new GroupAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string groupID = groupID_example; // string | group ID, 32 alpha numeric
        try {
            ServiceDocGroupGetSingle result = apiInstance.V1AccountAccountIDGroupGroupIDDelete(accountID, groupID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupAPI#V1AccountAccountIDGroupGroupIDDelete");
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
| **groupID** | **string**| group ID, 32 alpha numeric | |

### Return type

[**ServiceDocGroupGetSingle**](ServiceDocGroupGetSingle.md)

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


## V1AccountAccountIDGroupGroupIDGet

> ServiceDocGroupGetSingle V1AccountAccountIDGroupGroupIDGet(accountID, groupID)

Get Group Details

Access details about a single group within an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.GroupAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        GroupAPI apiInstance = new GroupAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string groupID = groupID_example; // string | Group ID, 32 alpha numeric
        try {
            ServiceDocGroupGetSingle result = apiInstance.V1AccountAccountIDGroupGroupIDGet(accountID, groupID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupAPI#V1AccountAccountIDGroupGroupIDGet");
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
| **groupID** | **string**| Group ID, 32 alpha numeric | |

### Return type

[**ServiceDocGroupGetSingle**](ServiceDocGroupGetSingle.md)

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


## V1AccountAccountIDGroupGroupIDPut

> ServiceDocGroupGetSingle V1AccountAccountIDGroupGroupIDPut(accountID, groupID, reqBody)

Update Group

Modify the name, settings and other information for a group within an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.GroupAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        GroupAPI apiInstance = new GroupAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string groupID = groupID_example; // string | Group ID, 32 alpha numeric
        ServiceVOIPGroupAddEdit2 reqBody = ; // ServiceVOIPGroupAddEdit2 | payload fields
        try {
            ServiceDocGroupGetSingle result = apiInstance.V1AccountAccountIDGroupGroupIDPut(accountID, groupID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupAPI#V1AccountAccountIDGroupGroupIDPut");
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
| **groupID** | **string**| Group ID, 32 alpha numeric | |
| **reqBody** | [**ServiceVOIPGroupAddEdit2**](ServiceVOIPGroupAddEdit2.md)| payload fields | |

### Return type

[**ServiceDocGroupGetSingle**](ServiceDocGroupGetSingle.md)

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


## V1AccountAccountIDGroupPost

> ServiceDocGroupGetSingle V1AccountAccountIDGroupPost(accountID, group)

Create Group

Provide an additional resource by adding a group list to an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.GroupAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        GroupAPI apiInstance = new GroupAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID
        ServiceVOIPGroupAddEdit2 group = ; // ServiceVOIPGroupAddEdit2 | group fields
        try {
            ServiceDocGroupGetSingle result = apiInstance.V1AccountAccountIDGroupPost(accountID, group);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupAPI#V1AccountAccountIDGroupPost");
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
| **group** | [**ServiceVOIPGroupAddEdit2**](ServiceVOIPGroupAddEdit2.md)| group fields | |

### Return type

[**ServiceDocGroupGetSingle**](ServiceDocGroupGetSingle.md)

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

