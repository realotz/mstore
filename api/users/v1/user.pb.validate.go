// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/users/v1/user.proto

package userV1

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
)

// Validate checks the field values on ListUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Keyword

	// no validation rules for Role

	if v, ok := interface{}(m.GetOption()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListUserReqValidationError{
				field:  "Option",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListUserReqValidationError is the validation error returned by
// ListUserReq.Validate if the designated constraints aren't met.
type ListUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReqValidationError) ErrorName() string { return "ListUserReqValidationError" }

// Error satisfies the builtin error interface
func (e ListUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReqValidationError{}

// Validate checks the field values on UserListReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserListReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	return nil
}

// UserListReplyValidationError is the validation error returned by
// UserListReply.Validate if the designated constraints aren't met.
type UserListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListReplyValidationError) ErrorName() string { return "UserListReplyValidationError" }

// Error satisfies the builtin error interface
func (e UserListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListReplyValidationError{}

// Validate checks the field values on CreateUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Account

	// no validation rules for Name

	// no validation rules for Password

	// no validation rules for Role

	// no validation rules for Phone

	// no validation rules for Email

	return nil
}

// CreateUserReqValidationError is the validation error returned by
// CreateUserReq.Validate if the designated constraints aren't met.
type CreateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReqValidationError) ErrorName() string { return "CreateUserReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReqValidationError{}

// Validate checks the field values on UpdateUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Phone

	// no validation rules for Email

	return nil
}

// UpdateUserReqValidationError is the validation error returned by
// UpdateUserReq.Validate if the designated constraints aren't met.
type UpdateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserReqValidationError) ErrorName() string { return "UpdateUserReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserReqValidationError{}
