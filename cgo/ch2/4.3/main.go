/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Sat 13 Jul 2019 03:54:38 PM CST
 */

package main

/*
#include <stdio.h>

struct A {
    int i;        //普通变量
    int type;     //go 关键字
    float _type;  //无法导出
};

struct A a = {1, 2, 3};

union B {
    int i;
    float f;
};

enum C {
    ONE,
    TWO,
};

void print_struct() {
    printf("i = %d, type = %d, _type = %f\n", a.i, a.type, a._type);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	C.print_struct()
	fmt.Printf("C.a.i = %v\n", C.a.i)             //对应i
	fmt.Printf("C.a._type = %v\n", C.a._type)     //对应_type
	fmt.Printf("C.a.__type = %v\n\n", C.a.__type) //对应type

	var b C.union_B
	fmt.Printf("C.union_B.i = %v\n", *(*C.int)(unsafe.Pointer(&b)))
	fmt.Printf("C.union_B.f = %v\n\n", *(*C.float)(unsafe.Pointer(&b)))

	var c C.enum_C = C.TWO
	fmt.Printf("C.enum_C = %v\n", c)
	fmt.Printf("C.ONE = %v\n", C.ONE)
	fmt.Printf("C.TWO = %v\n", C.TWO)
}
