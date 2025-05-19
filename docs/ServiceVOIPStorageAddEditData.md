# ServiceVOIPStorageAddEditData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **map[string]interface{}** |  | [optional] 
**Connections** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**ServiceStoragePlan**](ServiceStoragePlan.md) |  | [optional] 

## Methods

### NewServiceVOIPStorageAddEditData

`func NewServiceVOIPStorageAddEditData() *ServiceVOIPStorageAddEditData`

NewServiceVOIPStorageAddEditData instantiates a new ServiceVOIPStorageAddEditData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPStorageAddEditDataWithDefaults

`func NewServiceVOIPStorageAddEditDataWithDefaults() *ServiceVOIPStorageAddEditData`

NewServiceVOIPStorageAddEditDataWithDefaults instantiates a new ServiceVOIPStorageAddEditData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *ServiceVOIPStorageAddEditData) GetAttachments() map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ServiceVOIPStorageAddEditData) GetAttachmentsOk() (*map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ServiceVOIPStorageAddEditData) SetAttachments(v map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ServiceVOIPStorageAddEditData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetConnections

`func (o *ServiceVOIPStorageAddEditData) GetConnections() map[string]interface{}`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ServiceVOIPStorageAddEditData) GetConnectionsOk() (*map[string]interface{}, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ServiceVOIPStorageAddEditData) SetConnections(v map[string]interface{})`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *ServiceVOIPStorageAddEditData) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetId

`func (o *ServiceVOIPStorageAddEditData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceVOIPStorageAddEditData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceVOIPStorageAddEditData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceVOIPStorageAddEditData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlan

`func (o *ServiceVOIPStorageAddEditData) GetPlan() ServiceStoragePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ServiceVOIPStorageAddEditData) GetPlanOk() (*ServiceStoragePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ServiceVOIPStorageAddEditData) SetPlan(v ServiceStoragePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ServiceVOIPStorageAddEditData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


