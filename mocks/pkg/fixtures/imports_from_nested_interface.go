// Code generated by B2H-MockGen v0.0.0-dev. EDIT AT YOUR OWN PERIL.

package mocks

import (
	http "net/http"

	fixtureshttp "github.com/pendo-io/b2h-mockgen/pkg/fixtures/http"

	mock "github.com/stretchr/testify/mock"
)

type HasConflictingNestedImportsMock struct {
	mock.Mock
}

func (_m *HasConflictingNestedImportsMock) Get(path string) (http.Response, error) {
	args := _m.Called(path)
	return args.Get(0).(http.Response), args.Error(1)
}
func (_m *HasConflictingNestedImportsMock) Z() fixtureshttp.MyStruct {
	args := _m.Called()
	return args.Get(0).(fixtureshttp.MyStruct)
}
