# ServiceDeviceOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForward** | Pointer to [**ServiceCallForward**](ServiceCallForward.md) |  | [optional] 
**CallRecording** | Pointer to [**ServiceCallRecordingSettings**](ServiceCallRecordingSettings.md) |  | [optional] 
**CallerId** | Pointer to [**ServiceDeviceOutputFullCallerid**](ServiceDeviceOutputFullCallerid.md) |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**DoNotDisturb** | Pointer to [**ServiceVOIPSharedDoNotDisturb**](ServiceVOIPSharedDoNotDisturb.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**ServiceDeviceOutputFullMedia**](ServiceDeviceOutputFullMedia.md) |  | [optional] 
**MusicOnHold** | Pointer to [**ServiceMusicOnHold**](ServiceMusicOnHold.md) | Provision  *DeviceOutputFullProvision &#x60;json:\&quot;provision\&quot;&#x60; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Sip** | Pointer to [**ServiceDeviceOutputFullSIP**](ServiceDeviceOutputFullSIP.md) |  | [optional] 

## Methods

### NewServiceDeviceOutputFull

`func NewServiceDeviceOutputFull() *ServiceDeviceOutputFull`

NewServiceDeviceOutputFull instantiates a new ServiceDeviceOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeviceOutputFullWithDefaults

`func NewServiceDeviceOutputFullWithDefaults() *ServiceDeviceOutputFull`

NewServiceDeviceOutputFullWithDefaults instantiates a new ServiceDeviceOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallForward

`func (o *ServiceDeviceOutputFull) GetCallForward() ServiceCallForward`

GetCallForward returns the CallForward field if non-nil, zero value otherwise.

### GetCallForwardOk

`func (o *ServiceDeviceOutputFull) GetCallForwardOk() (*ServiceCallForward, bool)`

GetCallForwardOk returns a tuple with the CallForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForward

`func (o *ServiceDeviceOutputFull) SetCallForward(v ServiceCallForward)`

SetCallForward sets CallForward field to given value.

### HasCallForward

`func (o *ServiceDeviceOutputFull) HasCallForward() bool`

HasCallForward returns a boolean if a field has been set.

### GetCallRecording

`func (o *ServiceDeviceOutputFull) GetCallRecording() ServiceCallRecordingSettings`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ServiceDeviceOutputFull) GetCallRecordingOk() (*ServiceCallRecordingSettings, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ServiceDeviceOutputFull) SetCallRecording(v ServiceCallRecordingSettings)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ServiceDeviceOutputFull) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetCallerId

`func (o *ServiceDeviceOutputFull) GetCallerId() ServiceDeviceOutputFullCallerid`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *ServiceDeviceOutputFull) GetCallerIdOk() (*ServiceDeviceOutputFullCallerid, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *ServiceDeviceOutputFull) SetCallerId(v ServiceDeviceOutputFullCallerid)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *ServiceDeviceOutputFull) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDeviceType

`func (o *ServiceDeviceOutputFull) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ServiceDeviceOutputFull) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ServiceDeviceOutputFull) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ServiceDeviceOutputFull) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceDeviceOutputFull) GetDoNotDisturb() ServiceVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceDeviceOutputFull) GetDoNotDisturbOk() (*ServiceVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceDeviceOutputFull) SetDoNotDisturb(v ServiceVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceDeviceOutputFull) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceDeviceOutputFull) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceDeviceOutputFull) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceDeviceOutputFull) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceDeviceOutputFull) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ServiceDeviceOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceDeviceOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceDeviceOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceDeviceOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMacAddress

`func (o *ServiceDeviceOutputFull) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ServiceDeviceOutputFull) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ServiceDeviceOutputFull) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ServiceDeviceOutputFull) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMedia

`func (o *ServiceDeviceOutputFull) GetMedia() ServiceDeviceOutputFullMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ServiceDeviceOutputFull) GetMediaOk() (*ServiceDeviceOutputFullMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ServiceDeviceOutputFull) SetMedia(v ServiceDeviceOutputFullMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ServiceDeviceOutputFull) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ServiceDeviceOutputFull) GetMusicOnHold() ServiceMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ServiceDeviceOutputFull) GetMusicOnHoldOk() (*ServiceMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ServiceDeviceOutputFull) SetMusicOnHold(v ServiceMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ServiceDeviceOutputFull) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetName

`func (o *ServiceDeviceOutputFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDeviceOutputFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDeviceOutputFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceDeviceOutputFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *ServiceDeviceOutputFull) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceDeviceOutputFull) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceDeviceOutputFull) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceDeviceOutputFull) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetSip

`func (o *ServiceDeviceOutputFull) GetSip() ServiceDeviceOutputFullSIP`

GetSip returns the Sip field if non-nil, zero value otherwise.

### GetSipOk

`func (o *ServiceDeviceOutputFull) GetSipOk() (*ServiceDeviceOutputFullSIP, bool)`

GetSipOk returns a tuple with the Sip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSip

`func (o *ServiceDeviceOutputFull) SetSip(v ServiceDeviceOutputFullSIP)`

SetSip sets Sip field to given value.

### HasSip

`func (o *ServiceDeviceOutputFull) HasSip() bool`

HasSip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


