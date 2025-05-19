# ServiceMetaflowOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingDigit** | Pointer to **string** |  | [optional] 
**Numbers** | Pointer to [**map[string]ServiceMetaflowPattern**](ServiceMetaflowPattern.md) |  | [optional] 
**Patterns** | Pointer to [**map[string]ServiceMetaflowPattern**](ServiceMetaflowPattern.md) |  | [optional] 

## Methods

### NewServiceMetaflowOutput

`func NewServiceMetaflowOutput() *ServiceMetaflowOutput`

NewServiceMetaflowOutput instantiates a new ServiceMetaflowOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMetaflowOutputWithDefaults

`func NewServiceMetaflowOutputWithDefaults() *ServiceMetaflowOutput`

NewServiceMetaflowOutputWithDefaults instantiates a new ServiceMetaflowOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingDigit

`func (o *ServiceMetaflowOutput) GetBindingDigit() string`

GetBindingDigit returns the BindingDigit field if non-nil, zero value otherwise.

### GetBindingDigitOk

`func (o *ServiceMetaflowOutput) GetBindingDigitOk() (*string, bool)`

GetBindingDigitOk returns a tuple with the BindingDigit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingDigit

`func (o *ServiceMetaflowOutput) SetBindingDigit(v string)`

SetBindingDigit sets BindingDigit field to given value.

### HasBindingDigit

`func (o *ServiceMetaflowOutput) HasBindingDigit() bool`

HasBindingDigit returns a boolean if a field has been set.

### GetNumbers

`func (o *ServiceMetaflowOutput) GetNumbers() map[string]ServiceMetaflowPattern`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *ServiceMetaflowOutput) GetNumbersOk() (*map[string]ServiceMetaflowPattern, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *ServiceMetaflowOutput) SetNumbers(v map[string]ServiceMetaflowPattern)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *ServiceMetaflowOutput) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetPatterns

`func (o *ServiceMetaflowOutput) GetPatterns() map[string]ServiceMetaflowPattern`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *ServiceMetaflowOutput) GetPatternsOk() (*map[string]ServiceMetaflowPattern, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *ServiceMetaflowOutput) SetPatterns(v map[string]ServiceMetaflowPattern)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *ServiceMetaflowOutput) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


