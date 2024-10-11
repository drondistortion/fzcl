package main

import (
	"fmt"
	"testing"
	"time"
)

type Parameters struct {
	hours   int
	minutes int
	out     string
}

func (p Parameters) Name() string {
	return fmt.Sprintf("%d:%02d", p.hours, p.minutes)
}

func TestTime(t *testing.T) {
	var times = []Parameters{
		{0, 0, "midnight"},
		{23, 40, "quarter to midnight"},
		{0, 20, "quarter past midnight"},
		{12, 0, "twelve o'clock"},
		{12, 2, "twelve o'clock"},
		{5, 40, "quarter to six"},
		{11, 58, "twelve o'clock"},
		{3, 22, "quarter past three"},
		{3, 37, "half past three"},
		{1, 29, "half past one"},
		{1, 20, "quarter past one"},
	}
	for _, param := range times {
		t.Run(param.Name(), func(t *testing.T) {
			now, _ := time.Parse("3:04", param.Name())
			result := fuzzy(now, 2)
			if result != param.out {
				t.Errorf("got %s, want %s", result, param.out)
			}
		})
	}
}
