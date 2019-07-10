/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Wed 10 Jul 2019 11:06:56 PM CST
 */

package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

func main() {
	//FIXME: Memory leak
	C.SayHello(C.CString("Hello, World\n"))
}
