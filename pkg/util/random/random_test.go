package random

import "testing"

func TestInt64(t *testing.T) {
	tables := []struct {
		min int64
		max int64
	}{
		{1, 5},
		{2, 4},
		{3, 5},
		{4, 5},
	}

	for _, table := range tables {
		got := Int64(table.min, table.max)
		if got < table.min || got > table.max {
			t.Errorf("Incorrect, \nmin: %v,\n max: %v, \n  got: %v", table.min, table.max, got)
		}
	}
}
