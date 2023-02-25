package main

import (
	"fmt"
	"io"
	"os/exec"
)

// Exec validates the specified arguments and delegates to kubectl.
// This returns error caused by kuberta itself and exit status.
func Exec(args []string, stdOutWriter, stdErrWriter io.Writer) (error, int) {
	if len(args) == 0 {
		return Help(stdOutWriter), 0
	}

	resourceAliases, err := GetResourceAliasMap()
	if err != nil {
		return fmt.Errorf("FATAL: failed to create resource alias map: %w", err), 1
	}
	err = ValidateResourceNames(args, resourceAliases)
	if err != nil {
		fmt.Fprintf(stdErrWriter, "ERROR: %s\n", err)
		// NOTE: ignore error because error messages have already been shown
		return nil, 1
	}

	optionAliases, err := GetOptionAliasMap(args)
	if err != nil {
		return fmt.Errorf("FATAL: failed to create option alias map: %w", err), 1
	}
	err = ValidateOptions(args, optionAliases)
	if err != nil {
		fmt.Fprintf(stdErrWriter, "ERROR: %s\n", err)
		// NOTE: ignore error because error messages have already been shown
		return nil, 1
	}

	err = delegateToKubectl(args, stdOutWriter, stdErrWriter)
	if err != nil {
		// NOTE: ignore error because error messages have already been shown
		return nil, 1
	}

	return nil, 0
}

func delegateToKubectl(args []string, stdOutWriter, stdErrWriter io.Writer) error {
	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = stdOutWriter
	cmd.Stderr = stdErrWriter

	return cmd.Run()
}
