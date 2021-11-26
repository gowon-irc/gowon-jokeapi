package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	j = JSON{
		Contents: c,
	}

	c = Contents{
		Quotes: []Quote{q},
	}

	q = Quote{
		Quote:  "test quote",
		Author: "lorii",
	}

	s = "test quote\n- lorii"
)

func TestString(t *testing.T) {
	cases := []struct {
		name     string
		obj      fmt.Stringer
		expected string
	}{
		{
			name: "JSON",
			obj:  j,
		},
		{
			name: "Contents",
			obj:  c,
		},
		{
			name: "Quote",
			obj:  q,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.obj.String(), s)
		})
	}
}
