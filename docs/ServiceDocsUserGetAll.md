# ServiceDocsUserGetAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServiceUserOutputShort**](ServiceUserOutputShort.md) |  | [optional] 
**NextStartKey** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**StartKey** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceDocsUserGetAll

`func NewServiceDocsUserGetAll() *ServiceDocsUserGetAll`

NewServiceDocsUserGetAll instantiates a new ServiceDocsUserGetAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsUserGetAllWithDefaults

`func NewServiceDocsUserGetAllWithDefaults() *ServiceDocsUserGetAll`

NewServiceDocsUserGetAllWithDefaults instantiates a new ServiceDocsUserGetAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsUserGetAll) GetData() []ServiceUserOutputShort`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsUserGetAll) GetDataOk() (*[]ServiceUserOutputShort, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsUserGetAll) SetData(v []ServiceUserOutputShort)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsUserGetAll) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextStartKey

`func (o *ServiceDocsUserGetAll) GetNextStartKey() string`

GetNextStartKey returns the NextStartKey field if non-nil, zero value otherwise.

### GetNextStartKeyOk

`func (o *ServiceDocsUserGetAll) GetNextStartKeyOk() (*string, bool)`

GetNextStartKeyOk returns a tuple with the NextStartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartKey

`func (o *ServiceDocsUserGetAll) SetNextStartKey(v string)`

SetNextStartKey sets NextStartKey field to given value.

### HasNextStartKey

`func (o *ServiceDocsUserGetAll) HasNextStartKey() bool`

HasNextStartKey returns a boolean if a field has been set.

### GetPageSize

`func (o *ServiceDocsUserGetAll) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ServiceDocsUserGetAll) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ServiceDocsUserGetAll) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ServiceDocsUserGetAll) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsUserGetAll) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsUserGetAll) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsUserGetAll) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsUserGetAll) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStartKey

`func (o *ServiceDocsUserGetAll) GetStartKey() string`

GetStartKey returns the StartKey field if non-nil, zero value otherwise.

### GetStartKeyOk

`func (o *ServiceDocsUserGetAll) GetStartKeyOk() (*string, bool)`

GetStartKeyOk returns a tuple with the StartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartKey

`func (o *ServiceDocsUserGetAll) SetStartKey(v string)`

SetStartKey sets StartKey field to given value.

### HasStartKey

`func (o *ServiceDocsUserGetAll) HasStartKey() bool`

HasStartKey returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsUserGetAll) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsUserGetAll) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsUserGetAll) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsUserGetAll) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


