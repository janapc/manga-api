package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateANewId(t *testing.T) {
	id := NewID()
	assert.NotEmpty(t, id)
}

func TestShouldVerifyIfIdIsValid(t *testing.T) {
	id := NewID()
	assert.NotEmpty(t, id)
	_, err := ParseID("7fbe749c-d4c0-4962-bea0-514b02c7d5ab")
	assert.Nil(t, err)
}

func TestShouldVerifyIfIdIsInValid(t *testing.T) {
	id := NewID()
	assert.NotEmpty(t, id)
	_, err := ParseID("7fbe7")
	assert.NotNil(t, err)
}
