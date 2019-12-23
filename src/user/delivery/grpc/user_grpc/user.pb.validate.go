// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package user_grpc

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Email

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on CreateReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CreateReq) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 10 {
		return CreateReqValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return CreateReqValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if !_CreateReq_Password_Pattern.MatchString(m.GetPassword()) {
		return CreateReqValidationError{
			field:  "Password",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]{6,72}$\"",
		}
	}

	return nil
}

func (m *CreateReq) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateReq) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateReqValidationError is the validation error returned by
// CreateReq.Validate if the designated constraints aren't met.
type CreateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateReqValidationError) ErrorName() string { return "CreateReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReqValidationError{}

var _CreateReq_Password_Pattern = regexp.MustCompile("^[A-Za-z0-9]{6,72}$")

// Validate checks the field values on UpdateReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UpdateReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return UpdateReqValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 10 {
		return UpdateReqValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return UpdateReqValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if !_UpdateReq_Password_Pattern.MatchString(m.GetPassword()) {
		return UpdateReqValidationError{
			field:  "Password",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]{6,72}$\"",
		}
	}

	return nil
}

func (m *UpdateReq) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UpdateReq) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UpdateReqValidationError is the validation error returned by
// UpdateReq.Validate if the designated constraints aren't met.
type UpdateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateReqValidationError) ErrorName() string { return "UpdateReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateReqValidationError{}

var _UpdateReq_Password_Pattern = regexp.MustCompile("^[A-Za-z0-9]{6,72}$")

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LoginReq) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return LoginReqValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if !_LoginReq_Password_Pattern.MatchString(m.GetPassword()) {
		return LoginReqValidationError{
			field:  "Password",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]{6,72}$\"",
		}
	}

	return nil
}

func (m *LoginReq) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *LoginReq) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

var _LoginReq_Password_Pattern = regexp.MustCompile("^[A-Za-z0-9]{6,72}$")

// Validate checks the field values on DeleteRes with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DeleteRes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Deleted

	return nil
}

// DeleteResValidationError is the validation error returned by
// DeleteRes.Validate if the designated constraints aren't met.
type DeleteResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResValidationError) ErrorName() string { return "DeleteResValidationError" }

// Error satisfies the builtin error interface
func (e DeleteResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResValidationError{}

// Validate checks the field values on ID with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *ID) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return IDValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// IDValidationError is the validation error returned by ID.Validate if the
// designated constraints aren't met.
type IDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IDValidationError) ErrorName() string { return "IDValidationError" }

// Error satisfies the builtin error interface
func (e IDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IDValidationError{}

// Validate checks the field values on UserWithToken with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserWithToken) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserWithTokenValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTokenPair()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserWithTokenValidationError{
				field:  "TokenPair",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UserWithTokenValidationError is the validation error returned by
// UserWithToken.Validate if the designated constraints aren't met.
type UserWithTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserWithTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserWithTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserWithTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserWithTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserWithTokenValidationError) ErrorName() string { return "UserWithTokenValidationError" }

// Error satisfies the builtin error interface
func (e UserWithTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserWithToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserWithTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserWithTokenValidationError{}

// Validate checks the field values on RefreshToken with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RefreshToken) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RefreshToken

	return nil
}

// RefreshTokenValidationError is the validation error returned by
// RefreshToken.Validate if the designated constraints aren't met.
type RefreshTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenValidationError) ErrorName() string { return "RefreshTokenValidationError" }

// Error satisfies the builtin error interface
func (e RefreshTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenValidationError{}

// Validate checks the field values on TokenPair with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TokenPair) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IdToken

	// no validation rules for RefreshToken

	return nil
}

// TokenPairValidationError is the validation error returned by
// TokenPair.Validate if the designated constraints aren't met.
type TokenPairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenPairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenPairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenPairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenPairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenPairValidationError) ErrorName() string { return "TokenPairValidationError" }

// Error satisfies the builtin error interface
func (e TokenPairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokenPair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenPairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenPairValidationError{}
