// Code generated by B2H-MockGen v0.0.0-dev. EDIT AT YOUR OWN PERIL.

package mocks

import (
	test "github.com/pendo-io/b2h-mockgen/pkg/fixtures"
	mock "github.com/stretchr/testify/mock"
)

type UsesOtherPkgIfaceMock struct {
	mock.Mock
}

func (_m *UsesOtherPkgIfaceMock) DoSomethingElse(obj test.Sibling) {
	_m.Called(obj)
}
