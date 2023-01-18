package address

import (
	"strings"
)

// Threats a valid address.
func TypeAddress(address string) string {
	validAddresses := []string{"rua", "avenida", "estrada", "rodovia"}

	lowerCaseAddress := strings.ToLower(address)
	firstWordAddress := strings.Split(lowerCaseAddress, " ")[0]

	addressHasAvalidType := false

	for _, currentType := range validAddresses {
		if currentType == firstWordAddress {
			addressHasAvalidType = true
		}
	}

	if addressHasAvalidType {
		return strings.Title(firstWordAddress)
	}

	return "Invalid type!"
}
