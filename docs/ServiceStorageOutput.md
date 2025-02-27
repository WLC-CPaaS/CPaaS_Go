# ServiceStorageOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **map[string]interface{}** |  | [optional] 
**Connections** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**ServiceStoragePlan**](ServiceStoragePlan.md) |  | [optional] 

## Methods

### NewServiceStorageOutput

`func NewServiceStorageOutput() *ServiceStorageOutput`

NewServiceStorageOutput instantiates a new ServiceStorageOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStorageOutputWithDefaults

`func NewServiceStorageOutputWithDefaults() *ServiceStorageOutput`

NewServiceStorageOutputWithDefaults instantiates a new ServiceStorageOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *ServiceStorageOutput) GetAttachments() map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ServiceStorageOutput) GetAttachmentsOk() (*map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ServiceStorageOutput) SetAttachments(v map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ServiceStorageOutput) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetConnections

`func (o *ServiceStorageOutput) GetConnections() map[string]interface{}`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ServiceStorageOutput) GetConnectionsOk() (*map[string]interface{}, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ServiceStorageOutput) SetConnections(v map[string]interface{})`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *ServiceStorageOutput) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetId

`func (o *ServiceStorageOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceStorageOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceStorageOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceStorageOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlan

`func (o *ServiceStorageOutput) GetPlan() ServiceStoragePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ServiceStorageOutput) GetPlanOk() (*ServiceStoragePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ServiceStorageOutput) SetPlan(v ServiceStoragePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ServiceStorageOutput) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


