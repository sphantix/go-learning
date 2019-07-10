/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Wed 10 Jul 2019 11:12:37 PM CST
 */

package main

//void SayHello(const char* s);
import "C"

func main() {
	//FIXME: Memory leak
	C.SayHello(C.CString("Hello, World\n"))
}
