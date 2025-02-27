# CallQueueManagementAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDCallqueueGet**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueGet) | **Get** /v1/account/{accountID}/callqueue | Get Call Queues |
| [**V1AccountAccountIDCallqueuePost**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueuePost) | **Post** /v1/account/{accountID}/callqueue | Create Call Queue |
| [**V1AccountAccountIDCallqueueQueueIDDelete**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueQueueIDDelete) | **Delete** /v1/account/{accountID}/callqueue/{queueID} | Delete Call Queue |
| [**V1AccountAccountIDCallqueueQueueIDGet**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueQueueIDGet) | **Get** /v1/account/{accountID}/callqueue/{queueID} | Get Call Queue Details |
| [**V1AccountAccountIDCallqueueQueueIDPut**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueQueueIDPut) | **Put** /v1/account/{accountID}/callqueue/{queueID} | Update Call Queue |
| [**V1AccountAccountIDCallqueueQueueIDStatusGet**](CallQueueManagementAPI.md#V1AccountAccountIDCallqueueQueueIDStatusGet) | **Get** /v1/account/{accountID}/callqueue/{queueID}/status | Get Call Queue Status |
| [**V1AccountAccountIDQueuerolesGet**](CallQueueManagementAPI.md#V1AccountAccountIDQueuerolesGet) | **Get** /v1/account/{accountID}/queueroles | Get Queue Roles of Account |
| [**V1AccountAccountIDQueuerolesQueueIDPost**](CallQueueManagementAPI.md#V1AccountAccountIDQueuerolesQueueIDPost) | **Post** /v1/account/{accountID}/queueroles/{queueID} | Assign Queue Role to Call Queue |



## V1AccountAccountIDCallqueueGet

> ServiceDocsCallQueueGetAll V1AccountAccountIDCallqueueGet(accountID)

Get Call Queues

Retrieve call queue details for an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallQueueManagementAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallQueueManagementAPI apiInstance = new CallQueueManagementAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocsCallQueueGetAll result = apiInstance.V1AccountAccountIDCallqueueGet(accountID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallQueueManagementAPI#V1AccountAccountIDCallqueueGet");
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

[**ServiceDocsCallQueueGetAll**](ServiceDocsCallQueueGetAll.md)

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


## V1AccountAccountIDCallqueuePost

> ServiceDocsCallQueueGetSingle V1AccountAccountIDCallqueuePost(accountID, reqBody)

Create Call Queue

Set up a call queue in an account for specific inbound calls.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallQueueManagementAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallQueueManagementAPI apiInstance = new CallQueueManagementAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        ServiceVOIPCallQueueAddEditData reqBody = ; // ServiceVOIPCallQueueAddEditData | payload fields
        try {
            ServiceDocsCallQueueGetSingle result = apiInstance.V1AccountAccountIDCallqueuePost(accountID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallQueueManagementAPI#V1AccountAccountIDCallqueuePost");
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
| **reqBody** | [**ServiceVOIPCallQueueAddEditData**](ServiceVOIPCallQueueAddEditData.md)| payload fields | |

### Return type

[**ServiceDocsCallQueueGetSingle**](ServiceDocsCallQueueGetSingle.md)

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


## V1AccountAccountIDCallqueueQueueIDDelete

> ServiceDocsCallQueueGetSingle V1AccountAccountIDCallqueueQueueIDDelete(accountID, queueID)

Delete Call Queue

Remove the call queue from an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallQueueManagementAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallQueueManagementAPI apiInstance = new CallQueueManagementAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string queueID = queueID_example; // string | Queue ID, 32 alpha numeric
        try {
            ServiceDocsCallQueueGetSingle result = apiInstance.V1AccountAccountIDCallqueueQueueIDDelete(accountID, queueID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallQueueManagementAPI#V1AccountAccountIDCallqueueQueueIDDelete");
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
| **queueID** | **string**| Queue ID, 32 alpha numeric | |

### Return type

[**ServiceDocsCallQueueGetSingle**](ServiceDocsCallQueueGetSingle.md)

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


## V1AccountAccountIDCallqueueQueueIDGet

> ServiceDocsCallQueueGetSingle V1AccountAccountIDCallqueueQueueIDGet(accountID, queueID)

Get Call Queue Details

Capture metadata about a specific queue, such as queue_type and agent_wrapup_time.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallQueueManagementAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallQueueManagementAPI apiInstance = new CallQueueManagementAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string queueID = queueID_example; // string | Queue ID, 32 alpha numeric
        try {
            ServiceDocsCallQueueGetSingle result = apiInstance.V1AccountAccountIDCallqueueQueueIDGet(accountID, queueID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallQueueManagementAPI#V1AccountAccountIDCallqueueQueueIDGet");
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
| **queueID** | **string**| Queue ID, 32 alpha numeric | |

### Return type

[**ServiceDocsCallQueueGetSingle**](ServiceDocsCallQueueGetSingle.md)

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


## V1AccountAccountIDCallqueueQueueIDPut

> ServiceDocsCallQueueGetSingle V1AccountAccountIDCallqueueQueueIDPut(accountID, queueID, reqBody)

Update Call Queue

Update the metadata mentioned above.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallQueueManagementAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallQueueManagementAPI apiInstance = new CallQueueManagementAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string queueID = queueID_example; // string | Queue ID, 32 alpha numeric
        ServiceVOIPCallQueueAddEditData reqBody = ; // ServiceVOIPCallQueueAddEditData | payload fields
        try {
            ServiceDocsCallQueueGetSingle result = apiInstance.V1AccountAccountIDCallqueueQueueIDPut(accountID, queueID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallQueueManagementAPI#V1AccountAccountIDCallqueueQueueIDPut");
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
| **queueID** | **string**| Queue ID, 32 alpha numeric | |
| **reqBody** | [**ServiceVOIPCallQueueAddEditData**](ServiceVOIPCallQueueAddEditData.md)| payload fields | |

### Return type

[**ServiceDocsCallQueueGetSingle**](ServiceDocsCallQueueGetSingle.md)

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


## V1AccountAccountIDCallqueueQueueIDStatusGet

> ServiceDocsCallQueueGetSingleStatus V1AccountAccountIDCallqueueQueueIDStatusGet(accountID, queueID)

Get Call Queue Status

Access the status of a call queue in an account, such as the number of available agents (recipients), estimated wait time, and number of active sessions.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallQueueManagementAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallQueueManagementAPI apiInstance = new CallQueueManagementAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string queueID = queueID_example; // string | Queue ID, 32 alpha numeric
        try {
            ServiceDocsCallQueueGetSingleStatus result = apiInstance.V1AccountAccountIDCallqueueQueueIDStatusGet(accountID, queueID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallQueueManagementAPI#V1AccountAccountIDCallqueueQueueIDStatusGet");
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
| **queueID** | **string**| Queue ID, 32 alpha numeric | |

### Return type

[**ServiceDocsCallQueueGetSingleStatus**](ServiceDocsCallQueueGetSingleStatus.md)

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


## V1AccountAccountIDQueuerolesGet

> ServiceDocsCallQueueGetRoles V1AccountAccountIDQueuerolesGet(accountID)

Get Queue Roles of Account

Obtain data about each queue role in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallQueueManagementAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallQueueManagementAPI apiInstance = new CallQueueManagementAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocsCallQueueGetRoles result = apiInstance.V1AccountAccountIDQueuerolesGet(accountID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallQueueManagementAPI#V1AccountAccountIDQueuerolesGet");
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

[**ServiceDocsCallQueueGetRoles**](ServiceDocsCallQueueGetRoles.md)

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


## V1AccountAccountIDQueuerolesQueueIDPost

> ServiceAPIResponse V1AccountAccountIDQueuerolesQueueIDPost(accountID, queueID, reqBody)

Assign Queue Role to Call Queue

Assign roles to members in a call queue.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.CallQueueManagementAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        CallQueueManagementAPI apiInstance = new CallQueueManagementAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string queueID = queueID_example; // string | Queue ID, 32 alpha numeric
        ServiceVOIPCallQueueRoleAssignData reqBody = ; // ServiceVOIPCallQueueRoleAssignData | payload fields
        try {
            ServiceAPIResponse result = apiInstance.V1AccountAccountIDQueuerolesQueueIDPost(accountID, queueID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CallQueueManagementAPI#V1AccountAccountIDQueuerolesQueueIDPost");
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
| **queueID** | **string**| Queue ID, 32 alpha numeric | |
| **reqBody** | [**ServiceVOIPCallQueueRoleAssignData**](ServiceVOIPCallQueueRoleAssignData.md)| payload fields | |

### Return type

[**ServiceAPIResponse**](ServiceAPIResponse.md)

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

