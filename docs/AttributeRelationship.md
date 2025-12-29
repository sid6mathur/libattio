# AttributeRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**AttributeId**](AttributeId.md) |  | 
**ObjectSlug** | **string** | The slug of the object that the related attribute belongs to. | 
**Title** | **string** | The title of the related attribute. | 
**ApiSlug** | **string** | The API slug identifying the related attribute. | 
**IsMultiselect** | **bool** | Whether the related attribute supports selecting multiple values. Combined with the parent attribute&#39;s &#x60;is_multiselect&#x60;, this determines the relationship type: both &#x60;false&#x60; &#x3D; one-to-one, parent &#x60;true&#x60; + related &#x60;false&#x60; &#x3D; many-to-one, parent &#x60;false&#x60; + related &#x60;true&#x60; &#x3D; one-to-many, both &#x60;true&#x60; &#x3D; many-to-many. | 

## Methods

### NewAttributeRelationship

`func NewAttributeRelationship(id AttributeId, objectSlug string, title string, apiSlug string, isMultiselect bool, ) *AttributeRelationship`

NewAttributeRelationship instantiates a new AttributeRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeRelationshipWithDefaults

`func NewAttributeRelationshipWithDefaults() *AttributeRelationship`

NewAttributeRelationshipWithDefaults instantiates a new AttributeRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttributeRelationship) GetId() AttributeId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeRelationship) GetIdOk() (*AttributeId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeRelationship) SetId(v AttributeId)`

SetId sets Id field to given value.


### GetObjectSlug

`func (o *AttributeRelationship) GetObjectSlug() string`

GetObjectSlug returns the ObjectSlug field if non-nil, zero value otherwise.

### GetObjectSlugOk

`func (o *AttributeRelationship) GetObjectSlugOk() (*string, bool)`

GetObjectSlugOk returns a tuple with the ObjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlug

`func (o *AttributeRelationship) SetObjectSlug(v string)`

SetObjectSlug sets ObjectSlug field to given value.


### GetTitle

`func (o *AttributeRelationship) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttributeRelationship) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttributeRelationship) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetApiSlug

`func (o *AttributeRelationship) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *AttributeRelationship) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *AttributeRelationship) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.


### GetIsMultiselect

`func (o *AttributeRelationship) GetIsMultiselect() bool`

GetIsMultiselect returns the IsMultiselect field if non-nil, zero value otherwise.

### GetIsMultiselectOk

`func (o *AttributeRelationship) GetIsMultiselectOk() (*bool, bool)`

GetIsMultiselectOk returns a tuple with the IsMultiselect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiselect

`func (o *AttributeRelationship) SetIsMultiselect(v bool)`

SetIsMultiselect sets IsMultiselect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


