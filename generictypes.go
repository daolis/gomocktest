package test

import (
	_ "github.com/golang/mock/gomock"
)

type TestType string

type TestGenericType[T any] struct {
	blaa T
}

//go:generate mockgen -destination zz_generated_mock.go -package mockst . TestGenericInterface

type TestGenericInterface interface {
	TestMethod(TestGenericType[TestType]) TestGenericType[TestType]
}
