# ServiceVOIPMediaAddEditData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**MediaSource** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Tts** | Pointer to [**ServiceTTS**](ServiceTTS.md) |  | [optional] 

## Methods

### NewServiceVOIPMediaAddEditData

`func NewServiceVOIPMediaAddEditData(name string, ) *ServiceVOIPMediaAddEditData`

NewServiceVOIPMediaAddEditData instantiates a new ServiceVOIPMediaAddEditData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPMediaAddEditDataWithDefaults

`func NewServiceVOIPMediaAddEditDataWithDefaults() *ServiceVOIPMediaAddEditData`

NewServiceVOIPMediaAddEditDataWithDefaults instantiates a new ServiceVOIPMediaAddEditData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceVOIPMediaAddEditData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceVOIPMediaAddEditData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceVOIPMediaAddEditData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceVOIPMediaAddEditData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMediaSource

`func (o *ServiceVOIPMediaAddEditData) GetMediaSource() string`

GetMediaSource returns the MediaSource field if non-nil, zero value otherwise.

### GetMediaSourceOk

`func (o *ServiceVOIPMediaAddEditData) GetMediaSourceOk() (*string, bool)`

GetMediaSourceOk returns a tuple with the MediaSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSource

`func (o *ServiceVOIPMediaAddEditData) SetMediaSource(v string)`

SetMediaSource sets MediaSource field to given value.

### HasMediaSource

`func (o *ServiceVOIPMediaAddEditData) HasMediaSource() bool`

HasMediaSource returns a boolean if a field has been set.

### GetName

`func (o *ServiceVOIPMediaAddEditData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPMediaAddEditData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPMediaAddEditData) SetName(v string)`

SetName sets Name field to given value.


### GetTts

`func (o *ServiceVOIPMediaAddEditData) GetTts() ServiceTTS`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *ServiceVOIPMediaAddEditData) GetTtsOk() (*ServiceTTS, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *ServiceVOIPMediaAddEditData) SetTts(v ServiceTTS)`

SetTts sets Tts field to given value.

### HasTts

`func (o *ServiceVOIPMediaAddEditData) HasTts() bool`

HasTts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


