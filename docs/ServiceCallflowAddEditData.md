# ServiceCallflowAddEditData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Featurecode** | Pointer to [**ServiceFeatureCode**](ServiceFeatureCode.md) |  | [optional] 
**Flow** | Pointer to [**ServiceCallflowAddEditFlowData**](ServiceCallflowAddEditFlowData.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Numbers** | **[]string** |  | 
**Patterns** | **[]string** |  | 

## Methods

### NewServiceCallflowAddEditData

`func NewServiceCallflowAddEditData(numbers []string, patterns []string, ) *ServiceCallflowAddEditData`

NewServiceCallflowAddEditData instantiates a new ServiceCallflowAddEditData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallflowAddEditDataWithDefaults

`func NewServiceCallflowAddEditDataWithDefaults() *ServiceCallflowAddEditData`

NewServiceCallflowAddEditDataWithDefaults instantiates a new ServiceCallflowAddEditData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeaturecode

`func (o *ServiceCallflowAddEditData) GetFeaturecode() ServiceFeatureCode`

GetFeaturecode returns the Featurecode field if non-nil, zero value otherwise.

### GetFeaturecodeOk

`func (o *ServiceCallflowAddEditData) GetFeaturecodeOk() (*ServiceFeatureCode, bool)`

GetFeaturecodeOk returns a tuple with the Featurecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturecode

`func (o *ServiceCallflowAddEditData) SetFeaturecode(v ServiceFeatureCode)`

SetFeaturecode sets Featurecode field to given value.

### HasFeaturecode

`func (o *ServiceCallflowAddEditData) HasFeaturecode() bool`

HasFeaturecode returns a boolean if a field has been set.

### GetFlow

`func (o *ServiceCallflowAddEditData) GetFlow() ServiceCallflowAddEditFlowData`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *ServiceCallflowAddEditData) GetFlowOk() (*ServiceCallflowAddEditFlowData, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *ServiceCallflowAddEditData) SetFlow(v ServiceCallflowAddEditFlowData)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *ServiceCallflowAddEditData) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetName

`func (o *ServiceCallflowAddEditData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceCallflowAddEditData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceCallflowAddEditData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceCallflowAddEditData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumbers

`func (o *ServiceCallflowAddEditData) GetNumbers() []string`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *ServiceCallflowAddEditData) GetNumbersOk() (*[]string, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *ServiceCallflowAddEditData) SetNumbers(v []string)`

SetNumbers sets Numbers field to given value.


### GetPatterns

`func (o *ServiceCallflowAddEditData) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *ServiceCallflowAddEditData) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *ServiceCallflowAddEditData) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


