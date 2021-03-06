// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	entity "github.com/gifff/xenelectronic/entity"
	mock "github.com/stretchr/testify/mock"
)

// CategoryRepository is an autogenerated mock type for the CategoryRepository type
type CategoryRepository struct {
	mock.Mock
}

// ListAll provides a mock function with given fields:
func (_m *CategoryRepository) ListAll() ([]entity.Category, error) {
	ret := _m.Called()

	var r0 []entity.Category
	if rf, ok := ret.Get(0).(func() []entity.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProductsByCategoryID provides a mock function with given fields: categoryID, since, limit
func (_m *CategoryRepository) ListProductsByCategoryID(categoryID int64, since int64, limit int32) ([]entity.Product, error) {
	ret := _m.Called(categoryID, since, limit)

	var r0 []entity.Product
	if rf, ok := ret.Get(0).(func(int64, int64, int32) []entity.Product); ok {
		r0 = rf(categoryID, since, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int32) error); ok {
		r1 = rf(categoryID, since, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
