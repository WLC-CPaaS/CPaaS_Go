# ModelsDeviceOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForward** | Pointer to [**ModelsCallForward**](ModelsCallForward.md) |  | [optional] 
**CallRecording** | Pointer to [**ModelsCallRecordingSettings**](ModelsCallRecordingSettings.md) |  | [optional] 
**CallerId** | Pointer to [**ModelsDeviceOutputFullCallerid**](ModelsDeviceOutputFullCallerid.md) |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**DoNotDisturb** | Pointer to [**ModelsVOIPSharedDoNotDisturb**](ModelsVOIPSharedDoNotDisturb.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**ModelsDeviceOutputFullMedia**](ModelsDeviceOutputFullMedia.md) |  | [optional] 
**MusicOnHold** | Pointer to [**ModelsMusicOnHold**](ModelsMusicOnHold.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Provision** | Pointer to [**ModelsDeviceOutputFullProvision**](ModelsDeviceOutputFullProvision.md) |  | [optional] 
**Sip** | Pointer to [**ModelsDeviceOutputFullSIP**](ModelsDeviceOutputFullSIP.md) |  | [optional] 

## Methods

### NewModelsDeviceOutputFull

`func NewModelsDeviceOutputFull() *ModelsDeviceOutputFull`

NewModelsDeviceOutputFull instantiates a new ModelsDeviceOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsDeviceOutputFullWithDefaults

`func NewModelsDeviceOutputFullWithDefaults() *ModelsDeviceOutputFull`

NewModelsDeviceOutputFullWithDefaults instantiates a new ModelsDeviceOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallForward

`func (o *ModelsDeviceOutputFull) GetCallForward() ModelsCallForward`

GetCallForward returns the CallForward field if non-nil, zero value otherwise.

### GetCallForwardOk

`func (o *ModelsDeviceOutputFull) GetCallForwardOk() (*ModelsCallForward, bool)`

GetCallForwardOk returns a tuple with the CallForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForward

`func (o *ModelsDeviceOutputFull) SetCallForward(v ModelsCallForward)`

SetCallForward sets CallForward field to given value.

### HasCallForward

`func (o *ModelsDeviceOutputFull) HasCallForward() bool`

HasCallForward returns a boolean if a field has been set.

### GetCallRecording

`func (o *ModelsDeviceOutputFull) GetCallRecording() ModelsCallRecordingSettings`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ModelsDeviceOutputFull) GetCallRecordingOk() (*ModelsCallRecordingSettings, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ModelsDeviceOutputFull) SetCallRecording(v ModelsCallRecordingSettings)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ModelsDeviceOutputFull) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetCallerId

`func (o *ModelsDeviceOutputFull) GetCallerId() ModelsDeviceOutputFullCallerid`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *ModelsDeviceOutputFull) GetCallerIdOk() (*ModelsDeviceOutputFullCallerid, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *ModelsDeviceOutputFull) SetCallerId(v ModelsDeviceOutputFullCallerid)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *ModelsDeviceOutputFull) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDeviceType

`func (o *ModelsDeviceOutputFull) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ModelsDeviceOutputFull) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ModelsDeviceOutputFull) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ModelsDeviceOutputFull) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ModelsDeviceOutputFull) GetDoNotDisturb() ModelsVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ModelsDeviceOutputFull) GetDoNotDisturbOk() (*ModelsVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ModelsDeviceOutputFull) SetDoNotDisturb(v ModelsVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ModelsDeviceOutputFull) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelsDeviceOutputFull) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsDeviceOutputFull) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsDeviceOutputFull) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsDeviceOutputFull) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ModelsDeviceOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsDeviceOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsDeviceOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsDeviceOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMacAddress

`func (o *ModelsDeviceOutputFull) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ModelsDeviceOutputFull) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ModelsDeviceOutputFull) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ModelsDeviceOutputFull) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMedia

`func (o *ModelsDeviceOutputFull) GetMedia() ModelsDeviceOutputFullMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ModelsDeviceOutputFull) GetMediaOk() (*ModelsDeviceOutputFullMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ModelsDeviceOutputFull) SetMedia(v ModelsDeviceOutputFullMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ModelsDeviceOutputFull) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ModelsDeviceOutputFull) GetMusicOnHold() ModelsMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ModelsDeviceOutputFull) GetMusicOnHoldOk() (*ModelsMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ModelsDeviceOutputFull) SetMusicOnHold(v ModelsMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ModelsDeviceOutputFull) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetName

`func (o *ModelsDeviceOutputFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsDeviceOutputFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsDeviceOutputFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsDeviceOutputFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *ModelsDeviceOutputFull) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ModelsDeviceOutputFull) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ModelsDeviceOutputFull) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ModelsDeviceOutputFull) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetProvision

`func (o *ModelsDeviceOutputFull) GetProvision() ModelsDeviceOutputFullProvision`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *ModelsDeviceOutputFull) GetProvisionOk() (*ModelsDeviceOutputFullProvision, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *ModelsDeviceOutputFull) SetProvision(v ModelsDeviceOutputFullProvision)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *ModelsDeviceOutputFull) HasProvision() bool`

HasProvision returns a boolean if a field has been set.

### GetSip

`func (o *ModelsDeviceOutputFull) GetSip() ModelsDeviceOutputFullSIP`

GetSip returns the Sip field if non-nil, zero value otherwise.

### GetSipOk

`func (o *ModelsDeviceOutputFull) GetSipOk() (*ModelsDeviceOutputFullSIP, bool)`

GetSipOk returns a tuple with the Sip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSip

`func (o *ModelsDeviceOutputFull) SetSip(v ModelsDeviceOutputFullSIP)`

SetSip sets Sip field to given value.

### HasSip

`func (o *ModelsDeviceOutputFull) HasSip() bool`

HasSip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


