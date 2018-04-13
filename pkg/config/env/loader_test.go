package env

import (
	"os"
	"testing"
)

type aTest struct {
	Path string `env:"PATH"`
}

func TestInjectWithEnv(t *testing.T) {
	a := new(aTest)
	if err := InjectWithEnv(a); err != nil {
		t.Error(err)
	}

	if path := os.Getenv("PATH"); a.Path != path {
		t.Errorf("Fail to inject Path from env. Excepted: \"%s\", Receive: \"%s\"", path, a.Path)
	}
}
