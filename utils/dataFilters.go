package utils

import (
	"github.com/klassmann/cpfcnpj"
	"strings"
)

// FilterNullString checks if the given string is equal to "NULL" and returns the string or empty string
func FilterNullString(s string) string {
	if s != "NULL" {
		return s
	}

	return ""
}

// FilterComma checks if the given string is equal to "NULL" and returns the values without commas or empty string
func FilterComma(s string) string {
	if s != "NULL" {
		return strings.ReplaceAll(s, ",", "")
	}

	return ""
}

// FilterAndValidateCNPJ checks if the given string is equal to "NULL" and returns the CNPJ checked value or empty string
func FilterAndValidateCNPJ(s string) string {
	cnpj := cpfcnpj.NewCNPJ(s)

	if cnpj.IsValid() {
		return string(cnpj)
	}

	return ""
}
