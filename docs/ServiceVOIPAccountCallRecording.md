# ServiceVOIPAccountCallRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**ServiceCallRecordingSettings**](ServiceCallRecordingSettings.md) |  | [optional] 
**Endpoint** | Pointer to [**ServiceCallRecordingSettings**](ServiceCallRecordingSettings.md) |  | [optional] 

## Methods

### NewServiceVOIPAccountCallRecording

`func NewServiceVOIPAccountCallRecording() *ServiceVOIPAccountCallRecording`

NewServiceVOIPAccountCallRecording instantiates a new ServiceVOIPAccountCallRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPAccountCallRecordingWithDefaults

`func NewServiceVOIPAccountCallRecordingWithDefaults() *ServiceVOIPAccountCallRecording`

NewServiceVOIPAccountCallRecordingWithDefaults instantiates a new ServiceVOIPAccountCallRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ServiceVOIPAccountCallRecording) GetAccount() ServiceCallRecordingSettings`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ServiceVOIPAccountCallRecording) GetAccountOk() (*ServiceCallRecordingSettings, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ServiceVOIPAccountCallRecording) SetAccount(v ServiceCallRecordingSettings)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ServiceVOIPAccountCallRecording) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEndpoint

`func (o *ServiceVOIPAccountCallRecording) GetEndpoint() ServiceCallRecordingSettings`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ServiceVOIPAccountCallRecording) GetEndpointOk() (*ServiceCallRecordingSettings, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ServiceVOIPAccountCallRecording) SetEndpoint(v ServiceCallRecordingSettings)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ServiceVOIPAccountCallRecording) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


