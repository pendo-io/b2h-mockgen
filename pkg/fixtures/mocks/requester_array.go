// Code generated by B2H-MockGen v0.0.0-dev. EDIT AT YOUR OWN PERIL.

package mocks

import mock "github.com/stretchr/testify/mock"

type RequesterArrayMock struct {
	mock.Mock
}

func (_m *RequesterArrayMock) Get(path string) ([2]string, error) {
	args := _m.Called(path)
	return args.Get(0).([2]string), args.Error(1)
}
