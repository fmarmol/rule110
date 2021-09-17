package main

import "testing"

func TestRules(t *testing.T) {
	testCases := []struct {
		index             uint8
		nextStateExpected uint8
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 0},
		{5, 1},
		{6, 1},
		{7, 0},
	}

	for _, testCase := range testCases {
		nextState := nextState(testCase.index)
		if nextState != testCase.nextStateExpected {
			t.Errorf("next state for index %v expected %v got %v", testCase.index, testCase.nextStateExpected, nextState)
		}
	}
}

func TestCalculateIndext(t *testing.T) {
	testCases := []struct {
		prev, current, next uint8
		indexExpected       uint8
	}{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{0, 1, 0, 2},
		{0, 1, 1, 3},
		{1, 0, 0, 4},
		{1, 0, 1, 5},
		{1, 1, 0, 6},
		{1, 1, 1, 7},
	}

	for _, testCase := range testCases {
		index := calculateIndex(testCase.prev, testCase.current, testCase.next)
		if index != testCase.indexExpected {
			t.Errorf("state [%v%v%v] expected to be at index %v got %v", testCase.prev, testCase.current, testCase.next, testCase.indexExpected, index)
		}
	}
}
