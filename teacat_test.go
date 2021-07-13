package main

import "testing"

func TestMain(t *testing.T) {
	status := goMain([]string{"-h"})
	if status != 0 {
		t.Errorf("status code wont 0, but got %d", status)
	}
}
