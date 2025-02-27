# ServiceVOIPUserAdd2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForward** | Pointer to [**ServiceCallForward**](ServiceCallForward.md) |  | [optional] 
**CallRecording** | Pointer to [**ServiceCallRecordingSettings**](ServiceCallRecordingSettings.md) |  | [optional] 
**DoNotDisturb** | Pointer to [**ServiceVOIPSharedDoNotDisturb**](ServiceVOIPSharedDoNotDisturb.md) |  | [optional] 
**Email** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**MusicOnHold** | Pointer to [**ServiceMusicOnHold**](ServiceMusicOnHold.md) |  | [optional] 
**PresenceId** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceVOIPUserAdd2

`func NewServiceVOIPUserAdd2(email string, firstName string, lastName string, ) *ServiceVOIPUserAdd2`

NewServiceVOIPUserAdd2 instantiates a new ServiceVOIPUserAdd2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPUserAdd2WithDefaults

`func NewServiceVOIPUserAdd2WithDefaults() *ServiceVOIPUserAdd2`

NewServiceVOIPUserAdd2WithDefaults instantiates a new ServiceVOIPUserAdd2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallForward

`func (o *ServiceVOIPUserAdd2) GetCallForward() ServiceCallForward`

GetCallForward returns the CallForward field if non-nil, zero value otherwise.

### GetCallForwardOk

`func (o *ServiceVOIPUserAdd2) GetCallForwardOk() (*ServiceCallForward, bool)`

GetCallForwardOk returns a tuple with the CallForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForward

`func (o *ServiceVOIPUserAdd2) SetCallForward(v ServiceCallForward)`

SetCallForward sets CallForward field to given value.

### HasCallForward

`func (o *ServiceVOIPUserAdd2) HasCallForward() bool`

HasCallForward returns a boolean if a field has been set.

### GetCallRecording

`func (o *ServiceVOIPUserAdd2) GetCallRecording() ServiceCallRecordingSettings`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ServiceVOIPUserAdd2) GetCallRecordingOk() (*ServiceCallRecordingSettings, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ServiceVOIPUserAdd2) SetCallRecording(v ServiceCallRecordingSettings)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ServiceVOIPUserAdd2) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceVOIPUserAdd2) GetDoNotDisturb() ServiceVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceVOIPUserAdd2) GetDoNotDisturbOk() (*ServiceVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceVOIPUserAdd2) SetDoNotDisturb(v ServiceVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceVOIPUserAdd2) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEmail

`func (o *ServiceVOIPUserAdd2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceVOIPUserAdd2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceVOIPUserAdd2) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *ServiceVOIPUserAdd2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceVOIPUserAdd2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceVOIPUserAdd2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceVOIPUserAdd2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFirstName

`func (o *ServiceVOIPUserAdd2) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ServiceVOIPUserAdd2) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ServiceVOIPUserAdd2) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ServiceVOIPUserAdd2) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ServiceVOIPUserAdd2) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ServiceVOIPUserAdd2) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMusicOnHold

`func (o *ServiceVOIPUserAdd2) GetMusicOnHold() ServiceMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ServiceVOIPUserAdd2) GetMusicOnHoldOk() (*ServiceMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ServiceVOIPUserAdd2) SetMusicOnHold(v ServiceMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ServiceVOIPUserAdd2) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetPresenceId

`func (o *ServiceVOIPUserAdd2) GetPresenceId() string`

GetPresenceId returns the PresenceId field if non-nil, zero value otherwise.

### GetPresenceIdOk

`func (o *ServiceVOIPUserAdd2) GetPresenceIdOk() (*string, bool)`

GetPresenceIdOk returns a tuple with the PresenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceId

`func (o *ServiceVOIPUserAdd2) SetPresenceId(v string)`

SetPresenceId sets PresenceId field to given value.

### HasPresenceId

`func (o *ServiceVOIPUserAdd2) HasPresenceId() bool`

HasPresenceId returns a boolean if a field has been set.

### GetTimezone

`func (o *ServiceVOIPUserAdd2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServiceVOIPUserAdd2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServiceVOIPUserAdd2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ServiceVOIPUserAdd2) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetVerified

`func (o *ServiceVOIPUserAdd2) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ServiceVOIPUserAdd2) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ServiceVOIPUserAdd2) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ServiceVOIPUserAdd2) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


