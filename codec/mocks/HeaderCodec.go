// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	codec "goreplay/codec"

	mock "github.com/stretchr/testify/mock"
)

// HeaderCodec is an autogenerated mock type for the HeaderCodec type
type HeaderCodec struct {
	mock.Mock
}

// Decode provides a mock function with given fields: reqBuf, connectionID
func (_m *HeaderCodec) Decode(reqBuf []byte, connectionID string) (codec.ProtocolHeader, error) {
	ret := _m.Called(reqBuf, connectionID)

	var r0 codec.ProtocolHeader
	if rf, ok := ret.Get(0).(func([]byte, string) codec.ProtocolHeader); ok {
		r0 = rf(reqBuf, connectionID)
	} else {
		r0 = ret.Get(0).(codec.ProtocolHeader)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(reqBuf, connectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
