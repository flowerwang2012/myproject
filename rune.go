package main

import "fmt"

func main() {
	//Go语言中不允许隐式转换，所有类型转换必须显式声明，使用类型前置加括号的方式
	var b byte = 97
	fmt.Println(string(b)) //结果：a

	//字符串 本质上是一个字节数组
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	str := "abc是我" //这个字符串里面有unicode字符，所以用rune类型

	//结果：9，ascii字符集（abc），每个字符占用一个字节，unicode字符集（是我），因为golang默认utf-8编码，中文每个字符占用3个字节
	//utf-8是编码规则，是一种变长编码规则，从1到4个字节不等，如果是ascii字符，就占用一个字节，而unicode字符，中文每个字符占用3个字节
	fmt.Println(len(str))

	//字符串可以按切片方式进行操作
	s := str[0:5]
	fmt.Println(s) //结果：abc�，字节不全
	fmt.Println(str[0]) //结果：97，切片下标012-abc，345-'是'，789-'我'

	//明显不能直接str[3]拿到'是'字符，需要用到rune类型，通过rune类型处理unicode字符
	r := []rune(str) //rune类型代表unicode码
	for _, v := range r {
		fmt.Println(v, string(v)) //v是字符集里字符的ID
		//97 a
		//98 b
		//99 c
		//26159 是
		//25105 我
	}
	r[3] = 'd' //改变字符的值
	fmt.Println(string(r)) //结果：abcd我
}