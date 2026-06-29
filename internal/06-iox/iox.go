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
	"os"
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

// ReadFile 一次性读取文件全部内容并返回 string。
func ReadFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// WriteFile 一次性将内容写入文件（覆盖写入）。
func WriteFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// AppendToFile 将内容追加写入文件末尾，文件不存在时会自动创建。
func AppendToFile(path, content string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

// CopyFile 复制文件，返回复制的字节数。
func CopyFile(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

// FileExists 检查文件或目录是否存在。
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

// ListDir 读取目录下的所有条目名称（文件与子目录）。
func ListDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name())
	}
	return names, nil
}
