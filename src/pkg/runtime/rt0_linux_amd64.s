// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Darwin and Linux use the same linkage to main

TEXT _rt0_amd64_linux(SB),7,$-8  // // dlv list  入口函数 ,
	MOVQ	$_rt0_amd64(SB), AX  // -> src/pkg/runtime/asm_amd64.s:8
	MOVQ	SP, DI
	JMP	AX
