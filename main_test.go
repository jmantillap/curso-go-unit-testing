package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccess(t *testing.T) {
	result := Add(20, 2)

	expect := 22

	if result != expect {
		t.Errorf("got %q, expected %d", result, expect)
	}
}

func TestAddSuccess1(t *testing.T) {
	c := require.New(t)

	result := Add(20, 2)

	expect := 22

	c.Equal(expect, result)
}
