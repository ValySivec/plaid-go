/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkTokenCreateRequestDepositSwitch Specifies options for initializing Link for use with the Deposit Switch (beta) product. This field is required if `deposit_switch` is included in the `products` array.
type LinkTokenCreateRequestDepositSwitch struct {
	// The `deposit_switch_id` provided by the `/deposit_switch/create` endpoint.
	DepositSwitchId string `json:"deposit_switch_id"`
}

// NewLinkTokenCreateRequestDepositSwitch instantiates a new LinkTokenCreateRequestDepositSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestDepositSwitch(depositSwitchId string) *LinkTokenCreateRequestDepositSwitch {
	this := LinkTokenCreateRequestDepositSwitch{}
	this.DepositSwitchId = depositSwitchId
	return &this
}

// NewLinkTokenCreateRequestDepositSwitchWithDefaults instantiates a new LinkTokenCreateRequestDepositSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestDepositSwitchWithDefaults() *LinkTokenCreateRequestDepositSwitch {
	this := LinkTokenCreateRequestDepositSwitch{}
	return &this
}

// GetDepositSwitchId returns the DepositSwitchId field value
func (o *LinkTokenCreateRequestDepositSwitch) GetDepositSwitchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositSwitchId
}

// GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestDepositSwitch) GetDepositSwitchIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositSwitchId, true
}

// SetDepositSwitchId sets field value
func (o *LinkTokenCreateRequestDepositSwitch) SetDepositSwitchId(v string) {
	o.DepositSwitchId = v
}

func (o LinkTokenCreateRequestDepositSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deposit_switch_id"] = o.DepositSwitchId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestDepositSwitch struct {
	value *LinkTokenCreateRequestDepositSwitch
	isSet bool
}

func (v NullableLinkTokenCreateRequestDepositSwitch) Get() *LinkTokenCreateRequestDepositSwitch {
	return v.value
}

func (v *NullableLinkTokenCreateRequestDepositSwitch) Set(val *LinkTokenCreateRequestDepositSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestDepositSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestDepositSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestDepositSwitch(val *LinkTokenCreateRequestDepositSwitch) *NullableLinkTokenCreateRequestDepositSwitch {
	return &NullableLinkTokenCreateRequestDepositSwitch{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestDepositSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestDepositSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


