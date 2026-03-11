package timex_test

import (
	"testing"
	"time"

	"gojuniper/internal/timex"
)

func TestParseRFC3339(t *testing.T) {
	_, err := timex.ParseRFC3339("not-a-time")
	if err == nil {
		t.Fatalf("expected error")
	}

	got, err := timex.ParseRFC3339("2026-03-11T12:34:56Z")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.UTC().Format(time.RFC3339) != "2026-03-11T12:34:56Z" {
		t.Fatalf("unexpected time: %s", got.UTC().Format(time.RFC3339))
	}
}

func TestStartOfDay(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	in := time.Date(2026, 3, 11, 12, 34, 56, 0, loc)
	got := timex.StartOfDay(in)
	want := time.Date(2026, 3, 11, 0, 0, 0, 0, loc)
	if !got.Equal(want) {
		t.Fatalf("got=%s want=%s", got, want)
	}
}
