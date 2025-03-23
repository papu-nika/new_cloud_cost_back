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
	// LambdaTypeEdgeDuration is the 'edge-duration' lambda_type.
	LambdaTypeEdgeDuration LambdaType = 3
	// LambdaTypeEdgeRequest is the 'edge-request' lambda_type.
	LambdaTypeEdgeRequest LambdaType = 4
	// LambdaTypeProcessed is the 'processed' lambda_type.
	LambdaTypeProcessed LambdaType = 5
	// LambdaTypeRequests is the 'requests' lambda_type.
	LambdaTypeRequests LambdaType = 6
	// LambdaTypeStorageDutation is the 'storage-dutation' lambda_type.
	LambdaTypeStorageDutation LambdaType = 7
	// LambdaTypeProvisionedConcurrency is the 'provisioned-concurrency' lambda_type.
	LambdaTypeProvisionedConcurrency LambdaType = 8
	// LambdaTypeProvisionedDuration is the 'provisioned-duration' lambda_type.
	LambdaTypeProvisionedDuration LambdaType = 2
)

// String satisfies the [fmt.Stringer] interface.
func (lt LambdaType) String() string {
	switch lt {
	case LambdaTypeDuration:
		return "duration"
	case LambdaTypeEdgeDuration:
		return "edge-duration"
	case LambdaTypeEdgeRequest:
		return "edge-request"
	case LambdaTypeProcessed:
		return "processed"
	case LambdaTypeRequests:
		return "requests"
	case LambdaTypeStorageDutation:
		return "storage-dutation"
	case LambdaTypeProvisionedConcurrency:
		return "provisioned-concurrency"
	case LambdaTypeProvisionedDuration:
		return "provisioned-duration"
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
	case "edge-duration":
		*lt = LambdaTypeEdgeDuration
	case "edge-request":
		*lt = LambdaTypeEdgeRequest
	case "processed":
		*lt = LambdaTypeProcessed
	case "requests":
		*lt = LambdaTypeRequests
	case "storage-dutation":
		*lt = LambdaTypeStorageDutation
	case "provisioned-concurrency":
		*lt = LambdaTypeProvisionedConcurrency
	case "provisioned-duration":
		*lt = LambdaTypeProvisionedDuration
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
