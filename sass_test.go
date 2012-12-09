// Copyright 2012, Homin Lee. All rights reserved.                              
// Use of this source code is governed by a BSD-style                           
// license that can be found in the LICENSE file.

package sass

import "fmt"
import "testing"

func TestBasicCompile(t *testing.T) {
	sc, _ := NewSass()
	css, _ := sc.Compile("a { b { color: blue; } }")
	fmt.Println(css)
}
