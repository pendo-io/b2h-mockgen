// Code generated by B2H-MockGen v0.0.0-dev. EDIT AT YOUR OWN PERIL.

package b2hmock

import mock "github.com/stretchr/testify/mock"

type Requester2Mock struct {
	mock.Mock
}

func (_m *Requester2Mock) Get(path string) error {
	args := _m.Called(path)
	return args.Error(0)
}
