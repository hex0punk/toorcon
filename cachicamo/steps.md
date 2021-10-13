# Steps

## GCatch

```shell
export GO111MODULE=off
go get github.com/hex0punk/toorcon/cachicamo
cd ~/go/src/github.com/hex0punk/toorcon/cachicamo/
go get
export GOPATH=/home/vagrant/go/src
```

```shell
export GO111MODULE=on
cd ~/go/src/github.com/hex0punk/toorcon/cachicamo
GCatch -path=/home/vagrant/go/src/github.com/hex0punk/toorcon/cachicamo -include=github.com/hex0punk/toorcon/cachicamo -checker=BMOC -r -compile-error
```

## GoFuzz

Place go-fuzz code under `cachicamo/helpers`:

```shell script
go get github.com/trailofbits/go-fuzz-utils
go get github.com/dvyukov/go-fuzz/go-fuzz-dep
```

Then

```go
// +build gofuzz

package helpers

import (
	"github.com/trailofbits/go-fuzz-utils"
)

func Fuzz(data []byte) int {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return 0 // not enough data was supplied, exit gracefully for the next fuzzing iteration
	}

	payload, err := tp.GetString()
	if err != nil {
		return 0 // not enough data was supplied, exit gracefully for the next fuzzing iteration
	}

	if err := ParseSecretPass(payload); err != nil {
		return 0
	}
	return 1
}
```

Next:

```shell script
go-fuzz-build
go-fuzz -bin=helpers-fuzz.zip -workdir=workdir
```

Test:

```shell script
go test -v -run TestCanParseBadSecretPass
```

Add a new test to `helpers_test.go` with quoted crasher:

```go
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
```

Run:

```shell script
go test -v -run TestFuzz
```







