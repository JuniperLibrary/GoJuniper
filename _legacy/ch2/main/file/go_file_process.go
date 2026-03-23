package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	/*
		Go 语言文件处理
			在Go语言中，文件处理是一个非常重要的功能，它允许我们读取、写入和操作文件。无论是处理配置文件、日志文件还是进行数据持久化，文件处理都是不可或缺的一部分。
			Go 语言提供可丰富的标准库来支持i文件处理，包括文件的打开、关闭、读取、写入、复制、移动、删除、获取文件信息、遍历目录等等。
			1. os 是核心库：提供了底层文件操作 创建、读取删除 大多数场景优先使用
			2. io 提供通用接口：如Reader/Writer 可与文件、网络等数据源交互
			3. buffer 优化性能：通过缓冲减少I/O次数，适合频繁读写
			4. ioutil 已弃用：Go 1.16后其功能迁移到 os 和io包
			5. path/filepath 提供文件路径处理: 跨平台兼容
	*/

	/*
		1. 文件创建
			在Go 语言中，创建文件可以使用 os 包的 Create 函数。
			os.Create 函数用于创建一个文件，并返回一个 *os.File 类型的文件对象。
			创建文件后，我们通常需要调用 Close 方法来关闭文件，以释放系统资源
			关闭文件是一个重要的步骤，它可以防止文件描述符泄漏。在 Go 中，我们通常使用 defer 语句来确保文件在函数结束时被关闭。
	*/
	// ，如果文件已存在会被截断（清空）
	//file, err := os.Create("test.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("File created:", file.Name())

	/*
		2.文件的读取
			Go 语言提供了多种读取文件的方式，包括逐行读取、一次性读取整个文件等。我们可以使用 bufio 包来逐行读取文件，
			或者使用 ioutil 包来一次性读取整个文件。
	*/

	fileOpen, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileOpen.Close()

	scanner := bufio.NewScanner(fileOpen)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	/*
		3. 一次性读取整个文件
	*/
	// file, err := ioutil.ReadFile("test.txt")
	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(file))

	/*
		4. 文件写入
			Go 语言提供了 WriteString 和 Write 方法来写入文件。
			WriteString 方法将字符串写入文件，Write 方法将字节切片写入文件。
	*/
	//
	fileWrite, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer fileWrite.Close()
	_, err = fileWrite.WriteString("直接写入字符串\\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	data := []byte("写入字节切片\\n")
	_, err = fileWrite.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println(fileWrite, "格式化写入: %d\\n\", 123")

	fmt.Println("====================逐行文件写入=============================")

	/*
		5.逐行文件写入
	*/
	fileColumn, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(fileColumn *os.File) {
		_ = fileColumn.Close()
	}(fileColumn)

	writer := bufio.NewWriter(fileColumn)
	_, _ = fmt.Fprintln(writer, "Hello, World!")
	_ = writer.Flush()

	fmt.Println("=====================一次性写入文件=========================")

	/*
		6.一次性写入文件
	*/
	content := []byte("Hello, World!一次性写入文件")
	err = os.WriteFile("output.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully!")

	/*
		7.文件的追加写入
			有时候我们需要在文件的末尾追加内容，而不是覆盖原有内容。Go 语言提供了 os.OpenFile 函数来实现这一功能。
	*/
	fileAppend, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileAppend.Close()

	if _, err := fileAppend.WriteString("追加写入"); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File appended successfully!")

	/*
		8.文件的删除
	*/
	//err = os.Remove("test.txt")
	//if err != nil {
	//	fmt.Println("Error deleting file:", err)
	//	return
	//}
	//fmt.Println("File deleted successfully!")

	/*
		9.文件信息与操作
	*/
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size(), "字节")
	fmt.Println("权限:", fileInfo.Mode())
	fmt.Println("最后修改时间:", fileInfo.ModTime())
	fmt.Println("是目录吗:", fileInfo.IsDir())

	/*
		10.检查文件是否存在
	*/
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件存在")
	}

	/*
		11.重命名和移动文件
	*/
	err = os.Rename("old.txt", "new.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("重命名成功")

	/*
		11.目录操作
	*/

	err = os.Mkdir("new_dir", 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("创建目录成功")
	// 递归创建多级目录
	err = os.MkdirAll("path/to/newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	/*
		12.读取目录内容
	*/
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("%-20s %8d %v\n",
			entry.Name(),
			info.Size(),
			info.ModTime().Format("2006-01-02 15:04:05"))
	}

	/*
		13.删除目录
	*/
	// 删除空目录
	err = os.Remove("emptydir")
	if err != nil {
		log.Fatal(err)
	}

	// 递归删除目录及其内容
	err = os.RemoveAll("path/to/dir")
	if err != nil {
		log.Fatal(err)
	}

	/*
		14.文件复制
	*/
	srcFile, err := os.Open("source.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("destination.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("复制完成，共复制 %d 字节", bytesCopied)

	/*
		15.文件追加
	*/
	fileAndFile, err := os.OpenFile("log.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fileAndFile.Close()

	if _, err := fileAndFile.WriteString("新的日志内容\n"); err != nil {
		log.Fatal(err)
	}

	/*
		16.临时文件和目录
	*/
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // 清理

	fmt.Println("临时文件:", tmpFile.Name())

	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "example-*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir) // 清理

	fmt.Println("临时目录:", tmpDir)
}
