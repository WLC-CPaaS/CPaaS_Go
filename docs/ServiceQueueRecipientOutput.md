# ServiceQueueRecipientOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to [**ServiceQueueRecipientOutputRecipient**](ServiceQueueRecipientOutputRecipient.md) |  | [optional] 
**Roles** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceQueueRecipientOutput

`func NewServiceQueueRecipientOutput() *ServiceQueueRecipientOutput`

NewServiceQueueRecipientOutput instantiates a new ServiceQueueRecipientOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceQueueRecipientOutputWithDefaults

`func NewServiceQueueRecipientOutputWithDefaults() *ServiceQueueRecipientOutput`

NewServiceQueueRecipientOutputWithDefaults instantiates a new ServiceQueueRecipientOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ServiceQueueRecipientOutput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServiceQueueRecipientOutput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServiceQueueRecipientOutput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServiceQueueRecipientOutput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *ServiceQueueRecipientOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceQueueRecipientOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceQueueRecipientOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceQueueRecipientOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceQueueRecipientOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceQueueRecipientOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceQueueRecipientOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceQueueRecipientOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecipient

`func (o *ServiceQueueRecipientOutput) GetRecipient() ServiceQueueRecipientOutputRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ServiceQueueRecipientOutput) GetRecipientOk() (*ServiceQueueRecipientOutputRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ServiceQueueRecipientOutput) SetRecipient(v ServiceQueueRecipientOutputRecipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *ServiceQueueRecipientOutput) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetRoles

`func (o *ServiceQueueRecipientOutput) GetRoles() map[string]string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ServiceQueueRecipientOutput) GetRolesOk() (*map[string]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ServiceQueueRecipientOutput) SetRoles(v map[string]string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ServiceQueueRecipientOutput) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


