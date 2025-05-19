# ServiceVOIPAccountAddData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerId** | Pointer to [**ModelsVOIPAccountOutputFullCallerid**](ModelsVOIPAccountOutputFullCallerid.md) |  | [optional] 
**DoNotDisturb** | Pointer to [**ModelsVOIPSharedDoNotDisturb**](ModelsVOIPSharedDoNotDisturb.md) |  | [optional] 
**MusicOnHold** | Pointer to [**ModelsVOIPAccountMusicOnHold**](ModelsVOIPAccountMusicOnHold.md) |  | [optional] 
**Name** | **string** |  | 
**Realm** | Pointer to **string** |  | [optional] 
**Timezone** | **string** |  | 

## Methods

### NewServiceVOIPAccountAddData

`func NewServiceVOIPAccountAddData(name string, timezone string, ) *ServiceVOIPAccountAddData`

NewServiceVOIPAccountAddData instantiates a new ServiceVOIPAccountAddData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPAccountAddDataWithDefaults

`func NewServiceVOIPAccountAddDataWithDefaults() *ServiceVOIPAccountAddData`

NewServiceVOIPAccountAddDataWithDefaults instantiates a new ServiceVOIPAccountAddData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerId

`func (o *ServiceVOIPAccountAddData) GetCallerId() ModelsVOIPAccountOutputFullCallerid`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *ServiceVOIPAccountAddData) GetCallerIdOk() (*ModelsVOIPAccountOutputFullCallerid, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *ServiceVOIPAccountAddData) SetCallerId(v ModelsVOIPAccountOutputFullCallerid)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *ServiceVOIPAccountAddData) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceVOIPAccountAddData) GetDoNotDisturb() ModelsVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceVOIPAccountAddData) GetDoNotDisturbOk() (*ModelsVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceVOIPAccountAddData) SetDoNotDisturb(v ModelsVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceVOIPAccountAddData) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ServiceVOIPAccountAddData) GetMusicOnHold() ModelsVOIPAccountMusicOnHold`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ServiceVOIPAccountAddData) GetMusicOnHoldOk() (*ModelsVOIPAccountMusicOnHold, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ServiceVOIPAccountAddData) SetMusicOnHold(v ModelsVOIPAccountMusicOnHold)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ServiceVOIPAccountAddData) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetName

`func (o *ServiceVOIPAccountAddData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPAccountAddData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPAccountAddData) SetName(v string)`

SetName sets Name field to given value.


### GetRealm

`func (o *ServiceVOIPAccountAddData) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *ServiceVOIPAccountAddData) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *ServiceVOIPAccountAddData) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *ServiceVOIPAccountAddData) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetTimezone

`func (o *ServiceVOIPAccountAddData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServiceVOIPAccountAddData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServiceVOIPAccountAddData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


