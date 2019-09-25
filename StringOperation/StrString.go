/*
@author : Aeishen
@data :  19/09/25, 17:51
@description :
*/

/*
题目：
	实现 strStr() 函数。给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
示例 1:
	输入: haystack = "hello", needle = "ll"
	输出: 2
示例 2:
	输入: haystack = "aaaaa", needle = "bba"
	输出: -1
说明:
	当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
	对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"
	haystack = "a"
	needle = "a"
	haystack ="mississippi"
	needle = "issip"
	fmt.Println(strStr(haystack , needle))
}

func strStr(haystack string, needle string) int {
	len_h := len(haystack)
	len_n := len(needle)
	if len_n == 0{
		return 0
	}

	for i := 0;i <= len_h - len_n; i++{
		if haystack[i] == needle[0]{
			for j:=0;j<len_n;j++{
				if haystack[i] == needle[j]{
					i ++
				}else{
					i = i - j
					break
				}
				if j == len_n-1{
					return i - len_n
				}
			}
		}
	}
	return -1
}