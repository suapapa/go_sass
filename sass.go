// Copyright 2012, Homin Lee All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
  Package sass is cgo binding of libsass.
*/
package sass

// #cgo LDFLAGS: -lsass -lstdc++
// #include "libsass_wrapper.h"
import "C"
import "errors"
import "fmt"
import "os"
import "path/filepath"
import "strings"

type Compiler struct {
	OutputStyle    int
	SourceComments bool
	IncludePaths   string
	ImagePath      string
}

// Compile scss to css
func (c *Compiler) Compile(source string) (string, error) {
	ctx, err := C._sass_new_context()
	if err != nil {
		errStr := "failed to alloc ctx: " + err.Error()
		return "", errors.New(errStr)
	}
	defer C._sass_free_context(ctx)

	c.fillOptions((*C.sass_options_t)(&ctx.options))
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

// Compile scss file to css
func (c *Compiler) CompileFile(path string) (string, error) {
	ctx, err := C._sass_new_file_context()
	if err != nil {
		errStr := "failed to alloc file ctx: " + err.Error()
		return "", errors.New(errStr)
	}
	defer C._sass_free_file_context(ctx)

	c.fillOptions((*C.sass_options_t)(&ctx.options))
	ctx.input_path = C.CString(path)

	C._sass_compile_file(ctx)
	if ctx.error_status != 0 {
		errStr := fmt.Sprintf("failed to compile file: %d: %s",
			ctx.error_status, C.GoString(ctx.error_message))
		return "", errors.New(errStr)
	}

	return C.GoString(ctx.output_string), nil
}

// Compile Sass files in in srcPath to CSS in outPath
func (c *Compiler) CompileFolder(srcPath, outPath string) error {
	ctx, err := C._sass_new_file_context()
	if err != nil {
		errStr := "failed to alloc file ctx: " + err.Error()
		return errors.New(errStr)
	}
	defer C._sass_free_file_context(ctx)

	walkF := func(p string, f os.FileInfo, e error) error {
		if f.IsDir() {
			return nil
		}

		if !strings.HasSuffix(p, ".scss") {
			return nil
		}

		base := filepath.Base(p)
		if strings.HasPrefix(base, "_") {
			// skip "_*.scss". they will be imported by other.
			return nil
		}

		dir := filepath.Dir(p)
		outDir := strings.Replace(dir, srcPath, outPath, 1)
		err := os.MkdirAll(outDir, 0770)
		if err != nil {
			return err
		}

		c.fillOptions((*C.sass_options_t)(&ctx.options))
		ctx.input_path = C.CString(p)
		C._sass_compile_file(ctx)
		if ctx.error_status != 0 {
			errStr := fmt.Sprintf("failed to compile file: %d: %s",
				ctx.error_status, C.GoString(ctx.error_message))
			return errors.New(errStr)
		}

		base = strings.TrimRight(base, ".scss") + ".css"
		outPath := filepath.Join(outDir, base)
		wp, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer wp.Close()

		n, err := wp.Write([]byte(C.GoString(ctx.output_string)))
		if err != nil {
			return err
		}
		if n == 0 {
			return errors.New("nothing written to " + p)
		}

		return nil
	}

	return filepath.Walk(srcPath, walkF)
}

func (c *Compiler) fillOptions(o *C.sass_options_t) {
	o.output_style = C.int(c.OutputStyle)
	if c.SourceComments {
		o.source_comments = 1
	} else {
		o.source_comments = 0
	}
	o.include_paths = C.CString(c.IncludePaths)
	o.image_path = C.CString(c.ImagePath)
}
