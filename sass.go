// Copyright 2012, Homin Lee All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sass

// #cgo LDFLAGS: -lsass -lstdc++
// #include "libsass_wrapper.h"
import "C"
import "errors"
import "fmt"
import "log"

type Sass struct {
	options C.sass_options_t
}

func NewSass() (sass *Sass, err error) {
	sass = new(Sass)
	err = nil
	return
}

func (c *Sass) Compile(source []byte) ([]byte, error) {
	css, err := c.CompileToString(string(source))
	return []byte(css), err
}

func (c *Sass) CompileToString(source string) (string, error) {
	ctx, err := C._sass_new_context()
	if err != nil {
		errStr := "failed to alloc ctx: " + err.Error()
		return "", errors.New(errStr)
	}
	defer C._sass_free_context(ctx)

	// TODO: assign options!
	// ctx.options = c.options
	ctx.source_string = C.CString(source)

	_, err = C._sass_compile(ctx)
	if err != nil {
		errStr := "failed to compile: " + err.Error()
		return "", errors.New(errStr)
	}

	if ctx.error_status != 0 {
		errStr := fmt.Sprintf("failed to compile: %d: %s",
			ctx.error_status, C.GoString(ctx.error_message))
		return "", errors.New(errStr)
	}

	return C.GoString(ctx.output_string), nil
}

func (c *Sass) CompileFile(path string) ([]byte, error) {
	css, err := c.CompileFileToString(path)
	return []byte(css), err
}

func (c *Sass) CompileFileToString(path string) (string, error) {
	ctx, err := C._sass_new_file_context()
	if err != nil {
		errStr := "failed to alloc file ctx: " + err.Error()
		return "", errors.New(errStr)
	}
	defer C._sass_free_file_context(ctx)

	/* ctx.options = c.options */
	ctx.input_path = C.CString(path)

	_, err = C._sass_compile_file(ctx)
	if err != nil {
		// TODO: don't know why error here when scss has import.
		// even it came here, the ctx has no error and output is good.
		errStr := "failed to compile file: " + err.Error()
		log.Println("Can we ignore this?: " + errStr)
		// return "", errors.New(errStr)
	}

	if ctx.error_status != 0 {
		errStr := fmt.Sprintf("failed to compile file: %d: %s",
			ctx.error_status, C.GoString(ctx.error_message))
		return "", errors.New(errStr)
	}

	return C.GoString(ctx.output_string), nil
}
