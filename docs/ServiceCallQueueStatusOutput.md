# ServiceCallQueueStatusOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveRecipientCount** | Pointer to **int32** |  | [optional] 
**AvailableRecipientCount** | Pointer to **int32** |  | [optional] 
**Node** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**ServiceCallQueueStatusStats**](ServiceCallQueueStatusStats.md) |  | [optional] 

## Methods

### NewServiceCallQueueStatusOutput

`func NewServiceCallQueueStatusOutput() *ServiceCallQueueStatusOutput`

NewServiceCallQueueStatusOutput instantiates a new ServiceCallQueueStatusOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallQueueStatusOutputWithDefaults

`func NewServiceCallQueueStatusOutputWithDefaults() *ServiceCallQueueStatusOutput`

NewServiceCallQueueStatusOutputWithDefaults instantiates a new ServiceCallQueueStatusOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveRecipientCount

`func (o *ServiceCallQueueStatusOutput) GetActiveRecipientCount() int32`

GetActiveRecipientCount returns the ActiveRecipientCount field if non-nil, zero value otherwise.

### GetActiveRecipientCountOk

`func (o *ServiceCallQueueStatusOutput) GetActiveRecipientCountOk() (*int32, bool)`

GetActiveRecipientCountOk returns a tuple with the ActiveRecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRecipientCount

`func (o *ServiceCallQueueStatusOutput) SetActiveRecipientCount(v int32)`

SetActiveRecipientCount sets ActiveRecipientCount field to given value.

### HasActiveRecipientCount

`func (o *ServiceCallQueueStatusOutput) HasActiveRecipientCount() bool`

HasActiveRecipientCount returns a boolean if a field has been set.

### GetAvailableRecipientCount

`func (o *ServiceCallQueueStatusOutput) GetAvailableRecipientCount() int32`

GetAvailableRecipientCount returns the AvailableRecipientCount field if non-nil, zero value otherwise.

### GetAvailableRecipientCountOk

`func (o *ServiceCallQueueStatusOutput) GetAvailableRecipientCountOk() (*int32, bool)`

GetAvailableRecipientCountOk returns a tuple with the AvailableRecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRecipientCount

`func (o *ServiceCallQueueStatusOutput) SetAvailableRecipientCount(v int32)`

SetAvailableRecipientCount sets AvailableRecipientCount field to given value.

### HasAvailableRecipientCount

`func (o *ServiceCallQueueStatusOutput) HasAvailableRecipientCount() bool`

HasAvailableRecipientCount returns a boolean if a field has been set.

### GetNode

`func (o *ServiceCallQueueStatusOutput) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ServiceCallQueueStatusOutput) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ServiceCallQueueStatusOutput) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *ServiceCallQueueStatusOutput) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetStats

`func (o *ServiceCallQueueStatusOutput) GetStats() ServiceCallQueueStatusStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ServiceCallQueueStatusOutput) GetStatsOk() (*ServiceCallQueueStatusStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ServiceCallQueueStatusOutput) SetStats(v ServiceCallQueueStatusStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ServiceCallQueueStatusOutput) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


