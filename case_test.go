package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestCamel(test *testing.T) {
	testcases := []struct {
		in       string
		expected string
	}{
		{"a b", "AB"},
		{"a_b", "AB"},
		{"a-b", "AB"},
	}
	for index, testcase := range testcases {
		_case := gox.Case(testcase.in)
		if got := _case.Camel(true).String(); testcase.expected != got {
			test.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, testcase.expected)
		}
	}
}
