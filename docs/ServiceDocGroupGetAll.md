# ServiceDocGroupGetAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceGroupOutputShort**](ServiceGroupOutputShort.md) |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceDocGroupGetAll

`func NewServiceDocGroupGetAll() *ServiceDocGroupGetAll`

NewServiceDocGroupGetAll instantiates a new ServiceDocGroupGetAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocGroupGetAllWithDefaults

`func NewServiceDocGroupGetAllWithDefaults() *ServiceDocGroupGetAll`

NewServiceDocGroupGetAllWithDefaults instantiates a new ServiceDocGroupGetAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocGroupGetAll) GetData() ServiceGroupOutputShort`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocGroupGetAll) GetDataOk() (*ServiceGroupOutputShort, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocGroupGetAll) SetData(v ServiceGroupOutputShort)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocGroupGetAll) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocGroupGetAll) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocGroupGetAll) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocGroupGetAll) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocGroupGetAll) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocGroupGetAll) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocGroupGetAll) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocGroupGetAll) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocGroupGetAll) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


