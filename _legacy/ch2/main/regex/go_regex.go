package main

import (
	"fmt"
	"regexp"
)

func main() {

	/*
		1.Go 语言正则表达式
			正则表达式（Regular Expression，简称 regex 或 regexp）是一种用于匹配字符串的强大工具。
			正则表达式通过定义一种模式（pattern），可以快速搜索、替换或提取符合该模式的字符串，详细可以参见正则表达式教程。
			在 Go 语言中，正则表达式通过 regexp 包来实现

		2.Go 语言的标准库提供了 regexp 包，用于处理正则表达式。以下是 regexp 包中常用的函数和方法：
			Compile 和 MustCompile
				用于编译正则表达式。Compile 返回一个 *Regexp 对象和一个错误，而 MustCompile 在编译失败时会直接 panic。
			MatchString
				检查字符串是否匹配正则表达式。
			FindString 和 FindAllString
				用于查找匹配的字符串。FindString 返回第一个匹配项，FindAllString 返回所有匹配项。
			ReplaceAllString
				用于替换匹配的字符串。
			Split
				根据正则表达式分割字符串。
		3.正则表达式的基本语法
			.：匹配任意单个字符（除了换行符）。
			*：匹配前面的字符 0 次或多次。
			+：匹配前面的字符 1 次或多次。
			?：匹配前面的字符 0 次或 1 次。
			\d：匹配数字字符（等价于 [0-9]）。
			\w：匹配字母、数字或下划线（等价于 [a-zA-Z0-9_]）。
			\s：匹配空白字符（包括空格、制表符、换行符等）。
			[]：匹配括号内的任意一个字符（例如 [abc] 匹配 a、b 或 c）。
			^：匹配字符串的开头。
			$：匹配字符串的结尾。
	*/

	/*
		1.检查字符串是否匹配正则表达式
	*/
	pattern := `^[a-zA-Z0-9]+$`
	// 创建正则表达式对象
	regex := regexp.MustCompile(pattern)
	str := "Hello123"
	if regex.MatchString(str) {
		fmt.Println("字符串匹配正则表达式")
	} else {
		fmt.Println("字符串不匹配正则表达式")
	}

	/*
		2.查找匹配的字符串
	*/
	patternFindStr := `\d+`
	regexFindStr := regexp.MustCompile(patternFindStr)

	findStr := "我有 3 个苹果和 5 个香蕉"
	matches := regexFindStr.FindAllString(findStr, -1)
	fmt.Println("找到的数字：", matches)

	/*
		3.替换匹配的字符串
	*/
	patternReplaceStr := `\s+`
	regexReplaceStr := regexp.MustCompile(patternReplaceStr)

	strReplace := "Hello    World"
	result := regexReplaceStr.ReplaceAllString(strReplace, " ")
	fmt.Println("替换后的字符串：", result)

	/*
		4.分割字符串
	*/
	patternSplit := `,`
	regexSplit := regexp.MustCompile(patternSplit)

	strSplit := "apple,banana,orange"
	parts := regexSplit.Split(strSplit, -1)
	fmt.Println("分割后的字符串：", parts)

}
