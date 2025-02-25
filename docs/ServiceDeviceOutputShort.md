# ServiceDeviceOutputShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallRecording** | Pointer to [**ServiceCallRecordingSettings**](ServiceCallRecordingSettings.md) |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**DoNotDisturb** | Pointer to [**ServiceVOIPSharedDoNotDisturb**](ServiceVOIPSharedDoNotDisturb.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceDeviceOutputShort

`func NewServiceDeviceOutputShort() *ServiceDeviceOutputShort`

NewServiceDeviceOutputShort instantiates a new ServiceDeviceOutputShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeviceOutputShortWithDefaults

`func NewServiceDeviceOutputShortWithDefaults() *ServiceDeviceOutputShort`

NewServiceDeviceOutputShortWithDefaults instantiates a new ServiceDeviceOutputShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallRecording

`func (o *ServiceDeviceOutputShort) GetCallRecording() ServiceCallRecordingSettings`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *ServiceDeviceOutputShort) GetCallRecordingOk() (*ServiceCallRecordingSettings, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *ServiceDeviceOutputShort) SetCallRecording(v ServiceCallRecordingSettings)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *ServiceDeviceOutputShort) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetDeviceType

`func (o *ServiceDeviceOutputShort) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ServiceDeviceOutputShort) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ServiceDeviceOutputShort) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ServiceDeviceOutputShort) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDoNotDisturb

`func (o *ServiceDeviceOutputShort) GetDoNotDisturb() ServiceVOIPSharedDoNotDisturb`

GetDoNotDisturb returns the DoNotDisturb field if non-nil, zero value otherwise.

### GetDoNotDisturbOk

`func (o *ServiceDeviceOutputShort) GetDoNotDisturbOk() (*ServiceVOIPSharedDoNotDisturb, bool)`

GetDoNotDisturbOk returns a tuple with the DoNotDisturb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturb

`func (o *ServiceDeviceOutputShort) SetDoNotDisturb(v ServiceVOIPSharedDoNotDisturb)`

SetDoNotDisturb sets DoNotDisturb field to given value.

### HasDoNotDisturb

`func (o *ServiceDeviceOutputShort) HasDoNotDisturb() bool`

HasDoNotDisturb returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceDeviceOutputShort) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceDeviceOutputShort) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceDeviceOutputShort) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceDeviceOutputShort) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatures

`func (o *ServiceDeviceOutputShort) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ServiceDeviceOutputShort) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ServiceDeviceOutputShort) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ServiceDeviceOutputShort) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFlags

`func (o *ServiceDeviceOutputShort) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ServiceDeviceOutputShort) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ServiceDeviceOutputShort) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ServiceDeviceOutputShort) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetId

`func (o *ServiceDeviceOutputShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceDeviceOutputShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceDeviceOutputShort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceDeviceOutputShort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceDeviceOutputShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDeviceOutputShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDeviceOutputShort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceDeviceOutputShort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *ServiceDeviceOutputShort) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceDeviceOutputShort) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceDeviceOutputShort) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceDeviceOutputShort) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetUsername

`func (o *ServiceDeviceOutputShort) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceDeviceOutputShort) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceDeviceOutputShort) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServiceDeviceOutputShort) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


