package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// Os is the 'os' enum type from schema 'public'.
type Os uint16

// Os values.
const (
	// OsLinux is the 'Linux' os.
	OsLinux Os = 1
	// OsWindows is the 'Windows' os.
	OsWindows Os = 2
)

// String satisfies the [fmt.Stringer] interface.
func (o Os) String() string {
	switch o {
	case OsLinux:
		return "Linux"
	case OsWindows:
		return "Windows"
	}
	return fmt.Sprintf("Os(%d)", o)
}

// MarshalText marshals [Os] into text.
func (o Os) MarshalText() ([]byte, error) {
	return []byte(o.String()), nil
}

// UnmarshalText unmarshals [Os] from text.
func (o *Os) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "Linux":
		*o = OsLinux
	case "Windows":
		*o = OsWindows
	default:
		return ErrInvalidOs(str)
	}
	return nil
}

// Value satisfies the [driver.Valuer] interface.
func (o Os) Value() (driver.Value, error) {
	return o.String(), nil
}

// Scan satisfies the [sql.Scanner] interface.
func (o *Os) Scan(v interface{}) error {
	switch x := v.(type) {
	case []byte:
		return o.UnmarshalText(x)
	case string:
		return o.UnmarshalText([]byte(x))
	}
	return ErrInvalidOs(fmt.Sprintf("%T", v))
}

// NullOs represents a null 'os' enum for schema 'public'.
type NullOs struct {
	Os Os
	// Valid is true if [Os] is not null.
	Valid bool
}

// Value satisfies the [driver.Valuer] interface.
func (no NullOs) Value() (driver.Value, error) {
	if !no.Valid {
		return nil, nil
	}
	return no.Os.Value()
}

// Scan satisfies the [sql.Scanner] interface.
func (no *NullOs) Scan(v interface{}) error {
	if v == nil {
		no.Os, no.Valid = 0, false
		return nil
	}
	err := no.Os.Scan(v)
	no.Valid = err == nil
	return err
}

// ErrInvalidOs is the invalid [Os] error.
type ErrInvalidOs string

// Error satisfies the error interface.
func (err ErrInvalidOs) Error() string {
	return fmt.Sprintf("invalid Os(%s)", string(err))
}
