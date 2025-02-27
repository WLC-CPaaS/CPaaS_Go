# ServiceWebhookAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackMethod** | **string** |  | 
**CallbackUrl** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | **string** |  | 
**WebhookType** | **string** |  | 

## Methods

### NewServiceWebhookAdd

`func NewServiceWebhookAdd(callbackMethod string, callbackUrl string, title string, webhookType string, ) *ServiceWebhookAdd`

NewServiceWebhookAdd instantiates a new ServiceWebhookAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWebhookAddWithDefaults

`func NewServiceWebhookAddWithDefaults() *ServiceWebhookAdd`

NewServiceWebhookAddWithDefaults instantiates a new ServiceWebhookAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackMethod

`func (o *ServiceWebhookAdd) GetCallbackMethod() string`

GetCallbackMethod returns the CallbackMethod field if non-nil, zero value otherwise.

### GetCallbackMethodOk

`func (o *ServiceWebhookAdd) GetCallbackMethodOk() (*string, bool)`

GetCallbackMethodOk returns a tuple with the CallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackMethod

`func (o *ServiceWebhookAdd) SetCallbackMethod(v string)`

SetCallbackMethod sets CallbackMethod field to given value.


### GetCallbackUrl

`func (o *ServiceWebhookAdd) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ServiceWebhookAdd) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ServiceWebhookAdd) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetData

`func (o *ServiceWebhookAdd) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceWebhookAdd) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceWebhookAdd) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ServiceWebhookAdd) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTitle

`func (o *ServiceWebhookAdd) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceWebhookAdd) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceWebhookAdd) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWebhookType

`func (o *ServiceWebhookAdd) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *ServiceWebhookAdd) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *ServiceWebhookAdd) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


