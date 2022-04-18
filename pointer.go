// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package pointer

// To returns a pointer to the provided value.
func To[T any](t T) *T {
	return &t
}

// SafeDeref returns the value from the passed pointer or zero value for a nil pointer.
func SafeDeref[T any](t *T) T { //nolint:ireturn
	if t == nil {
		var zero T

		return zero
	}

	return *t
}
