package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	SetupTest()
	exitVal := m.Run()
	TeardownTest()
	os.Exit(exitVal)
}

func TestSignin(t *testing.T) {
	Signin()
}
