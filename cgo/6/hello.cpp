/*
 * File Name: hello.cpp
 * Author: Sphantix
 * Mail: sphantix@gmail.cn
 * Created Time: Wed 10 Jul 2019 11:19:32 PM CST
 */

#include <iostream>
extern "C" {
#include "hello.h"
}

void SayHello(const char* s){
    std::cout << s;
}
