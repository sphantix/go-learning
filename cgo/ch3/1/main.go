/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Sun 14 Jul 2019 10:57:24 PM CST
 */

package main

//int sum(int a, int b) {return a+b;}
import "C"

func main() {
	println(C.sum(1, 1))
}
