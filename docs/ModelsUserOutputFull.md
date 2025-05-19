# ModelsUserOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForward** | Pointer to [**ModelsCallForward**](ModelsCallForward.md) |  | [optional] 
**CallRecording** | Pointer to [**ModelsCallRecordingSettings**](ModelsCallRecordingSettings.md) |  | [optional] 
**CallerId** | Pointer to [**ModelsUserOutputFullCallerid**](ModelsUserOutputFullCallerid.md) |  | [optional] 
**DoNotDisturb** | Pointer to [**ModelsVOIPSharedDoNotDisturb**](ModelsVOIPSharedDoNotDisturb.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**MusicOnHold** | Pointer to [**ModelsMusicOnHold**](ModelsMusicOnHold.md) |  | [optional] 
**PresenceId** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsUserOutputFull

`func NewModelsUserOutputFull() *ModelsUserOutputFull`

NewModelsUserOutputFull instantiates a new ModelsUserOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUserOutputFullWithDefaults

`func NewModelsUserOutputFullWithDefaults() *ModelsUserOutputFull`

NewModelsUserOutputFullWithDefaults instantiates a new ModelsUserOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallForward

`func (o *ModelsUserOutputFull) GetCallForward() ModelsCallForward`

GetCallForward returns the CallForward field if non-nil, zero value otherwise.

### GetCallForwardOk

`func (o *ModelsUserOutputFull) GetCallForwardOk() (*ModelsCallForward, bool)`

GetCallForwardOk returns a tuple with the CallForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForward

`func (o *ModelsUserOutputFull) SetCallForward(v ModelsCallForward)`

SetCallForward sets CallForward field to given value.

### HasCallForward

`func (o *ModelsUserOutputFull) HasCallForward() bool`

HasCallForward returns a boolean if a field has been set.

### GetCallRecording

`func (o *ModelsUserOutputFull) GetCallRecording() ModelsCallRecordingSettings`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ModelsUserOutputFull) GetCallRecordingOk() (*ModelsCallRecordingSettings, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ModelsUserOutputFull) SetCallRecording(v ModelsCallRecordingSettings)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ModelsUserOutputFull) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetCallerId

`func (o *ModelsUserOutputFull) GetCallerId() ModelsUserOutputFullCallerid`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *ModelsUserOutputFull) GetCallerIdOk() (*ModelsUserOutputFullCallerid, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *ModelsUserOutputFull) SetCallerId(v ModelsUserOutputFullCallerid)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *ModelsUserOutputFull) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ModelsUserOutputFull) GetDoNotDisturb() ModelsVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ModelsUserOutputFull) GetDoNotDisturbOk() (*ModelsVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ModelsUserOutputFull) SetDoNotDisturb(v ModelsVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ModelsUserOutputFull) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEmail

`func (o *ModelsUserOutputFull) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsUserOutputFull) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsUserOutputFull) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsUserOutputFull) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelsUserOutputFull) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsUserOutputFull) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsUserOutputFull) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsUserOutputFull) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFirstName

`func (o *ModelsUserOutputFull) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ModelsUserOutputFull) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ModelsUserOutputFull) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ModelsUserOutputFull) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *ModelsUserOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsUserOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsUserOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsUserOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastName

`func (o *ModelsUserOutputFull) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ModelsUserOutputFull) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ModelsUserOutputFull) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ModelsUserOutputFull) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ModelsUserOutputFull) GetMusicOnHold() ModelsMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ModelsUserOutputFull) GetMusicOnHoldOk() (*ModelsMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ModelsUserOutputFull) SetMusicOnHold(v ModelsMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ModelsUserOutputFull) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetPresenceId

`func (o *ModelsUserOutputFull) GetPresenceId() string`

GetPresenceId returns the PresenceId field if non-nil, zero value otherwise.

### GetPresenceIdOk

`func (o *ModelsUserOutputFull) GetPresenceIdOk() (*string, bool)`

GetPresenceIdOk returns a tuple with the PresenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceId

`func (o *ModelsUserOutputFull) SetPresenceId(v string)`

SetPresenceId sets PresenceId field to given value.

### HasPresenceId

`func (o *ModelsUserOutputFull) HasPresenceId() bool`

HasPresenceId returns a boolean if a field has been set.

### GetTimezone

`func (o *ModelsUserOutputFull) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ModelsUserOutputFull) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ModelsUserOutputFull) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ModelsUserOutputFull) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetVerified

`func (o *ModelsUserOutputFull) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ModelsUserOutputFull) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ModelsUserOutputFull) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ModelsUserOutputFull) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


