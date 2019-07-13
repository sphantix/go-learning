/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Sat 13 Jul 2019 11:05:34 PM CST
 */

package main

//static void noreturn() {}
import "C"

import "fmt"

func main() {
	v, err := C.noreturn()
	fmt.Printf("%T\n", v)
	fmt.Printf("%#v\n", v)
	fmt.Printf("%v\n", v)
	fmt.Printf("%T\n", err)
	fmt.Printf("%v\n", err)
}
