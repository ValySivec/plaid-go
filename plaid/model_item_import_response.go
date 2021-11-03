/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.44.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ItemImportResponse ItemImportResponse defines the response schema for `/item/import`
type ItemImportResponse struct {
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ItemImportResponse ItemImportResponse

// NewItemImportResponse instantiates a new ItemImportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemImportResponse(accessToken string, requestId string) *ItemImportResponse {
	this := ItemImportResponse{}
	this.AccessToken = accessToken
	this.RequestId = requestId
	return &this
}

// NewItemImportResponseWithDefaults instantiates a new ItemImportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemImportResponseWithDefaults() *ItemImportResponse {
	this := ItemImportResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *ItemImportResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ItemImportResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ItemImportResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRequestId returns the RequestId field value
func (o *ItemImportResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ItemImportResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ItemImportResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ItemImportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemImportResponse) UnmarshalJSON(bytes []byte) (err error) {
	varItemImportResponse := _ItemImportResponse{}

	if err = json.Unmarshal(bytes, &varItemImportResponse); err == nil {
		*o = ItemImportResponse(varItemImportResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemImportResponse struct {
	value *ItemImportResponse
	isSet bool
}

func (v NullableItemImportResponse) Get() *ItemImportResponse {
	return v.value
}

func (v *NullableItemImportResponse) Set(val *ItemImportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemImportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemImportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemImportResponse(val *ItemImportResponse) *NullableItemImportResponse {
	return &NullableItemImportResponse{value: val, isSet: true}
}

func (v NullableItemImportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemImportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


