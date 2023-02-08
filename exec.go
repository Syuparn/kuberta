package main

import (
	"fmt"
	"io"
	"os/exec"
)

func Exec(args []string, w io.Writer) error {
	if len(args) == 0 {
		return Help(w)
	}

	resourceAliases, err := GetResourceAliasMap()
	if err != nil {
		return fmt.Errorf("FATAL: failed to create resource alias map: %w", err)
	}
	err = ValidateResourceNames(args, resourceAliases)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s\n", err)
		return nil
	}

	optionAliases, err := GetOptionAliasMap(args)
	if err != nil {
		return fmt.Errorf("FATAL: failed to create option alias map: %w", err)
	}
	err = ValidateOptions(args, optionAliases)
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s\n", err)
		return nil
	}

	delegateToKubectl(args, w)

	return nil
}

func delegateToKubectl(args []string, w io.Writer) error {
	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = w

	return cmd.Run()
}
