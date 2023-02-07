package main

import (
	"fmt"
	"io"
)

const helpMessage = `
kuberta - kubectl short-coding supporter

usage: This is just a thin wrapper of kubectl. If you forget to use shortcut alias, this raises an error.
`

func Help(w io.Writer) error {
	fmt.Fprint(w, helpMessage)
	return nil
}
