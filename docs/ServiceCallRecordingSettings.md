# ServiceCallRecordingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Any** | Pointer to [**ServiceCallRecordingSource**](ServiceCallRecordingSource.md) |  | [optional] 
**Inbound** | Pointer to [**ServiceCallRecordingSource**](ServiceCallRecordingSource.md) |  | [optional] 
**Outbound** | Pointer to [**ServiceCallRecordingSource**](ServiceCallRecordingSource.md) |  | [optional] 

## Methods

### NewServiceCallRecordingSettings

`func NewServiceCallRecordingSettings() *ServiceCallRecordingSettings`

NewServiceCallRecordingSettings instantiates a new ServiceCallRecordingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallRecordingSettingsWithDefaults

`func NewServiceCallRecordingSettingsWithDefaults() *ServiceCallRecordingSettings`

NewServiceCallRecordingSettingsWithDefaults instantiates a new ServiceCallRecordingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAny

`func (o *ServiceCallRecordingSettings) GetAny() ServiceCallRecordingSource`

GetAny returns the Any field if non-nil, zero value otherwise.

### GetAnyOk

`func (o *ServiceCallRecordingSettings) GetAnyOk() (*ServiceCallRecordingSource, bool)`

GetAnyOk returns a tuple with the Any field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAny

`func (o *ServiceCallRecordingSettings) SetAny(v ServiceCallRecordingSource)`

SetAny sets Any field to given value.

### HasAny

`func (o *ServiceCallRecordingSettings) HasAny() bool`

HasAny returns a boolean if a field has been set.

### GetInbound

`func (o *ServiceCallRecordingSettings) GetInbound() ServiceCallRecordingSource`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *ServiceCallRecordingSettings) GetInboundOk() (*ServiceCallRecordingSource, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *ServiceCallRecordingSettings) SetInbound(v ServiceCallRecordingSource)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *ServiceCallRecordingSettings) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *ServiceCallRecordingSettings) GetOutbound() ServiceCallRecordingSource`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *ServiceCallRecordingSettings) GetOutboundOk() (*ServiceCallRecordingSource, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *ServiceCallRecordingSettings) SetOutbound(v ServiceCallRecordingSource)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *ServiceCallRecordingSettings) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


