package maths_test

import (
	"github.com/aceakash/go-tdd-template/maths"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestName(t *testing.T) {
	assert.Equal(t, 3, maths.Add(1, 2))
}
