// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/proto/auth-service.proto

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

// Validate checks the field values on SignUpRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignUpRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignUpRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignUpRequestMultiError, or
// nil if none found.
func (m *SignUpRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignUpRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUsername()); l < 4 || l > 10 {
		err := SignUpRequestValidationError{
			field:  "Username",
			reason: "value length must be between 4 and 10 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_SignUpRequest_Username_Pattern.MatchString(m.GetUsername()) {
		err := SignUpRequestValidationError{
			field:  "Username",
			reason: "value does not match regex pattern \"(?i)^[A-Za-z0-9_]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 5 {
		err := SignUpRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _SignUpRequest_Role_InLookup[m.GetRole()]; !ok {
		err := SignUpRequestValidationError{
			field:  "Role",
			reason: "value must be in list [Guest]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Image

	if len(errors) > 0 {
		return SignUpRequestMultiError(errors)
	}

	return nil
}

// SignUpRequestMultiError is an error wrapping multiple validation errors
// returned by SignUpRequest.ValidateAll() if the designated constraints
// aren't met.
type SignUpRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignUpRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignUpRequestMultiError) AllErrors() []error { return m }

// SignUpRequestValidationError is the validation error returned by
// SignUpRequest.Validate if the designated constraints aren't met.
type SignUpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpRequestValidationError) ErrorName() string { return "SignUpRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignUpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpRequestValidationError{}

var _SignUpRequest_Username_Pattern = regexp.MustCompile("(?i)^[A-Za-z0-9_]+$")

var _SignUpRequest_Role_InLookup = map[string]struct{}{
	"Guest": {},
}

// Validate checks the field values on SignUpResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignUpResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignUpResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignUpResponseMultiError,
// or nil if none found.
func (m *SignUpResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SignUpResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignUpResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignUpResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignUpResponseValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Username

	if len(errors) > 0 {
		return SignUpResponseMultiError(errors)
	}

	return nil
}

// SignUpResponseMultiError is an error wrapping multiple validation errors
// returned by SignUpResponse.ValidateAll() if the designated constraints
// aren't met.
type SignUpResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignUpResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignUpResponseMultiError) AllErrors() []error { return m }

// SignUpResponseValidationError is the validation error returned by
// SignUpResponse.Validate if the designated constraints aren't met.
type SignUpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpResponseValidationError) ErrorName() string { return "SignUpResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignUpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpResponseValidationError{}

// Validate checks the field values on SignInRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInRequestMultiError, or
// nil if none found.
func (m *SignInRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUsername()); l < 4 || l > 10 {
		err := SignInRequestValidationError{
			field:  "Username",
			reason: "value length must be between 4 and 10 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_SignInRequest_Username_Pattern.MatchString(m.GetUsername()) {
		err := SignInRequestValidationError{
			field:  "Username",
			reason: "value does not match regex pattern \"(?i)^[A-Za-z0-9_]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 5 {
		err := SignInRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SignInRequestMultiError(errors)
	}

	return nil
}

// SignInRequestMultiError is an error wrapping multiple validation errors
// returned by SignInRequest.ValidateAll() if the designated constraints
// aren't met.
type SignInRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInRequestMultiError) AllErrors() []error { return m }

// SignInRequestValidationError is the validation error returned by
// SignInRequest.Validate if the designated constraints aren't met.
type SignInRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInRequestValidationError) ErrorName() string { return "SignInRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignInRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInRequestValidationError{}

var _SignInRequest_Username_Pattern = regexp.MustCompile("(?i)^[A-Za-z0-9_]+$")

// Validate checks the field values on SignInResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInResponseMultiError,
// or nil if none found.
func (m *SignInResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignInResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignInResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignInResponseValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTokens()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignInResponseValidationError{
					field:  "Tokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignInResponseValidationError{
					field:  "Tokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokens()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignInResponseValidationError{
				field:  "Tokens",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SignInResponseMultiError(errors)
	}

	return nil
}

// SignInResponseMultiError is an error wrapping multiple validation errors
// returned by SignInResponse.ValidateAll() if the designated constraints
// aren't met.
type SignInResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInResponseMultiError) AllErrors() []error { return m }

// SignInResponseValidationError is the validation error returned by
// SignInResponse.Validate if the designated constraints aren't met.
type SignInResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInResponseValidationError) ErrorName() string { return "SignInResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignInResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInResponseValidationError{}

// Validate checks the field values on CachedTokens with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CachedTokens) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CachedTokens with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CachedTokensMultiError, or
// nil if none found.
func (m *CachedTokens) ValidateAll() error {
	return m.validate(true)
}

func (m *CachedTokens) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for AccessTTL

	// no validation rules for RefreshToken

	// no validation rules for RefreshTTL

	if len(errors) > 0 {
		return CachedTokensMultiError(errors)
	}

	return nil
}

// CachedTokensMultiError is an error wrapping multiple validation errors
// returned by CachedTokens.ValidateAll() if the designated constraints aren't met.
type CachedTokensMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CachedTokensMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CachedTokensMultiError) AllErrors() []error { return m }

// CachedTokensValidationError is the validation error returned by
// CachedTokens.Validate if the designated constraints aren't met.
type CachedTokensValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CachedTokensValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CachedTokensValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CachedTokensValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CachedTokensValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CachedTokensValidationError) ErrorName() string { return "CachedTokensValidationError" }

// Error satisfies the builtin error interface
func (e CachedTokensValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCachedTokens.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CachedTokensValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CachedTokensValidationError{}

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

	// no validation rules for RequestStatus

	// no validation rules for ErrorMessage

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

// Validate checks the field values on LogoutRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutRequestMultiError, or
// nil if none found.
func (m *LogoutRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return LogoutRequestMultiError(errors)
	}

	return nil
}

// LogoutRequestMultiError is an error wrapping multiple validation errors
// returned by LogoutRequest.ValidateAll() if the designated constraints
// aren't met.
type LogoutRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutRequestMultiError) AllErrors() []error { return m }

// LogoutRequestValidationError is the validation error returned by
// LogoutRequest.Validate if the designated constraints aren't met.
type LogoutRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutRequestValidationError) ErrorName() string { return "LogoutRequestValidationError" }

// Error satisfies the builtin error interface
func (e LogoutRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutRequestValidationError{}

// Validate checks the field values on IdentifyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IdentifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdentifyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdentifyRequestMultiError, or nil if none found.
func (m *IdentifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IdentifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return IdentifyRequestMultiError(errors)
	}

	return nil
}

// IdentifyRequestMultiError is an error wrapping multiple validation errors
// returned by IdentifyRequest.ValidateAll() if the designated constraints
// aren't met.
type IdentifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdentifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdentifyRequestMultiError) AllErrors() []error { return m }

// IdentifyRequestValidationError is the validation error returned by
// IdentifyRequest.Validate if the designated constraints aren't met.
type IdentifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentifyRequestValidationError) ErrorName() string { return "IdentifyRequestValidationError" }

// Error satisfies the builtin error interface
func (e IdentifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentifyRequestValidationError{}

// Validate checks the field values on IdentifyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IdentifyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdentifyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdentifyResponseMultiError, or nil if none found.
func (m *IdentifyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IdentifyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for Role

	if len(errors) > 0 {
		return IdentifyResponseMultiError(errors)
	}

	return nil
}

// IdentifyResponseMultiError is an error wrapping multiple validation errors
// returned by IdentifyResponse.ValidateAll() if the designated constraints
// aren't met.
type IdentifyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdentifyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdentifyResponseMultiError) AllErrors() []error { return m }

// IdentifyResponseValidationError is the validation error returned by
// IdentifyResponse.Validate if the designated constraints aren't met.
type IdentifyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentifyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentifyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentifyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentifyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentifyResponseValidationError) ErrorName() string { return "IdentifyResponseValidationError" }

// Error satisfies the builtin error interface
func (e IdentifyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentifyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentifyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentifyResponseValidationError{}

// Validate checks the field values on UpdatingRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatingRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatingRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatingRoleRequestMultiError, or nil if none found.
func (m *UpdatingRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatingRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for RequestedRole

	if len(errors) > 0 {
		return UpdatingRoleRequestMultiError(errors)
	}

	return nil
}

// UpdatingRoleRequestMultiError is an error wrapping multiple validation
// errors returned by UpdatingRoleRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdatingRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatingRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatingRoleRequestMultiError) AllErrors() []error { return m }

// UpdatingRoleRequestValidationError is the validation error returned by
// UpdatingRoleRequest.Validate if the designated constraints aren't met.
type UpdatingRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatingRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatingRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatingRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatingRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatingRoleRequestValidationError) ErrorName() string {
	return "UpdatingRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatingRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatingRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatingRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatingRoleRequestValidationError{}

// Validate checks the field values on UpdatingRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatingRoleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatingRoleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatingRoleResponseMultiError, or nil if none found.
func (m *UpdatingRoleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatingRoleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return UpdatingRoleResponseMultiError(errors)
	}

	return nil
}

// UpdatingRoleResponseMultiError is an error wrapping multiple validation
// errors returned by UpdatingRoleResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdatingRoleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatingRoleResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatingRoleResponseMultiError) AllErrors() []error { return m }

// UpdatingRoleResponseValidationError is the validation error returned by
// UpdatingRoleResponse.Validate if the designated constraints aren't met.
type UpdatingRoleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatingRoleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatingRoleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatingRoleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatingRoleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatingRoleResponseValidationError) ErrorName() string {
	return "UpdatingRoleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatingRoleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatingRoleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatingRoleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatingRoleResponseValidationError{}
