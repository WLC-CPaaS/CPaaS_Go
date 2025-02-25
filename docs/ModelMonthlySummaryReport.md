# ModelMonthlySummaryReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]ModelMonthlySummaryComponentsData**](ModelMonthlySummaryComponentsData.md) |  | [optional] 
**TransactionMonth** | Pointer to **int32** |  | [optional] 
**TransactionYear** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelMonthlySummaryReport

`func NewModelMonthlySummaryReport() *ModelMonthlySummaryReport`

NewModelMonthlySummaryReport instantiates a new ModelMonthlySummaryReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMonthlySummaryReportWithDefaults

`func NewModelMonthlySummaryReportWithDefaults() *ModelMonthlySummaryReport`

NewModelMonthlySummaryReportWithDefaults instantiates a new ModelMonthlySummaryReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *ModelMonthlySummaryReport) GetComponents() []ModelMonthlySummaryComponentsData`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ModelMonthlySummaryReport) GetComponentsOk() (*[]ModelMonthlySummaryComponentsData, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ModelMonthlySummaryReport) SetComponents(v []ModelMonthlySummaryComponentsData)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ModelMonthlySummaryReport) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetTransactionMonth

`func (o *ModelMonthlySummaryReport) GetTransactionMonth() int32`

GetTransactionMonth returns the TransactionMonth field if non-nil, zero value otherwise.

### GetTransactionMonthOk

`func (o *ModelMonthlySummaryReport) GetTransactionMonthOk() (*int32, bool)`

GetTransactionMonthOk returns a tuple with the TransactionMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMonth

`func (o *ModelMonthlySummaryReport) SetTransactionMonth(v int32)`

SetTransactionMonth sets TransactionMonth field to given value.

### HasTransactionMonth

`func (o *ModelMonthlySummaryReport) HasTransactionMonth() bool`

HasTransactionMonth returns a boolean if a field has been set.

### GetTransactionYear

`func (o *ModelMonthlySummaryReport) GetTransactionYear() int32`

GetTransactionYear returns the TransactionYear field if non-nil, zero value otherwise.

### GetTransactionYearOk

`func (o *ModelMonthlySummaryReport) GetTransactionYearOk() (*int32, bool)`

GetTransactionYearOk returns a tuple with the TransactionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionYear

`func (o *ModelMonthlySummaryReport) SetTransactionYear(v int32)`

SetTransactionYear sets TransactionYear field to given value.

### HasTransactionYear

`func (o *ModelMonthlySummaryReport) HasTransactionYear() bool`

HasTransactionYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


