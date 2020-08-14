// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy

import(
	_ "github.com/agentabstract/go-fitz/include/mupdf/fitz"
)
