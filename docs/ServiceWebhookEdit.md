# ServiceWebhookEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackMethod** | **string** |  | 
**CallbackUrl** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Title** | **string** |  | 
**WebhookType** | **string** |  | 

## Methods

### NewServiceWebhookEdit

`func NewServiceWebhookEdit(callbackMethod string, callbackUrl string, title string, webhookType string, ) *ServiceWebhookEdit`

NewServiceWebhookEdit instantiates a new ServiceWebhookEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWebhookEditWithDefaults

`func NewServiceWebhookEditWithDefaults() *ServiceWebhookEdit`

NewServiceWebhookEditWithDefaults instantiates a new ServiceWebhookEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackMethod

`func (o *ServiceWebhookEdit) GetCallbackMethod() string`

GetCallbackMethod returns the CallbackMethod field if non-nil, zero value otherwise.

### GetCallbackMethodOk

`func (o *ServiceWebhookEdit) GetCallbackMethodOk() (*string, bool)`

GetCallbackMethodOk returns a tuple with the CallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackMethod

`func (o *ServiceWebhookEdit) SetCallbackMethod(v string)`

SetCallbackMethod sets CallbackMethod field to given value.


### GetCallbackUrl

`func (o *ServiceWebhookEdit) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ServiceWebhookEdit) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ServiceWebhookEdit) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetData

`func (o *ServiceWebhookEdit) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceWebhookEdit) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceWebhookEdit) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ServiceWebhookEdit) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIsActive

`func (o *ServiceWebhookEdit) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ServiceWebhookEdit) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ServiceWebhookEdit) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ServiceWebhookEdit) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTitle

`func (o *ServiceWebhookEdit) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceWebhookEdit) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceWebhookEdit) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWebhookType

`func (o *ServiceWebhookEdit) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *ServiceWebhookEdit) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *ServiceWebhookEdit) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


