package lifecatutil

import (
	"fmt"
	"testing"
)

func TestMockUnique(t *testing.T) {
	test := []string{"Apa", "Apa", "Napa", "Oca"}
	MockUnique(test)
}

func TestGetCWD(t *testing.T) {
	test, err := GetCWD()
	fmt.Println(test, err)
}

func TestEngLatTaxon(t *testing.T) {
	var test string
	var err error
	_, err = EngLatTaxon("monk")
	if err == nil {
		t.Error("Error propagation not correct!")
	}
	test, _ = EngLatTaxon("Class")
	if test != "Classis" {
		t.Error("Did not get correct translation!")
	}
}
