package itis

import (
	"testing"
)

func TestPrintTaxon(t *testing.T) {
	tox := ReadAndUnmarsh("Canis", "latrans")
	PrintTaxon(&tox)
}
