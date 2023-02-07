package main

import (
	"fmt"
	"io"
)

func Exec(args []string, w io.Writer) error {
	if len(args) == 0 {
		return Help(w)
	}

	// TODO: delegate to kubectl
	aliases, err := GetResourceAliasMap()
	if err != nil {
		return fmt.Errorf("FATAL: failed to create resource alias map: %w", err)
	}
	err = ValidateResourceNames(args, aliases)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s\n", err)
		return nil
	}

	return nil
}
