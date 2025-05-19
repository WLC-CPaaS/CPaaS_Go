# ServiceCallflowOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Featurecode** | Pointer to [**ServiceFeatureCode**](ServiceFeatureCode.md) |  | [optional] 
**Flow** | [**ServiceCallflowAddEditFlowData**](ServiceCallflowAddEditFlowData.md) |  | 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Numbers** | **[]string** |  | 
**Patterns** | **[]string** |  | 

## Methods

### NewServiceCallflowOutputFull

`func NewServiceCallflowOutputFull(flow ServiceCallflowAddEditFlowData, numbers []string, patterns []string, ) *ServiceCallflowOutputFull`

NewServiceCallflowOutputFull instantiates a new ServiceCallflowOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallflowOutputFullWithDefaults

`func NewServiceCallflowOutputFullWithDefaults() *ServiceCallflowOutputFull`

NewServiceCallflowOutputFullWithDefaults instantiates a new ServiceCallflowOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeaturecode

`func (o *ServiceCallflowOutputFull) GetFeaturecode() ServiceFeatureCode`

GetFeaturecode returns the Featurecode field if non-nil, zero value otherwise.

### GetFeaturecodeOk

`func (o *ServiceCallflowOutputFull) GetFeaturecodeOk() (*ServiceFeatureCode, bool)`

GetFeaturecodeOk returns a tuple with the Featurecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturecode

`func (o *ServiceCallflowOutputFull) SetFeaturecode(v ServiceFeatureCode)`

SetFeaturecode sets Featurecode field to given value.

### HasFeaturecode

`func (o *ServiceCallflowOutputFull) HasFeaturecode() bool`

HasFeaturecode returns a boolean if a field has been set.

### GetFlow

`func (o *ServiceCallflowOutputFull) GetFlow() ServiceCallflowAddEditFlowData`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *ServiceCallflowOutputFull) GetFlowOk() (*ServiceCallflowAddEditFlowData, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *ServiceCallflowOutputFull) SetFlow(v ServiceCallflowAddEditFlowData)`

SetFlow sets Flow field to given value.


### GetId

`func (o *ServiceCallflowOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceCallflowOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceCallflowOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceCallflowOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceCallflowOutputFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceCallflowOutputFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceCallflowOutputFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceCallflowOutputFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumbers

`func (o *ServiceCallflowOutputFull) GetNumbers() []string`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *ServiceCallflowOutputFull) GetNumbersOk() (*[]string, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *ServiceCallflowOutputFull) SetNumbers(v []string)`

SetNumbers sets Numbers field to given value.


### GetPatterns

`func (o *ServiceCallflowOutputFull) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *ServiceCallflowOutputFull) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *ServiceCallflowOutputFull) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


