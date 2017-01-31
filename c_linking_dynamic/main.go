package main

/**
Before go build:

export PKG_CONFIG_PATH=$GOPATH/src/github.com/alessandrozucca/linking_c_to_go/c_linking_dynamic/pkgconfig
export DYLD_LIBRARY_PATH=$GOPATH/src/github.com/alessandrozucca/linking_c_to_go/c_linking_dynamic/dylib

Before running program
export DYLD_LIBRARY_PATH=$GOPATH/src/github.com/alessandrozucca/linking_c_to_go/c_linking_dynamic/dylib
*/

/*
#cgo pkg-config: --define-variable=prefix=. helloworld
#include <helloworld.h>
*/
import "C"

func main() {
	C.print_hello_world()
}
