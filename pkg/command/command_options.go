package command

import (
	"io"
)

type Options struct {
	Executable string
	Args       []string
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
}


