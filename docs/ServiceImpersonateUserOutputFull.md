# ServiceImpersonateUserOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**AuthToken** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**IsMasterAccount** | Pointer to **bool** |  | [optional] 
**IsReseller** | Pointer to **bool** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**ResellerId** | Pointer to **string** |  | [optional] 
**UserInfo** | Pointer to [**ServiceImpersonatedUserInfo**](ServiceImpersonatedUserInfo.md) |  | [optional] 

## Methods

### NewServiceImpersonateUserOutputFull

`func NewServiceImpersonateUserOutputFull() *ServiceImpersonateUserOutputFull`

NewServiceImpersonateUserOutputFull instantiates a new ServiceImpersonateUserOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceImpersonateUserOutputFullWithDefaults

`func NewServiceImpersonateUserOutputFullWithDefaults() *ServiceImpersonateUserOutputFull`

NewServiceImpersonateUserOutputFullWithDefaults instantiates a new ServiceImpersonateUserOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ServiceImpersonateUserOutputFull) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ServiceImpersonateUserOutputFull) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ServiceImpersonateUserOutputFull) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ServiceImpersonateUserOutputFull) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *ServiceImpersonateUserOutputFull) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ServiceImpersonateUserOutputFull) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ServiceImpersonateUserOutputFull) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ServiceImpersonateUserOutputFull) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAuthToken

`func (o *ServiceImpersonateUserOutputFull) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *ServiceImpersonateUserOutputFull) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *ServiceImpersonateUserOutputFull) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *ServiceImpersonateUserOutputFull) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetClusterId

`func (o *ServiceImpersonateUserOutputFull) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ServiceImpersonateUserOutputFull) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ServiceImpersonateUserOutputFull) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ServiceImpersonateUserOutputFull) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetIsMasterAccount

`func (o *ServiceImpersonateUserOutputFull) GetIsMasterAccount() bool`

GetIsMasterAccount returns the IsMasterAccount field if non-nil, zero value otherwise.

### GetIsMasterAccountOk

`func (o *ServiceImpersonateUserOutputFull) GetIsMasterAccountOk() (*bool, bool)`

GetIsMasterAccountOk returns a tuple with the IsMasterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterAccount

`func (o *ServiceImpersonateUserOutputFull) SetIsMasterAccount(v bool)`

SetIsMasterAccount sets IsMasterAccount field to given value.

### HasIsMasterAccount

`func (o *ServiceImpersonateUserOutputFull) HasIsMasterAccount() bool`

HasIsMasterAccount returns a boolean if a field has been set.

### GetIsReseller

`func (o *ServiceImpersonateUserOutputFull) GetIsReseller() bool`

GetIsReseller returns the IsReseller field if non-nil, zero value otherwise.

### GetIsResellerOk

`func (o *ServiceImpersonateUserOutputFull) GetIsResellerOk() (*bool, bool)`

GetIsResellerOk returns a tuple with the IsReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReseller

`func (o *ServiceImpersonateUserOutputFull) SetIsReseller(v bool)`

SetIsReseller sets IsReseller field to given value.

### HasIsReseller

`func (o *ServiceImpersonateUserOutputFull) HasIsReseller() bool`

HasIsReseller returns a boolean if a field has been set.

### GetOwnerId

`func (o *ServiceImpersonateUserOutputFull) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServiceImpersonateUserOutputFull) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServiceImpersonateUserOutputFull) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServiceImpersonateUserOutputFull) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetResellerId

`func (o *ServiceImpersonateUserOutputFull) GetResellerId() string`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *ServiceImpersonateUserOutputFull) GetResellerIdOk() (*string, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *ServiceImpersonateUserOutputFull) SetResellerId(v string)`

SetResellerId sets ResellerId field to given value.

### HasResellerId

`func (o *ServiceImpersonateUserOutputFull) HasResellerId() bool`

HasResellerId returns a boolean if a field has been set.

### GetUserInfo

`func (o *ServiceImpersonateUserOutputFull) GetUserInfo() ServiceImpersonatedUserInfo`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *ServiceImpersonateUserOutputFull) GetUserInfoOk() (*ServiceImpersonatedUserInfo, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *ServiceImpersonateUserOutputFull) SetUserInfo(v ServiceImpersonatedUserInfo)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *ServiceImpersonateUserOutputFull) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


