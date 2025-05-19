# ServiceUserOutputShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallRecording** | Pointer to [**ModelsCallRecordingSettings**](ModelsCallRecordingSettings.md) |  | [optional] 
**DoNotDisturb** | Pointer to [**ModelsVOIPSharedDoNotDisturb**](ModelsVOIPSharedDoNotDisturb.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**PresenceId** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceUserOutputShort

`func NewServiceUserOutputShort() *ServiceUserOutputShort`

NewServiceUserOutputShort instantiates a new ServiceUserOutputShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceUserOutputShortWithDefaults

`func NewServiceUserOutputShortWithDefaults() *ServiceUserOutputShort`

NewServiceUserOutputShortWithDefaults instantiates a new ServiceUserOutputShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallRecording

`func (o *ServiceUserOutputShort) GetCallRecording() ModelsCallRecordingSettings`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ServiceUserOutputShort) GetCallRecordingOk() (*ModelsCallRecordingSettings, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ServiceUserOutputShort) SetCallRecording(v ModelsCallRecordingSettings)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ServiceUserOutputShort) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceUserOutputShort) GetDoNotDisturb() ModelsVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceUserOutputShort) GetDoNotDisturbOk() (*ModelsVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceUserOutputShort) SetDoNotDisturb(v ModelsVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceUserOutputShort) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEmail

`func (o *ServiceUserOutputShort) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceUserOutputShort) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceUserOutputShort) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServiceUserOutputShort) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceUserOutputShort) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceUserOutputShort) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceUserOutputShort) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceUserOutputShort) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatures

`func (o *ServiceUserOutputShort) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ServiceUserOutputShort) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ServiceUserOutputShort) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ServiceUserOutputShort) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFirstName

`func (o *ServiceUserOutputShort) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ServiceUserOutputShort) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ServiceUserOutputShort) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ServiceUserOutputShort) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFlags

`func (o *ServiceUserOutputShort) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ServiceUserOutputShort) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ServiceUserOutputShort) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ServiceUserOutputShort) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetId

`func (o *ServiceUserOutputShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceUserOutputShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceUserOutputShort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceUserOutputShort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastName

`func (o *ServiceUserOutputShort) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ServiceUserOutputShort) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ServiceUserOutputShort) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ServiceUserOutputShort) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPresenceId

`func (o *ServiceUserOutputShort) GetPresenceId() string`

GetPresenceId returns the PresenceId field if non-nil, zero value otherwise.

### GetPresenceIdOk

`func (o *ServiceUserOutputShort) GetPresenceIdOk() (*string, bool)`

GetPresenceIdOk returns a tuple with the PresenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceId

`func (o *ServiceUserOutputShort) SetPresenceId(v string)`

SetPresenceId sets PresenceId field to given value.

### HasPresenceId

`func (o *ServiceUserOutputShort) HasPresenceId() bool`

HasPresenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


