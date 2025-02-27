# ServiceCallflowAddEditFlowData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**map[string]ServiceCallflowAddEditFlowData**](ServiceCallflowAddEditFlowData.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Module** | **string** |  | 

## Methods

### NewServiceCallflowAddEditFlowData

`func NewServiceCallflowAddEditFlowData(module string, ) *ServiceCallflowAddEditFlowData`

NewServiceCallflowAddEditFlowData instantiates a new ServiceCallflowAddEditFlowData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallflowAddEditFlowDataWithDefaults

`func NewServiceCallflowAddEditFlowDataWithDefaults() *ServiceCallflowAddEditFlowData`

NewServiceCallflowAddEditFlowDataWithDefaults instantiates a new ServiceCallflowAddEditFlowData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ServiceCallflowAddEditFlowData) GetChildren() map[string]ServiceCallflowAddEditFlowData`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ServiceCallflowAddEditFlowData) GetChildrenOk() (*map[string]ServiceCallflowAddEditFlowData, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ServiceCallflowAddEditFlowData) SetChildren(v map[string]ServiceCallflowAddEditFlowData)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ServiceCallflowAddEditFlowData) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetData

`func (o *ServiceCallflowAddEditFlowData) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceCallflowAddEditFlowData) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceCallflowAddEditFlowData) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ServiceCallflowAddEditFlowData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModule

`func (o *ServiceCallflowAddEditFlowData) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ServiceCallflowAddEditFlowData) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ServiceCallflowAddEditFlowData) SetModule(v string)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


