// +build gofuzz

package helpers

import (
	"github.com/trailofbits/go-fuzz-utils"
)

func Fuzz(data []byte) int {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return 0
	}

	payload, err := tp.GetString()
	if err != nil {
		return 0
	}

	if err := ParseSecretPass(payload); err != nil {
		return 0
	}
	return 1
}
