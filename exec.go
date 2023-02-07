package main

import "io"

func Exec(args []string, w io.Writer) error {
	if len(args) == 0 {
		return Help(w)
	}

	// TODO: delegate to kubectl
	return nil
}
