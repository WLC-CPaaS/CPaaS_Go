# VoicemailAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDVoicemailGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailGet) | **Get** /v1/account/{accountID}/voicemail | Get Voicemail Box List |
| [**V1AccountAccountIDVoicemailPost**](VoicemailAPI.md#V1AccountAccountIDVoicemailPost) | **Post** /v1/account/{accountID}/voicemail | Create Voicemail Box |
| [**V1AccountAccountIDVoicemailVoicemailIDDelete**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDDelete) | **Delete** /v1/account/{accountID}/voicemail/{voicemailID} | Delete Voicemail Box |
| [**V1AccountAccountIDVoicemailVoicemailIDGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDGet) | **Get** /v1/account/{accountID}/voicemail/{voicemailID} | Get Voicemail Box Details |
| [**V1AccountAccountIDVoicemailVoicemailIDMessageGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageGet) | **Get** /v1/account/{accountID}/voicemail/{voicemailID}/message | Get Voicemail Message List |
| [**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete) | **Delete** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID} | Delete Voicemail Message |
| [**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet) | **Get** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID} | Get Voicemail Message Details |
| [**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut) | **Put** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID} | Update Voicemail Message |
| [**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet) | **Get** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID}/raw | Get Voicemail Message File |
| [**V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost) | **Post** /v1/account/{accountID}/voicemail/{voicemailID}/message/{messageID}/raw | Add Voicemail Message File |
| [**V1AccountAccountIDVoicemailVoicemailIDMessagePost**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDMessagePost) | **Post** /v1/account/{accountID}/voicemail/{voicemailID}/message | Create Voicemail Message |
| [**V1AccountAccountIDVoicemailVoicemailIDPut**](VoicemailAPI.md#V1AccountAccountIDVoicemailVoicemailIDPut) | **Put** /v1/account/{accountID}/voicemail/{voicemailID} | Update Voicemail Box |



## V1AccountAccountIDVoicemailGet

> ServiceDocsVoicemailGetAll V1AccountAccountIDVoicemailGet(accountID, startKey, pageSize)

Get Voicemail Box List

List all voicemail boxes in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsVoicemailGetAll result = apiInstance.V1AccountAccountIDVoicemailGet(accountID, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailGet");
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

[**ServiceDocsVoicemailGetAll**](ServiceDocsVoicemailGetAll.md)

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


## V1AccountAccountIDVoicemailPost

> ServiceDocsVoicemailGetSingle V1AccountAccountIDVoicemailPost(accountID, voicemail)

Create Voicemail Box

Create a voicemail box for receiving and storing voicemail messages.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | account ID, 32 alphanumeric
        ServiceVOIPVoicemailAddEditData voicemail = ; // ServiceVOIPVoicemailAddEditData | voicemail payload fields
        try {
            ServiceDocsVoicemailGetSingle result = apiInstance.V1AccountAccountIDVoicemailPost(accountID, voicemail);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailPost");
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
| **accountID** | **string**| account ID, 32 alphanumeric | |
| **voicemail** | [**ServiceVOIPVoicemailAddEditData**](ServiceVOIPVoicemailAddEditData.md)| voicemail payload fields | |

### Return type

[**ServiceDocsVoicemailGetSingle**](ServiceDocsVoicemailGetSingle.md)

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


## V1AccountAccountIDVoicemailVoicemailIDDelete

> ServiceDocsVoicemailGetSingle V1AccountAccountIDVoicemailVoicemailIDDelete(accountID, voicemailID)

Delete Voicemail Box

Delete a voicemail box in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string voicemailID = voicemailID_example; // string | Voicemail ID, 32 alpha numeric
        try {
            ServiceDocsVoicemailGetSingle result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDDelete(accountID, voicemailID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDDelete");
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
| **voicemailID** | **string**| Voicemail ID, 32 alpha numeric | |

### Return type

[**ServiceDocsVoicemailGetSingle**](ServiceDocsVoicemailGetSingle.md)

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


## V1AccountAccountIDVoicemailVoicemailIDGet

> ServiceDocsVoicemailGetSingle V1AccountAccountIDVoicemailVoicemailIDGet(accountID, voicemailID)

Get Voicemail Box Details

Get information about a single voicemail box.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string voicemailID = voicemailID_example; // string | Voicemail ID, 32 alpha numeric
        try {
            ServiceDocsVoicemailGetSingle result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDGet(accountID, voicemailID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDGet");
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
| **voicemailID** | **string**| Voicemail ID, 32 alpha numeric | |

### Return type

[**ServiceDocsVoicemailGetSingle**](ServiceDocsVoicemailGetSingle.md)

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


## V1AccountAccountIDVoicemailVoicemailIDMessageGet

> ServiceDocsVoicemailMessageGetAll V1AccountAccountIDVoicemailVoicemailIDMessageGet(accountID, voicemailID, startKey, pageSize)

Get Voicemail Message List

Get a list of voicemail messages from an account&#39;s voicemail box.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string voicemailID = voicemailID_example; // string | voicemail ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsVoicemailMessageGetAll result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDMessageGet(accountID, voicemailID, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDMessageGet");
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
| **voicemailID** | **string**| voicemail ID, 32 alpha numeric | |
| **startKey** | **string**| start_key for pagination that was returned as next_start_key from your previous call | [optional] |
| **pageSize** | **int32**| number of records to return, range 1 to 50 | [optional] |

### Return type

[**ServiceDocsVoicemailMessageGetAll**](ServiceDocsVoicemailMessageGetAll.md)

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


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete

> ServiceDocsVoicemailMessageGetSingle V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete(accountID, voicemailID, messageID)

Delete Voicemail Message

Delete a voicemail message from a voicemail box in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string voicemailID = voicemailID_example; // string | Voicemail ID, 32 alpha numeric
        string messageID = messageID_example; // string | message ID, 32 alpha numeric
        try {
            ServiceDocsVoicemailMessageGetSingle result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete(accountID, voicemailID, messageID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDDelete");
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
| **voicemailID** | **string**| Voicemail ID, 32 alpha numeric | |
| **messageID** | **string**| message ID, 32 alpha numeric | |

### Return type

[**ServiceDocsVoicemailMessageGetSingle**](ServiceDocsVoicemailMessageGetSingle.md)

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


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet

> ServiceDocsVoicemailMessageGetSingle V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet(accountID, voicemailID, messageID)

Get Voicemail Message Details

Retrieve the container details of an individual voicemail message. This includes a reference to the audio file, but not the message itself.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string voicemailID = voicemailID_example; // string | Voicemail ID, 32 alpha numeric
        string messageID = messageID_example; // string | Message ID, 39 (yyyymm-<32 id>)
        try {
            ServiceDocsVoicemailMessageGetSingle result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet(accountID, voicemailID, messageID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDGet");
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
| **voicemailID** | **string**| Voicemail ID, 32 alpha numeric | |
| **messageID** | **string**| Message ID, 39 (yyyymm-&lt;32 id&gt;) | |

### Return type

[**ServiceDocsVoicemailMessageGetSingle**](ServiceDocsVoicemailMessageGetSingle.md)

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


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut

> ServiceDocsVoicemailMessageGetSingle V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut(accountID, voicemailID, messageID, reqBody)

Update Voicemail Message

Copy or move a voicemail message to a different folder in the same voicemail box or move the message to a separate voicemail box.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string voicemailID = voicemailID_example; // string | Voicemail ID, 32 alpha numeric
        string messageID = messageID_example; // string | Message ID, 39 (yyyymm-<32 id>)
        ServiceVOIPVoicemailMessageChange reqBody = ; // ServiceVOIPVoicemailMessageChange | payload fields
        try {
            ServiceDocsVoicemailMessageGetSingle result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut(accountID, voicemailID, messageID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDPut");
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
| **voicemailID** | **string**| Voicemail ID, 32 alpha numeric | |
| **messageID** | **string**| Message ID, 39 (yyyymm-&lt;32 id&gt;) | |
| **reqBody** | [**ServiceVOIPVoicemailMessageChange**](ServiceVOIPVoicemailMessageChange.md)| payload fields | |

### Return type

[**ServiceDocsVoicemailMessageGetSingle**](ServiceDocsVoicemailMessageGetSingle.md)

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


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet

> *os.File V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet(accountID, voicemailID, messageID)

Get Voicemail Message File

Get the original audio content of a specific voicemail message identified by its unique ID within an account&#39;s voicemail box. URL Param \&quot;voicemailID\&quot; is a unique 32-character alphanumeric identifier assigned by the system, which refers to a specific voicemail box. URL Param \&quot;messageID\&quot; is a unique 32-character alphanumeric identifier assigned by the system, which refers to a specific message within a voicemail box.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, unique 32-character alphanumeric identifier
        string voicemailID = voicemailID_example; // string | Voicemail Box ID, unique 32-character alphanumeric identifier
        string messageID = messageID_example; // string | Message ID, unique 32-character alphanumeric identifier
        try {
            *os.File result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet(accountID, voicemailID, messageID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawGet");
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
| **accountID** | **string**| Account ID, unique 32-character alphanumeric identifier | |
| **voicemailID** | **string**| Voicemail Box ID, unique 32-character alphanumeric identifier | |
| **messageID** | **string**| Message ID, unique 32-character alphanumeric identifier | |

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |


## V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost

> map[string]interface{} V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost(accountID, voicemailID, messageID, file)

Add Voicemail Message File

Associate an audio recording file with the voicemail to fully complete the message.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alphanumeric characters
        string voicemailID = voicemailID_example; // string | Voicemail ID, 32 alphanumeric characters
        string messageID = messageID_example; // string | Message ID, 32 alphanumeric characters
        *os.File file = BINARY_DATA_HERE; // *os.File | Audio file to upload
        try {
            map[string]interface{} result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost(accountID, voicemailID, messageID, file);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDMessageMessageIDRawPost");
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
| **accountID** | **string**| Account ID, 32 alphanumeric characters | |
| **voicemailID** | **string**| Voicemail ID, 32 alphanumeric characters | |
| **messageID** | **string**| Message ID, 32 alphanumeric characters | |
| **file** | ***os.File**| Audio file to upload | |

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |


## V1AccountAccountIDVoicemailVoicemailIDMessagePost

> ServiceDocsVoicemailMessageGetSingle V1AccountAccountIDVoicemailVoicemailIDMessagePost(accountID, voicemailID, message)

Create Voicemail Message

Create the container information for a recorded voicemail message in a voicemail box.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | account ID, 32 alphanumeric
        string voicemailID = voicemailID_example; // string | voicemail ID, 32 alphanumeric
        ServiceVOIPVoicemailMessageAddData message = ; // ServiceVOIPVoicemailMessageAddData | voicemail message payload fields
        try {
            ServiceDocsVoicemailMessageGetSingle result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDMessagePost(accountID, voicemailID, message);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDMessagePost");
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
| **accountID** | **string**| account ID, 32 alphanumeric | |
| **voicemailID** | **string**| voicemail ID, 32 alphanumeric | |
| **message** | [**ServiceVOIPVoicemailMessageAddData**](ServiceVOIPVoicemailMessageAddData.md)| voicemail message payload fields | |

### Return type

[**ServiceDocsVoicemailMessageGetSingle**](ServiceDocsVoicemailMessageGetSingle.md)

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


## V1AccountAccountIDVoicemailVoicemailIDPut

> ServiceDocsVoicemailGetSingle V1AccountAccountIDVoicemailVoicemailIDPut(accountID, voicemailID, reqBody)

Update Voicemail Box

Update the settings in an individual voicemail box, such as the owner, PIN, etc.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.VoicemailAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        VoicemailAPI apiInstance = new VoicemailAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string voicemailID = voicemailID_example; // string | Voicemail ID, 32 alpha numeric
        ServiceVOIPVoicemailAddEditData reqBody = ; // ServiceVOIPVoicemailAddEditData | payload fields
        try {
            ServiceDocsVoicemailGetSingle result = apiInstance.V1AccountAccountIDVoicemailVoicemailIDPut(accountID, voicemailID, reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling VoicemailAPI#V1AccountAccountIDVoicemailVoicemailIDPut");
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
| **voicemailID** | **string**| Voicemail ID, 32 alpha numeric | |
| **reqBody** | [**ServiceVOIPVoicemailAddEditData**](ServiceVOIPVoicemailAddEditData.md)| payload fields | |

### Return type

[**ServiceDocsVoicemailGetSingle**](ServiceDocsVoicemailGetSingle.md)

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

