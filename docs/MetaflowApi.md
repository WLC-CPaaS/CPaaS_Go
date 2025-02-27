# MetaflowAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDDeviceDeviceIDMetaflowDelete**](MetaflowAPI.md#V1AccountAccountIDDeviceDeviceIDMetaflowDelete) | **Delete** /v1/account/{accountID}/device/{deviceID}/metaflow | Delete Device Metaflow |
| [**V1AccountAccountIDDeviceDeviceIDMetaflowGet**](MetaflowAPI.md#V1AccountAccountIDDeviceDeviceIDMetaflowGet) | **Get** /v1/account/{accountID}/device/{deviceID}/metaflow | Get Device Metaflow List |
| [**V1AccountAccountIDDeviceDeviceIDMetaflowPost**](MetaflowAPI.md#V1AccountAccountIDDeviceDeviceIDMetaflowPost) | **Post** /v1/account/{accountID}/device/{deviceID}/metaflow | Create Device Metaflow |
| [**V1AccountAccountIDMetaflowDelete**](MetaflowAPI.md#V1AccountAccountIDMetaflowDelete) | **Delete** /v1/account/{accountID}/metaflow | Delete Account Metaflow |
| [**V1AccountAccountIDMetaflowGet**](MetaflowAPI.md#V1AccountAccountIDMetaflowGet) | **Get** /v1/account/{accountID}/metaflow | Get Account Metaflow List |
| [**V1AccountAccountIDMetaflowPost**](MetaflowAPI.md#V1AccountAccountIDMetaflowPost) | **Post** /v1/account/{accountID}/metaflow | Create Account Metaflow |
| [**V1AccountAccountIDUserUserIDMetaflowDelete**](MetaflowAPI.md#V1AccountAccountIDUserUserIDMetaflowDelete) | **Delete** /v1/account/{accountID}/user/{userID}/metaflow | Delete User Metaflow |
| [**V1AccountAccountIDUserUserIDMetaflowGet**](MetaflowAPI.md#V1AccountAccountIDUserUserIDMetaflowGet) | **Get** /v1/account/{accountID}/user/{userID}/metaflow | Get User Metaflow List |
| [**V1AccountAccountIDUserUserIDMetaflowPost**](MetaflowAPI.md#V1AccountAccountIDUserUserIDMetaflowPost) | **Post** /v1/account/{accountID}/user/{userID}/metaflow | Create User Metaflow |



## V1AccountAccountIDDeviceDeviceIDMetaflowDelete

> ServiceDocMetaflowGet V1AccountAccountIDDeviceDeviceIDMetaflowDelete(accountID, deviceID)

Delete Device Metaflow

Delete all metaflows associated with a device.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string deviceID = deviceID_example; // string | Device ID, 32 alpha numeric
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDDeviceDeviceIDMetaflowDelete(accountID, deviceID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDDeviceDeviceIDMetaflowDelete");
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
| **deviceID** | **string**| Device ID, 32 alpha numeric | |

### Return type

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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


## V1AccountAccountIDDeviceDeviceIDMetaflowGet

> ServiceDocMetaflowGet V1AccountAccountIDDeviceDeviceIDMetaflowGet(accountID, deviceID)

Get Device Metaflow List

Get the list of metaflows for a device.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string deviceID = deviceID_example; // string | Device ID, 32 alpha numeric
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDDeviceDeviceIDMetaflowGet(accountID, deviceID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDDeviceDeviceIDMetaflowGet");
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
| **deviceID** | **string**| Device ID, 32 alpha numeric | |

### Return type

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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


## V1AccountAccountIDDeviceDeviceIDMetaflowPost

> ServiceDocMetaflowGet V1AccountAccountIDDeviceDeviceIDMetaflowPost(accountID, deviceID, reqBody)

Create Device Metaflow

Create a metaflow or multiple metaflows for a device.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string deviceID = deviceID_example; // string | Device ID, 32 alpha numeric
        ServiceVOIPMetaflowAddData reqBody = ; // ServiceVOIPMetaflowAddData | payload fields
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDDeviceDeviceIDMetaflowPost(accountID, deviceID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDDeviceDeviceIDMetaflowPost");
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
| **deviceID** | **string**| Device ID, 32 alpha numeric | |
| **reqBody** | [**ServiceVOIPMetaflowAddData**](ServiceVOIPMetaflowAddData.md)| payload fields | |

### Return type

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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


## V1AccountAccountIDMetaflowDelete

> ServiceDocMetaflowGet V1AccountAccountIDMetaflowDelete(accountID)

Delete Account Metaflow

Remove all metaflows from an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDMetaflowDelete(accountID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDMetaflowDelete");
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

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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


## V1AccountAccountIDMetaflowGet

> ServiceDocMetaflowGet V1AccountAccountIDMetaflowGet(accountID)

Get Account Metaflow List

Get an account&#39;s metaflow list.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDMetaflowGet(accountID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDMetaflowGet");
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

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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


## V1AccountAccountIDMetaflowPost

> ServiceDocMetaflowGet V1AccountAccountIDMetaflowPost(accountID, metaflow)

Create Account Metaflow

Generate a metaflow for an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID
        ServiceVOIPMetaflowAddData metaflow = ; // ServiceVOIPMetaflowAddData | Metaflow fields
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDMetaflowPost(accountID, metaflow);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDMetaflowPost");
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
| **metaflow** | [**ServiceVOIPMetaflowAddData**](ServiceVOIPMetaflowAddData.md)| Metaflow fields | |

### Return type

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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


## V1AccountAccountIDUserUserIDMetaflowDelete

> ServiceDocMetaflowGet V1AccountAccountIDUserUserIDMetaflowDelete(accountID, userID)

Delete User Metaflow

Delete all metaflows associated with a user.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string userID = userID_example; // string | user ID, 32 alpha numeric
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDUserUserIDMetaflowDelete(accountID, userID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDUserUserIDMetaflowDelete");
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
| **userID** | **string**| user ID, 32 alpha numeric | |

### Return type

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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


## V1AccountAccountIDUserUserIDMetaflowGet

> ServiceDocMetaflowGet V1AccountAccountIDUserUserIDMetaflowGet(accountID, userID)

Get User Metaflow List

Get the list of metaflows for a user.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string userID = userID_example; // string | user ID, 32 alpha numeric
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDUserUserIDMetaflowGet(accountID, userID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDUserUserIDMetaflowGet");
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
| **userID** | **string**| user ID, 32 alpha numeric | |

### Return type

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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


## V1AccountAccountIDUserUserIDMetaflowPost

> ServiceDocMetaflowGet V1AccountAccountIDUserUserIDMetaflowPost(accountID, userID, reqBody)

Create User Metaflow

Add a metaflow or multiple metaflows for a user in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MetaflowAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MetaflowAPI apiInstance = new MetaflowAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string userID = userID_example; // string | user ID, 32 alpha numeric
        ServiceVOIPMetaflowAddData reqBody = ; // ServiceVOIPMetaflowAddData | payload fields
        try {
            ServiceDocMetaflowGet result = apiInstance.V1AccountAccountIDUserUserIDMetaflowPost(accountID, userID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MetaflowAPI#V1AccountAccountIDUserUserIDMetaflowPost");
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
| **userID** | **string**| user ID, 32 alpha numeric | |
| **reqBody** | [**ServiceVOIPMetaflowAddData**](ServiceVOIPMetaflowAddData.md)| payload fields | |

### Return type

[**ServiceDocMetaflowGet**](ServiceDocMetaflowGet.md)

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

