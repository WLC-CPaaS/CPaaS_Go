# ServiceVOIPTemporalRuleAddEdit2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cycle** | **string** |  | 
**Days** | Pointer to **[]int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**Month** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Ordinal** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **int32** |  | [optional] 
**StartDateReq** | Pointer to **string** |  | [optional] 
**TimeWindowStart** | Pointer to **int32** |  | [optional] 
**TimeWindowStartReq** | Pointer to **string** |  | [optional] 
**TimeWindowStop** | Pointer to **int32** |  | [optional] 
**TimeWindowStopReq** | Pointer to **string** |  | [optional] 
**Wdays** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceVOIPTemporalRuleAddEdit2

`func NewServiceVOIPTemporalRuleAddEdit2(cycle string, name string, ) *ServiceVOIPTemporalRuleAddEdit2`

NewServiceVOIPTemporalRuleAddEdit2 instantiates a new ServiceVOIPTemporalRuleAddEdit2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPTemporalRuleAddEdit2WithDefaults

`func NewServiceVOIPTemporalRuleAddEdit2WithDefaults() *ServiceVOIPTemporalRuleAddEdit2`

NewServiceVOIPTemporalRuleAddEdit2WithDefaults instantiates a new ServiceVOIPTemporalRuleAddEdit2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycle

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetCycle() string`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetCycleOk() (*string, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetCycle(v string)`

SetCycle sets Cycle field to given value.


### GetDays

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetDays() []int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetDaysOk() (*[]int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetDays(v []int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInterval

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMonth

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetName

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetName(v string)`

SetName sets Name field to given value.


### GetOrdinal

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetOrdinal() string`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetOrdinalOk() (*string, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetOrdinal(v string)`

SetOrdinal sets Ordinal field to given value.

### HasOrdinal

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasOrdinal() bool`

HasOrdinal returns a boolean if a field has been set.

### GetStartDate

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartDateReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetStartDateReq() string`

GetStartDateReq returns the StartDateReq field if non-nil, zero value otherwise.

### GetStartDateReqOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetStartDateReqOk() (*string, bool)`

GetStartDateReqOk returns a tuple with the StartDateReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetStartDateReq(v string)`

SetStartDateReq sets StartDateReq field to given value.

### HasStartDateReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasStartDateReq() bool`

HasStartDateReq returns a boolean if a field has been set.

### GetTimeWindowStart

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetTimeWindowStart() int32`

GetTimeWindowStart returns the TimeWindowStart field if non-nil, zero value otherwise.

### GetTimeWindowStartOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetTimeWindowStartOk() (*int32, bool)`

GetTimeWindowStartOk returns a tuple with the TimeWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowStart

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetTimeWindowStart(v int32)`

SetTimeWindowStart sets TimeWindowStart field to given value.

### HasTimeWindowStart

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasTimeWindowStart() bool`

HasTimeWindowStart returns a boolean if a field has been set.

### GetTimeWindowStartReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetTimeWindowStartReq() string`

GetTimeWindowStartReq returns the TimeWindowStartReq field if non-nil, zero value otherwise.

### GetTimeWindowStartReqOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetTimeWindowStartReqOk() (*string, bool)`

GetTimeWindowStartReqOk returns a tuple with the TimeWindowStartReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowStartReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetTimeWindowStartReq(v string)`

SetTimeWindowStartReq sets TimeWindowStartReq field to given value.

### HasTimeWindowStartReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasTimeWindowStartReq() bool`

HasTimeWindowStartReq returns a boolean if a field has been set.

### GetTimeWindowStop

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetTimeWindowStop() int32`

GetTimeWindowStop returns the TimeWindowStop field if non-nil, zero value otherwise.

### GetTimeWindowStopOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetTimeWindowStopOk() (*int32, bool)`

GetTimeWindowStopOk returns a tuple with the TimeWindowStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowStop

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetTimeWindowStop(v int32)`

SetTimeWindowStop sets TimeWindowStop field to given value.

### HasTimeWindowStop

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasTimeWindowStop() bool`

HasTimeWindowStop returns a boolean if a field has been set.

### GetTimeWindowStopReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetTimeWindowStopReq() string`

GetTimeWindowStopReq returns the TimeWindowStopReq field if non-nil, zero value otherwise.

### GetTimeWindowStopReqOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetTimeWindowStopReqOk() (*string, bool)`

GetTimeWindowStopReqOk returns a tuple with the TimeWindowStopReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowStopReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetTimeWindowStopReq(v string)`

SetTimeWindowStopReq sets TimeWindowStopReq field to given value.

### HasTimeWindowStopReq

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasTimeWindowStopReq() bool`

HasTimeWindowStopReq returns a boolean if a field has been set.

### GetWdays

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetWdays() []string`

GetWdays returns the Wdays field if non-nil, zero value otherwise.

### GetWdaysOk

`func (o *ServiceVOIPTemporalRuleAddEdit2) GetWdaysOk() (*[]string, bool)`

GetWdaysOk returns a tuple with the Wdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdays

`func (o *ServiceVOIPTemporalRuleAddEdit2) SetWdays(v []string)`

SetWdays sets Wdays field to given value.

### HasWdays

`func (o *ServiceVOIPTemporalRuleAddEdit2) HasWdays() bool`

HasWdays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


