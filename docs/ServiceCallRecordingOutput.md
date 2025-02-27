# ServiceCallRecordingOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallId** | Pointer to **string** |  | [optional] 
**CalleeIdName** | Pointer to **string** |  | [optional] 
**CalleeIdNumber** | Pointer to **string** |  | [optional] 
**CallerIdName** | Pointer to **string** |  | [optional] 
**CallerIdNumber** | Pointer to **string** |  | [optional] 
**CdrId** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**CustomChannelVars** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InteractionId** | Pointer to **string** |  | [optional] 
**MediaSource** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Request** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceCallRecordingOutput

`func NewServiceCallRecordingOutput() *ServiceCallRecordingOutput`

NewServiceCallRecordingOutput instantiates a new ServiceCallRecordingOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallRecordingOutputWithDefaults

`func NewServiceCallRecordingOutputWithDefaults() *ServiceCallRecordingOutput`

NewServiceCallRecordingOutputWithDefaults instantiates a new ServiceCallRecordingOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallId

`func (o *ServiceCallRecordingOutput) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *ServiceCallRecordingOutput) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *ServiceCallRecordingOutput) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *ServiceCallRecordingOutput) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCalleeIdName

`func (o *ServiceCallRecordingOutput) GetCalleeIdName() string`

GetCalleeIdName returns the CalleeIdName field if non-nil, zero value otherwise.

### GetCalleeIdNameOk

`func (o *ServiceCallRecordingOutput) GetCalleeIdNameOk() (*string, bool)`

GetCalleeIdNameOk returns a tuple with the CalleeIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeIdName

`func (o *ServiceCallRecordingOutput) SetCalleeIdName(v string)`

SetCalleeIdName sets CalleeIdName field to given value.

### HasCalleeIdName

`func (o *ServiceCallRecordingOutput) HasCalleeIdName() bool`

HasCalleeIdName returns a boolean if a field has been set.

### GetCalleeIdNumber

`func (o *ServiceCallRecordingOutput) GetCalleeIdNumber() string`

GetCalleeIdNumber returns the CalleeIdNumber field if non-nil, zero value otherwise.

### GetCalleeIdNumberOk

`func (o *ServiceCallRecordingOutput) GetCalleeIdNumberOk() (*string, bool)`

GetCalleeIdNumberOk returns a tuple with the CalleeIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeIdNumber

`func (o *ServiceCallRecordingOutput) SetCalleeIdNumber(v string)`

SetCalleeIdNumber sets CalleeIdNumber field to given value.

### HasCalleeIdNumber

`func (o *ServiceCallRecordingOutput) HasCalleeIdNumber() bool`

HasCalleeIdNumber returns a boolean if a field has been set.

### GetCallerIdName

`func (o *ServiceCallRecordingOutput) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *ServiceCallRecordingOutput) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *ServiceCallRecordingOutput) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *ServiceCallRecordingOutput) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNumber

`func (o *ServiceCallRecordingOutput) GetCallerIdNumber() string`

GetCallerIdNumber returns the CallerIdNumber field if non-nil, zero value otherwise.

### GetCallerIdNumberOk

`func (o *ServiceCallRecordingOutput) GetCallerIdNumberOk() (*string, bool)`

GetCallerIdNumberOk returns a tuple with the CallerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNumber

`func (o *ServiceCallRecordingOutput) SetCallerIdNumber(v string)`

SetCallerIdNumber sets CallerIdNumber field to given value.

### HasCallerIdNumber

`func (o *ServiceCallRecordingOutput) HasCallerIdNumber() bool`

HasCallerIdNumber returns a boolean if a field has been set.

### GetCdrId

`func (o *ServiceCallRecordingOutput) GetCdrId() string`

GetCdrId returns the CdrId field if non-nil, zero value otherwise.

### GetCdrIdOk

`func (o *ServiceCallRecordingOutput) GetCdrIdOk() (*string, bool)`

GetCdrIdOk returns a tuple with the CdrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrId

`func (o *ServiceCallRecordingOutput) SetCdrId(v string)`

SetCdrId sets CdrId field to given value.

### HasCdrId

`func (o *ServiceCallRecordingOutput) HasCdrId() bool`

HasCdrId returns a boolean if a field has been set.

### GetContentType

`func (o *ServiceCallRecordingOutput) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ServiceCallRecordingOutput) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ServiceCallRecordingOutput) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ServiceCallRecordingOutput) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCustomChannelVars

`func (o *ServiceCallRecordingOutput) GetCustomChannelVars() map[string]interface{}`

GetCustomChannelVars returns the CustomChannelVars field if non-nil, zero value otherwise.

### GetCustomChannelVarsOk

`func (o *ServiceCallRecordingOutput) GetCustomChannelVarsOk() (*map[string]interface{}, bool)`

GetCustomChannelVarsOk returns a tuple with the CustomChannelVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomChannelVars

`func (o *ServiceCallRecordingOutput) SetCustomChannelVars(v map[string]interface{})`

SetCustomChannelVars sets CustomChannelVars field to given value.

### HasCustomChannelVars

`func (o *ServiceCallRecordingOutput) HasCustomChannelVars() bool`

HasCustomChannelVars returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceCallRecordingOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceCallRecordingOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceCallRecordingOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceCallRecordingOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *ServiceCallRecordingOutput) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ServiceCallRecordingOutput) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ServiceCallRecordingOutput) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ServiceCallRecordingOutput) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDuration

`func (o *ServiceCallRecordingOutput) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ServiceCallRecordingOutput) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ServiceCallRecordingOutput) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ServiceCallRecordingOutput) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationMs

`func (o *ServiceCallRecordingOutput) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ServiceCallRecordingOutput) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ServiceCallRecordingOutput) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ServiceCallRecordingOutput) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEndpointId

`func (o *ServiceCallRecordingOutput) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ServiceCallRecordingOutput) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ServiceCallRecordingOutput) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ServiceCallRecordingOutput) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetFrom

`func (o *ServiceCallRecordingOutput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ServiceCallRecordingOutput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ServiceCallRecordingOutput) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ServiceCallRecordingOutput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *ServiceCallRecordingOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceCallRecordingOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceCallRecordingOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceCallRecordingOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInteractionId

`func (o *ServiceCallRecordingOutput) GetInteractionId() string`

GetInteractionId returns the InteractionId field if non-nil, zero value otherwise.

### GetInteractionIdOk

`func (o *ServiceCallRecordingOutput) GetInteractionIdOk() (*string, bool)`

GetInteractionIdOk returns a tuple with the InteractionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionId

`func (o *ServiceCallRecordingOutput) SetInteractionId(v string)`

SetInteractionId sets InteractionId field to given value.

### HasInteractionId

`func (o *ServiceCallRecordingOutput) HasInteractionId() bool`

HasInteractionId returns a boolean if a field has been set.

### GetMediaSource

`func (o *ServiceCallRecordingOutput) GetMediaSource() string`

GetMediaSource returns the MediaSource field if non-nil, zero value otherwise.

### GetMediaSourceOk

`func (o *ServiceCallRecordingOutput) GetMediaSourceOk() (*string, bool)`

GetMediaSourceOk returns a tuple with the MediaSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSource

`func (o *ServiceCallRecordingOutput) SetMediaSource(v string)`

SetMediaSource sets MediaSource field to given value.

### HasMediaSource

`func (o *ServiceCallRecordingOutput) HasMediaSource() bool`

HasMediaSource returns a boolean if a field has been set.

### GetMediaType

`func (o *ServiceCallRecordingOutput) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ServiceCallRecordingOutput) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ServiceCallRecordingOutput) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *ServiceCallRecordingOutput) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetName

`func (o *ServiceCallRecordingOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceCallRecordingOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceCallRecordingOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceCallRecordingOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *ServiceCallRecordingOutput) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ServiceCallRecordingOutput) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ServiceCallRecordingOutput) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ServiceCallRecordingOutput) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOwnerId

`func (o *ServiceCallRecordingOutput) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceCallRecordingOutput) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceCallRecordingOutput) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceCallRecordingOutput) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRequest

`func (o *ServiceCallRecordingOutput) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ServiceCallRecordingOutput) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ServiceCallRecordingOutput) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ServiceCallRecordingOutput) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSourceType

`func (o *ServiceCallRecordingOutput) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ServiceCallRecordingOutput) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ServiceCallRecordingOutput) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ServiceCallRecordingOutput) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetStart

`func (o *ServiceCallRecordingOutput) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ServiceCallRecordingOutput) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ServiceCallRecordingOutput) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ServiceCallRecordingOutput) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTo

`func (o *ServiceCallRecordingOutput) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ServiceCallRecordingOutput) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ServiceCallRecordingOutput) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ServiceCallRecordingOutput) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUrl

`func (o *ServiceCallRecordingOutput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceCallRecordingOutput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceCallRecordingOutput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceCallRecordingOutput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


