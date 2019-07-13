/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Wed 10 Jul 2019 11:16:41 PM CST
 */

package main

//#include "hello.h"
import "C"

func main() {
	//FIXME: ...
	C.SayHello(C.CString("Hello, World\n"))
}
