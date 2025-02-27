# DeviceAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountidDeviceDeviceidDelete**](DeviceAPI.md#V1AccountAccountidDeviceDeviceidDelete) | **Delete** /v1/account/{accountid}/device/{deviceid} | Delete Device |
| [**V1AccountAccountidDeviceDeviceidGet**](DeviceAPI.md#V1AccountAccountidDeviceDeviceidGet) | **Get** /v1/account/{accountid}/device/{deviceid} | Get Device Details |
| [**V1AccountAccountidDeviceDeviceidPut**](DeviceAPI.md#V1AccountAccountidDeviceDeviceidPut) | **Put** /v1/account/{accountid}/device/{deviceid} | Update Device |
| [**V1AccountAccountidDeviceGet**](DeviceAPI.md#V1AccountAccountidDeviceGet) | **Get** /v1/account/{accountid}/device | Get Device List |
| [**V1AccountAccountidDevicePost**](DeviceAPI.md#V1AccountAccountidDevicePost) | **Post** /v1/account/{accountid}/device | Create Device |



## V1AccountAccountidDeviceDeviceidDelete

> ServiceDocsDeviceGetSingle V1AccountAccountidDeviceDeviceidDelete(accountid, deviceid)

Delete Device

Remove one device from a CPaaS account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.DeviceAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        DeviceAPI apiInstance = new DeviceAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string deviceid = deviceid_example; // string | Device ID, 32 alpha numeric
        try {
            ServiceDocsDeviceGetSingle result = apiInstance.V1AccountAccountidDeviceDeviceidDelete(accountid, deviceid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DeviceAPI#V1AccountAccountidDeviceDeviceidDelete");
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
| **deviceid** | **string**| Device ID, 32 alpha numeric | |

### Return type

[**ServiceDocsDeviceGetSingle**](ServiceDocsDeviceGetSingle.md)

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


## V1AccountAccountidDeviceDeviceidGet

> ServiceDocsDeviceGetSingle V1AccountAccountidDeviceDeviceidGet(accountid, deviceid)

Get Device Details

Permit a user to view specific device details.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.DeviceAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        DeviceAPI apiInstance = new DeviceAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string deviceid = deviceid_example; // string | Device ID, 32 alpha numeric
        try {
            ServiceDocsDeviceGetSingle result = apiInstance.V1AccountAccountidDeviceDeviceidGet(accountid, deviceid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DeviceAPI#V1AccountAccountidDeviceDeviceidGet");
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
| **deviceid** | **string**| Device ID, 32 alpha numeric | |

### Return type

[**ServiceDocsDeviceGetSingle**](ServiceDocsDeviceGetSingle.md)

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


## V1AccountAccountidDeviceDeviceidPut

> ServiceDocsDeviceGetSingle V1AccountAccountidDeviceDeviceidPut(accountid, deviceid, device)

Update Device

Edit specifics about the device, such as the device type, name, and owner.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.DeviceAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        DeviceAPI apiInstance = new DeviceAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string deviceid = deviceid_example; // string | Device ID, 32 alpha numeric
        ServiceVOIPDeviceAddEdit2 device = ; // ServiceVOIPDeviceAddEdit2 | device fields
        try {
            ServiceDocsDeviceGetSingle result = apiInstance.V1AccountAccountidDeviceDeviceidPut(accountid, deviceid, device);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DeviceAPI#V1AccountAccountidDeviceDeviceidPut");
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
| **deviceid** | **string**| Device ID, 32 alpha numeric | |
| **device** | [**ServiceVOIPDeviceAddEdit2**](ServiceVOIPDeviceAddEdit2.md)| device fields | |

### Return type

[**ServiceDocsDeviceGetSingle**](ServiceDocsDeviceGetSingle.md)

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


## V1AccountAccountidDeviceGet

> ServiceDocsDeviceGetAll V1AccountAccountidDeviceGet(accountid, startKey, pageSize)

Get Device List

Obtain a list of all devices associated with an account such as fax machines, cell phones, and soft phones.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.DeviceAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        DeviceAPI apiInstance = new DeviceAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsDeviceGetAll result = apiInstance.V1AccountAccountidDeviceGet(accountid, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DeviceAPI#V1AccountAccountidDeviceGet");
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

[**ServiceDocsDeviceGetAll**](ServiceDocsDeviceGetAll.md)

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


## V1AccountAccountidDevicePost

> ServiceDocsDeviceGetSingle V1AccountAccountidDevicePost(accountid, device)

Create Device

Connect a new device to an account to enhance communication methods.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.DeviceAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        DeviceAPI apiInstance = new DeviceAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        ServiceVOIPDeviceAddEdit2 device = ; // ServiceVOIPDeviceAddEdit2 | device fields
        try {
            ServiceDocsDeviceGetSingle result = apiInstance.V1AccountAccountidDevicePost(accountid, device);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DeviceAPI#V1AccountAccountidDevicePost");
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
| **device** | [**ServiceVOIPDeviceAddEdit2**](ServiceVOIPDeviceAddEdit2.md)| device fields | |

### Return type

[**ServiceDocsDeviceGetSingle**](ServiceDocsDeviceGetSingle.md)

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

