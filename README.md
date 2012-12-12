# About sass

It's Go Binding for [libsass][4] which written in cpp by [hcailin][5].
It compile [Sass][3] file (.scss) to CSS.

# Documentation

## Prerequisites

[Install Go][1]

## Installation

    $ go get github.com/suapapa/go_sass

If you don't have libsass in your system, run;

    $ cd $GOROOT/src/pkg/github.com/suapapa/go_sass
    $ install_libsass.sh

I'll ask admin password to install the libsass system widely.

## General Documentation

Use `go doc` to vew the documentation for sass

    go doc github.com/suapapa/go_sass

Or alternatively, refer [go.pkgdoc.org][2]

## Example

Compile Sass folder into CSS:

    package main

    import (
            "github.com/suapapa/go_sass"
    )

    func main() {
            var sc sass.Compiler
            sc.CompileFolder("_scss", "css")
    }

# Author

Homin Lee &lt;homin.lee@suapapa.net&gt;

# Copyright & License

Copyright (c) 2012, Homin Lee.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.

[1]: http://golang.org/doc/install
[2]: http://go.pkgdoc.org/github.com/suapapa/go_sass
[3]: http://sass-lang.com/
[4]: https://github.com/hcatlin/libsass
[5]: https://github.com/hcatlin
