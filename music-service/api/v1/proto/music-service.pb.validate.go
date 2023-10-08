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

// Validate checks the field values on NewAlbumRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NewAlbumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewAlbumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NewAlbumRequestMultiError, or nil if none found.
func (m *NewAlbumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *NewAlbumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := NewAlbumRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTracks()) < 1 {
		err := NewAlbumRequestValidationError{
			field:  "Tracks",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetTracks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NewAlbumRequestValidationError{
						field:  fmt.Sprintf("Tracks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NewAlbumRequestValidationError{
						field:  fmt.Sprintf("Tracks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NewAlbumRequestValidationError{
					field:  fmt.Sprintf("Tracks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return NewAlbumRequestMultiError(errors)
	}

	return nil
}

// NewAlbumRequestMultiError is an error wrapping multiple validation errors
// returned by NewAlbumRequest.ValidateAll() if the designated constraints
// aren't met.
type NewAlbumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewAlbumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewAlbumRequestMultiError) AllErrors() []error { return m }

// NewAlbumRequestValidationError is the validation error returned by
// NewAlbumRequest.Validate if the designated constraints aren't met.
type NewAlbumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewAlbumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewAlbumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewAlbumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewAlbumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewAlbumRequestValidationError) ErrorName() string { return "NewAlbumRequestValidationError" }

// Error satisfies the builtin error interface
func (e NewAlbumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewAlbumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewAlbumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewAlbumRequestValidationError{}

// Validate checks the field values on NewAlbumResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NewAlbumResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewAlbumResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NewAlbumResponseMultiError, or nil if none found.
func (m *NewAlbumResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *NewAlbumResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Status

	if len(errors) > 0 {
		return NewAlbumResponseMultiError(errors)
	}

	return nil
}

// NewAlbumResponseMultiError is an error wrapping multiple validation errors
// returned by NewAlbumResponse.ValidateAll() if the designated constraints
// aren't met.
type NewAlbumResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewAlbumResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewAlbumResponseMultiError) AllErrors() []error { return m }

// NewAlbumResponseValidationError is the validation error returned by
// NewAlbumResponse.Validate if the designated constraints aren't met.
type NewAlbumResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewAlbumResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewAlbumResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewAlbumResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewAlbumResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewAlbumResponseValidationError) ErrorName() string { return "NewAlbumResponseValidationError" }

// Error satisfies the builtin error interface
func (e NewAlbumResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewAlbumResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewAlbumResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewAlbumResponseValidationError{}

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

// Validate checks the field values on DeleteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteResponseMultiError,
// or nil if none found.
func (m *DeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Views

	if len(errors) > 0 {
		return DeleteResponseMultiError(errors)
	}

	return nil
}

// DeleteResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResponseMultiError) AllErrors() []error { return m }

// DeleteResponseValidationError is the validation error returned by
// DeleteResponse.Validate if the designated constraints aren't met.
type DeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResponseValidationError) ErrorName() string { return "DeleteResponseValidationError" }

// Error satisfies the builtin error interface
func (e DeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResponseValidationError{}

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteRequestMultiError, or
// nil if none found.
func (m *DeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for AuthorId

	if len(errors) > 0 {
		return DeleteRequestMultiError(errors)
	}

	return nil
}

// DeleteRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRequestMultiError) AllErrors() []error { return m }

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}
