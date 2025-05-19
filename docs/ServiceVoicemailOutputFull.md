# ServiceVoicemailOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Mailbox** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**ServiceVoicemailMedia**](ServiceVoicemailMedia.md) |  | [optional] 
**MediaExtension** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Pin** | Pointer to **string** |  | [optional] 
**RequirePin** | Pointer to **bool** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceVoicemailOutputFull

`func NewServiceVoicemailOutputFull() *ServiceVoicemailOutputFull`

NewServiceVoicemailOutputFull instantiates a new ServiceVoicemailOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVoicemailOutputFullWithDefaults

`func NewServiceVoicemailOutputFullWithDefaults() *ServiceVoicemailOutputFull`

NewServiceVoicemailOutputFullWithDefaults instantiates a new ServiceVoicemailOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceVoicemailOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceVoicemailOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceVoicemailOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceVoicemailOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMailbox

`func (o *ServiceVoicemailOutputFull) GetMailbox() string`

GetMailbox returns the Mailbox field if non-nil, zero value otherwise.

### GetMailboxOk

`func (o *ServiceVoicemailOutputFull) GetMailboxOk() (*string, bool)`

GetMailboxOk returns a tuple with the Mailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailbox

`func (o *ServiceVoicemailOutputFull) SetMailbox(v string)`

SetMailbox sets Mailbox field to given value.

### HasMailbox

`func (o *ServiceVoicemailOutputFull) HasMailbox() bool`

HasMailbox returns a boolean if a field has been set.

### GetMedia

`func (o *ServiceVoicemailOutputFull) GetMedia() ServiceVoicemailMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ServiceVoicemailOutputFull) GetMediaOk() (*ServiceVoicemailMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ServiceVoicemailOutputFull) SetMedia(v ServiceVoicemailMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ServiceVoicemailOutputFull) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetMediaExtension

`func (o *ServiceVoicemailOutputFull) GetMediaExtension() string`

GetMediaExtension returns the MediaExtension field if non-nil, zero value otherwise.

### GetMediaExtensionOk

`func (o *ServiceVoicemailOutputFull) GetMediaExtensionOk() (*string, bool)`

GetMediaExtensionOk returns a tuple with the MediaExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaExtension

`func (o *ServiceVoicemailOutputFull) SetMediaExtension(v string)`

SetMediaExtension sets MediaExtension field to given value.

### HasMediaExtension

`func (o *ServiceVoicemailOutputFull) HasMediaExtension() bool`

HasMediaExtension returns a boolean if a field has been set.

### GetMessages

`func (o *ServiceVoicemailOutputFull) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ServiceVoicemailOutputFull) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ServiceVoicemailOutputFull) SetMessages(v int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ServiceVoicemailOutputFull) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetName

`func (o *ServiceVoicemailOutputFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVoicemailOutputFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVoicemailOutputFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceVoicemailOutputFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *ServiceVoicemailOutputFull) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceVoicemailOutputFull) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceVoicemailOutputFull) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceVoicemailOutputFull) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPin

`func (o *ServiceVoicemailOutputFull) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ServiceVoicemailOutputFull) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ServiceVoicemailOutputFull) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *ServiceVoicemailOutputFull) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetRequirePin

`func (o *ServiceVoicemailOutputFull) GetRequirePin() bool`

GetRequirePin returns the RequirePin field if non-nil, zero value otherwise.

### GetRequirePinOk

`func (o *ServiceVoicemailOutputFull) GetRequirePinOk() (*bool, bool)`

GetRequirePinOk returns a tuple with the RequirePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePin

`func (o *ServiceVoicemailOutputFull) SetRequirePin(v bool)`

SetRequirePin sets RequirePin field to given value.

### HasRequirePin

`func (o *ServiceVoicemailOutputFull) HasRequirePin() bool`

HasRequirePin returns a boolean if a field has been set.

### GetTimezone

`func (o *ServiceVoicemailOutputFull) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServiceVoicemailOutputFull) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServiceVoicemailOutputFull) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ServiceVoicemailOutputFull) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


