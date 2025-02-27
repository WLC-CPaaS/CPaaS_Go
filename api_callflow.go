/*
 * White Label Communications CPaas API Documentation
 * A CPaaS platform API
 *
 * The version of the OpenAPI document: 1.1
 * Contact: support@whitelabelcomm.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi;

import com.sun.jersey.api.client.GenericType;

import .ApiException;
import .ApiClient;
import .Configuration;
import openapi.*;
import .Pair;

import strings;


import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class CallflowAPI {
  private ApiClient apiClient;

  public CallflowAPI() {
    this(Configuration.getDefaultApiClient());
  }

  public CallflowAPI(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  public ApiClient getApiClient() {
    return apiClient;
  }

  public void setApiClient(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  /**
   * Delete Call Group
   * Delete a callflow in an account.
   * @param accountID Account ID, 32 alpha numeric (required)
   * @param callflowID callflow ID, 32 alpha numeric (required)
   * @return ServiceDocsCallflowGetSingle
   * @throws ApiException if fails to make API call
   */
  public ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDDelete(string accountID, string callflowID) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'accountID' is set
    if (accountID == null) {
      throw new ApiException(400, "Missing the required parameter 'accountID' when calling V1AccountAccountIDCallflowCallflowIDDelete");
    }
    
    // verify the required parameter 'callflowID' is set
    if (callflowID == null) {
      throw new ApiException(400, "Missing the required parameter 'callflowID' when calling V1AccountAccountIDCallflowCallflowIDDelete");
    }
    
    // create path and map variables
    String localVarPath = "/v1/account/{accountID}/callflow/{callflowID}"
      .replaceAll("\\{" + "accountID" + "\\}", apiClient.escapeString(accountID.toString()))
      .replaceAll("\\{" + "callflowID" + "\\}", apiClient.escapeString(callflowID.toString()));

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();


    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "BearerAuth" };

    GenericType<ServiceDocsCallflowGetSingle> localVarReturnType = new GenericType<ServiceDocsCallflowGetSingle>() {};
    return apiClient.invokeAPI(localVarPath, "Delete", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * Get Call Group Details
   * Get the details for a single callflow in an account.
   * @param accountID Account ID, 32 alpha numeric (required)
   * @param callflowID Callflow ID, 32 alpha numeric (required)
   * @return ServiceDocsCallflowGetSingle
   * @throws ApiException if fails to make API call
   */
  public ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDGet(string accountID, string callflowID) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'accountID' is set
    if (accountID == null) {
      throw new ApiException(400, "Missing the required parameter 'accountID' when calling V1AccountAccountIDCallflowCallflowIDGet");
    }
    
    // verify the required parameter 'callflowID' is set
    if (callflowID == null) {
      throw new ApiException(400, "Missing the required parameter 'callflowID' when calling V1AccountAccountIDCallflowCallflowIDGet");
    }
    
    // create path and map variables
    String localVarPath = "/v1/account/{accountID}/callflow/{callflowID}"
      .replaceAll("\\{" + "accountID" + "\\}", apiClient.escapeString(accountID.toString()))
      .replaceAll("\\{" + "callflowID" + "\\}", apiClient.escapeString(callflowID.toString()));

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();


    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "BearerAuth" };

    GenericType<ServiceDocsCallflowGetSingle> localVarReturnType = new GenericType<ServiceDocsCallflowGetSingle>() {};
    return apiClient.invokeAPI(localVarPath, "Get", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * Update Call Group
   * Update the details for a single callflow in an account.
   * @param accountID Account ID, 32 alpha numeric (required)
   * @param callflowID Callflow ID, 32 alpha numeric (required)
   * @param reqBody payload fields (required)
   * @return ServiceDocsCallflowGetSingle
   * @throws ApiException if fails to make API call
   */
  public ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowCallflowIDPut(string accountID, string callflowID, ServiceCallflowAddEditData reqBody) throws ApiException {
    Object localVarPostBody = reqBody;
    
    // verify the required parameter 'accountID' is set
    if (accountID == null) {
      throw new ApiException(400, "Missing the required parameter 'accountID' when calling V1AccountAccountIDCallflowCallflowIDPut");
    }
    
    // verify the required parameter 'callflowID' is set
    if (callflowID == null) {
      throw new ApiException(400, "Missing the required parameter 'callflowID' when calling V1AccountAccountIDCallflowCallflowIDPut");
    }
    
    // verify the required parameter 'reqBody' is set
    if (reqBody == null) {
      throw new ApiException(400, "Missing the required parameter 'reqBody' when calling V1AccountAccountIDCallflowCallflowIDPut");
    }
    
    // create path and map variables
    String localVarPath = "/v1/account/{accountID}/callflow/{callflowID}"
      .replaceAll("\\{" + "accountID" + "\\}", apiClient.escapeString(accountID.toString()))
      .replaceAll("\\{" + "callflowID" + "\\}", apiClient.escapeString(callflowID.toString()));

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();


    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      "application/json"
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "BearerAuth" };

    GenericType<ServiceDocsCallflowGetSingle> localVarReturnType = new GenericType<ServiceDocsCallflowGetSingle>() {};
    return apiClient.invokeAPI(localVarPath, "Put", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * Get Callflow List
   * Permit a user to view the callflow details in an account.
   * @param accountID Account ID, 32 alpha numeric (required)
   * @param startKey start_key for pagination that was returned as next_start_key from your previous call (optional)
   * @param pageSize number of records to return, range 1 to 50 (optional)
   * @return ServiceDocsCallflowGetAll
   * @throws ApiException if fails to make API call
   */
  public ServiceDocsCallflowGetAll V1AccountAccountIDCallflowGet(string accountID, string startKey, int32 pageSize) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'accountID' is set
    if (accountID == null) {
      throw new ApiException(400, "Missing the required parameter 'accountID' when calling V1AccountAccountIDCallflowGet");
    }
    
    // create path and map variables
    String localVarPath = "/v1/account/{accountID}/callflow"
      .replaceAll("\\{" + "accountID" + "\\}", apiClient.escapeString(accountID.toString()));

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    localVarQueryParams.addAll(apiClient.parameterToPair("start_key", startKey));
    localVarQueryParams.addAll(apiClient.parameterToPair("page_size", pageSize));

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "BearerAuth" };

    GenericType<ServiceDocsCallflowGetAll> localVarReturnType = new GenericType<ServiceDocsCallflowGetAll>() {};
    return apiClient.invokeAPI(localVarPath, "Get", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * Create Call Group
   * Create instructions for routing a call to a user or system.
   * @param accountID Account ID, 32 alpha-numeric (required)
   * @param request Call flow configuration (required)
   * @return ServiceDocsCallflowGetSingle
   * @throws ApiException if fails to make API call
   */
  public ServiceDocsCallflowGetSingle V1AccountAccountIDCallflowPost(string accountID, ServiceCallflowAddEditData request) throws ApiException {
    Object localVarPostBody = request;
    
    // verify the required parameter 'accountID' is set
    if (accountID == null) {
      throw new ApiException(400, "Missing the required parameter 'accountID' when calling V1AccountAccountIDCallflowPost");
    }
    
    // verify the required parameter 'request' is set
    if (request == null) {
      throw new ApiException(400, "Missing the required parameter 'request' when calling V1AccountAccountIDCallflowPost");
    }
    
    // create path and map variables
    String localVarPath = "/v1/account/{accountID}/callflow"
      .replaceAll("\\{" + "accountID" + "\\}", apiClient.escapeString(accountID.toString()));

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();


    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      "application/json"
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "BearerAuth" };

    GenericType<ServiceDocsCallflowGetSingle> localVarReturnType = new GenericType<ServiceDocsCallflowGetSingle>() {};
    return apiClient.invokeAPI(localVarPath, "Post", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
}
