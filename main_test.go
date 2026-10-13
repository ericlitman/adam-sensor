package main

import "testing"

func TestGreeting(t *testing.T) {
	if got := greeting("Go"); got != "Hello, Go!" {
		t.Errorf("greeting(\"Go\") = %q, want %q", got, "Hello, Go!")
	}
}
