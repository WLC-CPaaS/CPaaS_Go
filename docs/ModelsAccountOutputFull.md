# ModelsAccountOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerId** | Pointer to [**ModelsVOIPAccountOutputFullCallerid**](ModelsVOIPAccountOutputFullCallerid.md) |  | [optional] 
**DoNotDisturb** | Pointer to [**ModelsVOIPSharedDoNotDisturb**](ModelsVOIPSharedDoNotDisturb.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MusicOnHold** | Pointer to [**ModelsVOIPAccountMusicOnHold**](ModelsVOIPAccountMusicOnHold.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsAccountOutputFull

`func NewModelsAccountOutputFull() *ModelsAccountOutputFull`

NewModelsAccountOutputFull instantiates a new ModelsAccountOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAccountOutputFullWithDefaults

`func NewModelsAccountOutputFullWithDefaults() *ModelsAccountOutputFull`

NewModelsAccountOutputFullWithDefaults instantiates a new ModelsAccountOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerId

`func (o *ModelsAccountOutputFull) GetCallerId() ModelsVOIPAccountOutputFullCallerid`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *ModelsAccountOutputFull) GetCallerIdOk() (*ModelsVOIPAccountOutputFullCallerid, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *ModelsAccountOutputFull) SetCallerId(v ModelsVOIPAccountOutputFullCallerid)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *ModelsAccountOutputFull) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ModelsAccountOutputFull) GetDoNotDisturb() ModelsVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ModelsAccountOutputFull) GetDoNotDisturbOk() (*ModelsVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ModelsAccountOutputFull) SetDoNotDisturb(v ModelsVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ModelsAccountOutputFull) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelsAccountOutputFull) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsAccountOutputFull) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsAccountOutputFull) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsAccountOutputFull) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ModelsAccountOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAccountOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAccountOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsAccountOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ModelsAccountOutputFull) GetMusicOnHold() ModelsVOIPAccountMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ModelsAccountOutputFull) GetMusicOnHoldOk() (*ModelsVOIPAccountMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ModelsAccountOutputFull) SetMusicOnHold(v ModelsVOIPAccountMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ModelsAccountOutputFull) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetName

`func (o *ModelsAccountOutputFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsAccountOutputFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsAccountOutputFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsAccountOutputFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealm

`func (o *ModelsAccountOutputFull) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *ModelsAccountOutputFull) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *ModelsAccountOutputFull) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *ModelsAccountOutputFull) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetTimezone

`func (o *ModelsAccountOutputFull) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ModelsAccountOutputFull) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ModelsAccountOutputFull) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ModelsAccountOutputFull) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


