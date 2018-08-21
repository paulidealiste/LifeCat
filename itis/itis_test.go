package itis

import (
	"fmt"
	"testing"
)

func TestReadAndUnmarsh(t *testing.T) {
	tox := ReadAndUnmarsh("Umbra", "krameri")
	fmt.Println(tox)
}
