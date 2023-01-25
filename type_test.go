package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("a", "b", "gson123")
	assert.Nil(t, err)
	log.Printf("%+v", acc)
}
