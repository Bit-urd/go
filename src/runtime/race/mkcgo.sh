#!/bin/bash

hdr='
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by mkcgo.sh. DO NOT EDIT.

//go:build race

'

convert() {
	(echo "$hdr"; go tool cgo -dynpackage race -dynimport $1) | gofmt
}

convert race_darwin_arm64.syso >race_darwin_arm64.go
convert internal/amd64v1/race_darwin.syso >race_darwin_amd64.go

