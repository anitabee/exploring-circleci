package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	x := 1
	y := 2
	swap(&x, &y)
	assert.EqualValues(t, x, 2)
	assert.EqualValues(t, y, 1)
}
