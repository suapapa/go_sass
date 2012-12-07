// Copyright 2012, Homin Lee All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* package sass */
package main

// #cgo LDFLAGS: -lsass -lstdc++
// #include "libsass_wrapper.h"
import "C"
import "errors"
import "fmt"
import "log"


type Sass struct {
	/* context C.sass_context */
	options C.sass_options_t
	/* options C.sass_options */
}

func NewSass() (sass *Sass, err error) {
	sass = new(Sass)
	err = nil
	return
}

func (c *Sass) Compile(source string) (string, error) {
	context, err := C._sass_new_context()
	if err != nil {
		e := errors.New("failed to make context for compile:" + err.Error())
		return "", e
	}
	defer C._sass_free_context(context)

	/* context.options = c.options */
	context.source_string = C.CString(source)

	_, err = C._sass_compile(context)
	if err != nil {
		e := errors.New("failed to compile:" + err.Error())
		return "", e
	}

	if context.error_status != 0 {
		e := errors.New(fmt.Sprintf("failed to compile: %d", context.error_status))
		return "", e
	}

	return C.GoString(context.output_string), nil
}

func main() {
	log.Println("TestCompile")
	sass, _ := NewSass()
	css, _ := sass.Compile("a { b { color: blue; } }")
	log.Println(css)
}
