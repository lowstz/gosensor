package main

import (
	"testing"
)

func TestShowCallerName(t *testing.T) {
	testData := []struct {
		name      string
		isCorrect bool
	}{
		{"github.com/lowstz/gosensor.TestShowCallerName", true},
		{"TooYoungTooSimple", false},
	}

	for _, td := range testData {
		if (showCallerName() == td.name) != td.isCorrect {
			t.Error(showCallerName())
			t.Errorf("%s showCallerName wrong", td.name)
		}
	}
}
