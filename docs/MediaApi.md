# MediaAPI

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1AccountAccountIDMediaMediaIDFileGet**](MediaAPI.md#V1AccountAccountIDMediaMediaIDFileGet) | **Get** /v1/account/{accountID}/media/{mediaID}/file | Get Media File |
| [**V1AccountAccountIDMediaMediaIDFilePost**](MediaAPI.md#V1AccountAccountIDMediaMediaIDFilePost) | **Post** /v1/account/{accountID}/media/{mediaID}/file | Add Media File |
| [**V1AccountAccountidMediaGet**](MediaAPI.md#V1AccountAccountidMediaGet) | **Get** /v1/account/{accountid}/media | Get Media List |
| [**V1AccountAccountidMediaMediaidDelete**](MediaAPI.md#V1AccountAccountidMediaMediaidDelete) | **Delete** /v1/account/{accountid}/media/{mediaid} | Delete Media |
| [**V1AccountAccountidMediaMediaidGet**](MediaAPI.md#V1AccountAccountidMediaMediaidGet) | **Get** /v1/account/{accountid}/media/{mediaid} | Get Media Details |
| [**V1AccountAccountidMediaPost**](MediaAPI.md#V1AccountAccountidMediaPost) | **Post** /v1/account/{accountid}/media | Create Media |



## V1AccountAccountIDMediaMediaIDFileGet

> *os.File V1AccountAccountIDMediaMediaIDFileGet(accountID, mediaID)

Get Media File

Gather data about the media objects in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MediaAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MediaAPI apiInstance = new MediaAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string mediaID = mediaID_example; // string | Media ID, 32 alpha numeric
        try {
            *os.File result = apiInstance.V1AccountAccountIDMediaMediaIDFileGet(accountID, mediaID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MediaAPI#V1AccountAccountIDMediaMediaIDFileGet");
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
| **mediaID** | **string**| Media ID, 32 alpha numeric | |

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, audio/mp3, audio/mpeg, audio/mpeg3, audio/x-wav, audio/wav, audio/ogg, video/x-flv, video/h264, video/mpeg, video/quicktime, video/mp4, video/webm


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |


## V1AccountAccountIDMediaMediaIDFilePost

> ServiceDocsMediaGetSingle V1AccountAccountIDMediaMediaIDFilePost(accountID, mediaID, file)

Add Media File

Include a media file that is connected to a media object in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MediaAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MediaAPI apiInstance = new MediaAPI(defaultClient);
        string accountID = accountID_example; // string | Account ID, 32 alpha numeric
        string mediaID = mediaID_example; // string | Media ID, 32 alpha numeric
        *os.File file = BINARY_DATA_HERE; // *os.File | Media file
        try {
            ServiceDocsMediaGetSingle result = apiInstance.V1AccountAccountIDMediaMediaIDFilePost(accountID, mediaID, file);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MediaAPI#V1AccountAccountIDMediaMediaIDFilePost");
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
| **mediaID** | **string**| Media ID, 32 alpha numeric | |
| **file** | ***os.File**| Media file | |

### Return type

[**ServiceDocsMediaGetSingle**](ServiceDocsMediaGetSingle.md)

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


## V1AccountAccountidMediaGet

> ServiceDocsMediaGetAll V1AccountAccountidMediaGet(accountid, startKey, pageSize)

Get Media List

View all media files for an account in your organization.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MediaAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MediaAPI apiInstance = new MediaAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string startKey = startKey_example; // string | start_key for pagination that was returned as next_start_key from your previous call
        int32 pageSize = 56; // int32 | number of records to return, range 1 to 50
        try {
            ServiceDocsMediaGetAll result = apiInstance.V1AccountAccountidMediaGet(accountid, startKey, pageSize);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MediaAPI#V1AccountAccountidMediaGet");
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

[**ServiceDocsMediaGetAll**](ServiceDocsMediaGetAll.md)

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


## V1AccountAccountidMediaMediaidDelete

> ServiceDocsMediaGetSingle V1AccountAccountidMediaMediaidDelete(accountid, mediaid)

Delete Media

Remove a media file that is no longer in use from an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MediaAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MediaAPI apiInstance = new MediaAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string mediaid = mediaid_example; // string | Device ID, 32 alpha numeric
        try {
            ServiceDocsMediaGetSingle result = apiInstance.V1AccountAccountidMediaMediaidDelete(accountid, mediaid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MediaAPI#V1AccountAccountidMediaMediaidDelete");
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
| **mediaid** | **string**| Device ID, 32 alpha numeric | |

### Return type

[**ServiceDocsMediaGetSingle**](ServiceDocsMediaGetSingle.md)

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


## V1AccountAccountidMediaMediaidGet

> ServiceDocsMediaGetSingle V1AccountAccountidMediaMediaidGet(accountid, mediaid)

Get Media Details

Permit users to view an account&#39;s specific media information.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MediaAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MediaAPI apiInstance = new MediaAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        string mediaid = mediaid_example; // string | Media ID, 32 alpha numeric
        try {
            ServiceDocsMediaGetSingle result = apiInstance.V1AccountAccountidMediaMediaidGet(accountid, mediaid);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MediaAPI#V1AccountAccountidMediaMediaidGet");
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
| **mediaid** | **string**| Media ID, 32 alpha numeric | |

### Return type

[**ServiceDocsMediaGetSingle**](ServiceDocsMediaGetSingle.md)

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


## V1AccountAccountidMediaPost

> ServiceDocsMediaGetSingle V1AccountAccountidMediaPost(accountid, media)

Create Media

Generate a media object to allow users to upload a media file in an account.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .auth.*;
import .models.*;
import openapi.MediaAPI;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");
        
        // Configure API key authorization: BearerAuth
        ApiKeyAuth BearerAuth = (ApiKeyAuth) defaultClient.getAuthentication("BearerAuth");
        BearerAuth.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //BearerAuth.setApiKeyPrefix("Token");

        MediaAPI apiInstance = new MediaAPI(defaultClient);
        string accountid = accountid_example; // string | Account ID, 32 alpha numeric
        ServiceVOIPMediaAddEditData media = ; // ServiceVOIPMediaAddEditData | Media creation or update payload
        try {
            ServiceDocsMediaGetSingle result = apiInstance.V1AccountAccountidMediaPost(accountid, media);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MediaAPI#V1AccountAccountidMediaPost");
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
| **media** | [**ServiceVOIPMediaAddEditData**](ServiceVOIPMediaAddEditData.md)| Media creation or update payload | |

### Return type

[**ServiceDocsMediaGetSingle**](ServiceDocsMediaGetSingle.md)

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

