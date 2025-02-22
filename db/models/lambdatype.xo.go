package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// LambdaType is the 'lambda_type' enum type from schema 'public'.
type LambdaType uint16

// LambdaType values.
const (
	// LambdaTypeDuration is the 'duration' lambda_type.
	LambdaTypeDuration LambdaType = 1
	// LambdaTypeProvisioned is the 'provisioned' lambda_type.
	LambdaTypeProvisioned LambdaType = 2
	// LambdaTypeEdgeDuration is the 'edge-duration' lambda_type.
	LambdaTypeEdgeDuration LambdaType = 3
	// LambdaTypeEdgeRequest is the 'edge-request' lambda_type.
	LambdaTypeEdgeRequest LambdaType = 4
	// LambdaTypeRequests is the 'requests' lambda_type.
	LambdaTypeRequests LambdaType = 5
)

// String satisfies the [fmt.Stringer] interface.
func (lt LambdaType) String() string {
	switch lt {
	case LambdaTypeDuration:
		return "duration"
	case LambdaTypeProvisioned:
		return "provisioned"
	case LambdaTypeEdgeDuration:
		return "edge-duration"
	case LambdaTypeEdgeRequest:
		return "edge-request"
	case LambdaTypeRequests:
		return "requests"
	}
	return fmt.Sprintf("LambdaType(%d)", lt)
}

// MarshalText marshals [LambdaType] into text.
func (lt LambdaType) MarshalText() ([]byte, error) {
	return []byte(lt.String()), nil
}

// UnmarshalText unmarshals [LambdaType] from text.
func (lt *LambdaType) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "duration":
		*lt = LambdaTypeDuration
	case "provisioned":
		*lt = LambdaTypeProvisioned
	case "edge-duration":
		*lt = LambdaTypeEdgeDuration
	case "edge-request":
		*lt = LambdaTypeEdgeRequest
	case "requests":
		*lt = LambdaTypeRequests
	default:
		return ErrInvalidLambdaType(str)
	}
	return nil
}

// Value satisfies the [driver.Valuer] interface.
func (lt LambdaType) Value() (driver.Value, error) {
	return lt.String(), nil
}

// Scan satisfies the [sql.Scanner] interface.
func (lt *LambdaType) Scan(v interface{}) error {
	switch x := v.(type) {
	case []byte:
		return lt.UnmarshalText(x)
	case string:
		return lt.UnmarshalText([]byte(x))
	}
	return ErrInvalidLambdaType(fmt.Sprintf("%T", v))
}

// NullLambdaType represents a null 'lambda_type' enum for schema 'public'.
type NullLambdaType struct {
	LambdaType LambdaType
	// Valid is true if [LambdaType] is not null.
	Valid bool
}

// Value satisfies the [driver.Valuer] interface.
func (nlt NullLambdaType) Value() (driver.Value, error) {
	if !nlt.Valid {
		return nil, nil
	}
	return nlt.LambdaType.Value()
}

// Scan satisfies the [sql.Scanner] interface.
func (nlt *NullLambdaType) Scan(v interface{}) error {
	if v == nil {
		nlt.LambdaType, nlt.Valid = 0, false
		return nil
	}
	err := nlt.LambdaType.Scan(v)
	nlt.Valid = err == nil
	return err
}

// ErrInvalidLambdaType is the invalid [LambdaType] error.
type ErrInvalidLambdaType string

// Error satisfies the error interface.
func (err ErrInvalidLambdaType) Error() string {
	return fmt.Sprintf("invalid LambdaType(%s)", string(err))
}
