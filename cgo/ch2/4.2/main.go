/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Sat 13 Jul 2019 03:47:28 PM CST
 */

// +build go1.10

package main

//void SayHello(_GoString_ s);
import "C"

import "fmt"

func main() {
	C.SayHello("Hello, World")
}

//export SayHello
func SayHello(s string) {
	fmt.Println(s)
}
