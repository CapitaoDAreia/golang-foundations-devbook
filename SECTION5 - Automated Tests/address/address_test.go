package address

import (
	"testing"
)

type testCases struct {
	actualCase     string
	expectedReturn string
}

func TestTypeAddress(t *testing.T) {

	testCases := []testCases{
		{"Avenida Inocêncio Seráfico", "Avenida"},
		{"Rua Altonia", "Rua"},
		{"Estrada da Fazendinha", "Estrada"},
		{"Lago da Fazendinha", "Invalid type!"},
	}

	for _, eachTestCase := range testCases {
		actualReturn := TypeAddress(eachTestCase.actualCase)

		if eachTestCase.expectedReturn != actualReturn {
			t.Error("Unexpected return!")
		}
	}
}
