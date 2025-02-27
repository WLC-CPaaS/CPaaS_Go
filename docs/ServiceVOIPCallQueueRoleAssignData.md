# ServiceVOIPCallQueueRoleAssignData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Recipients** | [**[]map[string][]string**](map[string][]string.md) |  | 
**SetMembership** | **bool** |  | 

## Methods

### NewServiceVOIPCallQueueRoleAssignData

`func NewServiceVOIPCallQueueRoleAssignData(action string, recipients []map[string][]string, setMembership bool, ) *ServiceVOIPCallQueueRoleAssignData`

NewServiceVOIPCallQueueRoleAssignData instantiates a new ServiceVOIPCallQueueRoleAssignData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPCallQueueRoleAssignDataWithDefaults

`func NewServiceVOIPCallQueueRoleAssignDataWithDefaults() *ServiceVOIPCallQueueRoleAssignData`

NewServiceVOIPCallQueueRoleAssignDataWithDefaults instantiates a new ServiceVOIPCallQueueRoleAssignData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ServiceVOIPCallQueueRoleAssignData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ServiceVOIPCallQueueRoleAssignData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ServiceVOIPCallQueueRoleAssignData) SetAction(v string)`

SetAction sets Action field to given value.


### GetRecipients

`func (o *ServiceVOIPCallQueueRoleAssignData) GetRecipients() []map[string][]string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ServiceVOIPCallQueueRoleAssignData) GetRecipientsOk() (*[]map[string][]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ServiceVOIPCallQueueRoleAssignData) SetRecipients(v []map[string][]string)`

SetRecipients sets Recipients field to given value.


### GetSetMembership

`func (o *ServiceVOIPCallQueueRoleAssignData) GetSetMembership() bool`

GetSetMembership returns the SetMembership field if non-nil, zero value otherwise.

### GetSetMembershipOk

`func (o *ServiceVOIPCallQueueRoleAssignData) GetSetMembershipOk() (*bool, bool)`

GetSetMembershipOk returns a tuple with the SetMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetMembership

`func (o *ServiceVOIPCallQueueRoleAssignData) SetSetMembership(v bool)`

SetSetMembership sets SetMembership field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


