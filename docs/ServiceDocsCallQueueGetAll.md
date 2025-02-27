# ServiceDocsCallQueueGetAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServiceCallQueueOutputShort**](ServiceCallQueueOutputShort.md) |  | [optional] 
**NextStartKey** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**StartKey** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceDocsCallQueueGetAll

`func NewServiceDocsCallQueueGetAll() *ServiceDocsCallQueueGetAll`

NewServiceDocsCallQueueGetAll instantiates a new ServiceDocsCallQueueGetAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsCallQueueGetAllWithDefaults

`func NewServiceDocsCallQueueGetAllWithDefaults() *ServiceDocsCallQueueGetAll`

NewServiceDocsCallQueueGetAllWithDefaults instantiates a new ServiceDocsCallQueueGetAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsCallQueueGetAll) GetData() []ServiceCallQueueOutputShort`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsCallQueueGetAll) GetDataOk() (*[]ServiceCallQueueOutputShort, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsCallQueueGetAll) SetData(v []ServiceCallQueueOutputShort)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsCallQueueGetAll) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextStartKey

`func (o *ServiceDocsCallQueueGetAll) GetNextStartKey() string`

GetNextStartKey returns the NextStartKey field if non-nil, zero value otherwise.

### GetNextStartKeyOk

`func (o *ServiceDocsCallQueueGetAll) GetNextStartKeyOk() (*string, bool)`

GetNextStartKeyOk returns a tuple with the NextStartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartKey

`func (o *ServiceDocsCallQueueGetAll) SetNextStartKey(v string)`

SetNextStartKey sets NextStartKey field to given value.

### HasNextStartKey

`func (o *ServiceDocsCallQueueGetAll) HasNextStartKey() bool`

HasNextStartKey returns a boolean if a field has been set.

### GetPageSize

`func (o *ServiceDocsCallQueueGetAll) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ServiceDocsCallQueueGetAll) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ServiceDocsCallQueueGetAll) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ServiceDocsCallQueueGetAll) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsCallQueueGetAll) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsCallQueueGetAll) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsCallQueueGetAll) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsCallQueueGetAll) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStartKey

`func (o *ServiceDocsCallQueueGetAll) GetStartKey() string`

GetStartKey returns the StartKey field if non-nil, zero value otherwise.

### GetStartKeyOk

`func (o *ServiceDocsCallQueueGetAll) GetStartKeyOk() (*string, bool)`

GetStartKeyOk returns a tuple with the StartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartKey

`func (o *ServiceDocsCallQueueGetAll) SetStartKey(v string)`

SetStartKey sets StartKey field to given value.

### HasStartKey

`func (o *ServiceDocsCallQueueGetAll) HasStartKey() bool`

HasStartKey returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsCallQueueGetAll) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsCallQueueGetAll) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsCallQueueGetAll) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsCallQueueGetAll) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


