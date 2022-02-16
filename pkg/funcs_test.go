package pkg

import (
	"testing"
)

func TestPadding(t *testing.T) {
	lcases := []struct {
		number   int
		zeros    int
		expected string
	}{
		{
			number:   1,
			zeros:    2,
			expected: "01",
		},
		{
			number:   195,
			zeros:    6,
			expected: "000195",
		},
	}
	for _, tcase := range lcases {
		if tcase.expected != leftPad(tcase.number, tcase.zeros) {
			t.Fatal()
		}
	}
}
