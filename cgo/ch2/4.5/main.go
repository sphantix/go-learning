/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Sat 13 Jul 2019 10:44:43 PM CST
 */

package main

/*
const char* p = "Hello, World!";
*/
import "C"

import "fmt"
import "unsafe"

func main() {
	r := uint64(uintptr(unsafe.Pointer(C.p)))
	fmt.Printf("0x%x %T\n", r, r)
}
