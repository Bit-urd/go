// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Darwin and Linux use the same linkage to main

TEXT _rt0_amd64_linux(SB),7,$-8
	MOVQ	$_rt0_amd64(SB), AX // 将_rt0_amd64地址移入ax寄存器
	MOVQ	SP, DI  // 将栈压入di寄存器
	JMP	AX
