package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenJWT(t *testing.T) {

	serviceAuth := ServiceAuth{}

	token, err := serviceAuth.GenerateTokenJWT()

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, token)

}
