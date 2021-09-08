/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.31.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferUserAddressInRequest The address associated with the account holder.
type TransferUserAddressInRequest struct {
	// The street number and name (i.e., \"100 Market St.\").
	Street *string `json:"street,omitempty"`
	// Ex. \"San Francisco\"
	City *string `json:"city,omitempty"`
	// The state or province (e.g., \"California\").
	Region *string `json:"region,omitempty"`
	// The postal code (e.g., \"94103\").
	PostalCode *string `json:"postal_code,omitempty"`
	// A two-letter country code (e.g., \"US\").
	Country *string `json:"country,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferUserAddressInRequest TransferUserAddressInRequest

// NewTransferUserAddressInRequest instantiates a new TransferUserAddressInRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferUserAddressInRequest() *TransferUserAddressInRequest {
	this := TransferUserAddressInRequest{}
	return &this
}

// NewTransferUserAddressInRequestWithDefaults instantiates a new TransferUserAddressInRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferUserAddressInRequestWithDefaults() *TransferUserAddressInRequest {
	this := TransferUserAddressInRequest{}
	return &this
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *TransferUserAddressInRequest) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserAddressInRequest) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *TransferUserAddressInRequest) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *TransferUserAddressInRequest) SetStreet(v string) {
	o.Street = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *TransferUserAddressInRequest) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserAddressInRequest) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *TransferUserAddressInRequest) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *TransferUserAddressInRequest) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *TransferUserAddressInRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserAddressInRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *TransferUserAddressInRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *TransferUserAddressInRequest) SetRegion(v string) {
	o.Region = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *TransferUserAddressInRequest) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserAddressInRequest) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *TransferUserAddressInRequest) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *TransferUserAddressInRequest) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *TransferUserAddressInRequest) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserAddressInRequest) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *TransferUserAddressInRequest) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *TransferUserAddressInRequest) SetCountry(v string) {
	o.Country = &v
}

func (o TransferUserAddressInRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
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

func (o *TransferUserAddressInRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTransferUserAddressInRequest := _TransferUserAddressInRequest{}

	if err = json.Unmarshal(bytes, &varTransferUserAddressInRequest); err == nil {
		*o = TransferUserAddressInRequest(varTransferUserAddressInRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "street")
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferUserAddressInRequest struct {
	value *TransferUserAddressInRequest
	isSet bool
}

func (v NullableTransferUserAddressInRequest) Get() *TransferUserAddressInRequest {
	return v.value
}

func (v *NullableTransferUserAddressInRequest) Set(val *TransferUserAddressInRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferUserAddressInRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferUserAddressInRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferUserAddressInRequest(val *TransferUserAddressInRequest) *NullableTransferUserAddressInRequest {
	return &NullableTransferUserAddressInRequest{value: val, isSet: true}
}

func (v NullableTransferUserAddressInRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferUserAddressInRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


