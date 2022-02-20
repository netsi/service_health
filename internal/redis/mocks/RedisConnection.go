// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	redis "github.com/go-redis/redis/v8"
)

// RedisConnection is an autogenerated mock type for the RedisConnection type
type RedisConnection struct {
	mock.Mock
}

// Do provides a mock function with given fields: ctx, args
func (_m *RedisConnection) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *redis.Cmd
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) *redis.Cmd); ok {
		r0 = rf(ctx, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.Cmd)
		}
	}

	return r0
}