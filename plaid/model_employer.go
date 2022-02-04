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

// Employer Data about the employer.
type Employer struct {
	// Plaid's unique identifier for the employer.
	EmployerId string `json:"employer_id"`
	// The name of the employer
	Name string `json:"name"`
	Address NullableAddressDataNullable `json:"address"`
	// A number from 0 to 1 indicating Plaid's level of confidence in the pairing between the employer and the institution (not yet implemented).
	ConfidenceScore float32 `json:"confidence_score"`
	AdditionalProperties map[string]interface{}
}

type _Employer Employer

// NewEmployer instantiates a new Employer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployer(employerId string, name string, address NullableAddressDataNullable, confidenceScore float32) *Employer {
	this := Employer{}
	this.EmployerId = employerId
	this.Name = name
	this.Address = address
	this.ConfidenceScore = confidenceScore
	return &this
}

// NewEmployerWithDefaults instantiates a new Employer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployerWithDefaults() *Employer {
	this := Employer{}
	return &this
}

// GetEmployerId returns the EmployerId field value
func (o *Employer) GetEmployerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployerId
}

// GetEmployerIdOk returns a tuple with the EmployerId field value
// and a boolean to check if the value has been set.
func (o *Employer) GetEmployerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployerId, true
}

// SetEmployerId sets field value
func (o *Employer) SetEmployerId(v string) {
	o.EmployerId = v
}

// GetName returns the Name field value
func (o *Employer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Employer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Employer) SetName(v string) {
	o.Name = v
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for AddressDataNullable will be returned
func (o *Employer) GetAddress() AddressDataNullable {
	if o == nil || o.Address.Get() == nil {
		var ret AddressDataNullable
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employer) GetAddressOk() (*AddressDataNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *Employer) SetAddress(v AddressDataNullable) {
	o.Address.Set(&v)
}

// GetConfidenceScore returns the ConfidenceScore field value
func (o *Employer) GetConfidenceScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ConfidenceScore
}

// GetConfidenceScoreOk returns a tuple with the ConfidenceScore field value
// and a boolean to check if the value has been set.
func (o *Employer) GetConfidenceScoreOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfidenceScore, true
}

// SetConfidenceScore sets field value
func (o *Employer) SetConfidenceScore(v float32) {
	o.ConfidenceScore = v
}

func (o Employer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["employer_id"] = o.EmployerId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["address"] = o.Address.Get()
	}
	if true {
		toSerialize["confidence_score"] = o.ConfidenceScore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Employer) UnmarshalJSON(bytes []byte) (err error) {
	varEmployer := _Employer{}

	if err = json.Unmarshal(bytes, &varEmployer); err == nil {
		*o = Employer(varEmployer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "employer_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "address")
		delete(additionalProperties, "confidence_score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmployer struct {
	value *Employer
	isSet bool
}

func (v NullableEmployer) Get() *Employer {
	return v.value
}

func (v *NullableEmployer) Set(val *Employer) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployer) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployer(val *Employer) *NullableEmployer {
	return &NullableEmployer{value: val, isSet: true}
}

func (v NullableEmployer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


