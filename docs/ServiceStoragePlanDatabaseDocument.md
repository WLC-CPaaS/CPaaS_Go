# ServiceStoragePlanDatabaseDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**ServiceStoragePlanDatabaseAttachment**](ServiceStoragePlanDatabaseAttachment.md) |  | [optional] 
**Connection** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceStoragePlanDatabaseDocument

`func NewServiceStoragePlanDatabaseDocument() *ServiceStoragePlanDatabaseDocument`

NewServiceStoragePlanDatabaseDocument instantiates a new ServiceStoragePlanDatabaseDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStoragePlanDatabaseDocumentWithDefaults

`func NewServiceStoragePlanDatabaseDocumentWithDefaults() *ServiceStoragePlanDatabaseDocument`

NewServiceStoragePlanDatabaseDocumentWithDefaults instantiates a new ServiceStoragePlanDatabaseDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *ServiceStoragePlanDatabaseDocument) GetAttachments() ServiceStoragePlanDatabaseAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ServiceStoragePlanDatabaseDocument) GetAttachmentsOk() (*ServiceStoragePlanDatabaseAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ServiceStoragePlanDatabaseDocument) SetAttachments(v ServiceStoragePlanDatabaseAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ServiceStoragePlanDatabaseDocument) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetConnection

`func (o *ServiceStoragePlanDatabaseDocument) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ServiceStoragePlanDatabaseDocument) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ServiceStoragePlanDatabaseDocument) SetConnection(v string)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ServiceStoragePlanDatabaseDocument) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


