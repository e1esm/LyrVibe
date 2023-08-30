// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/proto/music-service.proto

package proto

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

// Validate checks the field values on NewTrackRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NewTrackRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewTrackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NewTrackRequestMultiError, or nil if none found.
func (m *NewTrackRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *NewTrackRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetCover()) < 1 {
		err := NewTrackRequestValidationError{
			field:  "Cover",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 1 || l > 50 {
		err := NewTrackRequestValidationError{
			field:  "Title",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_NewTrackRequest_Title_Pattern.MatchString(m.GetTitle()) {
		err := NewTrackRequestValidationError{
			field:  "Title",
			reason: "value does not match regex pattern \"(?i)^[A-Za-z0-9]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ReleaseDate

	if _, ok := Genre_name[int32(m.GetGenre())]; !ok {
		err := NewTrackRequestValidationError{
			field:  "Genre",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_NewTrackRequest_Duration_Pattern.MatchString(m.GetDuration()) {
		err := NewTrackRequestValidationError{
			field:  "Duration",
			reason: "value does not match regex pattern \"^([0-9]+[mM])?([0-9]+[sS])?$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetCountry()); l < 1 || l > 50 {
		err := NewTrackRequestValidationError{
			field:  "Country",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_NewTrackRequest_Country_Pattern.MatchString(m.GetCountry()) {
		err := NewTrackRequestValidationError{
			field:  "Country",
			reason: "value does not match regex pattern \"(?i)^[A-Za-z]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for VideoLink

	// no validation rules for ArtistId

	if len(errors) > 0 {
		return NewTrackRequestMultiError(errors)
	}

	return nil
}

// NewTrackRequestMultiError is an error wrapping multiple validation errors
// returned by NewTrackRequest.ValidateAll() if the designated constraints
// aren't met.
type NewTrackRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewTrackRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewTrackRequestMultiError) AllErrors() []error { return m }

// NewTrackRequestValidationError is the validation error returned by
// NewTrackRequest.Validate if the designated constraints aren't met.
type NewTrackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewTrackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewTrackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewTrackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewTrackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewTrackRequestValidationError) ErrorName() string { return "NewTrackRequestValidationError" }

// Error satisfies the builtin error interface
func (e NewTrackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewTrackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewTrackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewTrackRequestValidationError{}

var _NewTrackRequest_Title_Pattern = regexp.MustCompile("(?i)^[A-Za-z0-9]+$")

var _NewTrackRequest_Duration_Pattern = regexp.MustCompile("^([0-9]+[mM])?([0-9]+[sS])?$")

var _NewTrackRequest_Country_Pattern = regexp.MustCompile("(?i)^[A-Za-z]+$")

// Validate checks the field values on NewTrackResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NewTrackResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewTrackResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NewTrackResponseMultiError, or nil if none found.
func (m *NewTrackResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *NewTrackResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Status

	if len(errors) > 0 {
		return NewTrackResponseMultiError(errors)
	}

	return nil
}

// NewTrackResponseMultiError is an error wrapping multiple validation errors
// returned by NewTrackResponse.ValidateAll() if the designated constraints
// aren't met.
type NewTrackResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewTrackResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewTrackResponseMultiError) AllErrors() []error { return m }

// NewTrackResponseValidationError is the validation error returned by
// NewTrackResponse.Validate if the designated constraints aren't met.
type NewTrackResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewTrackResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewTrackResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewTrackResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewTrackResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewTrackResponseValidationError) ErrorName() string { return "NewTrackResponseValidationError" }

// Error satisfies the builtin error interface
func (e NewTrackResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewTrackResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewTrackResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewTrackResponseValidationError{}

// Validate checks the field values on RequestStatus with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RequestStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestStatus with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RequestStatusMultiError, or
// nil if none found.
func (m *RequestStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return RequestStatusMultiError(errors)
	}

	return nil
}

// RequestStatusMultiError is an error wrapping multiple validation errors
// returned by RequestStatus.ValidateAll() if the designated constraints
// aren't met.
type RequestStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestStatusMultiError) AllErrors() []error { return m }

// RequestStatusValidationError is the validation error returned by
// RequestStatus.Validate if the designated constraints aren't met.
type RequestStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestStatusValidationError) ErrorName() string { return "RequestStatusValidationError" }

// Error satisfies the builtin error interface
func (e RequestStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestStatusValidationError{}
