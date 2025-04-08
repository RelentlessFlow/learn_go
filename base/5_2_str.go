package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	b := "world"
	c := "你好_世界"
	// 比较大小
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	// 是否包含
	fmt.Println(strings.Contains(a, "he"))
	// 字符串长度
	fmt.Println(len(a))
	fmt.Println(len([]rune(a)))
	// 子串出现次数
	fmt.Println(strings.Count(a, "l"))
	// 字符串分隔
	fmt.Println(strings.Split(a, "e"))
	fmt.Printf("%#v\n", strings.Split(a, "e")) // []string{"h", "llo"}
	fmt.Printf("%#v\n", strings.Split(a, "_")) // []string{"hello"}
	// 是否包含前缀，后缀
	fmt.Println(strings.HasPrefix(a, "hel"))
	fmt.Println(strings.HasSuffix(b, "ld"))
	// 查询子串出现的位置
	fmt.Println(strings.Index(a, "ll"))
	fmt.Println(strings.IndexRune(a, 'l'))
	fmt.Println(strings.IndexRune(c, '_')) // 6 表示是第六个字节
	// 子串替换
	fmt.Println(strings.Replace(b, "d", "d!", -1)) // world!
	fmt.Println(b)                                 // world 不影响原字符串
	// 大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))
	// 去掉特殊字符
	fmt.Println(strings.Trim("   hello ", " "))         // hello
	fmt.Println(strings.Trim("###hello#", "#"))         // hello
	fmt.Println(strings.Trim("$#v#hello#", "$#v"))      // hello
	fmt.Println(strings.TrimLeft("$#v#hello#", "$#v"))  // hello
	fmt.Println(strings.TrimRight("$#v#hello#", "$#v")) // hello
}
