package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	ty := &TypeConv{}
	assert.False(t, ty.Bool(""))
	assert.False(t, ty.Bool("asdf"))
	assert.False(t, ty.Bool("1234"))
	assert.False(t, ty.Bool("False"))
	assert.False(t, ty.Bool("0"))
	assert.False(t, ty.Bool("false"))
	assert.False(t, ty.Bool("F"))
	assert.False(t, ty.Bool("f"))
	assert.True(t, ty.Bool("true"))
	assert.True(t, ty.Bool("True"))
	assert.True(t, ty.Bool("t"))
	assert.True(t, ty.Bool("T"))
	assert.True(t, ty.Bool("1"))
}

func TestJSON(t *testing.T) {
	ty := new(TypeConv)
	expected := make(map[string]interface{})
	expected["foo"] = "bar"
	expected["one"] = 1.0
	expected["true"] = true

	actual := ty.JSON(`{"foo":"bar","one":1.0,"true":true}`)
	assert.Equal(t, expected["foo"], actual["foo"])
	assert.Equal(t, expected["one"], actual["one"])
	assert.Equal(t, expected["true"], actual["true"])
}
