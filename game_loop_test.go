package main

import (
	"strconv"
	"testing"
)

func TestCompatible(t *testing.T) {
	for i, tc := range []struct {
		guess     guess
		candidate string
		want      bool
	}{
		{
			guess:     guess{"LANES", "XXYXX"},
			candidate: "DRINK",
			want:      true,
		},
		{
			guess:     guess{"CRONY", "XGXGX"},
			candidate: "DRINK",
			want:      true,
		},
		{
			guess:     guess{"BUDGE", "XXYXX"},
			candidate: "DRINK",
			want:      true,
		},
		{
			guess:     guess{"SHALL", "XXXXX"},
			candidate: "DRINK",
			want:      true,
		},
		{
			guess:     guess{"ROATE", "XXYXY"},
			candidate: "ABBEY",
			want:      true,
		},
		{
			guess:     guess{"LENES", "XXXGX"},
			candidate: "ABBEY",
			want:      true,
		},
		{
			guess:     guess{"MEADS", "XYYXX"},
			candidate: "ABBEY",
			want:      true,
		},
		{
			guess:     guess{"HAPPY", "XYXXG"},
			candidate: "ABBEY",
			want:      true,
		},
		{
			guess:     guess{"ATOPY", "GXXXG"},
			candidate: "ABBEY",
			want:      true,
		},
		{
			guess:     guess{"ABBEY", "GGGGG"},
			candidate: "ABBEY",
			want:      true,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := compatible(tc.candidate, tc.guess)
			if got != tc.want {
				t.Logf("candidate=%q guess=%v", tc.candidate, tc.guess)
				t.Errorf("want=%t got=%t", tc.want, got)
			}
		})
	}
}
