# ServiceVOIPDeviceAddEdit2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForward** | Pointer to [**ServiceCallForward**](ServiceCallForward.md) |  | [optional] 
**CallRecording** | Pointer to [**ServiceCallRecordingSettings**](ServiceCallRecordingSettings.md) |  | [optional] 
**CallerId** | Pointer to [**ServiceVOIPDeviceAddEdit3c**](ServiceVOIPDeviceAddEdit3c.md) |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**DoNotDisturb** | Pointer to [**ServiceVOIPSharedDoNotDisturb**](ServiceVOIPSharedDoNotDisturb.md) |  | [optional] 
**Enabled** | Pointer to **bool** | cannot use required, else it has to be true and false is not allowed | [optional] 
**MacAddress** | Pointer to **string** | dont use mac, it enforces :, which voip does not like | [optional] 
**MusicOnHold** | Pointer to [**ServiceMusicOnHold**](ServiceMusicOnHold.md) |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | Pointer to **string** | json omitempty is needed else voip fails on \&quot;\&quot; for owner_id | [optional] 
**Sip** | [**ServiceVOIPDeviceAddEdit3a**](ServiceVOIPDeviceAddEdit3a.md) |  | 

## Methods

### NewServiceVOIPDeviceAddEdit2

`func NewServiceVOIPDeviceAddEdit2(name string, sip ServiceVOIPDeviceAddEdit3a, ) *ServiceVOIPDeviceAddEdit2`

NewServiceVOIPDeviceAddEdit2 instantiates a new ServiceVOIPDeviceAddEdit2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPDeviceAddEdit2WithDefaults

`func NewServiceVOIPDeviceAddEdit2WithDefaults() *ServiceVOIPDeviceAddEdit2`

NewServiceVOIPDeviceAddEdit2WithDefaults instantiates a new ServiceVOIPDeviceAddEdit2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallForward

`func (o *ServiceVOIPDeviceAddEdit2) GetCallForward() ServiceCallForward`

GetCallForward returns the CallForward field if non-nil, zero value otherwise.

### GetCallForwardOk

`func (o *ServiceVOIPDeviceAddEdit2) GetCallForwardOk() (*ServiceCallForward, bool)`

GetCallForwardOk returns a tuple with the CallForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForward

`func (o *ServiceVOIPDeviceAddEdit2) SetCallForward(v ServiceCallForward)`

SetCallForward sets CallForward field to given value.

### HasCallForward

`func (o *ServiceVOIPDeviceAddEdit2) HasCallForward() bool`

HasCallForward returns a boolean if a field has been set.

### GetCallRecording

`func (o *ServiceVOIPDeviceAddEdit2) GetCallRecording() ServiceCallRecordingSettings`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ServiceVOIPDeviceAddEdit2) GetCallRecordingOk() (*ServiceCallRecordingSettings, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ServiceVOIPDeviceAddEdit2) SetCallRecording(v ServiceCallRecordingSettings)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ServiceVOIPDeviceAddEdit2) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetCallerId

`func (o *ServiceVOIPDeviceAddEdit2) GetCallerId() ServiceVOIPDeviceAddEdit3c`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *ServiceVOIPDeviceAddEdit2) GetCallerIdOk() (*ServiceVOIPDeviceAddEdit3c, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *ServiceVOIPDeviceAddEdit2) SetCallerId(v ServiceVOIPDeviceAddEdit3c)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *ServiceVOIPDeviceAddEdit2) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDeviceType

`func (o *ServiceVOIPDeviceAddEdit2) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ServiceVOIPDeviceAddEdit2) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ServiceVOIPDeviceAddEdit2) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ServiceVOIPDeviceAddEdit2) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceVOIPDeviceAddEdit2) GetDoNotDisturb() ServiceVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceVOIPDeviceAddEdit2) GetDoNotDisturbOk() (*ServiceVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceVOIPDeviceAddEdit2) SetDoNotDisturb(v ServiceVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceVOIPDeviceAddEdit2) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceVOIPDeviceAddEdit2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceVOIPDeviceAddEdit2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceVOIPDeviceAddEdit2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceVOIPDeviceAddEdit2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMacAddress

`func (o *ServiceVOIPDeviceAddEdit2) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ServiceVOIPDeviceAddEdit2) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ServiceVOIPDeviceAddEdit2) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ServiceVOIPDeviceAddEdit2) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ServiceVOIPDeviceAddEdit2) GetMusicOnHold() ServiceMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ServiceVOIPDeviceAddEdit2) GetMusicOnHoldOk() (*ServiceMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ServiceVOIPDeviceAddEdit2) SetMusicOnHold(v ServiceMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ServiceVOIPDeviceAddEdit2) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetName

`func (o *ServiceVOIPDeviceAddEdit2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPDeviceAddEdit2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPDeviceAddEdit2) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *ServiceVOIPDeviceAddEdit2) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceVOIPDeviceAddEdit2) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceVOIPDeviceAddEdit2) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceVOIPDeviceAddEdit2) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetSip

`func (o *ServiceVOIPDeviceAddEdit2) GetSip() ServiceVOIPDeviceAddEdit3a`

GetSip returns the Sip field if non-nil, zero value otherwise.

### GetSipOk

`func (o *ServiceVOIPDeviceAddEdit2) GetSipOk() (*ServiceVOIPDeviceAddEdit3a, bool)`

GetSipOk returns a tuple with the Sip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSip

`func (o *ServiceVOIPDeviceAddEdit2) SetSip(v ServiceVOIPDeviceAddEdit3a)`

SetSip sets Sip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


