package itis

import (
	"testing"
)

func TestPrintTaxon(t *testing.T) {
	tox := ReadAndUnmarsh("Umbra", "krameri")
	PrintTaxon(&tox)
}
