# ServiceAccountOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallRecording** | Pointer to [**ServiceVOIPAccountCallRecording**](ServiceVOIPAccountCallRecording.md) |  | [optional] 
**DoNotDisturb** | Pointer to [**ServiceVOIPSharedDoNotDisturb**](ServiceVOIPSharedDoNotDisturb.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MusicOnHold** | Pointer to [**ServiceVOIPAccountMusicOnHold**](ServiceVOIPAccountMusicOnHold.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceAccountOutput

`func NewServiceAccountOutput() *ServiceAccountOutput`

NewServiceAccountOutput instantiates a new ServiceAccountOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountOutputWithDefaults

`func NewServiceAccountOutputWithDefaults() *ServiceAccountOutput`

NewServiceAccountOutputWithDefaults instantiates a new ServiceAccountOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallRecording

`func (o *ServiceAccountOutput) GetCallRecording() ServiceVOIPAccountCallRecording`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ServiceAccountOutput) GetCallRecordingOk() (*ServiceVOIPAccountCallRecording, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ServiceAccountOutput) SetCallRecording(v ServiceVOIPAccountCallRecording)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ServiceAccountOutput) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceAccountOutput) GetDoNotDisturb() ServiceVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceAccountOutput) GetDoNotDisturbOk() (*ServiceVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceAccountOutput) SetDoNotDisturb(v ServiceVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceAccountOutput) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceAccountOutput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceAccountOutput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceAccountOutput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceAccountOutput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ServiceAccountOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccountOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ServiceAccountOutput) GetMusicOnHold() ServiceVOIPAccountMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ServiceAccountOutput) GetMusicOnHoldOk() (*ServiceVOIPAccountMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ServiceAccountOutput) SetMusicOnHold(v ServiceVOIPAccountMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ServiceAccountOutput) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccountOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAccountOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealm

`func (o *ServiceAccountOutput) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *ServiceAccountOutput) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *ServiceAccountOutput) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *ServiceAccountOutput) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetTimezone

`func (o *ServiceAccountOutput) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServiceAccountOutput) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServiceAccountOutput) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ServiceAccountOutput) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


