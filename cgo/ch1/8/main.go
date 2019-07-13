/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Wed 10 Jul 2019 11:27:15 PM CST
 */

// +build go1.10

package main

//void SayHello(_GoString_ s);
import "C"

import "fmt"

func main() {
	C.SayHello("Hello, World\n")
}

//export SayHello
func SayHello(s string) {
	fmt.Println(s)
}
