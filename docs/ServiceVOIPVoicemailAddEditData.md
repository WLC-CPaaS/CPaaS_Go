# ServiceVOIPVoicemailAddEditData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailbox** | **string** |  | 
**Media** | Pointer to [**ServiceVoicemailMedia**](ServiceVoicemailMedia.md) |  | [optional] 
**MediaExtension** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | Pointer to **string** |  | [optional] 
**Pin** | Pointer to **string** |  | [optional] 
**RequirePin** | Pointer to **bool** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceVOIPVoicemailAddEditData

`func NewServiceVOIPVoicemailAddEditData(mailbox string, name string, ) *ServiceVOIPVoicemailAddEditData`

NewServiceVOIPVoicemailAddEditData instantiates a new ServiceVOIPVoicemailAddEditData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPVoicemailAddEditDataWithDefaults

`func NewServiceVOIPVoicemailAddEditDataWithDefaults() *ServiceVOIPVoicemailAddEditData`

NewServiceVOIPVoicemailAddEditDataWithDefaults instantiates a new ServiceVOIPVoicemailAddEditData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailbox

`func (o *ServiceVOIPVoicemailAddEditData) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *ServiceVOIPVoicemailAddEditData) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *ServiceVOIPVoicemailAddEditData) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.


### GetMedia

`func (o *ServiceVOIPVoicemailAddEditData) GetMedia() ServiceVoicemailMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ServiceVOIPVoicemailAddEditData) GetMediaOk() (*ServiceVoicemailMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ServiceVOIPVoicemailAddEditData) SetMedia(v ServiceVoicemailMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ServiceVOIPVoicemailAddEditData) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetMediaExtension

`func (o *ServiceVOIPVoicemailAddEditData) GetMediaExtension() string`

GetMediaExtension returns the MediaExtension field if non-nil, zero value otherwise.

### GetMediaExtensionOk

`func (o *ServiceVOIPVoicemailAddEditData) GetMediaExtensionOk() (*string, bool)`

GetMediaExtensionOk returns a tuple with the MediaExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaExtension

`func (o *ServiceVOIPVoicemailAddEditData) SetMediaExtension(v string)`

SetMediaExtension sets MediaExtension field to given value.

### HasMediaExtension

`func (o *ServiceVOIPVoicemailAddEditData) HasMediaExtension() bool`

HasMediaExtension returns a boolean if a field has been set.

### GetName

`func (o *ServiceVOIPVoicemailAddEditData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPVoicemailAddEditData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPVoicemailAddEditData) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *ServiceVOIPVoicemailAddEditData) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceVOIPVoicemailAddEditData) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceVOIPVoicemailAddEditData) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceVOIPVoicemailAddEditData) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPin

`func (o *ServiceVOIPVoicemailAddEditData) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ServiceVOIPVoicemailAddEditData) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ServiceVOIPVoicemailAddEditData) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *ServiceVOIPVoicemailAddEditData) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetRequirePin

`func (o *ServiceVOIPVoicemailAddEditData) GetRequirePin() bool`

GetRequirePin returns the RequirePin field if non-nil, zero value otherwise.

### GetRequirePinOk

`func (o *ServiceVOIPVoicemailAddEditData) GetRequirePinOk() (*bool, bool)`

GetRequirePinOk returns a tuple with the RequirePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePin

`func (o *ServiceVOIPVoicemailAddEditData) SetRequirePin(v bool)`

SetRequirePin sets RequirePin field to given value.

### HasRequirePin

`func (o *ServiceVOIPVoicemailAddEditData) HasRequirePin() bool`

HasRequirePin returns a boolean if a field has been set.

### GetTimezone

`func (o *ServiceVOIPVoicemailAddEditData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServiceVOIPVoicemailAddEditData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServiceVOIPVoicemailAddEditData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ServiceVOIPVoicemailAddEditData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


