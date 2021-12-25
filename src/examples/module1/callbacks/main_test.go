package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a,b int)int {
	return a+b;
}

func TestAdd(t *testing.T){
	t.Log("starting test")
	result := add(1,2)
	assert.Equal(t, result, 3)

}