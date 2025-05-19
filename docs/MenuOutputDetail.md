# MenuOutputDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MenuOutputDetailData**](MenuOutputDetailData.md) | Data Payload | [optional] 
**RequestId** | Pointer to **string** | Unique id for each request | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code | [optional] 

## Methods

### NewMenuOutputDetail

`func NewMenuOutputDetail() *MenuOutputDetail`

NewMenuOutputDetail instantiates a new MenuOutputDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuOutputDetailWithDefaults

`func NewMenuOutputDetailWithDefaults() *MenuOutputDetail`

NewMenuOutputDetailWithDefaults instantiates a new MenuOutputDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MenuOutputDetail) GetData() MenuOutputDetailData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MenuOutputDetail) GetDataOk() (*MenuOutputDetailData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MenuOutputDetail) SetData(v MenuOutputDetailData)`

SetData sets Data field to given value.

### HasData

`func (o *MenuOutputDetail) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *MenuOutputDetail) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *MenuOutputDetail) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *MenuOutputDetail) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *MenuOutputDetail) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *MenuOutputDetail) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *MenuOutputDetail) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *MenuOutputDetail) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *MenuOutputDetail) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


