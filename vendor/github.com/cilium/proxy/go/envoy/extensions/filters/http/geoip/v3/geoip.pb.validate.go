// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/geoip/v3/geoip.proto

package geoipv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Geoip with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Geoip) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Geoip with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in GeoipMultiError, or nil if none found.
func (m *Geoip) ValidateAll() error {
	return m.validate(true)
}

func (m *Geoip) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetXffConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GeoipValidationError{
					field:  "XffConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GeoipValidationError{
					field:  "XffConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetXffConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GeoipValidationError{
				field:  "XffConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetProvider() == nil {
		err := GeoipValidationError{
			field:  "Provider",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetProvider()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GeoipValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GeoipValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProvider()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GeoipValidationError{
				field:  "Provider",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GeoipMultiError(errors)
	}

	return nil
}

// GeoipMultiError is an error wrapping multiple validation errors returned by
// Geoip.ValidateAll() if the designated constraints aren't met.
type GeoipMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GeoipMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GeoipMultiError) AllErrors() []error { return m }

// GeoipValidationError is the validation error returned by Geoip.Validate if
// the designated constraints aren't met.
type GeoipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GeoipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GeoipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GeoipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GeoipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GeoipValidationError) ErrorName() string { return "GeoipValidationError" }

// Error satisfies the builtin error interface
func (e GeoipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGeoip.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GeoipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GeoipValidationError{}

// Validate checks the field values on Geoip_XffConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Geoip_XffConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Geoip_XffConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Geoip_XffConfigMultiError, or nil if none found.
func (m *Geoip_XffConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *Geoip_XffConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for XffNumTrustedHops

	if len(errors) > 0 {
		return Geoip_XffConfigMultiError(errors)
	}

	return nil
}

// Geoip_XffConfigMultiError is an error wrapping multiple validation errors
// returned by Geoip_XffConfig.ValidateAll() if the designated constraints
// aren't met.
type Geoip_XffConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Geoip_XffConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Geoip_XffConfigMultiError) AllErrors() []error { return m }

// Geoip_XffConfigValidationError is the validation error returned by
// Geoip_XffConfig.Validate if the designated constraints aren't met.
type Geoip_XffConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Geoip_XffConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Geoip_XffConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Geoip_XffConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Geoip_XffConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Geoip_XffConfigValidationError) ErrorName() string { return "Geoip_XffConfigValidationError" }

// Error satisfies the builtin error interface
func (e Geoip_XffConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGeoip_XffConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Geoip_XffConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Geoip_XffConfigValidationError{}