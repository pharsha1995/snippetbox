package main

import (
	"testing"
	"time"

	"github.com/pharsha1995/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 6, 2, 22, 15, 0, 0, time.UTC),
			want: "02 Jun 2024 at 22:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "IST",
			tm:   time.Date(2024, 6, 2, 17, 47, 0, 0, time.FixedZone("IST", 11*30*60)),
			want: "02 Jun 2024 at 12:17",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := humanDate(test.tm)

			assert.Equal(t, got, test.want)
		})
	}
}
