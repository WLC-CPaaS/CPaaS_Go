# ServiceDocsStorageGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceStorageOutput**](ServiceStorageOutput.md) |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceDocsStorageGet

`func NewServiceDocsStorageGet() *ServiceDocsStorageGet`

NewServiceDocsStorageGet instantiates a new ServiceDocsStorageGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsStorageGetWithDefaults

`func NewServiceDocsStorageGetWithDefaults() *ServiceDocsStorageGet`

NewServiceDocsStorageGetWithDefaults instantiates a new ServiceDocsStorageGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsStorageGet) GetData() ServiceStorageOutput`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsStorageGet) GetDataOk() (*ServiceStorageOutput, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsStorageGet) SetData(v ServiceStorageOutput)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsStorageGet) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsStorageGet) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsStorageGet) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsStorageGet) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsStorageGet) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsStorageGet) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsStorageGet) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsStorageGet) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsStorageGet) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


