package pipetest

import "testing"

func TestHelloPipes(t *testing.T) {
	want := "Hello Pipes"

	got := helloPipes()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
