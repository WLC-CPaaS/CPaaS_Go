# ServiceDocsAccountAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceAPIKey**](ServiceAPIKey.md) |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceDocsAccountAPIKey

`func NewServiceDocsAccountAPIKey() *ServiceDocsAccountAPIKey`

NewServiceDocsAccountAPIKey instantiates a new ServiceDocsAccountAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsAccountAPIKeyWithDefaults

`func NewServiceDocsAccountAPIKeyWithDefaults() *ServiceDocsAccountAPIKey`

NewServiceDocsAccountAPIKeyWithDefaults instantiates a new ServiceDocsAccountAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsAccountAPIKey) GetData() ServiceAPIKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsAccountAPIKey) GetDataOk() (*ServiceAPIKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsAccountAPIKey) SetData(v ServiceAPIKey)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsAccountAPIKey) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsAccountAPIKey) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsAccountAPIKey) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsAccountAPIKey) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsAccountAPIKey) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsAccountAPIKey) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsAccountAPIKey) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsAccountAPIKey) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsAccountAPIKey) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


