package stats

import (
	"testing"
	"time"
)

func TestMemoryJSON(t *testing.T) {
	stats := &Memory{
		RequestCount: 0,
		TotalTime:    0,
	}

	data, err := stats.JSON()

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	expectedDefault := `{"total":0,"average":0}`

	if string(data) != expectedDefault {
		t.Errorf("expected '%s', got '%s'", expectedDefault, string(data))
	}

	stats.AddPoint(2 * time.Second)
	stats.AddPoint(3 * time.Second)

	expectedAvg := `{"total":2,"average":2500000}`
	data, err = stats.JSON()

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if string(data) != expectedAvg {
		t.Errorf("expected '%s', got '%s'", expectedAvg, string(data))
	}
}

/*func TestCalculate(t *testing.T) {

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
*/
