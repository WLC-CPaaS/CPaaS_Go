# ServiceVOIPMenuAddEditData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Media** | Pointer to [**ServiceMenuOutputFullMedia**](ServiceMenuOutputFullMedia.md) |  | [optional] 
**Name** | **string** |  | 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceVOIPMenuAddEditData

`func NewServiceVOIPMenuAddEditData(name string, ) *ServiceVOIPMenuAddEditData`

NewServiceVOIPMenuAddEditData instantiates a new ServiceVOIPMenuAddEditData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPMenuAddEditDataWithDefaults

`func NewServiceVOIPMenuAddEditDataWithDefaults() *ServiceVOIPMenuAddEditData`

NewServiceVOIPMenuAddEditDataWithDefaults instantiates a new ServiceVOIPMenuAddEditData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedia

`func (o *ServiceVOIPMenuAddEditData) GetMedia() ServiceMenuOutputFullMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ServiceVOIPMenuAddEditData) GetMediaOk() (*ServiceMenuOutputFullMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ServiceVOIPMenuAddEditData) SetMedia(v ServiceMenuOutputFullMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ServiceVOIPMenuAddEditData) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetName

`func (o *ServiceVOIPMenuAddEditData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPMenuAddEditData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPMenuAddEditData) SetName(v string)`

SetName sets Name field to given value.


### GetTimeout

`func (o *ServiceVOIPMenuAddEditData) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ServiceVOIPMenuAddEditData) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ServiceVOIPMenuAddEditData) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ServiceVOIPMenuAddEditData) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


