# ServiceCdrOutputShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizingId** | Pointer to **string** |  | [optional] 
**BillingSeconds** | Pointer to **int32** |  | [optional] 
**BridgeId** | Pointer to **string** |  | [optional] 
**CallId** | Pointer to **string** |  | [optional] 
**CallPriority** | Pointer to **string** |  | [optional] 
**CallType** | Pointer to **string** |  | [optional] 
**CalleeIdName** | Pointer to **string** |  | [optional] 
**CalleeIdNumber** | Pointer to **string** |  | [optional] 
**CallerIdName** | Pointer to **string** |  | [optional] 
**CallerIdNumber** | Pointer to **string** |  | [optional] 
**CallingFrom** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **string** |  | [optional] 
**DialedNumber** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**DurationSeconds** | Pointer to **int32** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**HangupCause** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InteractionId** | Pointer to **string** |  | [optional] 
**MediaRecordings** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MediaServer** | Pointer to **string** |  | [optional] 
**OtherLegCallId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **string** |  | [optional] 
**RateName** | Pointer to **string** |  | [optional] 
**RecordingUrl** | Pointer to **string** |  | [optional] 
**Request** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**TimestampDatetime** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceCdrOutputShort

`func NewServiceCdrOutputShort() *ServiceCdrOutputShort`

NewServiceCdrOutputShort instantiates a new ServiceCdrOutputShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCdrOutputShortWithDefaults

`func NewServiceCdrOutputShortWithDefaults() *ServiceCdrOutputShort`

NewServiceCdrOutputShortWithDefaults instantiates a new ServiceCdrOutputShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizingId

`func (o *ServiceCdrOutputShort) GetAuthorizingId() string`

GetAuthorizingId returns the AuthorizingId field if non-nil, zero value otherwise.

### GetAuthorizingIdOk

`func (o *ServiceCdrOutputShort) GetAuthorizingIdOk() (*string, bool)`

GetAuthorizingIdOk returns a tuple with the AuthorizingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizingId

`func (o *ServiceCdrOutputShort) SetAuthorizingId(v string)`

SetAuthorizingId sets AuthorizingId field to given value.

### HasAuthorizingId

`func (o *ServiceCdrOutputShort) HasAuthorizingId() bool`

HasAuthorizingId returns a boolean if a field has been set.

### GetBillingSeconds

`func (o *ServiceCdrOutputShort) GetBillingSeconds() int32`

GetBillingSeconds returns the BillingSeconds field if non-nil, zero value otherwise.

### GetBillingSecondsOk

`func (o *ServiceCdrOutputShort) GetBillingSecondsOk() (*int32, bool)`

GetBillingSecondsOk returns a tuple with the BillingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSeconds

`func (o *ServiceCdrOutputShort) SetBillingSeconds(v int32)`

SetBillingSeconds sets BillingSeconds field to given value.

### HasBillingSeconds

`func (o *ServiceCdrOutputShort) HasBillingSeconds() bool`

HasBillingSeconds returns a boolean if a field has been set.

### GetBridgeId

`func (o *ServiceCdrOutputShort) GetBridgeId() string`

GetBridgeId returns the BridgeId field if non-nil, zero value otherwise.

### GetBridgeIdOk

`func (o *ServiceCdrOutputShort) GetBridgeIdOk() (*string, bool)`

GetBridgeIdOk returns a tuple with the BridgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeId

`func (o *ServiceCdrOutputShort) SetBridgeId(v string)`

SetBridgeId sets BridgeId field to given value.

### HasBridgeId

`func (o *ServiceCdrOutputShort) HasBridgeId() bool`

HasBridgeId returns a boolean if a field has been set.

### GetCallId

`func (o *ServiceCdrOutputShort) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *ServiceCdrOutputShort) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *ServiceCdrOutputShort) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *ServiceCdrOutputShort) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCallPriority

`func (o *ServiceCdrOutputShort) GetCallPriority() string`

GetCallPriority returns the CallPriority field if non-nil, zero value otherwise.

### GetCallPriorityOk

`func (o *ServiceCdrOutputShort) GetCallPriorityOk() (*string, bool)`

GetCallPriorityOk returns a tuple with the CallPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPriority

`func (o *ServiceCdrOutputShort) SetCallPriority(v string)`

SetCallPriority sets CallPriority field to given value.

### HasCallPriority

`func (o *ServiceCdrOutputShort) HasCallPriority() bool`

HasCallPriority returns a boolean if a field has been set.

### GetCallType

`func (o *ServiceCdrOutputShort) GetCallType() string`

GetCallType returns the CallType field if non-nil, zero value otherwise.

### GetCallTypeOk

`func (o *ServiceCdrOutputShort) GetCallTypeOk() (*string, bool)`

GetCallTypeOk returns a tuple with the CallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallType

`func (o *ServiceCdrOutputShort) SetCallType(v string)`

SetCallType sets CallType field to given value.

### HasCallType

`func (o *ServiceCdrOutputShort) HasCallType() bool`

HasCallType returns a boolean if a field has been set.

### GetCalleeIdName

`func (o *ServiceCdrOutputShort) GetCalleeIdName() string`

GetCalleeIdName returns the CalleeIdName field if non-nil, zero value otherwise.

### GetCalleeIdNameOk

`func (o *ServiceCdrOutputShort) GetCalleeIdNameOk() (*string, bool)`

GetCalleeIdNameOk returns a tuple with the CalleeIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeIdName

`func (o *ServiceCdrOutputShort) SetCalleeIdName(v string)`

SetCalleeIdName sets CalleeIdName field to given value.

### HasCalleeIdName

`func (o *ServiceCdrOutputShort) HasCalleeIdName() bool`

HasCalleeIdName returns a boolean if a field has been set.

### GetCalleeIdNumber

`func (o *ServiceCdrOutputShort) GetCalleeIdNumber() string`

GetCalleeIdNumber returns the CalleeIdNumber field if non-nil, zero value otherwise.

### GetCalleeIdNumberOk

`func (o *ServiceCdrOutputShort) GetCalleeIdNumberOk() (*string, bool)`

GetCalleeIdNumberOk returns a tuple with the CalleeIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeIdNumber

`func (o *ServiceCdrOutputShort) SetCalleeIdNumber(v string)`

SetCalleeIdNumber sets CalleeIdNumber field to given value.

### HasCalleeIdNumber

`func (o *ServiceCdrOutputShort) HasCalleeIdNumber() bool`

HasCalleeIdNumber returns a boolean if a field has been set.

### GetCallerIdName

`func (o *ServiceCdrOutputShort) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *ServiceCdrOutputShort) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *ServiceCdrOutputShort) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *ServiceCdrOutputShort) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNumber

`func (o *ServiceCdrOutputShort) GetCallerIdNumber() string`

GetCallerIdNumber returns the CallerIdNumber field if non-nil, zero value otherwise.

### GetCallerIdNumberOk

`func (o *ServiceCdrOutputShort) GetCallerIdNumberOk() (*string, bool)`

GetCallerIdNumberOk returns a tuple with the CallerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNumber

`func (o *ServiceCdrOutputShort) SetCallerIdNumber(v string)`

SetCallerIdNumber sets CallerIdNumber field to given value.

### HasCallerIdNumber

`func (o *ServiceCdrOutputShort) HasCallerIdNumber() bool`

HasCallerIdNumber returns a boolean if a field has been set.

### GetCallingFrom

`func (o *ServiceCdrOutputShort) GetCallingFrom() string`

GetCallingFrom returns the CallingFrom field if non-nil, zero value otherwise.

### GetCallingFromOk

`func (o *ServiceCdrOutputShort) GetCallingFromOk() (*string, bool)`

GetCallingFromOk returns a tuple with the CallingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingFrom

`func (o *ServiceCdrOutputShort) SetCallingFrom(v string)`

SetCallingFrom sets CallingFrom field to given value.

### HasCallingFrom

`func (o *ServiceCdrOutputShort) HasCallingFrom() bool`

HasCallingFrom returns a boolean if a field has been set.

### GetCost

`func (o *ServiceCdrOutputShort) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ServiceCdrOutputShort) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ServiceCdrOutputShort) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ServiceCdrOutputShort) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDialedNumber

`func (o *ServiceCdrOutputShort) GetDialedNumber() string`

GetDialedNumber returns the DialedNumber field if non-nil, zero value otherwise.

### GetDialedNumberOk

`func (o *ServiceCdrOutputShort) GetDialedNumberOk() (*string, bool)`

GetDialedNumberOk returns a tuple with the DialedNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialedNumber

`func (o *ServiceCdrOutputShort) SetDialedNumber(v string)`

SetDialedNumber sets DialedNumber field to given value.

### HasDialedNumber

`func (o *ServiceCdrOutputShort) HasDialedNumber() bool`

HasDialedNumber returns a boolean if a field has been set.

### GetDirection

`func (o *ServiceCdrOutputShort) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ServiceCdrOutputShort) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ServiceCdrOutputShort) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ServiceCdrOutputShort) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *ServiceCdrOutputShort) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *ServiceCdrOutputShort) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *ServiceCdrOutputShort) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *ServiceCdrOutputShort) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetFrom

`func (o *ServiceCdrOutputShort) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ServiceCdrOutputShort) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ServiceCdrOutputShort) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ServiceCdrOutputShort) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHangupCause

`func (o *ServiceCdrOutputShort) GetHangupCause() string`

GetHangupCause returns the HangupCause field if non-nil, zero value otherwise.

### GetHangupCauseOk

`func (o *ServiceCdrOutputShort) GetHangupCauseOk() (*string, bool)`

GetHangupCauseOk returns a tuple with the HangupCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHangupCause

`func (o *ServiceCdrOutputShort) SetHangupCause(v string)`

SetHangupCause sets HangupCause field to given value.

### HasHangupCause

`func (o *ServiceCdrOutputShort) HasHangupCause() bool`

HasHangupCause returns a boolean if a field has been set.

### GetId

`func (o *ServiceCdrOutputShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceCdrOutputShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceCdrOutputShort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceCdrOutputShort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInteractionId

`func (o *ServiceCdrOutputShort) GetInteractionId() string`

GetInteractionId returns the InteractionId field if non-nil, zero value otherwise.

### GetInteractionIdOk

`func (o *ServiceCdrOutputShort) GetInteractionIdOk() (*string, bool)`

GetInteractionIdOk returns a tuple with the InteractionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionId

`func (o *ServiceCdrOutputShort) SetInteractionId(v string)`

SetInteractionId sets InteractionId field to given value.

### HasInteractionId

`func (o *ServiceCdrOutputShort) HasInteractionId() bool`

HasInteractionId returns a boolean if a field has been set.

### GetMediaRecordings

`func (o *ServiceCdrOutputShort) GetMediaRecordings() []map[string]interface{}`

GetMediaRecordings returns the MediaRecordings field if non-nil, zero value otherwise.

### GetMediaRecordingsOk

`func (o *ServiceCdrOutputShort) GetMediaRecordingsOk() (*[]map[string]interface{}, bool)`

GetMediaRecordingsOk returns a tuple with the MediaRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaRecordings

`func (o *ServiceCdrOutputShort) SetMediaRecordings(v []map[string]interface{})`

SetMediaRecordings sets MediaRecordings field to given value.

### HasMediaRecordings

`func (o *ServiceCdrOutputShort) HasMediaRecordings() bool`

HasMediaRecordings returns a boolean if a field has been set.

### GetMediaServer

`func (o *ServiceCdrOutputShort) GetMediaServer() string`

GetMediaServer returns the MediaServer field if non-nil, zero value otherwise.

### GetMediaServerOk

`func (o *ServiceCdrOutputShort) GetMediaServerOk() (*string, bool)`

GetMediaServerOk returns a tuple with the MediaServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaServer

`func (o *ServiceCdrOutputShort) SetMediaServer(v string)`

SetMediaServer sets MediaServer field to given value.

### HasMediaServer

`func (o *ServiceCdrOutputShort) HasMediaServer() bool`

HasMediaServer returns a boolean if a field has been set.

### GetOtherLegCallId

`func (o *ServiceCdrOutputShort) GetOtherLegCallId() string`

GetOtherLegCallId returns the OtherLegCallId field if non-nil, zero value otherwise.

### GetOtherLegCallIdOk

`func (o *ServiceCdrOutputShort) GetOtherLegCallIdOk() (*string, bool)`

GetOtherLegCallIdOk returns a tuple with the OtherLegCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherLegCallId

`func (o *ServiceCdrOutputShort) SetOtherLegCallId(v string)`

SetOtherLegCallId sets OtherLegCallId field to given value.

### HasOtherLegCallId

`func (o *ServiceCdrOutputShort) HasOtherLegCallId() bool`

HasOtherLegCallId returns a boolean if a field has been set.

### GetOwnerId

`func (o *ServiceCdrOutputShort) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceCdrOutputShort) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceCdrOutputShort) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceCdrOutputShort) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRate

`func (o *ServiceCdrOutputShort) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ServiceCdrOutputShort) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ServiceCdrOutputShort) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ServiceCdrOutputShort) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRateName

`func (o *ServiceCdrOutputShort) GetRateName() string`

GetRateName returns the RateName field if non-nil, zero value otherwise.

### GetRateNameOk

`func (o *ServiceCdrOutputShort) GetRateNameOk() (*string, bool)`

GetRateNameOk returns a tuple with the RateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateName

`func (o *ServiceCdrOutputShort) SetRateName(v string)`

SetRateName sets RateName field to given value.

### HasRateName

`func (o *ServiceCdrOutputShort) HasRateName() bool`

HasRateName returns a boolean if a field has been set.

### GetRecordingUrl

`func (o *ServiceCdrOutputShort) GetRecordingUrl() string`

GetRecordingUrl returns the RecordingUrl field if non-nil, zero value otherwise.

### GetRecordingUrlOk

`func (o *ServiceCdrOutputShort) GetRecordingUrlOk() (*string, bool)`

GetRecordingUrlOk returns a tuple with the RecordingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingUrl

`func (o *ServiceCdrOutputShort) SetRecordingUrl(v string)`

SetRecordingUrl sets RecordingUrl field to given value.

### HasRecordingUrl

`func (o *ServiceCdrOutputShort) HasRecordingUrl() bool`

HasRecordingUrl returns a boolean if a field has been set.

### GetRequest

`func (o *ServiceCdrOutputShort) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ServiceCdrOutputShort) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ServiceCdrOutputShort) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ServiceCdrOutputShort) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetTimestamp

`func (o *ServiceCdrOutputShort) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceCdrOutputShort) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceCdrOutputShort) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServiceCdrOutputShort) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampDatetime

`func (o *ServiceCdrOutputShort) GetTimestampDatetime() string`

GetTimestampDatetime returns the TimestampDatetime field if non-nil, zero value otherwise.

### GetTimestampDatetimeOk

`func (o *ServiceCdrOutputShort) GetTimestampDatetimeOk() (*string, bool)`

GetTimestampDatetimeOk returns a tuple with the TimestampDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampDatetime

`func (o *ServiceCdrOutputShort) SetTimestampDatetime(v string)`

SetTimestampDatetime sets TimestampDatetime field to given value.

### HasTimestampDatetime

`func (o *ServiceCdrOutputShort) HasTimestampDatetime() bool`

HasTimestampDatetime returns a boolean if a field has been set.

### GetTo

`func (o *ServiceCdrOutputShort) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ServiceCdrOutputShort) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ServiceCdrOutputShort) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ServiceCdrOutputShort) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


