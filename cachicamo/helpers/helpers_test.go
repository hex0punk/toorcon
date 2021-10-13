package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanParseBadSecretPass(t *testing.T) {
	pass := "FIZZ:BUZZ"

	err := ParseSecretPass(pass)

	assert.NotNil(t, err)
}
