# ServiceUserOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForward** | Pointer to [**ServiceCallForward**](ServiceCallForward.md) |  | [optional] 
**CallRecording** | Pointer to [**ServiceCallRecordingSettings**](ServiceCallRecordingSettings.md) |  | [optional] 
**CallerId** | Pointer to [**ServiceUserOutputFullCallerid**](ServiceUserOutputFullCallerid.md) |  | [optional] 
**DoNotDisturb** | Pointer to [**ServiceVOIPSharedDoNotDisturb**](ServiceVOIPSharedDoNotDisturb.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**MusicOnHold** | Pointer to [**ServiceMusicOnHold**](ServiceMusicOnHold.md) |  | [optional] 
**PresenceId** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceUserOutputFull

`func NewServiceUserOutputFull() *ServiceUserOutputFull`

NewServiceUserOutputFull instantiates a new ServiceUserOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceUserOutputFullWithDefaults

`func NewServiceUserOutputFullWithDefaults() *ServiceUserOutputFull`

NewServiceUserOutputFullWithDefaults instantiates a new ServiceUserOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallForward

`func (o *ServiceUserOutputFull) GetCallForward() ServiceCallForward`

GetCallForward returns the CallForward field if non-nil, zero value otherwise.

### GetCallForwardOk

`func (o *ServiceUserOutputFull) GetCallForwardOk() (*ServiceCallForward, bool)`

GetCallForwardOk returns a tuple with the CallForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForward

`func (o *ServiceUserOutputFull) SetCallForward(v ServiceCallForward)`

SetCallForward sets CallForward field to given value.

### HasCallForward

`func (o *ServiceUserOutputFull) HasCallForward() bool`

HasCallForward returns a boolean if a field has been set.

### GetCallRecording

`func (o *ServiceUserOutputFull) GetCallRecording() ServiceCallRecordingSettings`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ServiceUserOutputFull) GetCallRecordingOk() (*ServiceCallRecordingSettings, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ServiceUserOutputFull) SetCallRecording(v ServiceCallRecordingSettings)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ServiceUserOutputFull) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetCallerId

`func (o *ServiceUserOutputFull) GetCallerId() ServiceUserOutputFullCallerid`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *ServiceUserOutputFull) GetCallerIdOk() (*ServiceUserOutputFullCallerid, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *ServiceUserOutputFull) SetCallerId(v ServiceUserOutputFullCallerid)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *ServiceUserOutputFull) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceUserOutputFull) GetDoNotDisturb() ServiceVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceUserOutputFull) GetDoNotDisturbOk() (*ServiceVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceUserOutputFull) SetDoNotDisturb(v ServiceVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceUserOutputFull) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEmail

`func (o *ServiceUserOutputFull) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceUserOutputFull) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceUserOutputFull) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServiceUserOutputFull) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceUserOutputFull) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceUserOutputFull) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceUserOutputFull) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceUserOutputFull) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFirstName

`func (o *ServiceUserOutputFull) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ServiceUserOutputFull) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ServiceUserOutputFull) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ServiceUserOutputFull) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *ServiceUserOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceUserOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceUserOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceUserOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastName

`func (o *ServiceUserOutputFull) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ServiceUserOutputFull) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ServiceUserOutputFull) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ServiceUserOutputFull) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ServiceUserOutputFull) GetMusicOnHold() ServiceMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ServiceUserOutputFull) GetMusicOnHoldOk() (*ServiceMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ServiceUserOutputFull) SetMusicOnHold(v ServiceMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ServiceUserOutputFull) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetPresenceId

`func (o *ServiceUserOutputFull) GetPresenceId() string`

GetPresenceId returns the PresenceId field if non-nil, zero value otherwise.

### GetPresenceIdOk

`func (o *ServiceUserOutputFull) GetPresenceIdOk() (*string, bool)`

GetPresenceIdOk returns a tuple with the PresenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceId

`func (o *ServiceUserOutputFull) SetPresenceId(v string)`

SetPresenceId sets PresenceId field to given value.

### HasPresenceId

`func (o *ServiceUserOutputFull) HasPresenceId() bool`

HasPresenceId returns a boolean if a field has been set.

### GetTimezone

`func (o *ServiceUserOutputFull) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServiceUserOutputFull) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServiceUserOutputFull) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ServiceUserOutputFull) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetVerified

`func (o *ServiceUserOutputFull) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ServiceUserOutputFull) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ServiceUserOutputFull) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ServiceUserOutputFull) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


