# ServiceCallQueueStatusStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbandonedSessions** | Pointer to **int32** |  | [optional] 
**ActiveSessionCount** | Pointer to **int32** |  | [optional] 
**AverageWait** | Pointer to **int32** |  | [optional] 
**EstimatedWait** | Pointer to **int32** |  | [optional] 
**LongestWait** | Pointer to **int32** |  | [optional] 
**MissedSessions** | Pointer to **int32** |  | [optional] 
**RecipientCount** | Pointer to **int32** |  | [optional] 
**TotalSessions** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceCallQueueStatusStats

`func NewServiceCallQueueStatusStats() *ServiceCallQueueStatusStats`

NewServiceCallQueueStatusStats instantiates a new ServiceCallQueueStatusStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallQueueStatusStatsWithDefaults

`func NewServiceCallQueueStatusStatsWithDefaults() *ServiceCallQueueStatusStats`

NewServiceCallQueueStatusStatsWithDefaults instantiates a new ServiceCallQueueStatusStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbandonedSessions

`func (o *ServiceCallQueueStatusStats) GetAbandonedSessions() int32`

GetAbandonedSessions returns the AbandonedSessions field if non-nil, zero value otherwise.

### GetAbandonedSessionsOk

`func (o *ServiceCallQueueStatusStats) GetAbandonedSessionsOk() (*int32, bool)`

GetAbandonedSessionsOk returns a tuple with the AbandonedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedSessions

`func (o *ServiceCallQueueStatusStats) SetAbandonedSessions(v int32)`

SetAbandonedSessions sets AbandonedSessions field to given value.

### HasAbandonedSessions

`func (o *ServiceCallQueueStatusStats) HasAbandonedSessions() bool`

HasAbandonedSessions returns a boolean if a field has been set.

### GetActiveSessionCount

`func (o *ServiceCallQueueStatusStats) GetActiveSessionCount() int32`

GetActiveSessionCount returns the ActiveSessionCount field if non-nil, zero value otherwise.

### GetActiveSessionCountOk

`func (o *ServiceCallQueueStatusStats) GetActiveSessionCountOk() (*int32, bool)`

GetActiveSessionCountOk returns a tuple with the ActiveSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessionCount

`func (o *ServiceCallQueueStatusStats) SetActiveSessionCount(v int32)`

SetActiveSessionCount sets ActiveSessionCount field to given value.

### HasActiveSessionCount

`func (o *ServiceCallQueueStatusStats) HasActiveSessionCount() bool`

HasActiveSessionCount returns a boolean if a field has been set.

### GetAverageWait

`func (o *ServiceCallQueueStatusStats) GetAverageWait() int32`

GetAverageWait returns the AverageWait field if non-nil, zero value otherwise.

### GetAverageWaitOk

`func (o *ServiceCallQueueStatusStats) GetAverageWaitOk() (*int32, bool)`

GetAverageWaitOk returns a tuple with the AverageWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWait

`func (o *ServiceCallQueueStatusStats) SetAverageWait(v int32)`

SetAverageWait sets AverageWait field to given value.

### HasAverageWait

`func (o *ServiceCallQueueStatusStats) HasAverageWait() bool`

HasAverageWait returns a boolean if a field has been set.

### GetEstimatedWait

`func (o *ServiceCallQueueStatusStats) GetEstimatedWait() int32`

GetEstimatedWait returns the EstimatedWait field if non-nil, zero value otherwise.

### GetEstimatedWaitOk

`func (o *ServiceCallQueueStatusStats) GetEstimatedWaitOk() (*int32, bool)`

GetEstimatedWaitOk returns a tuple with the EstimatedWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWait

`func (o *ServiceCallQueueStatusStats) SetEstimatedWait(v int32)`

SetEstimatedWait sets EstimatedWait field to given value.

### HasEstimatedWait

`func (o *ServiceCallQueueStatusStats) HasEstimatedWait() bool`

HasEstimatedWait returns a boolean if a field has been set.

### GetLongestWait

`func (o *ServiceCallQueueStatusStats) GetLongestWait() int32`

GetLongestWait returns the LongestWait field if non-nil, zero value otherwise.

### GetLongestWaitOk

`func (o *ServiceCallQueueStatusStats) GetLongestWaitOk() (*int32, bool)`

GetLongestWaitOk returns a tuple with the LongestWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestWait

`func (o *ServiceCallQueueStatusStats) SetLongestWait(v int32)`

SetLongestWait sets LongestWait field to given value.

### HasLongestWait

`func (o *ServiceCallQueueStatusStats) HasLongestWait() bool`

HasLongestWait returns a boolean if a field has been set.

### GetMissedSessions

`func (o *ServiceCallQueueStatusStats) GetMissedSessions() int32`

GetMissedSessions returns the MissedSessions field if non-nil, zero value otherwise.

### GetMissedSessionsOk

`func (o *ServiceCallQueueStatusStats) GetMissedSessionsOk() (*int32, bool)`

GetMissedSessionsOk returns a tuple with the MissedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissedSessions

`func (o *ServiceCallQueueStatusStats) SetMissedSessions(v int32)`

SetMissedSessions sets MissedSessions field to given value.

### HasMissedSessions

`func (o *ServiceCallQueueStatusStats) HasMissedSessions() bool`

HasMissedSessions returns a boolean if a field has been set.

### GetRecipientCount

`func (o *ServiceCallQueueStatusStats) GetRecipientCount() int32`

GetRecipientCount returns the RecipientCount field if non-nil, zero value otherwise.

### GetRecipientCountOk

`func (o *ServiceCallQueueStatusStats) GetRecipientCountOk() (*int32, bool)`

GetRecipientCountOk returns a tuple with the RecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCount

`func (o *ServiceCallQueueStatusStats) SetRecipientCount(v int32)`

SetRecipientCount sets RecipientCount field to given value.

### HasRecipientCount

`func (o *ServiceCallQueueStatusStats) HasRecipientCount() bool`

HasRecipientCount returns a boolean if a field has been set.

### GetTotalSessions

`func (o *ServiceCallQueueStatusStats) GetTotalSessions() int32`

GetTotalSessions returns the TotalSessions field if non-nil, zero value otherwise.

### GetTotalSessionsOk

`func (o *ServiceCallQueueStatusStats) GetTotalSessionsOk() (*int32, bool)`

GetTotalSessionsOk returns a tuple with the TotalSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSessions

`func (o *ServiceCallQueueStatusStats) SetTotalSessions(v int32)`

SetTotalSessions sets TotalSessions field to given value.

### HasTotalSessions

`func (o *ServiceCallQueueStatusStats) HasTotalSessions() bool`

HasTotalSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


