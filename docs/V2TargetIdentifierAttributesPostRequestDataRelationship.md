# V2TargetIdentifierAttributesPostRequestDataRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | The slug or UUID of the object to create the reverse relationship attribute on. | 
**Title** | **string** | The title for the reverse relationship attribute. | 
**ApiSlug** | **string** | The API slug for the reverse relationship attribute. | 
**IsMultiselect** | **bool** | Whether the related attribute supports selecting multiple values. Combined with the parent attribute&#39;s &#x60;is_multiselect&#x60;, this determines the relationship type: both &#x60;false&#x60; &#x3D; one-to-one, parent &#x60;true&#x60; + related &#x60;false&#x60; &#x3D; many-to-one, parent &#x60;false&#x60; + related &#x60;true&#x60; &#x3D; one-to-many, both &#x60;true&#x60; &#x3D; many-to-many. | 

## Methods

### NewV2TargetIdentifierAttributesPostRequestDataRelationship

`func NewV2TargetIdentifierAttributesPostRequestDataRelationship(object string, title string, apiSlug string, isMultiselect bool, ) *V2TargetIdentifierAttributesPostRequestDataRelationship`

NewV2TargetIdentifierAttributesPostRequestDataRelationship instantiates a new V2TargetIdentifierAttributesPostRequestDataRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TargetIdentifierAttributesPostRequestDataRelationshipWithDefaults

`func NewV2TargetIdentifierAttributesPostRequestDataRelationshipWithDefaults() *V2TargetIdentifierAttributesPostRequestDataRelationship`

NewV2TargetIdentifierAttributesPostRequestDataRelationshipWithDefaults instantiates a new V2TargetIdentifierAttributesPostRequestDataRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) SetObject(v string)`

SetObject sets Object field to given value.


### GetTitle

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetApiSlug

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.


### GetIsMultiselect

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) GetIsMultiselect() bool`

GetIsMultiselect returns the IsMultiselect field if non-nil, zero value otherwise.

### GetIsMultiselectOk

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) GetIsMultiselectOk() (*bool, bool)`

GetIsMultiselectOk returns a tuple with the IsMultiselect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiselect

`func (o *V2TargetIdentifierAttributesPostRequestDataRelationship) SetIsMultiselect(v bool)`

SetIsMultiselect sets IsMultiselect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


