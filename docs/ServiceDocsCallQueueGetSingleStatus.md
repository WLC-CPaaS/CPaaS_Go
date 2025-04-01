# ServiceDocsCallQueueGetSingleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceCallQueueStatusOutput**](ServiceCallQueueStatusOutput.md) |  | [optional] 
**RequestId** | Pointer to **string** | Unique id for each request | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code | [optional] 

## Methods

### NewServiceDocsCallQueueGetSingleStatus

`func NewServiceDocsCallQueueGetSingleStatus() *ServiceDocsCallQueueGetSingleStatus`

NewServiceDocsCallQueueGetSingleStatus instantiates a new ServiceDocsCallQueueGetSingleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsCallQueueGetSingleStatusWithDefaults

`func NewServiceDocsCallQueueGetSingleStatusWithDefaults() *ServiceDocsCallQueueGetSingleStatus`

NewServiceDocsCallQueueGetSingleStatusWithDefaults instantiates a new ServiceDocsCallQueueGetSingleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsCallQueueGetSingleStatus) GetData() ServiceCallQueueStatusOutput`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsCallQueueGetSingleStatus) GetDataOk() (*ServiceCallQueueStatusOutput, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsCallQueueGetSingleStatus) SetData(v ServiceCallQueueStatusOutput)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsCallQueueGetSingleStatus) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsCallQueueGetSingleStatus) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsCallQueueGetSingleStatus) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsCallQueueGetSingleStatus) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsCallQueueGetSingleStatus) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsCallQueueGetSingleStatus) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsCallQueueGetSingleStatus) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsCallQueueGetSingleStatus) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsCallQueueGetSingleStatus) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


