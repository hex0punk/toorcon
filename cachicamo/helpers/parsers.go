package helpers

import (
	"fmt"
	"strings"
)

func ParseSecretPass(phrase string) error {
	phraseParts := strings.Split(phrase, ":")
	if len(phraseParts) < 3 {
		return fmt.Errorf("invalid phrase")
	}

	first := []byte(phraseParts[0])

	if len(first) >= 4 {
		if first[0] == 'F' && first[1] == 'I' && first[2] == 'Z' && first[3] == 'Z' {
			second := []byte(phraseParts[1])
			if len(second) >= 4 {
				if second[0] == 'F' && second[1] == 'U' && second[2] == 'Z' && second[3] == 'Z' {
					third := []byte(phraseParts[2])
					if len(third) >= 3 {
						if third[0] == 'T' && third[1] == 'O' && third[2] == 'O' && third[3] == 'R'{
							return nil
						}
					}
				}
			}
		}
	}
	return fmt.Errorf("invalid phrase")
}

