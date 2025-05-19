# ServiceDocsTemporalRuleSetGetAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServiceTemporalRuleSetOutputShort**](ServiceTemporalRuleSetOutputShort.md) |  | [optional] 
**NextStartKey** | Pointer to **string** | List Pagination: Used to get the next page of results. Will not exist if this is the last page. | [optional] 
**PageSize** | Pointer to **int32** | List Pagination: The number of results returned in this page | [optional] 
**RequestId** | Pointer to **string** | Unique id for each request | [optional] 
**StartKey** | Pointer to **string** | List Pagination: Code for paged results | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code | [optional] 

## Methods

### NewServiceDocsTemporalRuleSetGetAll

`func NewServiceDocsTemporalRuleSetGetAll() *ServiceDocsTemporalRuleSetGetAll`

NewServiceDocsTemporalRuleSetGetAll instantiates a new ServiceDocsTemporalRuleSetGetAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsTemporalRuleSetGetAllWithDefaults

`func NewServiceDocsTemporalRuleSetGetAllWithDefaults() *ServiceDocsTemporalRuleSetGetAll`

NewServiceDocsTemporalRuleSetGetAllWithDefaults instantiates a new ServiceDocsTemporalRuleSetGetAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsTemporalRuleSetGetAll) GetData() []ServiceTemporalRuleSetOutputShort`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsTemporalRuleSetGetAll) GetDataOk() (*[]ServiceTemporalRuleSetOutputShort, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsTemporalRuleSetGetAll) SetData(v []ServiceTemporalRuleSetOutputShort)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsTemporalRuleSetGetAll) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextStartKey

`func (o *ServiceDocsTemporalRuleSetGetAll) GetNextStartKey() string`

GetNextStartKey returns the NextStartKey field if non-nil, zero value otherwise.

### GetNextStartKeyOk

`func (o *ServiceDocsTemporalRuleSetGetAll) GetNextStartKeyOk() (*string, bool)`

GetNextStartKeyOk returns a tuple with the NextStartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartKey

`func (o *ServiceDocsTemporalRuleSetGetAll) SetNextStartKey(v string)`

SetNextStartKey sets NextStartKey field to given value.

### HasNextStartKey

`func (o *ServiceDocsTemporalRuleSetGetAll) HasNextStartKey() bool`

HasNextStartKey returns a boolean if a field has been set.

### GetPageSize

`func (o *ServiceDocsTemporalRuleSetGetAll) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ServiceDocsTemporalRuleSetGetAll) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ServiceDocsTemporalRuleSetGetAll) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ServiceDocsTemporalRuleSetGetAll) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsTemporalRuleSetGetAll) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsTemporalRuleSetGetAll) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsTemporalRuleSetGetAll) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsTemporalRuleSetGetAll) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStartKey

`func (o *ServiceDocsTemporalRuleSetGetAll) GetStartKey() string`

GetStartKey returns the StartKey field if non-nil, zero value otherwise.

### GetStartKeyOk

`func (o *ServiceDocsTemporalRuleSetGetAll) GetStartKeyOk() (*string, bool)`

GetStartKeyOk returns a tuple with the StartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartKey

`func (o *ServiceDocsTemporalRuleSetGetAll) SetStartKey(v string)`

SetStartKey sets StartKey field to given value.

### HasStartKey

`func (o *ServiceDocsTemporalRuleSetGetAll) HasStartKey() bool`

HasStartKey returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsTemporalRuleSetGetAll) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsTemporalRuleSetGetAll) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsTemporalRuleSetGetAll) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsTemporalRuleSetGetAll) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


