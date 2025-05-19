# ServiceMetaflowPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**map[string]ServiceMetaflowPattern**](ServiceMetaflowPattern.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Module** | **string** |  | 

## Methods

### NewServiceMetaflowPattern

`func NewServiceMetaflowPattern(module string, ) *ServiceMetaflowPattern`

NewServiceMetaflowPattern instantiates a new ServiceMetaflowPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMetaflowPatternWithDefaults

`func NewServiceMetaflowPatternWithDefaults() *ServiceMetaflowPattern`

NewServiceMetaflowPatternWithDefaults instantiates a new ServiceMetaflowPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ServiceMetaflowPattern) GetChildren() map[string]ServiceMetaflowPattern`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ServiceMetaflowPattern) GetChildrenOk() (*map[string]ServiceMetaflowPattern, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ServiceMetaflowPattern) SetChildren(v map[string]ServiceMetaflowPattern)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ServiceMetaflowPattern) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetData

`func (o *ServiceMetaflowPattern) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceMetaflowPattern) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceMetaflowPattern) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ServiceMetaflowPattern) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModule

`func (o *ServiceMetaflowPattern) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ServiceMetaflowPattern) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ServiceMetaflowPattern) SetModule(v string)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


