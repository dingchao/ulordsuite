// Copyright (c) 2014 The ulordsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ulordjson_test

import (
	"testing"

	"github.com/ulordsuite/ulord/ulordjson"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   ulordjson.ErrorCode
		want string
	}{
		{ulordjson.ErrDuplicateMethod, "ErrDuplicateMethod"},
		{ulordjson.ErrInvalidUsageFlags, "ErrInvalidUsageFlags"},
		{ulordjson.ErrInvalidType, "ErrInvalidType"},
		{ulordjson.ErrEmbeddedType, "ErrEmbeddedType"},
		{ulordjson.ErrUnexportedField, "ErrUnexportedField"},
		{ulordjson.ErrUnsupportedFieldType, "ErrUnsupportedFieldType"},
		{ulordjson.ErrNonOptionalField, "ErrNonOptionalField"},
		{ulordjson.ErrNonOptionalDefault, "ErrNonOptionalDefault"},
		{ulordjson.ErrMismatchedDefault, "ErrMismatchedDefault"},
		{ulordjson.ErrUnregisteredMethod, "ErrUnregisteredMethod"},
		{ulordjson.ErrNumParams, "ErrNumParams"},
		{ulordjson.ErrMissingDescription, "ErrMissingDescription"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	// Detect additional error codes that don't have the stringer added.
	if len(tests)-1 != int(ulordjson.TstNumErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestError tests the error output for the Error type.
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   ulordjson.Error
		want string
	}{
		{
			ulordjson.Error{Description: "some error"},
			"some error",
		},
		{
			ulordjson.Error{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
