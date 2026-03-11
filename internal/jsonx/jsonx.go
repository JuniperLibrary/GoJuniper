// Package jsonx 提供 JSON 主题的基础练习：
// - struct tag（json:"..."）
// - marshal / unmarshal
// - omitempty
package jsonx

import (
	"encoding/json"
)

// Person 演示 JSON 编解码时常见的字段形式。
type Person struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Nickname *string `json:"nickname,omitempty"`
}

// EncodePerson 将 Person 编码为 JSON bytes。
func EncodePerson(p Person) ([]byte, error) {
	return json.Marshal(p)
}

// DecodePerson 从 JSON bytes 解码出 Person。
func DecodePerson(b []byte) (Person, error) {
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		return Person{}, err
	}
	return p, nil
}
