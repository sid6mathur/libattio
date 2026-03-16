# ScimV2UsersPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** |  | 
**Id** | **string** |  | 
**UserName** | **string** |  | 
**Name** | [**ScimV2UsersPost201ResponseName**](ScimV2UsersPost201ResponseName.md) |  | 
**Emails** | [**[]ScimV2UsersPost201ResponseEmailsInner**](ScimV2UsersPost201ResponseEmailsInner.md) |  | 
**Roles** | [**[]ScimV2UsersPost201ResponseRolesInner**](ScimV2UsersPost201ResponseRolesInner.md) |  | 
**ProfileUrl** | Pointer to **string** |  | [optional] 
**Active** | **bool** |  | 
**Meta** | [**ScimV2UsersPost201ResponseMeta**](ScimV2UsersPost201ResponseMeta.md) |  | 

## Methods

### NewScimV2UsersPost201Response

`func NewScimV2UsersPost201Response(schemas []string, id string, userName string, name ScimV2UsersPost201ResponseName, emails []ScimV2UsersPost201ResponseEmailsInner, roles []ScimV2UsersPost201ResponseRolesInner, active bool, meta ScimV2UsersPost201ResponseMeta, ) *ScimV2UsersPost201Response`

NewScimV2UsersPost201Response instantiates a new ScimV2UsersPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimV2UsersPost201ResponseWithDefaults

`func NewScimV2UsersPost201ResponseWithDefaults() *ScimV2UsersPost201Response`

NewScimV2UsersPost201ResponseWithDefaults instantiates a new ScimV2UsersPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimV2UsersPost201Response) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimV2UsersPost201Response) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimV2UsersPost201Response) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ScimV2UsersPost201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimV2UsersPost201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimV2UsersPost201Response) SetId(v string)`

SetId sets Id field to given value.


### GetUserName

`func (o *ScimV2UsersPost201Response) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimV2UsersPost201Response) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimV2UsersPost201Response) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *ScimV2UsersPost201Response) GetName() ScimV2UsersPost201ResponseName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimV2UsersPost201Response) GetNameOk() (*ScimV2UsersPost201ResponseName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimV2UsersPost201Response) SetName(v ScimV2UsersPost201ResponseName)`

SetName sets Name field to given value.


### GetEmails

`func (o *ScimV2UsersPost201Response) GetEmails() []ScimV2UsersPost201ResponseEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ScimV2UsersPost201Response) GetEmailsOk() (*[]ScimV2UsersPost201ResponseEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ScimV2UsersPost201Response) SetEmails(v []ScimV2UsersPost201ResponseEmailsInner)`

SetEmails sets Emails field to given value.


### GetRoles

`func (o *ScimV2UsersPost201Response) GetRoles() []ScimV2UsersPost201ResponseRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ScimV2UsersPost201Response) GetRolesOk() (*[]ScimV2UsersPost201ResponseRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ScimV2UsersPost201Response) SetRoles(v []ScimV2UsersPost201ResponseRolesInner)`

SetRoles sets Roles field to given value.


### GetProfileUrl

`func (o *ScimV2UsersPost201Response) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *ScimV2UsersPost201Response) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *ScimV2UsersPost201Response) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *ScimV2UsersPost201Response) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### GetActive

`func (o *ScimV2UsersPost201Response) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimV2UsersPost201Response) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimV2UsersPost201Response) SetActive(v bool)`

SetActive sets Active field to given value.


### GetMeta

`func (o *ScimV2UsersPost201Response) GetMeta() ScimV2UsersPost201ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimV2UsersPost201Response) GetMetaOk() (*ScimV2UsersPost201ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimV2UsersPost201Response) SetMeta(v ScimV2UsersPost201ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


