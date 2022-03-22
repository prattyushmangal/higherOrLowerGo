package src_test

import (
	src "anush22/cardgame/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

var options = [4]string{"A", "B", "a", "b"}

func TestGetInput(t *testing.T) {
	funcOutput := src.GetInput()
	assert.Contains(t, options, funcOutput)
}
