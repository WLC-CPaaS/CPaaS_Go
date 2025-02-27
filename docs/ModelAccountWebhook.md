# ModelAccountWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CallbackMethod** | Pointer to **string** |  | [optional] 
**CallbackUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**WebhookType** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAccountWebhook

`func NewModelAccountWebhook() *ModelAccountWebhook`

NewModelAccountWebhook instantiates a new ModelAccountWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAccountWebhookWithDefaults

`func NewModelAccountWebhookWithDefaults() *ModelAccountWebhook`

NewModelAccountWebhookWithDefaults instantiates a new ModelAccountWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ModelAccountWebhook) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ModelAccountWebhook) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ModelAccountWebhook) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ModelAccountWebhook) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCallbackMethod

`func (o *ModelAccountWebhook) GetCallbackMethod() string`

GetCallbackMethod returns the CallbackMethod field if non-nil, zero value otherwise.

### GetCallbackMethodOk

`func (o *ModelAccountWebhook) GetCallbackMethodOk() (*string, bool)`

GetCallbackMethodOk returns a tuple with the CallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackMethod

`func (o *ModelAccountWebhook) SetCallbackMethod(v string)`

SetCallbackMethod sets CallbackMethod field to given value.

### HasCallbackMethod

`func (o *ModelAccountWebhook) HasCallbackMethod() bool`

HasCallbackMethod returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *ModelAccountWebhook) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ModelAccountWebhook) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ModelAccountWebhook) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *ModelAccountWebhook) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelAccountWebhook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelAccountWebhook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelAccountWebhook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelAccountWebhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *ModelAccountWebhook) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelAccountWebhook) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelAccountWebhook) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ModelAccountWebhook) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *ModelAccountWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAccountWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAccountWebhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAccountWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *ModelAccountWebhook) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelAccountWebhook) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelAccountWebhook) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ModelAccountWebhook) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTitle

`func (o *ModelAccountWebhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelAccountWebhook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelAccountWebhook) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelAccountWebhook) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelAccountWebhook) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelAccountWebhook) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelAccountWebhook) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelAccountWebhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebhookType

`func (o *ModelAccountWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *ModelAccountWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *ModelAccountWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.

### HasWebhookType

`func (o *ModelAccountWebhook) HasWebhookType() bool`

HasWebhookType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


