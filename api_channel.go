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


// ChannelAPIService ChannelAPI service
type ChannelAPIService service

type ApiV1AccountAccountIDChannelChannelIDGetRequest struct {
	ctx context.Context
	ApiService *ChannelAPIService
	accountID string
	channelID string
}

func (r ApiV1AccountAccountIDChannelChannelIDGetRequest) Execute() (*ServiceDocsChannelGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDChannelChannelIDGetExecute(r)
}

/*
V1AccountAccountIDChannelChannelIDGet Get Channel Details

Access details about each channel in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param channelID Channel ID
 @return ApiV1AccountAccountIDChannelChannelIDGetRequest
*/
func (a *ChannelAPIService) V1AccountAccountIDChannelChannelIDGet(ctx context.Context, accountID string, channelID string) ApiV1AccountAccountIDChannelChannelIDGetRequest {
	return ApiV1AccountAccountIDChannelChannelIDGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		channelID: channelID,
	}
}

// Execute executes the request
//  @return ServiceDocsChannelGetSingle
func (a *ChannelAPIService) V1AccountAccountIDChannelChannelIDGetExecute(r ApiV1AccountAccountIDChannelChannelIDGetRequest) (*ServiceDocsChannelGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsChannelGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelAPIService.V1AccountAccountIDChannelChannelIDGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/channel/{channelID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"channelID"+"}", url.PathEscape(parameterValueToString(r.channelID, "channelID")), -1)

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

type ApiV1AccountAccountIDChannelChannelIDPostRequest struct {
	ctx context.Context
	ApiService *ChannelAPIService
	accountID string
	channelID string
	reqBody *ServiceChannelRunActionData
}

// payload fields
func (r ApiV1AccountAccountIDChannelChannelIDPostRequest) ReqBody(reqBody ServiceChannelRunActionData) ApiV1AccountAccountIDChannelChannelIDPostRequest {
	r.reqBody = &reqBody
	return r
}

func (r ApiV1AccountAccountIDChannelChannelIDPostRequest) Execute() (*ServiceAPIResponse, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDChannelChannelIDPostExecute(r)
}

/*
V1AccountAccountIDChannelChannelIDPost Associate Action to Channel

Link an action, such as transfer or hangup to a channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param channelID Channel ID
 @return ApiV1AccountAccountIDChannelChannelIDPostRequest
*/
func (a *ChannelAPIService) V1AccountAccountIDChannelChannelIDPost(ctx context.Context, accountID string, channelID string) ApiV1AccountAccountIDChannelChannelIDPostRequest {
	return ApiV1AccountAccountIDChannelChannelIDPostRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		channelID: channelID,
	}
}

// Execute executes the request
//  @return ServiceAPIResponse
func (a *ChannelAPIService) V1AccountAccountIDChannelChannelIDPostExecute(r ApiV1AccountAccountIDChannelChannelIDPostRequest) (*ServiceAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelAPIService.V1AccountAccountIDChannelChannelIDPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/channel/{channelID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"channelID"+"}", url.PathEscape(parameterValueToString(r.channelID, "channelID")), -1)

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

type ApiV1AccountAccountIDChannelChannelIDPutRequest struct {
	ctx context.Context
	ApiService *ChannelAPIService
	accountID string
	channelID string
	reqBody *ServiceChannelRunMetaflowData
}

// payload fields
func (r ApiV1AccountAccountIDChannelChannelIDPutRequest) ReqBody(reqBody ServiceChannelRunMetaflowData) ApiV1AccountAccountIDChannelChannelIDPutRequest {
	r.reqBody = &reqBody
	return r
}

func (r ApiV1AccountAccountIDChannelChannelIDPutRequest) Execute() (*ServiceAPIResponse, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDChannelChannelIDPutExecute(r)
}

/*
V1AccountAccountIDChannelChannelIDPut Associate Metaflow to Channel

Link a metaflow to an active channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param channelID Channel ID
 @return ApiV1AccountAccountIDChannelChannelIDPutRequest
*/
func (a *ChannelAPIService) V1AccountAccountIDChannelChannelIDPut(ctx context.Context, accountID string, channelID string) ApiV1AccountAccountIDChannelChannelIDPutRequest {
	return ApiV1AccountAccountIDChannelChannelIDPutRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		channelID: channelID,
	}
}

// Execute executes the request
//  @return ServiceAPIResponse
func (a *ChannelAPIService) V1AccountAccountIDChannelChannelIDPutExecute(r ApiV1AccountAccountIDChannelChannelIDPutRequest) (*ServiceAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelAPIService.V1AccountAccountIDChannelChannelIDPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/channel/{channelID}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"channelID"+"}", url.PathEscape(parameterValueToString(r.channelID, "channelID")), -1)

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

type ApiV1AccountAccountIDChannelGetRequest struct {
	ctx context.Context
	ApiService *ChannelAPIService
	accountID string
}

func (r ApiV1AccountAccountIDChannelGetRequest) Execute() (*ServiceDocsChannelGet, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDChannelGetExecute(r)
}

/*
V1AccountAccountIDChannelGet Get Account Channel List

Get a list of active channels for an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @return ApiV1AccountAccountIDChannelGetRequest
*/
func (a *ChannelAPIService) V1AccountAccountIDChannelGet(ctx context.Context, accountID string) ApiV1AccountAccountIDChannelGetRequest {
	return ApiV1AccountAccountIDChannelGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
	}
}

// Execute executes the request
//  @return ServiceDocsChannelGet
func (a *ChannelAPIService) V1AccountAccountIDChannelGetExecute(r ApiV1AccountAccountIDChannelGetRequest) (*ServiceDocsChannelGet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsChannelGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelAPIService.V1AccountAccountIDChannelGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/channel"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)

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

type ApiV1AccountAccountIDDeviceDeviceIDChannelGetRequest struct {
	ctx context.Context
	ApiService *ChannelAPIService
	accountID string
	deviceID string
}

func (r ApiV1AccountAccountIDDeviceDeviceIDChannelGetRequest) Execute() (*ServiceDocsChannelGet, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDDeviceDeviceIDChannelGetExecute(r)
}

/*
V1AccountAccountIDDeviceDeviceIDChannelGet Get Device Channel List

Get the list of active channels for a device.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param deviceID Device ID, 32 alpha numeric
 @return ApiV1AccountAccountIDDeviceDeviceIDChannelGetRequest
*/
func (a *ChannelAPIService) V1AccountAccountIDDeviceDeviceIDChannelGet(ctx context.Context, accountID string, deviceID string) ApiV1AccountAccountIDDeviceDeviceIDChannelGetRequest {
	return ApiV1AccountAccountIDDeviceDeviceIDChannelGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		deviceID: deviceID,
	}
}

// Execute executes the request
//  @return ServiceDocsChannelGet
func (a *ChannelAPIService) V1AccountAccountIDDeviceDeviceIDChannelGetExecute(r ApiV1AccountAccountIDDeviceDeviceIDChannelGetRequest) (*ServiceDocsChannelGet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsChannelGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelAPIService.V1AccountAccountIDDeviceDeviceIDChannelGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/device/{deviceID}/channel"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceID"+"}", url.PathEscape(parameterValueToString(r.deviceID, "deviceID")), -1)

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

type ApiV1AccountAccountIDUserUserIDChannelGetRequest struct {
	ctx context.Context
	ApiService *ChannelAPIService
	accountID string
	userID string
}

func (r ApiV1AccountAccountIDUserUserIDChannelGetRequest) Execute() (*ServiceDocsChannelGet, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDUserUserIDChannelGetExecute(r)
}

/*
V1AccountAccountIDUserUserIDChannelGet Get User Channel List

Get the list of active channels for a user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param userID User ID, 32 alpha numeric
 @return ApiV1AccountAccountIDUserUserIDChannelGetRequest
*/
func (a *ChannelAPIService) V1AccountAccountIDUserUserIDChannelGet(ctx context.Context, accountID string, userID string) ApiV1AccountAccountIDUserUserIDChannelGetRequest {
	return ApiV1AccountAccountIDUserUserIDChannelGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		userID: userID,
	}
}

// Execute executes the request
//  @return ServiceDocsChannelGet
func (a *ChannelAPIService) V1AccountAccountIDUserUserIDChannelGetExecute(r ApiV1AccountAccountIDUserUserIDChannelGetRequest) (*ServiceDocsChannelGet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsChannelGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelAPIService.V1AccountAccountIDUserUserIDChannelGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/user/{userID}/channel"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)

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
