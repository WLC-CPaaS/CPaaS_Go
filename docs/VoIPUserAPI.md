# VoIPUserAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountidUserGet**](VoIPUserAPI.md#V1AccountAccountidUserGet) | **Get** /v1/account/{accountid}/user | Get User List |
| [**V1AccountAccountidUserPost**](VoIPUserAPI.md#V1AccountAccountidUserPost) | **Post** /v1/account/{accountid}/user | Create User |
| [**V1AccountAccountidUserUseridDelete**](VoIPUserAPI.md#V1AccountAccountidUserUseridDelete) | **Delete** /v1/account/{accountid}/user/{userid} | Delete User |
| [**V1AccountAccountidUserUseridGet**](VoIPUserAPI.md#V1AccountAccountidUserUseridGet) | **Get** /v1/account/{accountid}/user/{userid} | Get User Details |
| [**V1AccountAccountidUserUseridPut**](VoIPUserAPI.md#V1AccountAccountidUserUseridPut) | **Put** /v1/account/{accountid}/user/{userid} | Update User |



## V1AccountAccountidUserGet

> ServiceDocsUserGetAll V1AccountAccountidUserGet(accountid, startKey, pageSize)

Get User List

Get a list of all VoIP users that includes first and last names, email addresses, extensions, and account statuses.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoIPUserAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoIPUserAPI apiInstance = new VoIPUserAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsUserGetAll result = apiInstance.V1AccountAccountidUserGet(accountid, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoIPUserAPI#V1AccountAccountidUserGet");
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
| **accountid** | **string**| Account ID, 32 alpha numeric | |
| **startKey** | **string**| start_key for pagination that was returned as next_start_key from your previous call | [optional] |
| **pageSize** | **int32**| number of records to return, range 1 to 50 | [optional] |

### Return type

[**ServiceDocsUserGetAll**](ServiceDocsUserGetAll.md)

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


## V1AccountAccountidUserPost

> ServiceDocsUserGetSingle V1AccountAccountidUserPost(accountid, user)

Create User

Add new users to the account. When a user is added, the system generates their unique 32 alpha numeric ID.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoIPUserAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoIPUserAPI apiInstance = new VoIPUserAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        ServiceVOIPUserAdd2 user = ; // ServiceVOIPUserAdd2 | user fields
        try {
            ServiceDocsUserGetSingle result = apiInstance.V1AccountAccountidUserPost(accountid, user);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoIPUserAPI#V1AccountAccountidUserPost");
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
| **accountid** | **string**| Account ID, 32 alpha numeric | |
| **user** | [**ServiceVOIPUserAdd2**](ServiceVOIPUserAdd2.md)| user fields | |

### Return type

[**ServiceDocsUserGetSingle**](ServiceDocsUserGetSingle.md)

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


## V1AccountAccountidUserUseridDelete

> ServiceDocsUserGetSingle V1AccountAccountidUserUseridDelete(accountid, userid)

Delete User

Delete VoIP user access to maintain the security of your accounts.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoIPUserAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoIPUserAPI apiInstance = new VoIPUserAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string userid = userid_example; // string | User ID, 32 alpha numeric
        try {
            ServiceDocsUserGetSingle result = apiInstance.V1AccountAccountidUserUseridDelete(accountid, userid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoIPUserAPI#V1AccountAccountidUserUseridDelete");
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
| **accountid** | **string**| Account ID, 32 alpha numeric | |
| **userid** | **string**| User ID, 32 alpha numeric | |

### Return type

[**ServiceDocsUserGetSingle**](ServiceDocsUserGetSingle.md)

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


## V1AccountAccountidUserUseridGet

> ServiceDocsUserGetSingle V1AccountAccountidUserUseridGet(accountid, userid)

Get User Details

View specific user details.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoIPUserAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoIPUserAPI apiInstance = new VoIPUserAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string userid = userid_example; // string | User ID, 32 alpha numeric
        try {
            ServiceDocsUserGetSingle result = apiInstance.V1AccountAccountidUserUseridGet(accountid, userid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoIPUserAPI#V1AccountAccountidUserUseridGet");
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
| **accountid** | **string**| Account ID, 32 alpha numeric | |
| **userid** | **string**| User ID, 32 alpha numeric | |

### Return type

[**ServiceDocsUserGetSingle**](ServiceDocsUserGetSingle.md)

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


## V1AccountAccountidUserUseridPut

> ServiceDocsUserGetSingle V1AccountAccountidUserUseridPut(accountid, userid, user)

Update User

Keep user information current. Modify the first and last name, extension, and other pertinent information.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoIPUserAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoIPUserAPI apiInstance = new VoIPUserAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string userid = userid_example; // string | User ID, 32 alpha numeric
        ServiceVOIPUserAdd2 user = ; // ServiceVOIPUserAdd2 | user fields
        try {
            ServiceDocsUserGetSingle result = apiInstance.V1AccountAccountidUserUseridPut(accountid, userid, user);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoIPUserAPI#V1AccountAccountidUserUseridPut");
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
| **accountid** | **string**| Account ID, 32 alpha numeric | |
| **userid** | **string**| User ID, 32 alpha numeric | |
| **user** | [**ServiceVOIPUserAdd2**](ServiceVOIPUserAdd2.md)| user fields | |

### Return type

[**ServiceDocsUserGetSingle**](ServiceDocsUserGetSingle.md)

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

