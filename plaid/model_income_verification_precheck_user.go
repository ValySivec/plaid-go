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

// IncomeVerificationPrecheckUser struct for IncomeVerificationPrecheckUser
type IncomeVerificationPrecheckUser struct {
	// The user's first name
	FirstName NullableString `json:"first_name,omitempty"`
	// The user's last name
	LastName NullableString `json:"last_name,omitempty"`
	// The user's email address
	EmailAddress NullableString `json:"email_address,omitempty"`
	HomeAddress NullableSignalAddressData `json:"home_address,omitempty"`
}

// NewIncomeVerificationPrecheckUser instantiates a new IncomeVerificationPrecheckUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPrecheckUser() *IncomeVerificationPrecheckUser {
	this := IncomeVerificationPrecheckUser{}
	return &this
}

// NewIncomeVerificationPrecheckUserWithDefaults instantiates a new IncomeVerificationPrecheckUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPrecheckUserWithDefaults() *IncomeVerificationPrecheckUser {
	this := IncomeVerificationPrecheckUser{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckUser) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckUser) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckUser) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *IncomeVerificationPrecheckUser) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *IncomeVerificationPrecheckUser) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *IncomeVerificationPrecheckUser) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckUser) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckUser) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckUser) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *IncomeVerificationPrecheckUser) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *IncomeVerificationPrecheckUser) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *IncomeVerificationPrecheckUser) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckUser) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckUser) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckUser) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *IncomeVerificationPrecheckUser) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *IncomeVerificationPrecheckUser) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *IncomeVerificationPrecheckUser) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetHomeAddress returns the HomeAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckUser) GetHomeAddress() SignalAddressData {
	if o == nil || o.HomeAddress.Get() == nil {
		var ret SignalAddressData
		return ret
	}
	return *o.HomeAddress.Get()
}

// GetHomeAddressOk returns a tuple with the HomeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckUser) GetHomeAddressOk() (*SignalAddressData, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HomeAddress.Get(), o.HomeAddress.IsSet()
}

// HasHomeAddress returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckUser) HasHomeAddress() bool {
	if o != nil && o.HomeAddress.IsSet() {
		return true
	}

	return false
}

// SetHomeAddress gets a reference to the given NullableSignalAddressData and assigns it to the HomeAddress field.
func (o *IncomeVerificationPrecheckUser) SetHomeAddress(v SignalAddressData) {
	o.HomeAddress.Set(&v)
}
// SetHomeAddressNil sets the value for HomeAddress to be an explicit nil
func (o *IncomeVerificationPrecheckUser) SetHomeAddressNil() {
	o.HomeAddress.Set(nil)
}

// UnsetHomeAddress ensures that no value is present for HomeAddress, not even an explicit nil
func (o *IncomeVerificationPrecheckUser) UnsetHomeAddress() {
	o.HomeAddress.Unset()
}

func (o IncomeVerificationPrecheckUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if o.HomeAddress.IsSet() {
		toSerialize["home_address"] = o.HomeAddress.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationPrecheckUser struct {
	value *IncomeVerificationPrecheckUser
	isSet bool
}

func (v NullableIncomeVerificationPrecheckUser) Get() *IncomeVerificationPrecheckUser {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckUser) Set(val *IncomeVerificationPrecheckUser) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckUser) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckUser(val *IncomeVerificationPrecheckUser) *NullableIncomeVerificationPrecheckUser {
	return &NullableIncomeVerificationPrecheckUser{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


