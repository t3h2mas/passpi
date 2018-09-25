package hash

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	seedText := "angryMonkey"
	expected := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="

	hash := &Hash{}

	result := hash.Calculate(seedText)

	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}
