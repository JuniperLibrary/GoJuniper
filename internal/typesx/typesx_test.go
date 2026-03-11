package typesx_test

import (
	"errors"
	"testing"

	"gojuniper/internal/typesx"
)

func TestNewUser(t *testing.T) {
	t.Run("invalid id", func(t *testing.T) {
		_, err := typesx.NewUser(0, "alice")
		if !errors.Is(err, typesx.ErrInvalidID) {
			t.Fatalf("expected ErrInvalidID, got %v", err)
		}
	})

	t.Run("empty name", func(t *testing.T) {
		_, err := typesx.NewUser(1, "   ")
		if !errors.Is(err, typesx.ErrEmptyName) {
			t.Fatalf("expected ErrEmptyName, got %v", err)
		}
	})

	t.Run("trims name", func(t *testing.T) {
		u, err := typesx.NewUser(1, "  alice  ")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.Name != "alice" {
			t.Fatalf("name=%q, want %q", u.Name, "alice")
		}
	})
}

func TestUser_SetName(t *testing.T) {
	u, err := typesx.NewUser(1, "alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := u.SetName("  bob "); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Name != "bob" {
		t.Fatalf("name=%q, want %q", u.Name, "bob")
	}
}

func TestAdmin_Embedding(t *testing.T) {
	u, err := typesx.NewUser(1, "alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	a := typesx.Admin{User: u, Level: 10}

	// embedding 的效果：可以直接访问 a.Name / a.Greeting() 等。
	if a.Name != "alice" {
		t.Fatalf("name=%q, want %q", a.Name, "alice")
	}
	if !a.IsSuper() {
		t.Fatalf("expected super admin")
	}
}
