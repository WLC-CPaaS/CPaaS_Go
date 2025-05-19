# ServiceVOIPAccountEditData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerId** | Pointer to [**ModelsVOIPAccountOutputFullCallerid**](ModelsVOIPAccountOutputFullCallerid.md) |  | [optional] 
**DoNotDisturb** | Pointer to [**ModelsVOIPSharedDoNotDisturb**](ModelsVOIPSharedDoNotDisturb.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MusicOnHold** | Pointer to [**ModelsVOIPAccountMusicOnHold**](ModelsVOIPAccountMusicOnHold.md) |  | [optional] 
**Name** | **string** |  | 
**Timezone** | **string** |  | 

## Methods

### NewServiceVOIPAccountEditData

`func NewServiceVOIPAccountEditData(name string, timezone string, ) *ServiceVOIPAccountEditData`

NewServiceVOIPAccountEditData instantiates a new ServiceVOIPAccountEditData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPAccountEditDataWithDefaults

`func NewServiceVOIPAccountEditDataWithDefaults() *ServiceVOIPAccountEditData`

NewServiceVOIPAccountEditDataWithDefaults instantiates a new ServiceVOIPAccountEditData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerId

`func (o *ServiceVOIPAccountEditData) GetCallerId() ModelsVOIPAccountOutputFullCallerid`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *ServiceVOIPAccountEditData) GetCallerIdOk() (*ModelsVOIPAccountOutputFullCallerid, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *ServiceVOIPAccountEditData) SetCallerId(v ModelsVOIPAccountOutputFullCallerid)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *ServiceVOIPAccountEditData) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceVOIPAccountEditData) GetDoNotDisturb() ModelsVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceVOIPAccountEditData) GetDoNotDisturbOk() (*ModelsVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceVOIPAccountEditData) SetDoNotDisturb(v ModelsVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceVOIPAccountEditData) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceVOIPAccountEditData) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceVOIPAccountEditData) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceVOIPAccountEditData) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceVOIPAccountEditData) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ServiceVOIPAccountEditData) GetMusicOnHold() ModelsVOIPAccountMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ServiceVOIPAccountEditData) GetMusicOnHoldOk() (*ModelsVOIPAccountMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ServiceVOIPAccountEditData) SetMusicOnHold(v ModelsVOIPAccountMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ServiceVOIPAccountEditData) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetName

`func (o *ServiceVOIPAccountEditData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPAccountEditData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPAccountEditData) SetName(v string)`

SetName sets Name field to given value.


### GetTimezone

`func (o *ServiceVOIPAccountEditData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServiceVOIPAccountEditData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServiceVOIPAccountEditData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


