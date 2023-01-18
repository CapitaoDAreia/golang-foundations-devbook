package app

import (
	"reflect"
	"testing"

	CliPackage "github.com/urfave/cli"
)

// This is a test function to Generate in app.
func TestGenerate(t *testing.T) {

	expectedReturn := CliPackage.NewApp()
	actualReturn := Generate()

	if reflect.TypeOf(actualReturn) != reflect.TypeOf(expectedReturn) {
		t.Fail()
	}
}
