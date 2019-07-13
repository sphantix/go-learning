/*
 * File Name: hello.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Wed 10 Jul 2019 11:23:03 PM CST
 */

package main

import "C"

import "fmt"

//export SayHello
func SayHello(s *C.char) {
	fmt.Println(C.GoString(s))
}
