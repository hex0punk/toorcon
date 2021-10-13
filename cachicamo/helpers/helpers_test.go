package helpers

import (
	"fmt"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanParseBadSecretPass(t *testing.T) {
	pass := "FIZZ:BUZZ"

	err := ParseSecretPass(pass)

	assert.NotNil(t, err)
}


func TestFuzz(t *testing.T) {
	var crashers = []string{
		"AP4FI\x1bR0FIZZ:FUZZ0:TOO",
	}

	for _, crasher := range crashers {
		data := []byte(crasher)
		tp, _ := go_fuzz_utils.NewTypeProvider(data)
		payload, _ := tp.GetString()

		fmt.Println("Testing with string: ", payload)
		if err := ParseSecretPass(payload); err != nil {
			fmt.Println(err.Error())
		}
	}
}
