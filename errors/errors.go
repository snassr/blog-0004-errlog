// Package errors defines the error handling used by all blog-004-errlog software.
package errors

import (
	"encoding/json"
)

// Op describes an operation, by package and method.
// example: usecases/fi.GetAvailablelocations
type Op string

// Kind defines the kind of error.
type Kind uint8

// Kinds of errors.
//
// The values of the error kinds are common between both
// clients and servers (order is important, list is append-only)
const (
	Other      Kind = iota // Unclassified error. this value is not printed in the error message.
	Invalid                // Invalid operation for this type of item.
	Permission             // Permission denied.
)

// String returns a string repesentation of the error kinds.
func (k Kind) String() string {
	switch k {
	case Other:
		return "other error"
	case Invalid:
		return "invalid operation"
	case Permission:
		return "permission denied"
	}
	return "unknown error kind"
}

// Error is the type that implements the error interface.
type Error struct {
	// Op is the operation being performed, usually the name of the method
	// being invoked.
	Op Op `json:"op"`
	// Kind is the class of error, such as permission failiure,
	// or "Other" if class is unknown or irrelevant.
	Kind Kind `json:"kind"`
	// Err is the underlying error that triggered this one, if any.
	Err string `json:"err"`
}

// JSON returns the error as a JSON object in bytes.
func (e Error) JSON() []byte {
	b, err := json.Marshal(e)
	if err != nil {
		b, _ := json.Marshal(Error{
			Op:   "monitor/errors.JSON",
			Kind: Other,
			Err:  err.Error(),
		})
		return b
	}
	return b
}
