/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Sat 13 Jul 2019 03:41:22 PM CST
 */

package main

/*
#include <stdio.h>

void printlong(long l) {
    printf("pringlong: %ld\n", l);
}
*/
import "C"

func main() {
	l := 25
	var ll int32 = 26
	C.printlong(C.long(l))
	C.printlong(C.long(ll))
}
