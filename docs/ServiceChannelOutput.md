# ServiceChannelOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answered** | Pointer to **bool** |  | [optional] 
**AuthorizingId** | Pointer to **string** |  | [optional] 
**AuthorizingType** | Pointer to **string** |  | [optional] 
**CallflowId** | Pointer to **string** |  | [optional] 
**ChannelAuthorized** | Pointer to **bool** |  | [optional] 
**CustomApplicationVars** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomAuthHeaders** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomChannelVars** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomSipHeaders** | Pointer to **map[string]interface{}** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**ElapsedS** | Pointer to **int32** |  | [optional] 
**FromTag** | Pointer to **string** |  | [optional] 
**InteractionId** | Pointer to **string** |  | [optional] 
**IsLoopback** | Pointer to **bool** |  | [optional] 
**IsOnhold** | Pointer to **bool** |  | [optional] 
**OtherLeg** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**PresenceId** | Pointer to **string** |  | [optional] 
**Request** | Pointer to **string** |  | [optional] 
**ResellerId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**ToTag** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceChannelOutput

`func NewServiceChannelOutput() *ServiceChannelOutput`

NewServiceChannelOutput instantiates a new ServiceChannelOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceChannelOutputWithDefaults

`func NewServiceChannelOutputWithDefaults() *ServiceChannelOutput`

NewServiceChannelOutputWithDefaults instantiates a new ServiceChannelOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswered

`func (o *ServiceChannelOutput) GetAnswered() bool`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *ServiceChannelOutput) GetAnsweredOk() (*bool, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *ServiceChannelOutput) SetAnswered(v bool)`

SetAnswered sets Answered field to given value.

### HasAnswered

`func (o *ServiceChannelOutput) HasAnswered() bool`

HasAnswered returns a boolean if a field has been set.

### GetAuthorizingId

`func (o *ServiceChannelOutput) GetAuthorizingId() string`

GetAuthorizingId returns the AuthorizingId field if non-nil, zero value otherwise.

### GetAuthorizingIdOk

`func (o *ServiceChannelOutput) GetAuthorizingIdOk() (*string, bool)`

GetAuthorizingIdOk returns a tuple with the AuthorizingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizingId

`func (o *ServiceChannelOutput) SetAuthorizingId(v string)`

SetAuthorizingId sets AuthorizingId field to given value.

### HasAuthorizingId

`func (o *ServiceChannelOutput) HasAuthorizingId() bool`

HasAuthorizingId returns a boolean if a field has been set.

### GetAuthorizingType

`func (o *ServiceChannelOutput) GetAuthorizingType() string`

GetAuthorizingType returns the AuthorizingType field if non-nil, zero value otherwise.

### GetAuthorizingTypeOk

`func (o *ServiceChannelOutput) GetAuthorizingTypeOk() (*string, bool)`

GetAuthorizingTypeOk returns a tuple with the AuthorizingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizingType

`func (o *ServiceChannelOutput) SetAuthorizingType(v string)`

SetAuthorizingType sets AuthorizingType field to given value.

### HasAuthorizingType

`func (o *ServiceChannelOutput) HasAuthorizingType() bool`

HasAuthorizingType returns a boolean if a field has been set.

### GetCallflowId

`func (o *ServiceChannelOutput) GetCallflowId() string`

GetCallflowId returns the CallflowId field if non-nil, zero value otherwise.

### GetCallflowIdOk

`func (o *ServiceChannelOutput) GetCallflowIdOk() (*string, bool)`

GetCallflowIdOk returns a tuple with the CallflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallflowId

`func (o *ServiceChannelOutput) SetCallflowId(v string)`

SetCallflowId sets CallflowId field to given value.

### HasCallflowId

`func (o *ServiceChannelOutput) HasCallflowId() bool`

HasCallflowId returns a boolean if a field has been set.

### GetChannelAuthorized

`func (o *ServiceChannelOutput) GetChannelAuthorized() bool`

GetChannelAuthorized returns the ChannelAuthorized field if non-nil, zero value otherwise.

### GetChannelAuthorizedOk

`func (o *ServiceChannelOutput) GetChannelAuthorizedOk() (*bool, bool)`

GetChannelAuthorizedOk returns a tuple with the ChannelAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAuthorized

`func (o *ServiceChannelOutput) SetChannelAuthorized(v bool)`

SetChannelAuthorized sets ChannelAuthorized field to given value.

### HasChannelAuthorized

`func (o *ServiceChannelOutput) HasChannelAuthorized() bool`

HasChannelAuthorized returns a boolean if a field has been set.

### GetCustomApplicationVars

`func (o *ServiceChannelOutput) GetCustomApplicationVars() map[string]interface{}`

GetCustomApplicationVars returns the CustomApplicationVars field if non-nil, zero value otherwise.

### GetCustomApplicationVarsOk

`func (o *ServiceChannelOutput) GetCustomApplicationVarsOk() (*map[string]interface{}, bool)`

GetCustomApplicationVarsOk returns a tuple with the CustomApplicationVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomApplicationVars

`func (o *ServiceChannelOutput) SetCustomApplicationVars(v map[string]interface{})`

SetCustomApplicationVars sets CustomApplicationVars field to given value.

### HasCustomApplicationVars

`func (o *ServiceChannelOutput) HasCustomApplicationVars() bool`

HasCustomApplicationVars returns a boolean if a field has been set.

### GetCustomAuthHeaders

`func (o *ServiceChannelOutput) GetCustomAuthHeaders() map[string]interface{}`

GetCustomAuthHeaders returns the CustomAuthHeaders field if non-nil, zero value otherwise.

### GetCustomAuthHeadersOk

`func (o *ServiceChannelOutput) GetCustomAuthHeadersOk() (*map[string]interface{}, bool)`

GetCustomAuthHeadersOk returns a tuple with the CustomAuthHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAuthHeaders

`func (o *ServiceChannelOutput) SetCustomAuthHeaders(v map[string]interface{})`

SetCustomAuthHeaders sets CustomAuthHeaders field to given value.

### HasCustomAuthHeaders

`func (o *ServiceChannelOutput) HasCustomAuthHeaders() bool`

HasCustomAuthHeaders returns a boolean if a field has been set.

### GetCustomChannelVars

`func (o *ServiceChannelOutput) GetCustomChannelVars() map[string]interface{}`

GetCustomChannelVars returns the CustomChannelVars field if non-nil, zero value otherwise.

### GetCustomChannelVarsOk

`func (o *ServiceChannelOutput) GetCustomChannelVarsOk() (*map[string]interface{}, bool)`

GetCustomChannelVarsOk returns a tuple with the CustomChannelVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomChannelVars

`func (o *ServiceChannelOutput) SetCustomChannelVars(v map[string]interface{})`

SetCustomChannelVars sets CustomChannelVars field to given value.

### HasCustomChannelVars

`func (o *ServiceChannelOutput) HasCustomChannelVars() bool`

HasCustomChannelVars returns a boolean if a field has been set.

### GetCustomSipHeaders

`func (o *ServiceChannelOutput) GetCustomSipHeaders() map[string]interface{}`

GetCustomSipHeaders returns the CustomSipHeaders field if non-nil, zero value otherwise.

### GetCustomSipHeadersOk

`func (o *ServiceChannelOutput) GetCustomSipHeadersOk() (*map[string]interface{}, bool)`

GetCustomSipHeadersOk returns a tuple with the CustomSipHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSipHeaders

`func (o *ServiceChannelOutput) SetCustomSipHeaders(v map[string]interface{})`

SetCustomSipHeaders sets CustomSipHeaders field to given value.

### HasCustomSipHeaders

`func (o *ServiceChannelOutput) HasCustomSipHeaders() bool`

HasCustomSipHeaders returns a boolean if a field has been set.

### GetDestination

`func (o *ServiceChannelOutput) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ServiceChannelOutput) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ServiceChannelOutput) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ServiceChannelOutput) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDirection

`func (o *ServiceChannelOutput) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ServiceChannelOutput) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ServiceChannelOutput) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ServiceChannelOutput) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetElapsedS

`func (o *ServiceChannelOutput) GetElapsedS() int32`

GetElapsedS returns the ElapsedS field if non-nil, zero value otherwise.

### GetElapsedSOk

`func (o *ServiceChannelOutput) GetElapsedSOk() (*int32, bool)`

GetElapsedSOk returns a tuple with the ElapsedS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedS

`func (o *ServiceChannelOutput) SetElapsedS(v int32)`

SetElapsedS sets ElapsedS field to given value.

### HasElapsedS

`func (o *ServiceChannelOutput) HasElapsedS() bool`

HasElapsedS returns a boolean if a field has been set.

### GetFromTag

`func (o *ServiceChannelOutput) GetFromTag() string`

GetFromTag returns the FromTag field if non-nil, zero value otherwise.

### GetFromTagOk

`func (o *ServiceChannelOutput) GetFromTagOk() (*string, bool)`

GetFromTagOk returns a tuple with the FromTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTag

`func (o *ServiceChannelOutput) SetFromTag(v string)`

SetFromTag sets FromTag field to given value.

### HasFromTag

`func (o *ServiceChannelOutput) HasFromTag() bool`

HasFromTag returns a boolean if a field has been set.

### GetInteractionId

`func (o *ServiceChannelOutput) GetInteractionId() string`

GetInteractionId returns the InteractionId field if non-nil, zero value otherwise.

### GetInteractionIdOk

`func (o *ServiceChannelOutput) GetInteractionIdOk() (*string, bool)`

GetInteractionIdOk returns a tuple with the InteractionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionId

`func (o *ServiceChannelOutput) SetInteractionId(v string)`

SetInteractionId sets InteractionId field to given value.

### HasInteractionId

`func (o *ServiceChannelOutput) HasInteractionId() bool`

HasInteractionId returns a boolean if a field has been set.

### GetIsLoopback

`func (o *ServiceChannelOutput) GetIsLoopback() bool`

GetIsLoopback returns the IsLoopback field if non-nil, zero value otherwise.

### GetIsLoopbackOk

`func (o *ServiceChannelOutput) GetIsLoopbackOk() (*bool, bool)`

GetIsLoopbackOk returns a tuple with the IsLoopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoopback

`func (o *ServiceChannelOutput) SetIsLoopback(v bool)`

SetIsLoopback sets IsLoopback field to given value.

### HasIsLoopback

`func (o *ServiceChannelOutput) HasIsLoopback() bool`

HasIsLoopback returns a boolean if a field has been set.

### GetIsOnhold

`func (o *ServiceChannelOutput) GetIsOnhold() bool`

GetIsOnhold returns the IsOnhold field if non-nil, zero value otherwise.

### GetIsOnholdOk

`func (o *ServiceChannelOutput) GetIsOnholdOk() (*bool, bool)`

GetIsOnholdOk returns a tuple with the IsOnhold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnhold

`func (o *ServiceChannelOutput) SetIsOnhold(v bool)`

SetIsOnhold sets IsOnhold field to given value.

### HasIsOnhold

`func (o *ServiceChannelOutput) HasIsOnhold() bool`

HasIsOnhold returns a boolean if a field has been set.

### GetOtherLeg

`func (o *ServiceChannelOutput) GetOtherLeg() string`

GetOtherLeg returns the OtherLeg field if non-nil, zero value otherwise.

### GetOtherLegOk

`func (o *ServiceChannelOutput) GetOtherLegOk() (*string, bool)`

GetOtherLegOk returns a tuple with the OtherLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherLeg

`func (o *ServiceChannelOutput) SetOtherLeg(v string)`

SetOtherLeg sets OtherLeg field to given value.

### HasOtherLeg

`func (o *ServiceChannelOutput) HasOtherLeg() bool`

HasOtherLeg returns a boolean if a field has been set.

### GetOwnerId

`func (o *ServiceChannelOutput) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceChannelOutput) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceChannelOutput) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceChannelOutput) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPresenceId

`func (o *ServiceChannelOutput) GetPresenceId() string`

GetPresenceId returns the PresenceId field if non-nil, zero value otherwise.

### GetPresenceIdOk

`func (o *ServiceChannelOutput) GetPresenceIdOk() (*string, bool)`

GetPresenceIdOk returns a tuple with the PresenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceId

`func (o *ServiceChannelOutput) SetPresenceId(v string)`

SetPresenceId sets PresenceId field to given value.

### HasPresenceId

`func (o *ServiceChannelOutput) HasPresenceId() bool`

HasPresenceId returns a boolean if a field has been set.

### GetRequest

`func (o *ServiceChannelOutput) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ServiceChannelOutput) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ServiceChannelOutput) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ServiceChannelOutput) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResellerId

`func (o *ServiceChannelOutput) GetResellerId() string`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *ServiceChannelOutput) GetResellerIdOk() (*string, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *ServiceChannelOutput) SetResellerId(v string)`

SetResellerId sets ResellerId field to given value.

### HasResellerId

`func (o *ServiceChannelOutput) HasResellerId() bool`

HasResellerId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ServiceChannelOutput) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceChannelOutput) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceChannelOutput) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServiceChannelOutput) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetToTag

`func (o *ServiceChannelOutput) GetToTag() string`

GetToTag returns the ToTag field if non-nil, zero value otherwise.

### GetToTagOk

`func (o *ServiceChannelOutput) GetToTagOk() (*string, bool)`

GetToTagOk returns a tuple with the ToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTag

`func (o *ServiceChannelOutput) SetToTag(v string)`

SetToTag sets ToTag field to given value.

### HasToTag

`func (o *ServiceChannelOutput) HasToTag() bool`

HasToTag returns a boolean if a field has been set.

### GetUsername

`func (o *ServiceChannelOutput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceChannelOutput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceChannelOutput) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServiceChannelOutput) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceChannelOutput) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceChannelOutput) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceChannelOutput) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceChannelOutput) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


