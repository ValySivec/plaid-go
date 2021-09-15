/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.33.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Category Information describing a transaction category
type Category struct {
	// An identifying number for the category. `category_id` is a Plaid-specific identifier and does not necessarily correspond to merchant category codes.
	CategoryId string `json:"category_id"`
	// `place` for physical transactions or `special` for other transactions such as bank charges.
	Group string `json:"group"`
	// A hierarchical array of the categories to which this `category_id` belongs.
	Hierarchy []string `json:"hierarchy"`
	AdditionalProperties map[string]interface{}
}

type _Category Category

// NewCategory instantiates a new Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategory(categoryId string, group string, hierarchy []string) *Category {
	this := Category{}
	this.CategoryId = categoryId
	this.Group = group
	this.Hierarchy = hierarchy
	return &this
}

// NewCategoryWithDefaults instantiates a new Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryWithDefaults() *Category {
	this := Category{}
	return &this
}

// GetCategoryId returns the CategoryId field value
func (o *Category) GetCategoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *Category) GetCategoryIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *Category) SetCategoryId(v string) {
	o.CategoryId = v
}

// GetGroup returns the Group field value
func (o *Category) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *Category) GetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *Category) SetGroup(v string) {
	o.Group = v
}

// GetHierarchy returns the Hierarchy field value
func (o *Category) GetHierarchy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hierarchy
}

// GetHierarchyOk returns a tuple with the Hierarchy field value
// and a boolean to check if the value has been set.
func (o *Category) GetHierarchyOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hierarchy, true
}

// SetHierarchy sets field value
func (o *Category) SetHierarchy(v []string) {
	o.Hierarchy = v
}

func (o Category) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["category_id"] = o.CategoryId
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["hierarchy"] = o.Hierarchy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Category) UnmarshalJSON(bytes []byte) (err error) {
	varCategory := _Category{}

	if err = json.Unmarshal(bytes, &varCategory); err == nil {
		*o = Category(varCategory)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "category_id")
		delete(additionalProperties, "group")
		delete(additionalProperties, "hierarchy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCategory struct {
	value *Category
	isSet bool
}

func (v NullableCategory) Get() *Category {
	return v.value
}

func (v *NullableCategory) Set(val *Category) {
	v.value = val
	v.isSet = true
}

func (v NullableCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategory(val *Category) *NullableCategory {
	return &NullableCategory{value: val, isSet: true}
}

func (v NullableCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


