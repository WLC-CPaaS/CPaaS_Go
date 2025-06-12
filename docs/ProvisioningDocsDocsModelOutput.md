# ProvisioningDocsDocsModelOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ModelsModel**](ModelsModel.md) |  | [optional] 
**NextStartKey** | Pointer to **string** | List Pagination: Used to get the next page of results. Will not exist if this is the last page. | [optional] 
**PageSize** | Pointer to **int32** | List Pagination: The number of results returned in this page | [optional] 
**RequestId** | Pointer to **string** | Unique id for each request | [optional] 
**StartKey** | Pointer to **string** | List Pagination: Code for paged results | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code | [optional] 

## Methods

### NewProvisioningDocsDocsModelOutput

`func NewProvisioningDocsDocsModelOutput() *ProvisioningDocsDocsModelOutput`

NewProvisioningDocsDocsModelOutput instantiates a new ProvisioningDocsDocsModelOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningDocsDocsModelOutputWithDefaults

`func NewProvisioningDocsDocsModelOutputWithDefaults() *ProvisioningDocsDocsModelOutput`

NewProvisioningDocsDocsModelOutputWithDefaults instantiates a new ProvisioningDocsDocsModelOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProvisioningDocsDocsModelOutput) GetData() []ModelsModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProvisioningDocsDocsModelOutput) GetDataOk() (*[]ModelsModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProvisioningDocsDocsModelOutput) SetData(v []ModelsModel)`

SetData sets Data field to given value.

### HasData

`func (o *ProvisioningDocsDocsModelOutput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextStartKey

`func (o *ProvisioningDocsDocsModelOutput) GetNextStartKey() string`

GetNextStartKey returns the NextStartKey field if non-nil, zero value otherwise.

### GetNextStartKeyOk

`func (o *ProvisioningDocsDocsModelOutput) GetNextStartKeyOk() (*string, bool)`

GetNextStartKeyOk returns a tuple with the NextStartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartKey

`func (o *ProvisioningDocsDocsModelOutput) SetNextStartKey(v string)`

SetNextStartKey sets NextStartKey field to given value.

### HasNextStartKey

`func (o *ProvisioningDocsDocsModelOutput) HasNextStartKey() bool`

HasNextStartKey returns a boolean if a field has been set.

### GetPageSize

`func (o *ProvisioningDocsDocsModelOutput) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ProvisioningDocsDocsModelOutput) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ProvisioningDocsDocsModelOutput) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ProvisioningDocsDocsModelOutput) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetRequestId

`func (o *ProvisioningDocsDocsModelOutput) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProvisioningDocsDocsModelOutput) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProvisioningDocsDocsModelOutput) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ProvisioningDocsDocsModelOutput) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStartKey

`func (o *ProvisioningDocsDocsModelOutput) GetStartKey() string`

GetStartKey returns the StartKey field if non-nil, zero value otherwise.

### GetStartKeyOk

`func (o *ProvisioningDocsDocsModelOutput) GetStartKeyOk() (*string, bool)`

GetStartKeyOk returns a tuple with the StartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartKey

`func (o *ProvisioningDocsDocsModelOutput) SetStartKey(v string)`

SetStartKey sets StartKey field to given value.

### HasStartKey

`func (o *ProvisioningDocsDocsModelOutput) HasStartKey() bool`

HasStartKey returns a boolean if a field has been set.

### GetStatusCode

`func (o *ProvisioningDocsDocsModelOutput) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ProvisioningDocsDocsModelOutput) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ProvisioningDocsDocsModelOutput) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ProvisioningDocsDocsModelOutput) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


