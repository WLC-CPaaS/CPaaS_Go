# ServiceMenuOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**ServiceMenuOutputFullMedia**](ServiceMenuOutputFullMedia.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceMenuOutputFull

`func NewServiceMenuOutputFull() *ServiceMenuOutputFull`

NewServiceMenuOutputFull instantiates a new ServiceMenuOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMenuOutputFullWithDefaults

`func NewServiceMenuOutputFullWithDefaults() *ServiceMenuOutputFull`

NewServiceMenuOutputFullWithDefaults instantiates a new ServiceMenuOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceMenuOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceMenuOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceMenuOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceMenuOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMedia

`func (o *ServiceMenuOutputFull) GetMedia() ServiceMenuOutputFullMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ServiceMenuOutputFull) GetMediaOk() (*ServiceMenuOutputFullMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ServiceMenuOutputFull) SetMedia(v ServiceMenuOutputFullMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ServiceMenuOutputFull) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetName

`func (o *ServiceMenuOutputFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceMenuOutputFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceMenuOutputFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceMenuOutputFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeout

`func (o *ServiceMenuOutputFull) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ServiceMenuOutputFull) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ServiceMenuOutputFull) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ServiceMenuOutputFull) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


