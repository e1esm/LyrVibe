// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/proto/artist-service.proto

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

// Validate checks the field values on VerificationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerificationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerificationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerificationRequestMultiError, or nil if none found.
func (m *VerificationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerificationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUsername()); l < 4 || l > 10 {
		err := VerificationRequestValidationError{
			field:  "Username",
			reason: "value length must be between 4 and 10 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_VerificationRequest_Username_Pattern.MatchString(m.GetUsername()) {
		err := VerificationRequestValidationError{
			field:  "Username",
			reason: "value does not match regex pattern \"(?i)^[A-Za-z0-9_]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCountry()) < 2 {
		err := VerificationRequestValidationError{
			field:  "Country",
			reason: "value length must be at least 2 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_VerificationRequest_Country_Pattern.MatchString(m.GetCountry()) {
		err := VerificationRequestValidationError{
			field:  "Country",
			reason: "value does not match regex pattern \"(?i)^[A-Z-a-z]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetFirstName()); l < 1 || l > 10 {
		err := VerificationRequestValidationError{
			field:  "FirstName",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_VerificationRequest_FirstName_Pattern.MatchString(m.GetFirstName()) {
		err := VerificationRequestValidationError{
			field:  "FirstName",
			reason: "value does not match regex pattern \"(?i)^[A-Za-z]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetSecondName()); l < 1 || l > 20 {
		err := VerificationRequestValidationError{
			field:  "SecondName",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_VerificationRequest_SecondName_Pattern.MatchString(m.GetSecondName()) {
		err := VerificationRequestValidationError{
			field:  "SecondName",
			reason: "value does not match regex pattern \"(?i)^[A-Za-z]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Id

	if len(errors) > 0 {
		return VerificationRequestMultiError(errors)
	}

	return nil
}

// VerificationRequestMultiError is an error wrapping multiple validation
// errors returned by VerificationRequest.ValidateAll() if the designated
// constraints aren't met.
type VerificationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerificationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerificationRequestMultiError) AllErrors() []error { return m }

// VerificationRequestValidationError is the validation error returned by
// VerificationRequest.Validate if the designated constraints aren't met.
type VerificationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerificationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerificationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerificationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerificationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerificationRequestValidationError) ErrorName() string {
	return "VerificationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerificationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerificationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerificationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerificationRequestValidationError{}

var _VerificationRequest_Username_Pattern = regexp.MustCompile("(?i)^[A-Za-z0-9_]+$")

var _VerificationRequest_Country_Pattern = regexp.MustCompile("(?i)^[A-Z-a-z]+$")

var _VerificationRequest_FirstName_Pattern = regexp.MustCompile("(?i)^[A-Za-z]+$")

var _VerificationRequest_SecondName_Pattern = regexp.MustCompile("(?i)^[A-Za-z]+$")

// Validate checks the field values on VerificationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerificationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerificationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerificationResponseMultiError, or nil if none found.
func (m *VerificationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VerificationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsVerified

	// no validation rules for RequestStatus

	if len(errors) > 0 {
		return VerificationResponseMultiError(errors)
	}

	return nil
}

// VerificationResponseMultiError is an error wrapping multiple validation
// errors returned by VerificationResponse.ValidateAll() if the designated
// constraints aren't met.
type VerificationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerificationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerificationResponseMultiError) AllErrors() []error { return m }

// VerificationResponseValidationError is the validation error returned by
// VerificationResponse.Validate if the designated constraints aren't met.
type VerificationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerificationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerificationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerificationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerificationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerificationResponseValidationError) ErrorName() string {
	return "VerificationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e VerificationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerificationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerificationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerificationResponseValidationError{}

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

	// no validation rules for RequestStatus

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

	// no validation rules for RequestStatus

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

// Validate checks the field values on DeleteTrackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTrackRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTrackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTrackRequestMultiError, or nil if none found.
func (m *DeleteTrackRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTrackRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AuthorId

	// no validation rules for TrackTitle

	if len(errors) > 0 {
		return DeleteTrackRequestMultiError(errors)
	}

	return nil
}

// DeleteTrackRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTrackRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTrackRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTrackRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTrackRequestMultiError) AllErrors() []error { return m }

// DeleteTrackRequestValidationError is the validation error returned by
// DeleteTrackRequest.Validate if the designated constraints aren't met.
type DeleteTrackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTrackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTrackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTrackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTrackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTrackRequestValidationError) ErrorName() string {
	return "DeleteTrackRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTrackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTrackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTrackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTrackRequestValidationError{}

// Validate checks the field values on DeleteTrackResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTrackResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTrackResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTrackResponseMultiError, or nil if none found.
func (m *DeleteTrackResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTrackResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for RequestStatus

	if len(errors) > 0 {
		return DeleteTrackResponseMultiError(errors)
	}

	return nil
}

// DeleteTrackResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteTrackResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteTrackResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTrackResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTrackResponseMultiError) AllErrors() []error { return m }

// DeleteTrackResponseValidationError is the validation error returned by
// DeleteTrackResponse.Validate if the designated constraints aren't met.
type DeleteTrackResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTrackResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTrackResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTrackResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTrackResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTrackResponseValidationError) ErrorName() string {
	return "DeleteTrackResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTrackResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTrackResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTrackResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTrackResponseValidationError{}
