package fitz

// this file exists purely to prevent the golang toolchain from stripping
// away the c source directories and files.
//
// how it works:
//  - every directory which only includes c source files receives a dummy.go file.
//  - every directory we want to preserve is included here as a _ import.
//  - this file is given a build to exclude it from the regular build.

import(
	//// prevent go tooling from stripping out the c source files.
	_ "github.com/agentabstract/go-fitz/include/mupdf"
	_ "github.com/agentabstract/go-fitz/include"
	_ "github.com/agentabstract/go-fitz/include/libs"
	_ "github.com/agentabstract/go-fitz/include"
)
