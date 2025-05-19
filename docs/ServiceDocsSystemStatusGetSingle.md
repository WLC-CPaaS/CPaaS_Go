# ServiceDocsSystemStatusGetSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceSystemStatusOutput**](ServiceSystemStatusOutput.md) |  | [optional] 
**RequestId** | Pointer to **string** | Unique id for each request | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code | [optional] 
**SystemStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceDocsSystemStatusGetSingle

`func NewServiceDocsSystemStatusGetSingle() *ServiceDocsSystemStatusGetSingle`

NewServiceDocsSystemStatusGetSingle instantiates a new ServiceDocsSystemStatusGetSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsSystemStatusGetSingleWithDefaults

`func NewServiceDocsSystemStatusGetSingleWithDefaults() *ServiceDocsSystemStatusGetSingle`

NewServiceDocsSystemStatusGetSingleWithDefaults instantiates a new ServiceDocsSystemStatusGetSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsSystemStatusGetSingle) GetData() ServiceSystemStatusOutput`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsSystemStatusGetSingle) GetDataOk() (*ServiceSystemStatusOutput, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsSystemStatusGetSingle) SetData(v ServiceSystemStatusOutput)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsSystemStatusGetSingle) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsSystemStatusGetSingle) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsSystemStatusGetSingle) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsSystemStatusGetSingle) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsSystemStatusGetSingle) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsSystemStatusGetSingle) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsSystemStatusGetSingle) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsSystemStatusGetSingle) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsSystemStatusGetSingle) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetSystemStatus

`func (o *ServiceDocsSystemStatusGetSingle) GetSystemStatus() string`

GetSystemStatus returns the SystemStatus field if non-nil, zero value otherwise.

### GetSystemStatusOk

`func (o *ServiceDocsSystemStatusGetSingle) GetSystemStatusOk() (*string, bool)`

GetSystemStatusOk returns a tuple with the SystemStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemStatus

`func (o *ServiceDocsSystemStatusGetSingle) SetSystemStatus(v string)`

SetSystemStatus sets SystemStatus field to given value.

### HasSystemStatus

`func (o *ServiceDocsSystemStatusGetSingle) HasSystemStatus() bool`

HasSystemStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


