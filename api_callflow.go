/*
White Label Communications CPaas API Documentation

A CPaaS platform API

API version: 1.1
Contact: support@whitelabelcomm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CallflowAPIService CallflowAPI service
type CallflowAPIService service

type ApiV1AccountAccountIDCallflowCallflowIDDeleteRequest struct {
	ctx context.Context
	ApiService *CallflowAPIService
	accountID string
	callflowID string
}

func (r ApiV1AccountAccountIDCallflowCallflowIDDeleteRequest) Execute() (*ServiceDocsCallflowGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDCallflowCallflowIDDeleteExecute(r)
}

/*
V1AccountAccountIDCallflowCallflowIDDelete Delete Call Group

Delete a callflow in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param callflowID callflow ID, 32 alpha numeric
 @return ApiV1AccountAccountIDCallflowCallflowIDDeleteRequest
*/
func (a *CallflowAPIService) V1AccountAccountIDCallflowCallflowIDDelete(ctx context.Context, accountID string, callflowID string) ApiV1AccountAccountIDCallflowCallflowIDDeleteRequest {
	return ApiV1AccountAccountIDCallflowCallflowIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		callflowID: callflowID,
	}
}

// Execute executes the request
//  @return ServiceDocsCallflowGetSingle
func (a *CallflowAPIService) V1AccountAccountIDCallflowCallflowIDDeleteExecute(r ApiV1AccountAccountIDCallflowCallflowIDDeleteRequest) (*ServiceDocsCallflowGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCallflowGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallflowAPIService.V1AccountAccountIDCallflowCallflowIDDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/callflow/{callflowID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"callflowID"+"}", url.PathEscape(parameterValueToString(r.callflowID, "callflowID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountIDCallflowCallflowIDGetRequest struct {
	ctx context.Context
	ApiService *CallflowAPIService
	accountID string
	callflowID string
}

func (r ApiV1AccountAccountIDCallflowCallflowIDGetRequest) Execute() (*ServiceDocsCallflowGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDCallflowCallflowIDGetExecute(r)
}

/*
V1AccountAccountIDCallflowCallflowIDGet Get Call Group Details

Get the details for a single callflow in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param callflowID Callflow ID, 32 alpha numeric
 @return ApiV1AccountAccountIDCallflowCallflowIDGetRequest
*/
func (a *CallflowAPIService) V1AccountAccountIDCallflowCallflowIDGet(ctx context.Context, accountID string, callflowID string) ApiV1AccountAccountIDCallflowCallflowIDGetRequest {
	return ApiV1AccountAccountIDCallflowCallflowIDGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		callflowID: callflowID,
	}
}

// Execute executes the request
//  @return ServiceDocsCallflowGetSingle
func (a *CallflowAPIService) V1AccountAccountIDCallflowCallflowIDGetExecute(r ApiV1AccountAccountIDCallflowCallflowIDGetRequest) (*ServiceDocsCallflowGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCallflowGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallflowAPIService.V1AccountAccountIDCallflowCallflowIDGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/callflow/{callflowID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"callflowID"+"}", url.PathEscape(parameterValueToString(r.callflowID, "callflowID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountIDCallflowCallflowIDPutRequest struct {
	ctx context.Context
	ApiService *CallflowAPIService
	accountID string
	callflowID string
	reqBody *ServiceCallflowAddEditData
}

// payload fields
func (r ApiV1AccountAccountIDCallflowCallflowIDPutRequest) ReqBody(reqBody ServiceCallflowAddEditData) ApiV1AccountAccountIDCallflowCallflowIDPutRequest {
	r.reqBody = &reqBody
	return r
}

func (r ApiV1AccountAccountIDCallflowCallflowIDPutRequest) Execute() (*ServiceDocsCallflowGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDCallflowCallflowIDPutExecute(r)
}

/*
V1AccountAccountIDCallflowCallflowIDPut Update Call Group

Update the details for a single callflow in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param callflowID Callflow ID, 32 alpha numeric
 @return ApiV1AccountAccountIDCallflowCallflowIDPutRequest
*/
func (a *CallflowAPIService) V1AccountAccountIDCallflowCallflowIDPut(ctx context.Context, accountID string, callflowID string) ApiV1AccountAccountIDCallflowCallflowIDPutRequest {
	return ApiV1AccountAccountIDCallflowCallflowIDPutRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		callflowID: callflowID,
	}
}

// Execute executes the request
//  @return ServiceDocsCallflowGetSingle
func (a *CallflowAPIService) V1AccountAccountIDCallflowCallflowIDPutExecute(r ApiV1AccountAccountIDCallflowCallflowIDPutRequest) (*ServiceDocsCallflowGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCallflowGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallflowAPIService.V1AccountAccountIDCallflowCallflowIDPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/callflow/{callflowID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"callflowID"+"}", url.PathEscape(parameterValueToString(r.callflowID, "callflowID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reqBody == nil {
		return localVarReturnValue, nil, reportError("reqBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.reqBody
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountIDCallflowGetRequest struct {
	ctx context.Context
	ApiService *CallflowAPIService
	accountID string
	startKey *string
	pageSize *int32
}

// start_key for pagination that was returned as next_start_key from your previous call
func (r ApiV1AccountAccountIDCallflowGetRequest) StartKey(startKey string) ApiV1AccountAccountIDCallflowGetRequest {
	r.startKey = &startKey
	return r
}

// number of records to return, range 1 to 50
func (r ApiV1AccountAccountIDCallflowGetRequest) PageSize(pageSize int32) ApiV1AccountAccountIDCallflowGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiV1AccountAccountIDCallflowGetRequest) Execute() (*ServiceDocsCallflowGetAll, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDCallflowGetExecute(r)
}

/*
V1AccountAccountIDCallflowGet Get Callflow List

Permit a user to view the callflow details in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @return ApiV1AccountAccountIDCallflowGetRequest
*/
func (a *CallflowAPIService) V1AccountAccountIDCallflowGet(ctx context.Context, accountID string) ApiV1AccountAccountIDCallflowGetRequest {
	return ApiV1AccountAccountIDCallflowGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
	}
}

// Execute executes the request
//  @return ServiceDocsCallflowGetAll
func (a *CallflowAPIService) V1AccountAccountIDCallflowGetExecute(r ApiV1AccountAccountIDCallflowGetRequest) (*ServiceDocsCallflowGetAll, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCallflowGetAll
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallflowAPIService.V1AccountAccountIDCallflowGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/callflow"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_key", r.startKey, "", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountIDCallflowPostRequest struct {
	ctx context.Context
	ApiService *CallflowAPIService
	accountID string
	request *ServiceCallflowAddEditData
}

// Call flow configuration
func (r ApiV1AccountAccountIDCallflowPostRequest) Request(request ServiceCallflowAddEditData) ApiV1AccountAccountIDCallflowPostRequest {
	r.request = &request
	return r
}

func (r ApiV1AccountAccountIDCallflowPostRequest) Execute() (*ServiceDocsCallflowGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDCallflowPostExecute(r)
}

/*
V1AccountAccountIDCallflowPost Create Call Group

Create instructions for routing a call to a user or system.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha-numeric
 @return ApiV1AccountAccountIDCallflowPostRequest
*/
func (a *CallflowAPIService) V1AccountAccountIDCallflowPost(ctx context.Context, accountID string) ApiV1AccountAccountIDCallflowPostRequest {
	return ApiV1AccountAccountIDCallflowPostRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
	}
}

// Execute executes the request
//  @return ServiceDocsCallflowGetSingle
func (a *CallflowAPIService) V1AccountAccountIDCallflowPostExecute(r ApiV1AccountAccountIDCallflowPostRequest) (*ServiceDocsCallflowGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCallflowGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallflowAPIService.V1AccountAccountIDCallflowPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/callflow"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, reportError("request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.request
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
