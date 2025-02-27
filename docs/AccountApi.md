# AccountAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountidChildrenGet**](AccountAPI.md#V1AccountAccountidChildrenGet) | **Get** /v1/account/{accountid}/children | Get Sub Account List |
| [**V1AccountAccountidDelete**](AccountAPI.md#V1AccountAccountidDelete) | **Delete** /v1/account/{accountid} | Delete Account |
| [**V1AccountAccountidDnsrecordGet**](AccountAPI.md#V1AccountAccountidDnsrecordGet) | **Get** /v1/account/{accountid}/dnsrecord |  |
| [**V1AccountAccountidDnsrecordPost**](AccountAPI.md#V1AccountAccountidDnsrecordPost) | **Post** /v1/account/{accountid}/dnsrecord |  |
| [**V1AccountAccountidDnsrecordPut**](AccountAPI.md#V1AccountAccountidDnsrecordPut) | **Put** /v1/account/{accountid}/dnsrecord |  |
| [**V1AccountAccountidGet**](AccountAPI.md#V1AccountAccountidGet) | **Get** /v1/account/{accountid} | Get Account Details |
| [**V1AccountAccountidLimitGet**](AccountAPI.md#V1AccountAccountidLimitGet) | **Get** /v1/account/{accountid}/limit | Get Account Limits |
| [**V1AccountAccountidLimitPut**](AccountAPI.md#V1AccountAccountidLimitPut) | **Put** /v1/account/{accountid}/limit | Set Account Limits |
| [**V1AccountAccountidPost**](AccountAPI.md#V1AccountAccountidPost) | **Post** /v1/account/{accountid} | Create Sub Account |
| [**V1AccountAccountidPut**](AccountAPI.md#V1AccountAccountidPut) | **Put** /v1/account/{accountid} | Update Account |
| [**V1AccountApikeyGet**](AccountAPI.md#V1AccountApikeyGet) | **Get** /v1/account/apikey |  |
| [**V1AccountGet**](AccountAPI.md#V1AccountGet) | **Get** /v1/account | Get Account List |
| [**V1AccountPost**](AccountAPI.md#V1AccountPost) | **Post** /v1/account | Create Account |



## V1AccountAccountidChildrenGet

> ServiceDocsAccountGetAll V1AccountAccountidChildrenGet(accountid, startKey, pageSize)

Get Sub Account List

Conveniently access the list of children accounts.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsAccountGetAll result = apiInstance.V1AccountAccountidChildrenGet(accountid, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidChildrenGet");
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

[**ServiceDocsAccountGetAll**](ServiceDocsAccountGetAll.md)

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


## V1AccountAccountidDelete

> ServiceDocsAccountGetSingle V1AccountAccountidDelete(accountid)

Delete Account

Delete an account within your organization.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocsAccountGetSingle result = apiInstance.V1AccountAccountidDelete(accountid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidDelete");
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

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

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


## V1AccountAccountidDnsrecordGet

> ServiceDocsAccountGetSingle V1AccountAccountidDnsrecordGet(accountid)



Get the DNS record of an account from the Route 53 entry.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocsAccountGetSingle result = apiInstance.V1AccountAccountidDnsrecordGet(accountid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidDnsrecordGet");
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

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

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


## V1AccountAccountidDnsrecordPost

> ServiceDocsAccountGetSingle V1AccountAccountidDnsrecordPost(accountid)



Create the DNS record of an account with the help realm in the Route 53 entry.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocsAccountGetSingle result = apiInstance.V1AccountAccountidDnsrecordPost(accountid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidDnsrecordPost");
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

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

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


## V1AccountAccountidDnsrecordPut

> ServiceDocsAccountGetSingle V1AccountAccountidDnsrecordPut(accountid, dnsrecord)



Toggle the realm DNS record between srv and cname.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        ServiceUpdateRecordTypeForAccount dnsrecord = ; // ServiceUpdateRecordTypeForAccount | record type fields with value SRV, CNAME
        try {
            ServiceDocsAccountGetSingle result = apiInstance.V1AccountAccountidDnsrecordPut(accountid, dnsrecord);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidDnsrecordPut");
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
| **dnsrecord** | [**ServiceUpdateRecordTypeForAccount**](ServiceUpdateRecordTypeForAccount.md)| record type fields with value SRV, CNAME | |

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

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


## V1AccountAccountidGet

> ServiceDocsAccountGetSingle V1AccountAccountidGet(accountid)

Get Account Details

This endpoint will not allow for modifying or making updates, it will only allow users to view/retrieve details.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocsAccountGetSingle result = apiInstance.V1AccountAccountidGet(accountid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidGet");
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

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

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


## V1AccountAccountidLimitGet

> ServiceDocsAccountLimit V1AccountAccountidLimitGet(accountid)

Get Account Limits

Check the maximum number of inbound, outbound, and two-way trunks.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        try {
            ServiceDocsAccountLimit result = apiInstance.V1AccountAccountidLimitGet(accountid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidLimitGet");
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

### Return type

[**ServiceDocsAccountLimit**](ServiceDocsAccountLimit.md)

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


## V1AccountAccountidLimitPut

> ServiceDocsAccountLimit V1AccountAccountidLimitPut(accountid, limit)

Set Account Limits

Apply parameters to restrict access to inbound, outbound, and two-way trunks.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        ServiceVOIPAccountLimit2 limit = ; // ServiceVOIPAccountLimit2 | account fields
        try {
            ServiceDocsAccountLimit result = apiInstance.V1AccountAccountidLimitPut(accountid, limit);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidLimitPut");
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
| **limit** | [**ServiceVOIPAccountLimit2**](ServiceVOIPAccountLimit2.md)| account fields | |

### Return type

[**ServiceDocsAccountLimit**](ServiceDocsAccountLimit.md)

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


## V1AccountAccountidPost

> ServiceDocsAccountGetSingle V1AccountAccountidPost(accountid, account)

Create Sub Account

Establish a sub account to enable an administrator within your organization to create accounts.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        ServiceVOIPAccountAddData account = ; // ServiceVOIPAccountAddData | account fields
        try {
            ServiceDocsAccountGetSingle result = apiInstance.V1AccountAccountidPost(accountid, account);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidPost");
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
| **account** | [**ServiceVOIPAccountAddData**](ServiceVOIPAccountAddData.md)| account fields | |

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

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


## V1AccountAccountidPut

> ServiceDocsAccountGetSingle V1AccountAccountidPut(accountid, account)

Update Account

Modify pertinent account data.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        ServiceVOIPAccountEditData account = ; // ServiceVOIPAccountEditData | account fields
        try {
            ServiceDocsAccountGetSingle result = apiInstance.V1AccountAccountidPut(accountid, account);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountAccountidPut");
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
| **account** | [**ServiceVOIPAccountEditData**](ServiceVOIPAccountEditData.md)| account fields | |

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

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


## V1AccountApikeyGet

> ServiceDocsAccountAPIKey V1AccountApikeyGet()



Authenticate an application or user request to get the client ID and client secret for a CPaaS account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        try {
            ServiceDocsAccountAPIKey result = apiInstance.V1AccountApikeyGet();
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountApikeyGet");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceDocsAccountAPIKey**](ServiceDocsAccountAPIKey.md)

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


## V1AccountGet

> ServiceDocsAccountGetAll V1AccountGet(startKey, pageSize)

Get Account List

Retrieve a list of all CPaaS accounts that exist within your organization.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsAccountGetAll result = apiInstance.V1AccountGet(startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountGet");
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
| **startKey** | **string**| start_key for pagination that was returned as next_start_key from your previous call | [optional] |
| **pageSize** | **int32**| number of records to return, range 1 to 50 | [optional] |

### Return type

[**ServiceDocsAccountGetAll**](ServiceDocsAccountGetAll.md)

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


## V1AccountPost

> ServiceDocsAccountGetSingle V1AccountPost(account)

Create Account

Create an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.AccountAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        AccountAPI apiInstance = new AccountAPI(defaultClient);
        ServiceVOIPAccountAddData account = ; // ServiceVOIPAccountAddData | account fields
        try {
            ServiceDocsAccountGetSingle result = apiInstance.V1AccountPost(account);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountAPI#V1AccountPost");
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
| **account** | [**ServiceVOIPAccountAddData**](ServiceVOIPAccountAddData.md)| account fields | |

### Return type

[**ServiceDocsAccountGetSingle**](ServiceDocsAccountGetSingle.md)

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

