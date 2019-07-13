/*
 * File Name: main.go
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Sat 13 Jul 2019 01:37:11 PM CST
 */

package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
    char* os = "windows";
#elif defined(CGO_OS_DARWIN)
    char* os = "darwin";
#elif defined(CGO_OS_LINUX)
    char* os = "linux";
#else
#   error(unknown os)
#endif
*/
import "C"
import "fmt"

func main() {
	fmt.Printf("%s\n", C.GoString(C.os))
}
