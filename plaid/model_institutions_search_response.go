/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.39.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InstitutionsSearchResponse InstitutionsSearchResponse defines the response schema for `/institutions/search`
type InstitutionsSearchResponse struct {
	// An array of institutions matching the search criteria
	Institutions []Institution `json:"institutions"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _InstitutionsSearchResponse InstitutionsSearchResponse

// NewInstitutionsSearchResponse instantiates a new InstitutionsSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsSearchResponse(institutions []Institution, requestId string) *InstitutionsSearchResponse {
	this := InstitutionsSearchResponse{}
	this.Institutions = institutions
	this.RequestId = requestId
	return &this
}

// NewInstitutionsSearchResponseWithDefaults instantiates a new InstitutionsSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsSearchResponseWithDefaults() *InstitutionsSearchResponse {
	this := InstitutionsSearchResponse{}
	return &this
}

// GetInstitutions returns the Institutions field value
func (o *InstitutionsSearchResponse) GetInstitutions() []Institution {
	if o == nil {
		var ret []Institution
		return ret
	}

	return o.Institutions
}

// GetInstitutionsOk returns a tuple with the Institutions field value
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchResponse) GetInstitutionsOk() (*[]Institution, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Institutions, true
}

// SetInstitutions sets field value
func (o *InstitutionsSearchResponse) SetInstitutions(v []Institution) {
	o.Institutions = v
}

// GetRequestId returns the RequestId field value
func (o *InstitutionsSearchResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InstitutionsSearchResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o InstitutionsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["institutions"] = o.Institutions
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InstitutionsSearchResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInstitutionsSearchResponse := _InstitutionsSearchResponse{}

	if err = json.Unmarshal(bytes, &varInstitutionsSearchResponse); err == nil {
		*o = InstitutionsSearchResponse(varInstitutionsSearchResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "institutions")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionsSearchResponse struct {
	value *InstitutionsSearchResponse
	isSet bool
}

func (v NullableInstitutionsSearchResponse) Get() *InstitutionsSearchResponse {
	return v.value
}

func (v *NullableInstitutionsSearchResponse) Set(val *InstitutionsSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsSearchResponse(val *InstitutionsSearchResponse) *NullableInstitutionsSearchResponse {
	return &NullableInstitutionsSearchResponse{value: val, isSet: true}
}

func (v NullableInstitutionsSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


