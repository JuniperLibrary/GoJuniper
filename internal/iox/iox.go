// Package iox 提供 I/O 主题的基础练习：
// - io.Reader / io.Writer 抽象
// - bufio 用于按行读取与缓冲写
// - io.ReadAll 读取全部内容
package iox

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// ReadAllString 读取 r 的全部内容并返回 string。
func ReadAllString(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReadLines 从 r 读取所有行。
// 注意：Scanner 默认 token 上限较小，这里显式提高 buffer，适合“学习场景”的输入。
func ReadLines(r io.Reader) ([]string, error) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// WriteLines 将每一行写入 w，并用 '\n' 分隔。
func WriteLines(w io.Writer, lines []string) error {
	bw := bufio.NewWriter(w)
	for _, line := range lines {
		if _, err := fmt.Fprintln(bw, line); err != nil {
			return err
		}
	}
	return bw.Flush()
}

// JoinLines 是一个小工具：把 lines 用 '\n' 拼起来（用于测试与示例）。
func JoinLines(lines []string) string {
	var b bytes.Buffer
	for i, s := range lines {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(s)
	}
	return b.String()
}
