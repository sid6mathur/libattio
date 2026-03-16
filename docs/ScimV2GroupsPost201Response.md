# ScimV2GroupsPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** |  | 
**Id** | **string** |  | 
**DisplayName** | **string** |  | 
**Members** | [**[]ScimV2GroupsPost201ResponseMembersInner**](ScimV2GroupsPost201ResponseMembersInner.md) |  | 
**Meta** | [**ScimV2UsersPost201ResponseMeta**](ScimV2UsersPost201ResponseMeta.md) |  | 

## Methods

### NewScimV2GroupsPost201Response

`func NewScimV2GroupsPost201Response(schemas []string, id string, displayName string, members []ScimV2GroupsPost201ResponseMembersInner, meta ScimV2UsersPost201ResponseMeta, ) *ScimV2GroupsPost201Response`

NewScimV2GroupsPost201Response instantiates a new ScimV2GroupsPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimV2GroupsPost201ResponseWithDefaults

`func NewScimV2GroupsPost201ResponseWithDefaults() *ScimV2GroupsPost201Response`

NewScimV2GroupsPost201ResponseWithDefaults instantiates a new ScimV2GroupsPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimV2GroupsPost201Response) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimV2GroupsPost201Response) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimV2GroupsPost201Response) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ScimV2GroupsPost201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimV2GroupsPost201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimV2GroupsPost201Response) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *ScimV2GroupsPost201Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimV2GroupsPost201Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimV2GroupsPost201Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMembers

`func (o *ScimV2GroupsPost201Response) GetMembers() []ScimV2GroupsPost201ResponseMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ScimV2GroupsPost201Response) GetMembersOk() (*[]ScimV2GroupsPost201ResponseMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ScimV2GroupsPost201Response) SetMembers(v []ScimV2GroupsPost201ResponseMembersInner)`

SetMembers sets Members field to given value.


### GetMeta

`func (o *ScimV2GroupsPost201Response) GetMeta() ScimV2UsersPost201ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimV2GroupsPost201Response) GetMetaOk() (*ScimV2UsersPost201ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimV2GroupsPost201Response) SetMeta(v ScimV2UsersPost201ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


