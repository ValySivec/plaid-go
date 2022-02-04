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

// IncomeVerificationPrecheckEmployerAddressData Data about the components comprising an address.
type IncomeVerificationPrecheckEmployerAddressData struct {
	// The full city name
	City *string `json:"city,omitempty"`
	// The region or state. In API versions 2018-05-22 and earlier, this field is called `state`. Example: `\"NC\"`
	Region *string `json:"region,omitempty"`
	// The full street address Example: `\"564 Main Street, APT 15\"`
	Street *string `json:"street,omitempty"`
	// The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.
	PostalCode *string `json:"postal_code,omitempty"`
	// The ISO 3166-1 alpha-2 country code
	Country *string `json:"country,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationPrecheckEmployerAddressData IncomeVerificationPrecheckEmployerAddressData

// NewIncomeVerificationPrecheckEmployerAddressData instantiates a new IncomeVerificationPrecheckEmployerAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPrecheckEmployerAddressData() *IncomeVerificationPrecheckEmployerAddressData {
	this := IncomeVerificationPrecheckEmployerAddressData{}
	return &this
}

// NewIncomeVerificationPrecheckEmployerAddressDataWithDefaults instantiates a new IncomeVerificationPrecheckEmployerAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPrecheckEmployerAddressDataWithDefaults() *IncomeVerificationPrecheckEmployerAddressData {
	this := IncomeVerificationPrecheckEmployerAddressData{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *IncomeVerificationPrecheckEmployerAddressData) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *IncomeVerificationPrecheckEmployerAddressData) SetRegion(v string) {
	o.Region = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *IncomeVerificationPrecheckEmployerAddressData) SetStreet(v string) {
	o.Street = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *IncomeVerificationPrecheckEmployerAddressData) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddressData) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *IncomeVerificationPrecheckEmployerAddressData) SetCountry(v string) {
	o.Country = &v
}

func (o IncomeVerificationPrecheckEmployerAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationPrecheckEmployerAddressData) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationPrecheckEmployerAddressData := _IncomeVerificationPrecheckEmployerAddressData{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationPrecheckEmployerAddressData); err == nil {
		*o = IncomeVerificationPrecheckEmployerAddressData(varIncomeVerificationPrecheckEmployerAddressData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationPrecheckEmployerAddressData struct {
	value *IncomeVerificationPrecheckEmployerAddressData
	isSet bool
}

func (v NullableIncomeVerificationPrecheckEmployerAddressData) Get() *IncomeVerificationPrecheckEmployerAddressData {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckEmployerAddressData) Set(val *IncomeVerificationPrecheckEmployerAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckEmployerAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckEmployerAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckEmployerAddressData(val *IncomeVerificationPrecheckEmployerAddressData) *NullableIncomeVerificationPrecheckEmployerAddressData {
	return &NullableIncomeVerificationPrecheckEmployerAddressData{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckEmployerAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckEmployerAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


