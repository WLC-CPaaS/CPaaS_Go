# ServiceStoragePlanDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**ServiceStoragePlanDatabaseAttachment**](ServiceStoragePlanDatabaseAttachment.md) |  | [optional] 
**Connection** | Pointer to **string** |  | [optional] 
**Database** | Pointer to [**ServiceStoragePlanDatabaseProperties**](ServiceStoragePlanDatabaseProperties.md) |  | [optional] 
**Types** | Pointer to [**ServiceStoragePlanDatabaseTypes**](ServiceStoragePlanDatabaseTypes.md) |  | [optional] 

## Methods

### NewServiceStoragePlanDatabase

`func NewServiceStoragePlanDatabase() *ServiceStoragePlanDatabase`

NewServiceStoragePlanDatabase instantiates a new ServiceStoragePlanDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStoragePlanDatabaseWithDefaults

`func NewServiceStoragePlanDatabaseWithDefaults() *ServiceStoragePlanDatabase`

NewServiceStoragePlanDatabaseWithDefaults instantiates a new ServiceStoragePlanDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *ServiceStoragePlanDatabase) GetAttachments() ServiceStoragePlanDatabaseAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ServiceStoragePlanDatabase) GetAttachmentsOk() (*ServiceStoragePlanDatabaseAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ServiceStoragePlanDatabase) SetAttachments(v ServiceStoragePlanDatabaseAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ServiceStoragePlanDatabase) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetConnection

`func (o *ServiceStoragePlanDatabase) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ServiceStoragePlanDatabase) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ServiceStoragePlanDatabase) SetConnection(v string)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ServiceStoragePlanDatabase) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetDatabase

`func (o *ServiceStoragePlanDatabase) GetDatabase() ServiceStoragePlanDatabaseProperties`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ServiceStoragePlanDatabase) GetDatabaseOk() (*ServiceStoragePlanDatabaseProperties, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ServiceStoragePlanDatabase) SetDatabase(v ServiceStoragePlanDatabaseProperties)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *ServiceStoragePlanDatabase) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetTypes

`func (o *ServiceStoragePlanDatabase) GetTypes() ServiceStoragePlanDatabaseTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ServiceStoragePlanDatabase) GetTypesOk() (*ServiceStoragePlanDatabaseTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ServiceStoragePlanDatabase) SetTypes(v ServiceStoragePlanDatabaseTypes)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ServiceStoragePlanDatabase) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


