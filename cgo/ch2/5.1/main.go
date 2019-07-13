/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Sat 13 Jul 2019 10:57:49 PM CST
 */

package main

/*
#include <errno.h>

static int div(int a, int b) {
    if (b == 0 ) {
        errno = EINVAL;
        return 0;
    }

    return a/b;
}
*/
import "C"
import "fmt"

func main() {
	v0, err0 := C.div(6, 3)
	fmt.Println(v0, err0)

	v1, err1 := C.div(2, 0)
	fmt.Println(v1, err1)
}
