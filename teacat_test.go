package main

import "testing"

func TestParseArgs(t *testing.T) {
	testdata := []struct {
		giveArgs   []string
		wontStatus int
	}{
		{[]string{"teacat", "-w", "teacat.go"}, 0},
		{[]string{"teacat", "-b", "teacat.go"}, 0},
		{[]string{"teacat", "-l", "teacat.go"}, 0},
		{[]string{"teacat", "-c", "teacat.go"}, 0},
		{[]string{"teacat", "-n", "teacat.go"}, 0},
		{[]string{"teacat", "-h"}, 0},
		{[]string{"teacat"}, 0}, // required parameters missing
		{[]string{"teacat", "-w", "teacat.txt"}, 0},
		{[]string{"teacat", "teacat.go"}, 0},
		{[]string{"teacat", "-p", "teacat.txt"}, 1},
		{[]string{"teacat", "-a"}, 1},
	}
	for _, td := range testdata {
		gotStatus := goMain(td.giveArgs)
		if gotStatus != td.wontStatus {
			t.Errorf("goMain(%v) status code did not match, wont %d, got %d", td.giveArgs, td.wontStatus, gotStatus)
		}
	}
}
