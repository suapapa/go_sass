// Copyright 2012, Homin Lee. All rights reserved.                              
// Use of this source code is governed by a BSD-style                           
// license that can be found in the LICENSE file.

package sass

import "fmt"
import "log"
import "testing"

func TestBasicCompile(t *testing.T) {
	log.Println("TestBasicCompile")
	sc, _ := NewSass()
	css, err := sc.CompileToString("a { b { color: blue; } }")
	fmt.Println(css, err)
}

func TestBasicFile(t *testing.T) {
	log.Println("TestBasicFile")
	sc, _ := NewSass()
	css, err := sc.CompileFileToString("_scss/simple.scss")
	fmt.Println(css, err)
}

func TestCompileFileWithInclude(t *testing.T) {
	log.Println("TestCompileFileWithInclude")
	sc, _ := NewSass()
	css, err := sc.CompileFileToString("_scss/style.scss")
	fmt.Println(css, err)
}

func TestCompileFolder(t *testing.T) {
	sc, _ := NewSass()
	err := sc.CompileFolder("_scss", "css")
	fmt.Println(err)
}
