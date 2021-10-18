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

// IncomeSummaryFieldString struct for IncomeSummaryFieldString
type IncomeSummaryFieldString struct {
	// The value of the field.
	Value string `json:"value"`
	VerificationStatus VerificationStatus `json:"verification_status"`
	AdditionalProperties map[string]interface{}
}

type _IncomeSummaryFieldString IncomeSummaryFieldString

// NewIncomeSummaryFieldString instantiates a new IncomeSummaryFieldString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeSummaryFieldString(value string, verificationStatus VerificationStatus) *IncomeSummaryFieldString {
	this := IncomeSummaryFieldString{}
	this.Value = value
	this.VerificationStatus = verificationStatus
	return &this
}

// NewIncomeSummaryFieldStringWithDefaults instantiates a new IncomeSummaryFieldString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeSummaryFieldStringWithDefaults() *IncomeSummaryFieldString {
	this := IncomeSummaryFieldString{}
	return &this
}

// GetValue returns the Value field value
func (o *IncomeSummaryFieldString) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IncomeSummaryFieldString) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IncomeSummaryFieldString) SetValue(v string) {
	o.Value = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *IncomeSummaryFieldString) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *IncomeSummaryFieldString) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *IncomeSummaryFieldString) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

func (o IncomeSummaryFieldString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeSummaryFieldString) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeSummaryFieldString := _IncomeSummaryFieldString{}

	if err = json.Unmarshal(bytes, &varIncomeSummaryFieldString); err == nil {
		*o = IncomeSummaryFieldString(varIncomeSummaryFieldString)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "verification_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeSummaryFieldString struct {
	value *IncomeSummaryFieldString
	isSet bool
}

func (v NullableIncomeSummaryFieldString) Get() *IncomeSummaryFieldString {
	return v.value
}

func (v *NullableIncomeSummaryFieldString) Set(val *IncomeSummaryFieldString) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeSummaryFieldString) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeSummaryFieldString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeSummaryFieldString(val *IncomeSummaryFieldString) *NullableIncomeSummaryFieldString {
	return &NullableIncomeSummaryFieldString{value: val, isSet: true}
}

func (v NullableIncomeSummaryFieldString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeSummaryFieldString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


