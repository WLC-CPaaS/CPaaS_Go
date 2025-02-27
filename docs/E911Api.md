# E911API

All URIs are relative to *http://api.cpaaslabs.net*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**V1E911Get**](E911API.md#V1E911Get) | **Get** /v1/e911 | Get E911 List |
| [**V1E911LocationLocationIDActivatePut**](E911API.md#V1E911LocationLocationIDActivatePut) | **Put** /v1/e911/location/{locationID}/activate | Activate E911 Location |
| [**V1E911LocationLocationIDDelete**](E911API.md#V1E911LocationLocationIDDelete) | **Delete** /v1/e911/location/{locationID} | Delete E911 Location |
| [**V1E911LocationValidatePut**](E911API.md#V1E911LocationValidatePut) | **Put** /v1/e911/location/validate | Validate a Location |
| [**V1E911PhoneNumberDelete**](E911API.md#V1E911PhoneNumberDelete) | **Delete** /v1/e911/{phoneNumber} | Delete E911 Phone Number |
| [**V1E911PhoneNumberLocationActiveGet**](E911API.md#V1E911PhoneNumberLocationActiveGet) | **Get** /v1/e911/{phoneNumber}/location/active | Get Actvie Location for a Phone Number |
| [**V1E911PhoneNumberLocationGet**](E911API.md#V1E911PhoneNumberLocationGet) | **Get** /v1/e911/{phoneNumber}/location | Get Location List for Phone Number |
| [**V1E911Post**](E911API.md#V1E911Post) | **Post** /v1/e911 | Create an E911 Location |



## V1E911Get

> ServiceDocE911URIsApiOutput V1E911Get()

Get E911 List

Obtain e911 URIs associated with the provided account ID.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .models.*;
import openapi.E911API;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");

        E911API apiInstance = new E911API(defaultClient);
        try {
            ServiceDocE911URIsApiOutput result = apiInstance.V1E911Get();
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling E911API#V1E911Get");
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

[**ServiceDocE911URIsApiOutput**](ServiceDocE911URIsApiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response with e911 URIs |  -  |
| **403** | Authorization failed or root account not allowed |  -  |
| **500** | Internal server error, including environment credential issues, HTTP request failures, or XML unmarshaling errors |  -  |


## V1E911LocationLocationIDActivatePut

> ServiceDocE911ActiveLocationOutput V1E911LocationLocationIDActivatePut(locationID)

Activate E911 Location

Edit the provision location.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .models.*;
import openapi.E911API;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");

        E911API apiInstance = new E911API(defaultClient);
        string locationID = locationID_example; // string | Location ID
        try {
            ServiceDocE911ActiveLocationOutput result = apiInstance.V1E911LocationLocationIDActivatePut(locationID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling E911API#V1E911LocationLocationIDActivatePut");
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
| **locationID** | **string**| Location ID | |

### Return type

[**ServiceDocE911ActiveLocationOutput**](ServiceDocE911ActiveLocationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response with location activate status |  -  |
| **403** | Authorization failed or root account not allowed |  -  |
| **500** | Internal server error, including environment credential issues, HTTP request failures, or XML unmarshaling errors |  -  |


## V1E911LocationLocationIDDelete

> ServiceDocE911RemoveLocationOutput V1E911LocationLocationIDDelete(locationID)

Delete E911 Location

Remove the location.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .models.*;
import openapi.E911API;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");

        E911API apiInstance = new E911API(defaultClient);
        string locationID = locationID_example; // string | Location ID
        try {
            ServiceDocE911RemoveLocationOutput result = apiInstance.V1E911LocationLocationIDDelete(locationID);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling E911API#V1E911LocationLocationIDDelete");
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
| **locationID** | **string**| Location ID | |

### Return type

[**ServiceDocE911RemoveLocationOutput**](ServiceDocE911RemoveLocationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response with location remove status |  -  |
| **403** | Authorization failed or root account not allowed |  -  |
| **500** | Internal server error, including environment credential issues, HTTP request failures, or XML unmarshaling errors |  -  |


## V1E911LocationValidatePut

> ServiceDocE911ValidateLocationOutput V1E911LocationValidatePut(reqBody)

Validate a Location

Validate the location details.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .models.*;
import openapi.E911API;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");

        E911API apiInstance = new E911API(defaultClient);
        ServiceE911ValidateLocationInput reqBody = ; // ServiceE911ValidateLocationInput | location details
        try {
            ServiceDocE911ValidateLocationOutput result = apiInstance.V1E911LocationValidatePut(reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling E911API#V1E911LocationValidatePut");
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
| **reqBody** | [**ServiceE911ValidateLocationInput**](ServiceE911ValidateLocationInput.md)| location details | |

### Return type

[**ServiceDocE911ValidateLocationOutput**](ServiceDocE911ValidateLocationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response with location details |  -  |
| **403** | Authorization failed or root account not allowed |  -  |
| **500** | Internal server error, including environment credential issues, HTTP request failures, or XML unmarshaling errors |  -  |


## V1E911PhoneNumberDelete

> ServiceDocE911RemoveURIApiOutput V1E911PhoneNumberDelete(phoneNumber)

Delete E911 Phone Number

Delete the e911 URI connected with the account URI.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .models.*;
import openapi.E911API;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");

        E911API apiInstance = new E911API(defaultClient);
        string phoneNumber = phoneNumber_example; // string | Phone Number
        try {
            ServiceDocE911RemoveURIApiOutput result = apiInstance.V1E911PhoneNumberDelete(phoneNumber);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling E911API#V1E911PhoneNumberDelete");
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
| **phoneNumber** | **string**| Phone Number | |

### Return type

[**ServiceDocE911RemoveURIApiOutput**](ServiceDocE911RemoveURIApiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **403** | Authorization failed or root account not allowed |  -  |
| **500** | Internal server error, including environment credential issues, HTTP request failures, or XML unmarshaling errors |  -  |


## V1E911PhoneNumberLocationActiveGet

> ServiceDocE911ActiveLocationURIApiOutput V1E911PhoneNumberLocationActiveGet(phoneNumber)

Get Actvie Location for a Phone Number

Get the e911 location connected with the URI.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .models.*;
import openapi.E911API;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");

        E911API apiInstance = new E911API(defaultClient);
        string phoneNumber = phoneNumber_example; // string | Phone Number
        try {
            ServiceDocE911ActiveLocationURIApiOutput result = apiInstance.V1E911PhoneNumberLocationActiveGet(phoneNumber);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling E911API#V1E911PhoneNumberLocationActiveGet");
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
| **phoneNumber** | **string**| Phone Number | |

### Return type

[**ServiceDocE911ActiveLocationURIApiOutput**](ServiceDocE911ActiveLocationURIApiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response with e911 Active Location URI |  -  |
| **403** | Authorization failed or root account not allowed |  -  |
| **500** | Internal server error, including environment credential issues, HTTP request failures, or XML unmarshaling errors |  -  |


## V1E911PhoneNumberLocationGet

> ServiceDocE911LocationsURIApiOutput V1E911PhoneNumberLocationGet(phoneNumber)

Get Location List for Phone Number

Access a list of the e911 locations associated with the provided URI.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .models.*;
import openapi.E911API;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");

        E911API apiInstance = new E911API(defaultClient);
        string phoneNumber = phoneNumber_example; // string | Phone Number
        try {
            ServiceDocE911LocationsURIApiOutput result = apiInstance.V1E911PhoneNumberLocationGet(phoneNumber);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling E911API#V1E911PhoneNumberLocationGet");
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
| **phoneNumber** | **string**| Phone Number | |

### Return type

[**ServiceDocE911LocationsURIApiOutput**](ServiceDocE911LocationsURIApiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response with e911 Locations URI |  -  |
| **403** | Authorization failed or root account not allowed |  -  |
| **500** | Internal server error, including environment credential issues, HTTP request failures, or XML unmarshaling errors |  -  |


## V1E911Post

> ServiceDocE911AddLocationOutput V1E911Post(reqBody)

Create an E911 Location

Enter new location details.

### Example

```java
// Import classes:
import .ApiClient;
import .ApiException;
import .Configuration;
import .models.*;
import openapi.E911API;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://api.cpaaslabs.net");

        E911API apiInstance = new E911API(defaultClient);
        ServiceE911AddLocationInput reqBody = ; // ServiceE911AddLocationInput | location details
        try {
            ServiceDocE911AddLocationOutput result = apiInstance.V1E911Post(reqBody);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling E911API#V1E911Post");
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
| **reqBody** | [**ServiceE911AddLocationInput**](ServiceE911AddLocationInput.md)| location details | |

### Return type

[**ServiceDocE911AddLocationOutput**](ServiceDocE911AddLocationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response with location details |  -  |
| **403** | Authorization failed or root account not allowed |  -  |
| **500** | Internal server error, including environment credential issues, HTTP request failures, or XML unmarshaling errors |  -  |

