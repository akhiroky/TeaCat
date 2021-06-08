package main

import "testing"

func Test_main(t *testing.T) {
	status := goMain([]string{"teacat", "-h"})
	if status != 0 {
		t.Errorf("status code wont 0, but got %d", status)
	}
}