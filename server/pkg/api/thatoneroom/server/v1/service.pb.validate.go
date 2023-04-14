// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: thatoneroom/server/v1/service.proto

package serverv1

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

// Validate checks the field values on ConnectRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConnectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConnectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConnectRequestMultiError,
// or nil if none found.
func (m *ConnectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConnectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ConnectRequestMultiError(errors)
	}

	return nil
}

// ConnectRequestMultiError is an error wrapping multiple validation errors
// returned by ConnectRequest.ValidateAll() if the designated constraints
// aren't met.
type ConnectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectRequestMultiError) AllErrors() []error { return m }

// ConnectRequestValidationError is the validation error returned by
// ConnectRequest.Validate if the designated constraints aren't met.
type ConnectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectRequestValidationError) ErrorName() string { return "ConnectRequestValidationError" }

// Error satisfies the builtin error interface
func (e ConnectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectRequestValidationError{}

// Validate checks the field values on ConnectResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConnectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConnectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConnectResponseMultiError, or nil if none found.
func (m *ConnectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ConnectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Token

	if len(errors) > 0 {
		return ConnectResponseMultiError(errors)
	}

	return nil
}

// ConnectResponseMultiError is an error wrapping multiple validation errors
// returned by ConnectResponse.ValidateAll() if the designated constraints
// aren't met.
type ConnectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectResponseMultiError) AllErrors() []error { return m }

// ConnectResponseValidationError is the validation error returned by
// ConnectResponse.Validate if the designated constraints aren't met.
type ConnectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectResponseValidationError) ErrorName() string { return "ConnectResponseValidationError" }

// Error satisfies the builtin error interface
func (e ConnectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectResponseValidationError{}

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RequestMultiError, or nil if none found.
func (m *Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Action.(type) {
	case *Request_Move:
		if v == nil {
			err := RequestValidationError{
				field:  "Action",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetMove()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RequestValidationError{
						field:  "Move",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RequestValidationError{
						field:  "Move",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMove()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestValidationError{
					field:  "Move",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return RequestMultiError(errors)
	}

	return nil
}

// RequestMultiError is an error wrapping multiple validation errors returned
// by Request.ValidateAll() if the designated constraints aren't met.
type RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestMultiError) AllErrors() []error { return m }

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on Response with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResponseMultiError, or nil
// if none found.
func (m *Response) ValidateAll() error {
	return m.validate(true)
}

func (m *Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Action.(type) {
	case *Response_UpdateEntity:
		if v == nil {
			err := ResponseValidationError{
				field:  "Action",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetUpdateEntity()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "UpdateEntity",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "UpdateEntity",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdateEntity()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  "UpdateEntity",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ResponseMultiError(errors)
	}

	return nil
}

// ResponseMultiError is an error wrapping multiple validation errors returned
// by Response.ValidateAll() if the designated constraints aren't met.
type ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResponseMultiError) AllErrors() []error { return m }

// ResponseValidationError is the validation error returned by
// Response.Validate if the designated constraints aren't met.
type ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseValidationError) ErrorName() string { return "ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseValidationError{}

// Validate checks the field values on Move with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Move) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Move with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MoveMultiError, or nil if none found.
func (m *Move) ValidateAll() error {
	return m.validate(true)
}

func (m *Move) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Direction

	if len(errors) > 0 {
		return MoveMultiError(errors)
	}

	return nil
}

// MoveMultiError is an error wrapping multiple validation errors returned by
// Move.ValidateAll() if the designated constraints aren't met.
type MoveMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MoveMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MoveMultiError) AllErrors() []error { return m }

// MoveValidationError is the validation error returned by Move.Validate if the
// designated constraints aren't met.
type MoveValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoveValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoveValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoveValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoveValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoveValidationError) ErrorName() string { return "MoveValidationError" }

// Error satisfies the builtin error interface
func (e MoveValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMove.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoveValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoveValidationError{}

// Validate checks the field values on Entity with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Entity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Entity with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EntityMultiError, or nil if none found.
func (m *Entity) ValidateAll() error {
	return m.validate(true)
}

func (m *Entity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Entity.(type) {
	case *Entity_Player:
		if v == nil {
			err := EntityValidationError{
				field:  "Entity",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetPlayer()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  "Player",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  "Player",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPlayer()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityValidationError{
					field:  "Player",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return EntityMultiError(errors)
	}

	return nil
}

// EntityMultiError is an error wrapping multiple validation errors returned by
// Entity.ValidateAll() if the designated constraints aren't met.
type EntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntityMultiError) AllErrors() []error { return m }

// EntityValidationError is the validation error returned by Entity.Validate if
// the designated constraints aren't met.
type EntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityValidationError) ErrorName() string { return "EntityValidationError" }

// Error satisfies the builtin error interface
func (e EntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityValidationError{}

// Validate checks the field values on Player with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Player) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Player with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PlayerMultiError, or nil if none found.
func (m *Player) ValidateAll() error {
	return m.validate(true)
}

func (m *Player) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetPosition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PlayerValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PlayerValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPosition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlayerValidationError{
				field:  "Position",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PlayerMultiError(errors)
	}

	return nil
}

// PlayerMultiError is an error wrapping multiple validation errors returned by
// Player.ValidateAll() if the designated constraints aren't met.
type PlayerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlayerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlayerMultiError) AllErrors() []error { return m }

// PlayerValidationError is the validation error returned by Player.Validate if
// the designated constraints aren't met.
type PlayerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlayerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlayerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlayerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlayerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlayerValidationError) ErrorName() string { return "PlayerValidationError" }

// Error satisfies the builtin error interface
func (e PlayerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlayer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlayerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlayerValidationError{}

// Validate checks the field values on Coordinate with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Coordinate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Coordinate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CoordinateMultiError, or
// nil if none found.
func (m *Coordinate) ValidateAll() error {
	return m.validate(true)
}

func (m *Coordinate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for X

	// no validation rules for Y

	if len(errors) > 0 {
		return CoordinateMultiError(errors)
	}

	return nil
}

// CoordinateMultiError is an error wrapping multiple validation errors
// returned by Coordinate.ValidateAll() if the designated constraints aren't met.
type CoordinateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CoordinateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CoordinateMultiError) AllErrors() []error { return m }

// CoordinateValidationError is the validation error returned by
// Coordinate.Validate if the designated constraints aren't met.
type CoordinateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CoordinateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CoordinateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CoordinateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CoordinateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CoordinateValidationError) ErrorName() string { return "CoordinateValidationError" }

// Error satisfies the builtin error interface
func (e CoordinateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCoordinate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CoordinateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CoordinateValidationError{}

// Validate checks the field values on UpdateEntity with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateEntity with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateEntityMultiError, or
// nil if none found.
func (m *UpdateEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEntity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateEntityValidationError{
					field:  "Entity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateEntityValidationError{
					field:  "Entity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEntityValidationError{
				field:  "Entity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateEntityMultiError(errors)
	}

	return nil
}

// UpdateEntityMultiError is an error wrapping multiple validation errors
// returned by UpdateEntity.ValidateAll() if the designated constraints aren't met.
type UpdateEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateEntityMultiError) AllErrors() []error { return m }

// UpdateEntityValidationError is the validation error returned by
// UpdateEntity.Validate if the designated constraints aren't met.
type UpdateEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEntityValidationError) ErrorName() string { return "UpdateEntityValidationError" }

// Error satisfies the builtin error interface
func (e UpdateEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEntityValidationError{}
