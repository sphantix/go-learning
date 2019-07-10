/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Wed 10 Jul 2019 11:04:36 PM CST
 */

package main

//#include <stdio.h>
import "C"

func main() {
	//FIXME: Memory leak
	C.puts(C.CString("Hello, Wrold\n"))
}
