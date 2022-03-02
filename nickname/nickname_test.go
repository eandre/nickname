package nickname

import (
	"context"
	"testing"
)

func TestName(t *testing.T) {
	resp, err := World(context.Background(), "Jane Doe")
	if err != nil {
		t.Fatal(err)
	}
	want := "Hello, Jane Doe!"
	if got := resp.Message; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
