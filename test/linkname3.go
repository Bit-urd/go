// errorcheck

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tests that errors are reported for misuse of linkname.
package p

import _ "unsafe"

type t int

var x, y int

//go:linkname x ok

// ERROR "//go:linkname must refer to declared function or variable"
// ERROR "//go:linkname must refer to declared function or variable"
// ERROR "duplicate //go:linkname for x"

//line linkname3.go:18
//go:linkname nonexist nonexist
//go:linkname t notvarfunc
//go:linkname x duplicate
