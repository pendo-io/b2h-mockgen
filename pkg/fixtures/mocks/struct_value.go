// Code generated by B2H-MockGen v0.0.0-dev. EDIT AT YOUR OWN PERIL.

package mocks

import (
	test "github.com/pendo-io/b2h-mockgen/pkg/fixtures"
	mock "github.com/stretchr/testify/mock"
)

type AMock struct {
	mock.Mock
}

func (_m *AMock) Call() (test.B, error) {
	args := _m.Called()
	return args.Get(0).(test.B), args.Error(1)
}
