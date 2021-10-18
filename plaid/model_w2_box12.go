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

// W2Box12 struct for W2Box12
type W2Box12 struct {
	// W2 Box 12 code.
	Code NullableString `json:"code,omitempty"`
	// W2 Box 12 amount.
	Amount NullableString `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _W2Box12 W2Box12

// NewW2Box12 instantiates a new W2Box12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewW2Box12() *W2Box12 {
	this := W2Box12{}
	return &this
}

// NewW2Box12WithDefaults instantiates a new W2Box12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewW2Box12WithDefaults() *W2Box12 {
	this := W2Box12{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2Box12) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2Box12) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *W2Box12) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *W2Box12) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *W2Box12) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *W2Box12) UnsetCode() {
	o.Code.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2Box12) GetAmount() string {
	if o == nil || o.Amount.Get() == nil {
		var ret string
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2Box12) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *W2Box12) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableString and assigns it to the Amount field.
func (o *W2Box12) SetAmount(v string) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *W2Box12) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *W2Box12) UnsetAmount() {
	o.Amount.Unset()
}

func (o W2Box12) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *W2Box12) UnmarshalJSON(bytes []byte) (err error) {
	varW2Box12 := _W2Box12{}

	if err = json.Unmarshal(bytes, &varW2Box12); err == nil {
		*o = W2Box12(varW2Box12)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableW2Box12 struct {
	value *W2Box12
	isSet bool
}

func (v NullableW2Box12) Get() *W2Box12 {
	return v.value
}

func (v *NullableW2Box12) Set(val *W2Box12) {
	v.value = val
	v.isSet = true
}

func (v NullableW2Box12) IsSet() bool {
	return v.isSet
}

func (v *NullableW2Box12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableW2Box12(val *W2Box12) *NullableW2Box12 {
	return &NullableW2Box12{value: val, isSet: true}
}

func (v NullableW2Box12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableW2Box12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


