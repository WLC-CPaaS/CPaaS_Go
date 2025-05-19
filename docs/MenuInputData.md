# MenuInputData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Media** | Pointer to [**MenuOutputDetailMedia**](MenuOutputDetailMedia.md) | The media (prompt) parameters | [optional] 
**Name** | **string** | A friendly name for the menu | 
**Timeout** | Pointer to **int32** | The amount of time (in milliseconds) to wait for the caller to begin entering digits | [optional] 

## Methods

### NewMenuInputData

`func NewMenuInputData(name string, ) *MenuInputData`

NewMenuInputData instantiates a new MenuInputData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuInputDataWithDefaults

`func NewMenuInputDataWithDefaults() *MenuInputData`

NewMenuInputDataWithDefaults instantiates a new MenuInputData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedia

`func (o *MenuInputData) GetMedia() MenuOutputDetailMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MenuInputData) GetMediaOk() (*MenuOutputDetailMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MenuInputData) SetMedia(v MenuOutputDetailMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *MenuInputData) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetName

`func (o *MenuInputData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuInputData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuInputData) SetName(v string)`

SetName sets Name field to given value.


### GetTimeout

`func (o *MenuInputData) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MenuInputData) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MenuInputData) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MenuInputData) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


