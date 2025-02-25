# ServiceTTS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Voice** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceTTS

`func NewServiceTTS(text string, ) *ServiceTTS`

NewServiceTTS instantiates a new ServiceTTS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTTSWithDefaults

`func NewServiceTTSWithDefaults() *ServiceTTS`

NewServiceTTSWithDefaults instantiates a new ServiceTTS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ServiceTTS) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ServiceTTS) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ServiceTTS) SetText(v string)`

SetText sets Text field to given value.


### GetVoice

`func (o *ServiceTTS) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *ServiceTTS) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *ServiceTTS) SetVoice(v string)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *ServiceTTS) HasVoice() bool`

HasVoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


