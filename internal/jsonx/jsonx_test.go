package jsonx_test

import (
	"strings"
	"testing"

	"gojuniper/internal/jsonx"
)

// 这组测试验证 JSON 编解码的两个常见点：
// - omitempty：空值字段不输出
// - roundtrip：Encode 后再 Decode 能还原出预期结构体
func TestEncodeDecodePerson(t *testing.T) {
	t.Run("omitempty", func(t *testing.T) {
		b, err := jsonx.EncodePerson(jsonx.Person{ID: 1, Name: "alice"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		s := string(b)
		if strings.Contains(s, "nickname") {
			t.Fatalf("expected nickname omitted, got %s", s)
		}
	})

	t.Run("roundtrip", func(t *testing.T) {
		nick := "ali"
		in := jsonx.Person{ID: 1, Name: "alice", Nickname: &nick}

		b, err := jsonx.EncodePerson(in)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		out, err := jsonx.DecodePerson(b)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if out.ID != in.ID || out.Name != in.Name || out.Nickname == nil || *out.Nickname != nick {
			t.Fatalf("unexpected decoded value: %#v", out)
		}
	})
}
