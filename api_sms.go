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


// SMSAPIService SMSAPI service
type SMSAPIService service

type ApiV1SmsAccountAccountIDCampaignCampaignIDImportGetRequest struct {
	ctx context.Context
	ApiService *SMSAPIService
	accountID string
	campaignID string
}

func (r ApiV1SmsAccountAccountIDCampaignCampaignIDImportGetRequest) Execute() (*ServiceDocsCampaignImportOutput, *http.Response, error) {
	return r.ApiService.V1SmsAccountAccountIDCampaignCampaignIDImportGetExecute(r)
}

/*
V1SmsAccountAccountIDCampaignCampaignIDImportGet Method for V1SmsAccountAccountIDCampaignCampaignIDImportGet

Get details about a single imported campaign in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param campaignID Campaign ID
 @return ApiV1SmsAccountAccountIDCampaignCampaignIDImportGetRequest
*/
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignCampaignIDImportGet(ctx context.Context, accountID string, campaignID string) ApiV1SmsAccountAccountIDCampaignCampaignIDImportGetRequest {
	return ApiV1SmsAccountAccountIDCampaignCampaignIDImportGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		campaignID: campaignID,
	}
}

// Execute executes the request
//  @return ServiceDocsCampaignImportOutput
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignCampaignIDImportGetExecute(r ApiV1SmsAccountAccountIDCampaignCampaignIDImportGetRequest) (*ServiceDocsCampaignImportOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCampaignImportOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSAPIService.V1SmsAccountAccountIDCampaignCampaignIDImportGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sms/account/{accountID}/campaign/{campaignID}/import"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"campaignID"+"}", url.PathEscape(parameterValueToString(r.campaignID, "campaignID")), -1)

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

type ApiV1SmsAccountAccountIDCampaignCampaignIDImportPostRequest struct {
	ctx context.Context
	ApiService *SMSAPIService
	accountID string
	campaignID string
}

func (r ApiV1SmsAccountAccountIDCampaignCampaignIDImportPostRequest) Execute() (*ServiceDocsCampaignImportOutput, *http.Response, error) {
	return r.ApiService.V1SmsAccountAccountIDCampaignCampaignIDImportPostExecute(r)
}

/*
V1SmsAccountAccountIDCampaignCampaignIDImportPost Method for V1SmsAccountAccountIDCampaignCampaignIDImportPost

Import campaign

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param campaignID Campaign ID
 @return ApiV1SmsAccountAccountIDCampaignCampaignIDImportPostRequest
*/
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignCampaignIDImportPost(ctx context.Context, accountID string, campaignID string) ApiV1SmsAccountAccountIDCampaignCampaignIDImportPostRequest {
	return ApiV1SmsAccountAccountIDCampaignCampaignIDImportPostRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		campaignID: campaignID,
	}
}

// Execute executes the request
//  @return ServiceDocsCampaignImportOutput
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignCampaignIDImportPostExecute(r ApiV1SmsAccountAccountIDCampaignCampaignIDImportPostRequest) (*ServiceDocsCampaignImportOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCampaignImportOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSAPIService.V1SmsAccountAccountIDCampaignCampaignIDImportPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sms/account/{accountID}/campaign/{campaignID}/import"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"campaignID"+"}", url.PathEscape(parameterValueToString(r.campaignID, "campaignID")), -1)

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

type ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest struct {
	ctx context.Context
	ApiService *SMSAPIService
	accountID string
	campaignID string
	pageNum *float32
	pageSize *float32
}

// Page number
func (r ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest) PageNum(pageNum float32) ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest {
	r.pageNum = &pageNum
	return r
}

// Page size
func (r ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest) PageSize(pageSize float32) ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest) Execute() (*ServiceDocsCampaignPhoneNumberOutput, *http.Response, error) {
	return r.ApiService.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetExecute(r)
}

/*
V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet Method for V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet

Get telephone numbers associated with a campaign.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param campaignID Campaign ID
 @return ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest
*/
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet(ctx context.Context, accountID string, campaignID string) ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest {
	return ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		campaignID: campaignID,
	}
}

// Execute executes the request
//  @return ServiceDocsCampaignPhoneNumberOutput
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetExecute(r ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberGetRequest) (*ServiceDocsCampaignPhoneNumberOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCampaignPhoneNumberOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSAPIService.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sms/account/{accountID}/campaign/{campaignID}/phonenumber"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"campaignID"+"}", url.PathEscape(parameterValueToString(r.campaignID, "campaignID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_num", r.pageNum, "", "")
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

type ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest struct {
	ctx context.Context
	ApiService *SMSAPIService
	accountID string
	campaignID string
	reqBody *ServiceCampaignTagDetagPhonenumbers
}

// payload fields
func (r ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest) ReqBody(reqBody ServiceCampaignTagDetagPhonenumbers) ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest {
	r.reqBody = &reqBody
	return r
}

func (r ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest) Execute() (*ServiceDocsCampaignTagDetagPhonenumbersOutput, *http.Response, error) {
	return r.ApiService.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutExecute(r)
}

/*
V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut Method for V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut

Associate or dissociate telephone numbers with a campaign.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param campaignID Campaign ID, 32 alpha numeric
 @return ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest
*/
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut(ctx context.Context, accountID string, campaignID string) ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest {
	return ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		campaignID: campaignID,
	}
}

// Execute executes the request
//  @return ServiceDocsCampaignTagDetagPhonenumbersOutput
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutExecute(r ApiV1SmsAccountAccountIDCampaignCampaignIDPhonenumberPutRequest) (*ServiceDocsCampaignTagDetagPhonenumbersOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCampaignTagDetagPhonenumbersOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSAPIService.V1SmsAccountAccountIDCampaignCampaignIDPhonenumberPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sms/account/{accountID}/campaign/{campaignID}/phonenumber"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"campaignID"+"}", url.PathEscape(parameterValueToString(r.campaignID, "campaignID")), -1)

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

type ApiV1SmsAccountAccountIDCampaignImportGetRequest struct {
	ctx context.Context
	ApiService *SMSAPIService
	accountID string
	pageNum *float32
	pageSize *float32
}

// Page number
func (r ApiV1SmsAccountAccountIDCampaignImportGetRequest) PageNum(pageNum float32) ApiV1SmsAccountAccountIDCampaignImportGetRequest {
	r.pageNum = &pageNum
	return r
}

// Page size
func (r ApiV1SmsAccountAccountIDCampaignImportGetRequest) PageSize(pageSize float32) ApiV1SmsAccountAccountIDCampaignImportGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiV1SmsAccountAccountIDCampaignImportGetRequest) Execute() (*ServiceDocsCampaignImportedGetAllOutput, *http.Response, error) {
	return r.ApiService.V1SmsAccountAccountIDCampaignImportGetExecute(r)
}

/*
V1SmsAccountAccountIDCampaignImportGet Method for V1SmsAccountAccountIDCampaignImportGet

Get a list of all imported campaigns in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @return ApiV1SmsAccountAccountIDCampaignImportGetRequest
*/
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignImportGet(ctx context.Context, accountID string) ApiV1SmsAccountAccountIDCampaignImportGetRequest {
	return ApiV1SmsAccountAccountIDCampaignImportGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
	}
}

// Execute executes the request
//  @return ServiceDocsCampaignImportedGetAllOutput
func (a *SMSAPIService) V1SmsAccountAccountIDCampaignImportGetExecute(r ApiV1SmsAccountAccountIDCampaignImportGetRequest) (*ServiceDocsCampaignImportedGetAllOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsCampaignImportedGetAllOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSAPIService.V1SmsAccountAccountIDCampaignImportGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sms/account/{accountID}/campaign/import"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_num", r.pageNum, "", "")
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
