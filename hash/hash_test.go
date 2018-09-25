package hash

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		inp      string
		expected string
	}{
		{
			"angryMonkey",
			"ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==",
		},
		{
			"",
			"z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==",
		},
	}

	hash := &HashSha512{}

	for _, tt := range tests {
		result := hash.Calculate(tt.inp)
		if result != tt.expected {
			t.Errorf("expected '%s', got '%s'", tt.expected, result)
		}

	}
}
